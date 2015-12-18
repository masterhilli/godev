package data

import (
	. "work.com/timetracking/jira/UrlDate"
	. "work.com/timetracking/jira/Timeentry"
)

type ProjectReportSetting struct {
	Prj          string
	Id           int
	Query        string
	Startdate    UrlDate
	Enddate      UrlDate
	ProductOwner string

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

func (this *ProjectReportSetting) GetTimeEntry() TimeEntry{
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
