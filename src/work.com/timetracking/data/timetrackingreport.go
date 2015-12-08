package data

import (
	"encoding/csv"
	"strings"
	. "work.com/timetracking/helper"
)

type TimeTrackingReport struct {
	Settings  []Prjinfo //ProjectReportSettings
	Seperator rune
}

func (p *TimeTrackingReport) Initialize(path string, seperator rune) {
	p.Seperator = seperator
	content := ReadInFile(path)
	p.parseProjectsFromByteStream(content)
}

func (p *TimeTrackingReport) parseProjectsFromByteStream(content []byte) {
	records := p.readRecordsFromContent(string(content))

	p.Settings = make([]Prjinfo, len(records))
	for i := 0; i < len(records); i++ {
		p.setPrjInfoAtPosition(i, records[i])
	}
}

func (p *TimeTrackingReport) readRecordsFromContent(content string) [][]string {
	r := csv.NewReader(strings.NewReader(content))
	r.Comma = p.Seperator
	r.Comment = '#'

	records, err := r.ReadAll()
	PanicOnError(err)
	return records
}

func (p *TimeTrackingReport) setPrjInfoAtPosition(position int, record []string) {
	if len(record) != 6 {
		p.Settings[position].Prj = "Length of items not enough, we need 6 items"
		return
	}
	lastPos := &p.Settings[position]
	lastPos.Prj = setStringValue(record[0])
	lastPos.Id = setIntValue(record[1])
	lastPos.Query = setStringValue(record[2])
	lastPos.Startdate = setJiraDateValue(record[3])
	lastPos.Enddate = setJiraDateValue(record[4])
	lastPos.ProductOwner = setStringValue(record[5])
}
