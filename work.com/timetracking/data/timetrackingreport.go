package data

import (
	"encoding/csv"
	"strconv"
	"fmt"
	"strings"
	. "../jira/Timeentry"
	. "../helper"
	. "../jira/UrlDate"
	. "../arguments"
)

type TimeTrackingReport struct {
	settings  map[string]ProjectReportSetting
	SumOfAllProjectTimes float64
	teammembers map[string]bool
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
	this.settings[strings.ToLower(prjRepSettings.Prj)] = prjRepSettings
}

func (this *TimeTrackingReport) SetTeamMembers(tm map[string]bool) {
	this.teammembers = make(map[string]bool)
	for i := range tm {
		this.teammembers[strings.ToLower(i)] = true
	}
}

func (this *TimeTrackingReport) Initialize(path string) {
	content := ReadInFile(path)
	this.parseProjectsFromByteStream(content)
}

func (this *TimeTrackingReport) Finish() {
	this.calculateSumOfAllTimes()
	this.finishPrjSettings()
}

func (this *TimeTrackingReport) calculateSumOfAllTimes() {
	for i := range this.settings {
		entry := this.settings[i]
		var retTotalTime TimeEntry
		retTotalTime = this.createTotalOfPrj(this.settings[i].Prj, entry)
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
				lastname = lastname[strings.IndexRune(lastname, '.') + 1:]
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

func (this *TimeTrackingReport) parseProjectsFromByteStream(content []byte) {
	records := this.readRecordsFromContent(string(content))

	this.settings = make(map[string]ProjectReportSetting, len(records))
	for i := 0; i < len(records); i++ {
		this.setPrjInfoAtPosition(i, records[i])
	}
}

func (this TimeTrackingReport) readRecordsFromContent(content string) [][]string {
	r := csv.NewReader(strings.NewReader(content))
	r.Comma = ','
	r.Comment = '#'

	records, err := r.ReadAll()
	PanicOnError(err)
	return records
}

func (this *TimeTrackingReport) setPrjInfoAtPosition(position int, record []string) {
	if len(record) != 6 {
		panic("Length of items not enough, we need 6 items")
		return
	}
	var newElement ProjectReportSetting
	newElement.Prj = this.setStringValue(record[0])
	newElement.Id = this.setIntValue(record[1])
	newElement.Query = this.setStringValue(record[2])
	newElement.Startdate = this.setUrlDateValue(record[3])
	newElement.Enddate = this.setUrlDateValue(record[4])
	newElement.productOwner = this.setStringValue(record[5])
	this.settings[strings.ToLower(newElement.Prj)] = newElement
}

func (this TimeTrackingReport) setStringValue(value string) string {
	return strings.TrimSpace(value)
}

func (this TimeTrackingReport) setIntValue(value string) int {
	k, parseErr := strconv.Atoi(strings.TrimSpace(value))
	if parseErr != nil {
		return -1
	} else {
		return k
	}
}

func (this TimeTrackingReport) setUrlDateValue(value string) UrlDate {
	var jiraDate UrlDate
	jiraDate.Initialize(strings.TrimSpace(value))

	return jiraDate
}
