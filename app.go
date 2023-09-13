package csv_cli_tool

import "fmt"

const defaultOutput = "out"
const defaultInput = "in"
const defaultTrackingFileName = "log.txt"
const defaultHeaderCSV = "header.csv"
const defaultStartRow = -1
const defaultEndRow = -1

func App(headerFile string, input string, outputFolder string, trackingFileName string, startRow int, endRow int) (err error) {
	if input == "" {
		input = defaultInput
	}
	if outputFolder == "" {
		outputFolder = defaultOutput
	}
	if trackingFileName == "" {
		trackingFileName = defaultTrackingFileName
	}
	if headerFile == "" {
		headerFile = defaultHeaderCSV
	}
	if startRow < 1 {
		startRow = defaultStartRow
	}
	if endRow < 1 {
		endRow = defaultEndRow
	}

	head, err := getHead(headerFile)
	if err != nil {
		return err
	}
	fmt.Println("Got header: ", head)

	files, err := getCSVFileOrFilesInPath(input, startRow, endRow)
	if err != nil {
		return err
	}

	fmt.Println("Got files: ", files)

	return mergeHeadWithFiles(head, files, outputFolder, trackingFileName)
}
