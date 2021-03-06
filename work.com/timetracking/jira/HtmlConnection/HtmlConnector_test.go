package jira

import (
	. "../../data"
	. "../Config"
	. "gopkg.in/check.v1"
	"testing"
)

type YamlTestEngine struct {
	ProjectInfo ProjectReportSetting
	jc          HtmlConnector
}

func TestYamlEngine(t *testing.T) {
	var yte YamlTestEngine
	yte.ProjectInfo.SetProject("SOLUT")
	yte.ProjectInfo.SetIdFromString("10941")
	yte.ProjectInfo.SetQuery("")

	yte.ProjectInfo.SetStartEndDateFromString("01.09.2015", "11.11.2015")
	yte.jc = NewHtmlConnector(Reader.Read("../../__testdata/jira.yaml"))
	Suite(&yte)
	TestingT(t)
}

var jiraUrl string = "http://10.99.11.333/j/secure/ConfigureReport.jspa?startDateId=1%2FSep%2F15&endDateId=11%2FNov%2F15&projectId=10941&jqlQueryId=&selectedProjectId=10941&reportKey=com.synergyapps.plugins.jira.timepo-timesheet-plugin%3Aissues-report&Next=Next"

func (y *YamlTestEngine) TestGenerateUrlToJira(c *C) {
	c.Assert(y.jc.generateUrlToConnect(y.ProjectInfo), Equals, jiraUrl)
}

// we ignore that one, because we must be @Schenker so that it would work!
func (y *YamlTestEngine) IgnoreGetContentOverJira(c *C) {
	content := y.jc.GetReportContentForProjectInTimeframe(y.ProjectInfo)
	if len(content) >= 1000 {
		content = content[0:1000]
	}
	c.Assert(content, HasLen, 1000)
}
