package report

import (
	"strconv"
	"fmt"
	. "work.com/timetracking/data"
	. "work.com/timetracking/jira/Timeentry"
)

type CCSReporter struct {
	separator rune
}

func (this *CCSReporter) SetSeparator(separator rune) {
	this.separator = separator
}

func (r CCSReporter) ExportReport(pi TimeTrackingReport) {
	pi.Finish()
	r.printValuesInCSVFormatSSS("Team", "Members", "Projectname", "Hours", "Percent")
	r.printAllProjectTimeEntries(pi)
	r.printValuesInCSVFormatSFF("LCC eServices Region South-East", "All", "OVERALLTIME", pi.SumOfAllProjectTimes, 100.0)
}

func (this CCSReporter) printAllProjectTimeEntries(reportData TimeTrackingReport) {
	for i := range reportData.GetAllSettings() {
		entry := reportData.GetEntry(i)
		timeEntry := entry.GetTimeEntry()
		if timeEntry.ToFloat64InHours() > 0.0 {
			this.printValuesInCSVFormatPersTime(timeEntry)
		}
	}
}

func (r CCSReporter) printValuesInCSVFormatSFF(teamName string, teamMembers string, projectname string, hours float64, percent float64) {
	r.printValuesInCSVFormatSSS(teamName, teamMembers, projectname, strconv.FormatFloat(hours, 'f', 2, 64), strconv.FormatFloat(percent, 'f', 1, 64)+"%")
}
func (r CCSReporter) printValuesInCSVFormatSSS(teamName string, teamMembers string, projectname string, hours string, percent string) {
	fmt.Printf("%s%c%s%c%s%c%s%c%s%c\n", teamName, r.separator, teamMembers, r.separator, projectname, r.separator, hours, r.separator, percent, r.separator)
}

func (r CCSReporter) printValuesInCSVFormatPersTime(prjTime TimeEntry) {
	fmt.Printf("%s\n", prjTime.ToCsvFormat(r.separator))
}
