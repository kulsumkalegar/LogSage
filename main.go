package main

import (
	"fmt"
	"logsage/analyser"
	"logsage/reader"
	"logsage/report"
)

func main() {
	lines, err := reader.ReadFile("server.log")
	if err != nil {
		fmt.Println(err)
		return
	}
	entries, ERROR, INFO, WARNING, UNKNOWN, analyserErr := analyser.Analyse(lines)
	if analyserErr != nil {
		fmt.Println(analyserErr)
		return
	}
	report.Generate(entries, ERROR, INFO, WARNING, UNKNOWN)

}
