package data

import (
	"strconv"
	"strings"
	"time"
	. "work.com/timetracking/helper"
)

type ProjectReportSetting struct {
	Prj          string
	Id           int
	Query        string
	Startdate    JiraDate
	Enddate      JiraDate
	ProductOwner string

	names []string
	times []string
}

type JiraDate struct {
	t time.Time
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

func setJiraDateValue(value string) JiraDate {
	var jiraDate JiraDate
	jiraDate.Initialize(strings.TrimSpace(value))

	return jiraDate
}

func (jd *JiraDate) Initialize(date string) {
	if len(date) == 0 {
		jd.t = time.Now()
		return
	}
	tmp, err := time.Parse("02.01.2006", date)
	PanicOnError(err)
	jd.t = tmp
}

func (jd *JiraDate) GetTimeForUrl() string {
	urldate := jd.t.Format("2") + "%2F" + jd.t.Format("Jan") + "%2F" + jd.t.Format("06")
	return urldate
}
