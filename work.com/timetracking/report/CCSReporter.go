package report

import (
	. "../data"
	"fmt"
)

var reporter CCSReporter

func GetCCSReporter() CCSReporter {
	return reporter
}

type CCSReporter struct {
	separator rune
}

func (this *CCSReporter) setSeparator(separator rune) {
	this.separator = separator
}

func (r CCSReporter) ExportReport(pi TimeTrackingReport) {
	r.printValuesInCSVFormatSSS("Team", "Members", "Projectname", "Hours", "Percent")
	r.printAllProjectTimeEntries(pi)
	r.printValuesInCSVFormatSFF(pi.GetReportName(), "All", "OVERALLTIME", pi.SumOfAllProjectTimes, 100.0)
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
	r.printValuesInCSVFormatSSS(teamName, teamMembers, projectname, convertFloat64ToString(hours), convertFloat64ToString(percent)+"%")
}
func (r CCSReporter) printValuesInCSVFormatSSS(teamName string, teamMembers string, projectname string, hours string, percent string) {
	fmt.Printf("%s%c%s%c%s%c%s%c%s%c\n", teamName, r.separator, teamMembers, r.separator, projectname, r.separator, hours, r.separator, percent, r.separator)
}

func (r CCSReporter) printValuesInCSVFormatPersTime(prjTime ProjectReportSetting, reportName string) {
	productOwner := prjTime.GetProductOwner()
	timeEntry := prjTime.GetTimeEntry()
	teamMembers := timeEntry.GetTeamMembersCommaSeperated(productOwner)
	r.printValuesInCSVFormatSFF(reportName, teamMembers+","+productOwner, prjTime.GetProject(), timeEntry.ToFloat64InHours(), timeEntry.GetInPercent())
}
