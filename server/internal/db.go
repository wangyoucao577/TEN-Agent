package internal

import (
	"encoding/csv"
	"fmt"
	"io"
	"log/slog"
	"os"
)

type DB struct {
	m map[string]UserInfo

	path string // .csv
}

func NewDatabase(path string) *DB {
	return &DB{
		m:    map[string]UserInfo{},
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

	// read csv per line
	var _headers []string
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("read line failed, err %v", err)
		}
		if len(_headers) == 0 {
			_headers = record // ignore headers line
			continue
		}

		u, err := NewUserInfoFromCSVRecord(record)
		if err != nil {
			return fmt.Errorf("create userinfo from csv record failed, err %v", err)
		}

		d.m[u.ID] = u
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
