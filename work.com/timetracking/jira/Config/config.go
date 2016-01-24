package jira

import (
	. "./Data"
	. "./Project"
	data "../../data"
	"strings"
)

type Config struct {
	Jiradata    JiraData
	Dates       TimeFrame
	Reportname  string
	Projects    map[string]Project
	Teammembers []string
}

type TimeFrame struct {
	Startdate string
	Enddate   string
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
	if project.Excludeothers {
		platformsNotPartOfQuery = make([]string, 0, len(this.Projects))
		for i := range this.Projects {
			platform := this.Projects[i].Platform
			if len(platform) > 0 {
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