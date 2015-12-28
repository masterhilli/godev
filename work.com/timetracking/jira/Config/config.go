package jira

import (
	data "../../data"
	. "../../helper"
	"gopkg.in/yaml.v2"
	"net/url"
	"strings"
)

const reportName string = "ConfigureReport.jspa?"
const startDate string = "startDateId="
const endDate string = "&endDateId="
const prjId string = "&projectId="
const query string = "&jqlQueryId="
const selectedPrjId string = "&selectedProjectId="
const reportKey string = "&reportKey=com.synergyapps.plugins.jira.timepo-timesheet-plugin%3Aissues-report&Next=Next"

type Config struct {
	Jiradata    JiraData
	Dates       TimeFrame
	Reportname  string
	Projects    map[string]Project
	Teammembers []string
}

type JiraData struct {
	Username string
	Password string
	Url      string
}

type TimeFrame struct {
	Startdate string
	Enddate   string
}

type Project struct {
	Project       string
	Platform      string
	Productowner  string
	Excludeothers bool
}

var Reader configReader

type configReader struct{}

func (cr configReader) Read(pathToConfig string) Config {
	content := ReadInFile(pathToConfig)
	config := cr.unmarshalToConfig(content)
	if len(config.Teammembers) == 1 {
		panic("Sorry, I do not allow you to track the times of single persons, use more then 1 person in the team!")
	}
	return config
}

func (cr configReader) unmarshalToConfig(content []byte) Config {
	var config Config
	err := yaml.Unmarshal(content, &config)
	PanicOnError(err)
	return config
}

func (this Config) GetTeammembersAsMap() map[string]bool {
	var retVal map[string]bool = make(map[string]bool)
	for i := range this.Teammembers {
		retVal[strings.ToLower(strings.TrimSpace(this.Teammembers[i]))] = true
	}

	return retVal
}

func (this Config) GetTimeTrackingReportData() data.TimeTrackingReport {
	var dataForReport data.TimeTrackingReport = data.NewTimeTrackingReport(len(this.Projects))
	dataForReport.SetTeamMembers(this.GetTeammembersAsMap())
	dataForReport.SetReportName(this.Reportname)
	for i := range this.Projects {
		dataForReport.SetEntry(this.CreateProjectReportSetting(i))
	}
	return dataForReport
}

func (this Config) CreateProjectReportSetting(key string) data.ProjectReportSetting {
	var newProjectReportData data.ProjectReportSetting
	project := this.Projects[key]

	var platformsNotPartOfQuery []string = nil
	if (project.Excludeothers) {
		platformsNotPartOfQuery = make([]string, 0, len(this.Projects))
		for i := range this.Projects {
			platform := this.Projects[i].Platform
			if  len(platform) > 0 {
				platformsNotPartOfQuery = append(platformsNotPartOfQuery, platform)
			}
		}
	}
	newProjectReportData.SetProductOwner(project.Productowner)
	newProjectReportData.SetProject(strings.ToUpper(key))
	newProjectReportData.SetIdFromString("")
	newProjectReportData.SetQuery(project.GetQuery(platformsNotPartOfQuery))
	newProjectReportData.SetStartEndDateFromString(this.Dates.Startdate, this.Dates.Enddate)

	return newProjectReportData
}

func (this JiraData) GetReportName() string {
	return reportName
}

func (this JiraData) GetStartDate() string {
	return startDate
}

func (this JiraData) GetEndDate() string {
	return endDate
}

func (this JiraData) GetPrjId() string {
	return prjId
}

func (this JiraData) GetSelectedPrjId() string {
	return selectedPrjId
}

func (this JiraData) GetReportKey() string {
	return reportKey
}

func (this JiraData) GetQuery() string {
	return query
}

func (this Project) GetQuery(platforms []string) string {
	var sqlQuery string
	if len(this.Platform) > 0 {
		sqlQuery = "Platform = \"" + this.Platform + "\""
	}
	if len(this.Project) > 0 && len(this.Platform) > 0 {
		sqlQuery = sqlQuery + " OR "
	}
	if len(this.Project) > 0 {
		sqlQuery = sqlQuery + "project = \"" + this.Project + "\""
	}

	sqlQuery = "(" + sqlQuery + ")"

	if (platforms != nil) {
		notInPart := " AND Platform not in ("
		for i := range platforms {
			if (len(platforms[i]) > 0) {
				if i > 0 {
					notInPart = notInPart + ","
				}
				notInPart = notInPart + "\"" + platforms[i] + "\""
			}
		}
		notInPart = notInPart + ")"
		sqlQuery = sqlQuery + notInPart
	}

	return url.QueryEscape(sqlQuery)
}
