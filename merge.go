package csv_cli_tool

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func mergeHeadWithFiles(head []string, dataFileNames []string, outputFolder string, trackingFileName string) error {
	if !fileExists(trackingFileName) {
		f, _ := os.Create(trackingFileName)
		f.Close()
	}

	trackingFile, err := os.OpenFile(trackingFileName, os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		return err
	}
	defer trackingFile.Close()

	_, err = trackingFile.WriteString(fmt.Sprintf("MERGE STARTING... at: %s\\n", time.Now().String()))
	if err != nil {
		return err
	}

	for _, dataFileName := range dataFileNames {
		fmt.Println("Merging file: " + dataFileName)
		if err = mergeHeadWithFile(head, dataFileName, outputFolder); err != nil {
			trackingFile.WriteString(fmt.Sprintf("filename: %s - could not be merged, err: %s\\n", dataFileName, err.Error()))
			continue
		}
		trackingFile.WriteString(fmt.Sprintf("filename: %s - merged successfully\\n", dataFileName))
	}

	trackingFile.WriteString(fmt.Sprintf("MERGE FINISHED... files processed:%d at: %s\\n", len(dataFileNames), time.Now().String()))

	trackingFile.Sync()

	return nil
}

func mergeHeadWithFile(head []string, dataFileName string, outputFolder string) error {
	outputFileName := fmt.Sprintf("%s/%s", outputFolder, filepath.Base(dataFileName))
	if fileExists(outputFileName) {
		return fmt.Errorf("output file already exists, program does not overwrite data, filename: %s", outputFileName)
	}

	dataFile, err := os.Open(dataFileName)
	if err != nil {
		return err
	}

	defer dataFile.Close()

	reader := csv.NewReader(dataFile)
	reader.FieldsPerRecord = -1
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	if len(records) == 0 {
		return fmt.Errorf("file does not have any data")
	}

	firstRecord := records[0]
	if len(head) != len(firstRecord) {
		return fmt.Errorf("number of columns in head and data file do not match head count: %d record count: %d", len(head), len(firstRecord))
	}

	header := [][]string{head}

	records = append(header, records...)

	outputFile, err := os.Create(outputFileName)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	writer := csv.NewWriter(outputFile)
	for _, record := range records {
		if err = writer.Write(record); err != nil {
			return err
		}
	}
	writer.Flush()

	return nil
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
