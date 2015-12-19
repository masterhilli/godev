package report

import (
	. "../arguments"
	. "../data"
	"strconv"
	"fmt"
)

type ReporterInterface interface {
	ExportReport(pi TimeTrackingReport)
}

func GetReporter(separator rune) ReporterInterface {
	args := GetArguments()
	if args.GetReporterId() == 0 {
		return GetSimpleProjectReporter()
	}
	ccsRep :=  GetCCSReporter()
	ccsRep.setSeparator(separator)
	return ccsRep
}


func convertFloat64ToString(floatToConvert float64) string {
	return strconv.FormatFloat(floatToConvert , 'f', 2, 64)
}