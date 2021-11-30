package cert

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ParseCSV(filename string) ([]*Cert, error) {
	certs := make([]*Cert, 0)
	f, err := os.Open(filename)
	if err != nil {
		return certs, err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Printf("Error while closing, %v\n", err)
		}
	}(f)

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		return certs, err
	}

	for _, rec := range records {
		course := rec[0]
		name := rec[1]
		date := rec[2]

		c, err := New(course, name, date)
		if err != nil {
			return certs, err
		}
		certs = append(certs, c)
	}

	return certs, nil
}
