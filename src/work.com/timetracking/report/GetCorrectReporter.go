package report

import (
	. "work.com/timetracking/arguments"
	. "work.com/timetracking/data"
	"strconv"
	"fmt"
)

type ReporterInterface interface {
	ExportReport(pi TimeTrackingReport)
}

func GetReporter(separator rune) ReporterInterface {
	args := GetArguments()
	if args.GetReporterId() == 0 {
		fmt.Println("Simple Report created!")
		return GetSimpleProjectReporter()
	}
	ccsRep :=  GetCCSReporter()
	ccsRep.setSeparator(separator)
	return ccsRep
}


func convertFloat64ToString(floatToConvert float64) string {
	return strconv.FormatFloat(floatToConvert , 'f', 2, 64)
}