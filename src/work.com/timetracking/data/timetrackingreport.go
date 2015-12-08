package data

import (
	"encoding/csv"
	"strconv"
	"strings"
	. "work.com/timetracking/helper"
	. "work.com/timetracking/jira/UrlDate"
)

type TimeTrackingReport struct {
	Settings  []ProjectReportSetting
	Seperator rune
}

func (this *TimeTrackingReport) Initialize(path string, seperator rune) {
	this.Seperator = seperator
	content := ReadInFile(path)
	this.parseProjectsFromByteStream(content)
}

func (this *TimeTrackingReport) parseProjectsFromByteStream(content []byte) {
	records := this.readRecordsFromContent(string(content))

	this.Settings = make([]ProjectReportSetting, len(records))
	for i := 0; i < len(records); i++ {
		this.setPrjInfoAtPosition(i, records[i])
	}
}

func (this *TimeTrackingReport) readRecordsFromContent(content string) [][]string {
	r := csv.NewReader(strings.NewReader(content))
	r.Comma = this.Seperator
	r.Comment = '#'

	records, err := r.ReadAll()
	PanicOnError(err)
	return records
}

func (this *TimeTrackingReport) setPrjInfoAtPosition(position int, record []string) {
	if len(record) != 6 {
		this.Settings[position].Prj = "Length of items not enough, we need 6 items"
		return
	}
	lastPos := &this.Settings[position]
	lastPos.Prj = this.setStringValue(record[0])
	lastPos.Id = this.setIntValue(record[1])
	lastPos.Query = this.setStringValue(record[2])
	lastPos.Startdate = this.setJiraDateValue(record[3])
	lastPos.Enddate = this.setJiraDateValue(record[4])
	lastPos.ProductOwner = this.setStringValue(record[5])
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

func (this TimeTrackingReport) setJiraDateValue(value string) UrlDate {
	var jiraDate UrlDate
	jiraDate.Initialize(strings.TrimSpace(value))

	return jiraDate
}
