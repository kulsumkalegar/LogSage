package report

import (
	"fmt"
	"logsage/analyser"
)

func Generate(entries []analyser.LogEntry, ERROR, INFO, WARNING, UNKNOWN int) {
	fmt.Println("=======Log Report======")
	fmt.Println("Entries: ", len(entries))
	fmt.Println("ERROR: ", ERROR)
	fmt.Println("INFO: ", INFO)
	fmt.Println("WARNING: ", WARNING)
	fmt.Println("UNKNOWN: ", UNKNOWN)
	fmt.Println("=======================")

}
