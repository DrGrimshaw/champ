package csv_cli_tool

import (
	"encoding/csv"
	"os"
)

func getHead(filePath string) ([]string, error) {
	if !fileExists(filePath) {
		return nil, os.ErrNotExist
	}

	csvFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1
	head, err := reader.Read()
	if err != nil {
		return nil, err
	}

	return head, nil
}
