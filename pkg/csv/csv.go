package csv

import (
	"encoding/csv"
	"os"
)

// Read reads a csv file into an array of maps, where each row in the csv
// corresponds to an element in the array. Requires that the csv file contains
// a header row. Header names are the keys of the map.
func Read(filename string) ([]map[string]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	headers, err := reader.Read()
	if err != nil {
		return nil, err
	}

	var rows []map[string]string
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		row := make(map[string]string)
		for i, value := range record {
			row[headers[i]] = value
		}

		rows = append(rows, row)
	}

	return rows, nil
}
