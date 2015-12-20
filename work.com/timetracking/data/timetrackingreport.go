package data

import (
	. "../arguments"
	. "../jira/Timeentry"
	"fmt"
	"strings"
)


func NewTimeTrackingReport(countProjects int) TimeTrackingReport {
	var retVal TimeTrackingReport
	retVal.settings = make(map[string]ProjectReportSetting, countProjects)
	return retVal
}

type TimeTrackingReport struct {
	settings             map[string]ProjectReportSetting
	SumOfAllProjectTimes float64
	teammembers          map[string]bool
}

func (this *TimeTrackingReport) GetAllSettings() map[string]ProjectReportSetting {
	return this.settings
}

func (this *TimeTrackingReport) GetSettingsLen() int {
	return len(this.settings)
}

func (this *TimeTrackingReport) GetEntry(index string) ProjectReportSetting {
	retVal := this.settings[strings.ToLower(index)]
	return retVal
}

func (this *TimeTrackingReport) SetEntry(prjRepSettings ProjectReportSetting) {
	this.settings[strings.ToLower(prjRepSettings.prj)] = prjRepSettings
}

func (this *TimeTrackingReport) SetTeamMembers(tm map[string]bool) {
	this.teammembers = make(map[string]bool)
	for i := range tm {
		this.teammembers[strings.ToLower(i)] = true
	}
}

func (this *TimeTrackingReport) Finish() {
	this.calculateSumOfAllTimes()
	this.finishPrjSettings()
}

func (this *TimeTrackingReport) calculateSumOfAllTimes() {
	for i := range this.settings {
		entry := this.settings[i]
		var retTotalTime TimeEntry
		retTotalTime = this.createTotalOfPrj(this.settings[i].prj, entry)
		this.SumOfAllProjectTimes = this.SumOfAllProjectTimes + retTotalTime.ToFloat64InHours()
		entry.SetTimeEntry(retTotalTime)
		this.settings[i] = entry
	}
}

func (this *TimeTrackingReport) finishPrjSettings() {
	for i := range this.settings {
		entry := this.settings[i]
		entry.timeEntry.SetOverallTime(this.SumOfAllProjectTimes)
		this.settings[i] = entry
	}
}

func (this TimeTrackingReport) createTotalOfPrj(prjName string, prjSpecificSettings ProjectReportSetting) TimeEntry {
	var total TimeEntry
	var sumOfTimes float64 = 0.0

	var i int = 0
	var personsTimes []TimeEntry = make([]TimeEntry, len(prjSpecificSettings.GetNames())+1)
	var personsWithTime []string = make([]string, 0, len(prjSpecificSettings.GetNames())+1)

	for i = 0; i < len(prjSpecificSettings.GetNames()); i++ {
		var person TimeEntry
		person.InitializeFromString(prjSpecificSettings.GetNames()[i], prjSpecificSettings.GetTimes()[i])
		personsTimes[i] = person
	}

	args := GetArguments()
	for key := range prjSpecificSettings.GetNames() {
		if this.teammembers[strings.ToLower(personsTimes[key].GetName())] == true {
			sumOfTimes = sumOfTimes + personsTimes[key].ToFloat64InHours()
			if personsTimes[key].ToFloat64InHours() > 0.0 {
				lastname := personsTimes[key].GetName()
				lastname = lastname[strings.IndexRune(lastname, '.')+1:]
				personsWithTime = append(personsWithTime, lastname)
			}
			if args.IsTesting() {
				fmt.Printf("Team Member: %s : %f\n", personsTimes[key].ToCsvFormat(';'), personsTimes[key].ToFloat64InHours())
			}
		} else if args.IsTesting() {
			fmt.Printf("Non-Team Member: %s : %f\n", personsTimes[key].ToCsvFormat(';'), personsTimes[key].ToFloat64InHours())
		}
	}

	total.InitializeFromFloat(prjName, sumOfTimes, personsWithTime)
	return total
}

