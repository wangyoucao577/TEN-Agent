package internal

import (
	"encoding/csv"
	"fmt"
	"io"
	"log/slog"
	"os"
)

const (
	CustomerFieldKeyID = "id"
)

type CustomerFields map[string]string // fields name mapping, EN->CN

func (c CustomerFields) Get() map[string]string {
	return c
}

func (c CustomerFields) GetReversed() map[string]string {
	r := CustomerFields{}
	for k, v := range c {
		r[v] = k
	}
	return r
}

func LoadCustomerFieldsFromCSV(path string) (CustomerFields, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open csv file %s failed, err %v", path, err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	customerFields := CustomerFields{}

	// read csv per line
	var headers []string
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("read line failed, err %v", err)
		}
		if len(headers) == 0 {
			headers = record
			if len(headers) < 2 {
				return nil, fmt.Errorf("invalid csv format, at least two columns")
			}
			if headers[0] != "EN" || headers[1] != "CN" {
				return nil, fmt.Errorf("invalid csv format, columns MUST be 'EN,CN'")
			}
			continue
		}

		// EN -> CN
		customerFields[record[0]] = record[1]
	}

	slog.Info("customer fields loaded", slog.String("path", path), slog.Int("count", len(customerFields)), slog.Any("fields", customerFields))
	return customerFields, nil
}
