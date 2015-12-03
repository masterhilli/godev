package prjinfo

import (
	"encoding/csv"
	"strconv"
	"strings"
	"time"
	. "work.com/timetracking/helper"
)

type Projects struct {
	Data      []Prjinfo
	Seperator rune
}

type Prjinfo struct {
	Prj          string
	Id           int
	Query        string
	Startdate    JiraDate
	Enddate      JiraDate
	ProductOwner string
}

type JiraDate struct {
	t time.Time
}

func (p *Projects) Initialize(path string, seperator rune) {
	p.Seperator = seperator
	content := ReadInFile(path)
	p.parseProjectsFromByteStream(content)
}

func (p *Projects) parseProjectsFromByteStream(content []byte) {
	records := p.readRecordsFromContent(string(content))

	p.Data = make([]Prjinfo, len(records))
	for i := 0; i < len(records); i++ {
		p.setPrjInfoAtPosition(i, records[i])
	}
}

func (p *Projects) readRecordsFromContent(content string) [][]string {
	r := csv.NewReader(strings.NewReader(content))
	r.Comma = p.Seperator
	r.Comment = '#'

	records, err := r.ReadAll()
	PanicOnError(err)
	return records
}

func (p *Projects) setPrjInfoAtPosition(position int, record []string) {
	if len(record) != 6 {
		p.Data[position].Prj = "Length of items not enough, we need 6 items"
		return
	}
	p.Data[position].Prj = setStringValue(record[0])
	p.Data[position].Id = setIntValue(record[1])
	p.Data[position].Query = setStringValue(record[2])
	p.Data[position].Startdate = setJiraDateValue(record[3])
	p.Data[position].Enddate = setJiraDateValue(record[4])
	p.Data[position].ProductOwner = setStringValue(record[5])
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
