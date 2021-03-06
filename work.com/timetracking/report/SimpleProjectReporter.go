package report

import (
	. "../data"
	"fmt"
)

type SimpleProjectReporter struct {
}

var simpleReporter SimpleProjectReporter

func GetSimpleProjectReporter() SimpleProjectReporter {
	return simpleReporter
}

func (this SimpleProjectReporter) ExportReport(pi TimeTrackingReport) {
	fmt.Println("ProjectName; Hours; Percent")
	for i := range pi.GetAllSettings() {
		entry := pi.GetEntry(i)
		timeEntry := entry.GetTimeEntry()
		if timeEntry.ToFloat64InHours() > 0 {
			hours := convertFloat64ToString(timeEntry.ToFloat64InHours())
			percents := convertFloat64ToString(timeEntry.GetInPercent())
			fmt.Printf("%s; %s; %s\n", entry.GetProject(), hours, percents+"%")
		}
	}

	fmt.Println("OVERALL; " + convertFloat64ToString(pi.SumOfAllProjectTimes) + "; " + convertFloat64ToString(100.0))
}
