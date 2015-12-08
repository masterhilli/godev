package data

import (
	. "work.com/timetracking/jira/UrlDate"
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
}

func (pi *ProjectReportSetting) GetNames() []string {
	return pi.names
}

func (pi *ProjectReportSetting) SetNames(names []string) {
	pi.names = names
}

func (pi *ProjectReportSetting) GetTimes() []string {
	return pi.times
}

func (pi *ProjectReportSetting) SetTimes(times []string) {
	pi.times = times
}
