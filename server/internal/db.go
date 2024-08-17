package internal

import (
	"encoding/csv"
	"fmt"
	"io"
	"log/slog"
	"os"
)

type DB struct {
	m map[string]CustomerInfo

	fields  CustomerFields
	headers []string

	infoPath          string // customerinfo.csv
	fieldsPath        string // customerfields.csv
	generatedInfoPath string // customerinfo.csv.generated.csv write generated in a seperated file
}

func NewDatabase(infoPath, fieldsPath string) *DB {
	return &DB{
		m: map[string]CustomerInfo{},

		fields:  nil,
		headers: []string{},

		infoPath:          infoPath,
		fieldsPath:        fieldsPath,
		generatedInfoPath: infoPath + ".generated.csv",
	}
}

func (d *DB) Load() error {
	if fields, err := LoadCustomerFieldsFromCSV(d.fieldsPath); err != nil {
		return fmt.Errorf("load customer fields from csv file %s failed, err %v", d.fieldsPath, err)
	} else {
		d.fields = fields
	}

	file, err := os.Open(d.infoPath)
	if err != nil {
		return fmt.Errorf("open csv file %s failed, err %v", d.infoPath, err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// get cn->en fields key mapping
	customerInfoFieldsMappingCN2EN := d.fields.GetReversed()

	// read csv per line
	var headers []string
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("read line failed, err %v", err)
		}
		if len(headers) == 0 {
			headers = record
			for _, h := range headers {
				if en, ok := customerInfoFieldsMappingCN2EN[h]; ok {
					d.headers = append(d.headers, en)
				} else {
					d.headers = append(d.headers, "") // append "" as placeholder to make sure correct columns
				}
			}
			continue
		}

		// generate user info
		user := CustomerInfo{}
		for i, h := range headers {
			if en, ok := customerInfoFieldsMappingCN2EN[h]; ok {
				user[en] = record[i]
			} else {
				slog.Warn("ignored unknown header", slog.String("header", h))
			}
		}

		// generate user id
		user[CustomerFieldKeyID] = GenerateCustomerID(user)

		d.m[user[CustomerFieldKeyID]] = user
	}

	slog.Info("database loaded", slog.String("path", d.infoPath), slog.Int("count", len(d.m)))
	i := 0
	for k, v := range d.m {
		slog.Info("record peek", slog.Any(k, v))
		i++
		if i >= 3 {
			break
		}
	}
	return nil
}

func (d *DB) List(fields []string) []CustomerInfo {
	if len(fields) == 0 {
		return nil
	}

	users := []CustomerInfo{}

	fieldsSet := map[string]struct{}{}
	for _, f := range fields {
		fieldsSet[f] = struct{}{}
	}

	for _, v := range d.m {
		outputUser := CustomerInfo{}
		for k, v := range v {
			if _, ok := fieldsSet[k]; ok { // only return sepcific keys
				outputUser[k] = v
			}
		}

		if !outputUser.Empty() {
			users = append(users, outputUser)
		}
	}

	return users
}

func (d *DB) Get(id string) CustomerInfo {
	if u, ok := d.m[id]; ok {
		return u
	}
	return nil
}

func (d *DB) Write(u CustomerInfo) error {
	// save in memory
	d.m[u[CustomerFieldKeyID]] = u

	file, err := os.OpenFile(d.generatedInfoPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("open csv file %s failed, err %v", d.infoPath, err)
	}
	defer file.Close()

	records := []string{}
	for _, h := range d.headers {
		if v, ok := u[h]; ok {
			records = append(records, v)
		} else {
			records = append(records, "") // append "" as placeholder to make sure correct columns
		}
	}

	writer := csv.NewWriter(file)
	if err := writer.Write(records); err != nil {
		return fmt.Errorf("write csv file %s failed, err %v", d.generatedInfoPath, err)
	}
	writer.Flush()
	return nil
}
