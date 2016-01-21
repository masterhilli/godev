package report

import (
	. "../arguments"
	. "../data"
	. "../jira/Config"
	"strconv"
)

type ReporterInterface interface {
	ExportReport(pi TimeTrackingReport)
}

func GetReporter(separator rune, config Config) ReporterInterface {
	args := GetArguments()
	if args.GetReporterId() == 0 {
		return GetSimpleProjectReporter()
	}
	ccsRep := GetCCSReporter(separator, config.Reportname)
	return ccsRep
}

func convertFloat64ToString(floatToConvert float64) string {
	return strconv.FormatFloat(floatToConvert, 'f', 2, 64)
}
