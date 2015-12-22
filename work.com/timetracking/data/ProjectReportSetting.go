package data

import (
	. "../jira/Timeentry"
	. "../jira/UrlDate"
	"strconv"
	"strings"
)

type ProjectReportSetting struct {
	prj          string
	id           int
	query        string
	startdate    UrlDate
	enddate      UrlDate
	productOwner string

	names []string
	times []string

	timeEntry TimeEntry
}

func (pi *ProjectReportSetting) GetNames() []string {
	return pi.names
}

func (pi *ProjectReportSetting) SetNames(names []string) {
	pi.names = names
}

func (this *ProjectReportSetting) GetTimeEntry() TimeEntry {
	return this.timeEntry
}

func (this *ProjectReportSetting) SetTimeEntry(timeEntry TimeEntry) {
	this.timeEntry = timeEntry
}

func (pi *ProjectReportSetting) GetTimes() []string {
	return pi.times
}

func (pi *ProjectReportSetting) SetTimes(times []string) {
	pi.times = times
}

func (this *ProjectReportSetting) GetProductOwner() string {
	return this.productOwner
}

func (this *ProjectReportSetting) SetProductOwner(po string) {
	this.productOwner = this.setStringValue(po)
}

func (this ProjectReportSetting) GetProject() string {
	return this.prj
}

func (this *ProjectReportSetting) SetProject(prj string) {
	this.prj = this.setStringValue(prj)
}

func (this ProjectReportSetting) GetId() int {
	return this.id
}

func (this ProjectReportSetting) GetIdAsString() string {
	return this.emptyStringForItoA(this.id)
}

func (this *ProjectReportSetting) SetIdFromString(id string) {
	this.id = this.setIntValue(id)
}

func (this ProjectReportSetting) GetQuery() string {
	return this.query
}

func (this *ProjectReportSetting) SetQuery(query string) {
	this.query = this.setStringValue(query)
}

func (this ProjectReportSetting) GetStartDate() UrlDate {
	return this.startdate
}

func (this ProjectReportSetting) GetEndDate() UrlDate {
	return this.enddate
}

func (this *ProjectReportSetting) SetStartEndDateFromString(start, end string) {
	this.startdate = this.setUrlDateValue(start)
	this.enddate = this.setUrlDateValue(end)
}

func (this ProjectReportSetting) setStringValue(value string) string {
	return strings.TrimSpace(value)
}

func (this ProjectReportSetting) setIntValue(value string) int {
	k, parseErr := strconv.Atoi(strings.TrimSpace(value))
	if parseErr != nil {
		return -1
	} else {
		return k
	}
}

func (this ProjectReportSetting) setUrlDateValue(value string) UrlDate {
	var jiraDate UrlDate
	jiraDate.Initialize(strings.TrimSpace(value))

	return jiraDate
}

func (this ProjectReportSetting) emptyStringForItoA(val int) string {
	if val >= 0 {
		return strconv.Itoa(val)
	} else {
		return ""
	}
}
