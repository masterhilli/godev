package jira

import (
	. "gopkg.in/check.v1"
	"testing"
	. "work.com/timetracking/jira/Config"
	prj "work.com/timetracking/prjinfo"
)

type YamlTestEngine struct {
	ProjectInfo prj.Prjinfo
	jc          HtmlConnector
}

func TestYamlEngine(t *testing.T) {
	var yte YamlTestEngine
	yte.ProjectInfo.Prj = "SOLUT"
	yte.ProjectInfo.Id = 10941
	yte.ProjectInfo.Query = ""
	var myStartDate prj.JiraDate
	var myEndDate prj.JiraDate

	myStartDate.Initialize("01.09.2015")
	myEndDate.Initialize("11.11.2015")
	yte.ProjectInfo.Startdate = myStartDate
	yte.ProjectInfo.Enddate = myEndDate
	yte.jc = NewHtmlConnector(Reader.Read("./jira.yaml"))
	Suite(&yte)
	TestingT(t)
}

var jiraUrl string = "http://10.207.121.181/j/secure/ConfigureReport.jspa?startDateId=1%2FSep%2F15&endDateId=11%2FNov%2F15&projectId=10941&jqlQueryId=&selectedProjectId=10941&reportKey=com.synergyapps.plugins.jira.timepo-timesheet-plugin%3Aissues-report&Next=Next"

func (y *YamlTestEngine) TestGenerateUrlToJira(c *C) {
	c.Assert(y.jc.generateUrlToConnect(y.ProjectInfo), Equals, jiraUrl)
}

func (y *YamlTestEngine) TestGetContentOverJira(c *C) {
	content := y.jc.GetReportContentForProjectInTimeframe(y.ProjectInfo)
	if len(content) >= 1000 {
		content = content[0:1000]
	}
	c.Assert(content, HasLen, 1000)
}
