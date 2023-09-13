package main

import (
	"flag"
	"fmt"
	csvclitool "github.com/DrGrimshaw/csv-cli-tool"
)

func main() {
	headerFile := flag.String("headerFile", "", "This is the file that contains the headers for the CSV file(s)")
	inputPath := flag.String("inputPath", "", "This is the file or directory that contains the CSV file(s)")
	outputFolder := flag.String("outputFolder", "", "This is the directory that will contain the output files with the header")
	trackingFileName := flag.String("trackingFileName", "", "This is the file that becomes your log file for the merge")
	startingRow := flag.Int("startingRow", 0, "This determines the file the merge starts at")
	endingRow := flag.Int("endingRow", 0, "This determines the file the merge ends at")
	help := flag.Bool("help", false, "This is to show you what the options for the program are")
	helpShort := flag.Bool("h", false, "This is to show you what the options for the program are")
	flag.Parse()

	if *help || *helpShort {
		flag.Usage()
		return
	}

	if (*startingRow == 0 && *endingRow != 0) || (*startingRow != 0 && *endingRow == 0) {
		fmt.Println("You must specify both startingRow and endingRow not just one")
		fmt.Println("Exiting application")
		return
	}

	err := csvclitool.App(
		*headerFile,
		*inputPath,
		*outputFolder,
		*trackingFileName,
		*startingRow,
		*endingRow)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Merge Application Ran Without Issue")
	fmt.Println("Exiting application")
}
