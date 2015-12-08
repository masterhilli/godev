package data

import (
	"strconv"
	"strings"
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

func setStringValue(value string) string {
	return strings.TrimSpace(value)
}

func setIntValue(value string) int {
	k, parseErr := strconv.Atoi(strings.TrimSpace(value))
	if parseErr != nil {
		return -1
	} else {
		return k
	}
}

func setJiraDateValue(value string) UrlDate {
	var jiraDate UrlDate
	jiraDate.Initialize(strings.TrimSpace(value))

	return jiraDate
}
