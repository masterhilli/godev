package report

import (
	. "../data"
)

var reporter CCSReporter

func GetCCSReporter(separator rune) CCSReporter {
	reporter.Initialize(separator)
	return reporter
}

type CCSReporter struct {
	writer        Writer
	isInitialized bool
}

func (this *CCSReporter) Initialize(separator rune) {
	if !this.isInitialized {
		this.writer = NewExcelWriter()
		//this.writer = NewCmdLineWriter()
		this.writer.Initialize([]string{string(separator)})
		this.isInitialized = true
	}
}

func (this CCSReporter) ExportReport(pi TimeTrackingReport) {
	this.writer.PrintLine("Team", "Members", "Projectname", "Hours", "Percent")
	this.printAllProjectTimeEntries(pi)
	this.printValuesInCSVFormatSFF(pi.GetReportName(), "All", "OVERALLTIME", pi.SumOfAllProjectTimes, 100.0)
	this.writer.Close()
}

func (this CCSReporter) printAllProjectTimeEntries(reportData TimeTrackingReport) {
	for i := range reportData.GetAllSettings() {
		entry := reportData.GetEntry(i)
		timeEntry := entry.GetTimeEntry()
		if timeEntry.ToFloat64InHours() > 0.0 {
			this.printValuesInCSVFormatPersTime(entry, reportData.GetReportName())
		}
	}
}

func (r CCSReporter) printValuesInCSVFormatSFF(teamName string, teamMembers string, projectname string, hours float64, percent float64) {
	r.writer.PrintLine(teamName, teamMembers, projectname, convertFloat64ToString(hours), convertFloat64ToString(percent)+"%")
}

func (r CCSReporter) printValuesInCSVFormatPersTime(prjTime ProjectReportSetting, reportName string) {
	productOwner := prjTime.GetProductOwner()
	timeEntry := prjTime.GetTimeEntry()
	teamMembers := timeEntry.GetTeamMembersCommaSeperated(productOwner)
	r.printValuesInCSVFormatSFF(reportName, teamMembers+","+productOwner, prjTime.GetProject(), timeEntry.ToFloat64InHours(), timeEntry.GetInPercent())
}
