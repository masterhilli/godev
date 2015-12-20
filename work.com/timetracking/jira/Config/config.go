package jira

import (
	. "../../helper"
	"gopkg.in/yaml.v2"
)

const reportName string = "ConfigureReport.jspa?"
const startDate string = "startDateId="
const endDate string = "&endDateId="
const prjId string = "&projectId="
const query string = "&jqlQueryId="
const selectedPrjId string = "&selectedProjectId="
const reportKey string = "&reportKey=com.synergyapps.plugins.jira.timepo-timesheet-plugin%3Aissues-report&Next=Next"

type Config struct {
	JiraLogin   LoginData
	JiraUrl     UrlInformation
	Projects    map[string]Project
	Teammembers []string
}

type LoginData struct {
	Username string
	Password string
}

type UrlInformation struct {
	Url string
}

type Project struct {
	Project      string
	Platform     string
	Productowner string
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

func (this UrlInformation) GetReportName() string {
	return reportName
}

func (this UrlInformation) GetStartDate() string {
	return startDate
}

func (this UrlInformation) GetEndDate() string {
	return endDate
}

func (this UrlInformation) GetPrjId() string {
	return prjId
}

func (this UrlInformation) GetSelectedPrjId() string {
	return selectedPrjId
}

func (this UrlInformation) GetReportKey() string {
	return reportKey
}

func (this UrlInformation) GetQuery() string {
	return query
}
