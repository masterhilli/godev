package jira

import (
	"gopkg.in/yaml.v2"
	. "../../helper"
)

type Config struct {
	JiraLogin LoginData
	JiraUrl   UrlInformation
}

type LoginData struct {
	Username string
	Password string
}

type UrlInformation struct {
	Url           string
	Reportname    string // "ConfigureReport.jspa?"
	Startdate     string //  "startDateId="
	Enddate       string //  "&endDateId="
	Prjid         string //  "&projectId="
	Query         string //  "&jqlQueryId="
	Selectedprjid string //  "&selectedProjectId="
	Prefix        string //  "&reportKey=com.synergyapps.plugins.jira.timepo-timesheet-plugin%3Aissues-report&Next=Next"
}

var Reader configReader

type configReader struct{}

func (cr configReader) Read(pathToConfig string) Config {
	content := ReadInFile(pathToConfig)
	return cr.unmarshalToConfig(content)
}

func (cr configReader) unmarshalToConfig(content []byte) Config {
	var config Config
	err := yaml.Unmarshal(content, &config)
	PanicOnError(err)
	return config
}
