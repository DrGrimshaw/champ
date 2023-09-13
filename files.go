package csv_cli_tool

import (
	"fmt"
	"os"
	"path/filepath"
)

func getCSVFileOrFilesInPath(path string, startRow int, endRow int) ([]string, error) {
	if !isDir(path) {
		fileName := path
		if err := isCSVFile(fileName); err != nil {
			return nil, err
		}

		return []string{fileName}, nil
	}

	files, err := findCSVFiles(path, startRow, endRow)
	if err != nil {
		return nil, err
	}
	return files, nil
}

func isCSVFile(fileName string) error {
	matched, err := filepath.Match("*.csv", filepath.Base(fileName))
	if err != nil {
		return err
	}

	if !matched {
		return fmt.Errorf("%s is not a csv file", fileName)
	}

	return nil
}

func isDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		fmt.Println(err)
	}
	return err == nil && info.IsDir()
}

func findCSVFiles(root string, startRow int, endRow int) ([]string, error) {
	var matches []string
	i := 1
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if err := isCSVFile(path); err != nil {
			fmt.Println(err)
			return nil
		}

		if (i >= startRow && i <= endRow) || (startRow == -1 && endRow == -1) {
			matches = append(matches, path)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return matches, nil
}
