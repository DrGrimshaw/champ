package main

import (
	"fmt"
	csvclitool "github.com/DrGrimshaw/csv-cli-tool"
)

func main() {
	err := csvclitool.App(
		"./test_data/header.csv",
		"./test_data/in",
		"./test_data/out",
		"./test_data/log.txt",
		0,
		0)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Merge Application Ran Without Issue")
}
