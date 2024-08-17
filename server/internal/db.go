package internal

import (
	"crypto/sha1"
	"encoding/csv"
	"encoding/hex"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"
)

type DB struct {
	m map[string]CustomerInfo

	path string // .csv
}

func NewDatabase(path string) *DB {
	return &DB{
		m:    map[string]CustomerInfo{},
		path: path,
	}
}

func (d *DB) Load() error {

	file, err := os.Open(d.path)
	if err != nil {
		return fmt.Errorf("open csv file %s failed, err %v", d.path, err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// get cn->en fields key mapping
	customerInfoFieldsMappingCN2EN := map[string]string{}
	for k, v := range customerInfoFieldsMappingEN2CN {
		customerInfoFieldsMappingCN2EN[v] = k
	}

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
		sum := sha1.Sum([]byte(strings.Join(record, ",")))
		user["id"] = hex.EncodeToString(sum[:])

		d.m[user["id"]] = user
	}

	slog.Info("database loaded", slog.String("path", d.path), slog.Int("count", len(d.m)))
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
