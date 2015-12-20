package jira

import (
	. "../../data"
	. "../../helper"
	. "../Config"
	"io/ioutil"
	"net/http"
	"strconv"
)

type HtmlConnector struct {
	config Config
}

func NewHtmlConnector(config Config) HtmlConnector {
	var htmlConnector HtmlConnector
	htmlConnector.config = config
	return htmlConnector
}

func (jc *HtmlConnector) GetReportContentForProjectInTimeframe(projectInfo ProjectReportSetting) string {
	requ := jc.generateRequest(projectInfo)
	return jc.getHTMLBodyFromRequest(requ)
}

func (jc *HtmlConnector) generateRequest(projectInfo ProjectReportSetting) *http.Request {
	requ, err := http.NewRequest("GET", jc.generateUrlToConnect(projectInfo), nil)
	PanicOnError(err)
	requ.SetBasicAuth(jc.config.JiraLogin.Username, jc.config.JiraLogin.Password)
	return requ
}

func (jc *HtmlConnector) generateUrlToConnect(projectInfo ProjectReportSetting) string {
	return jc.config.JiraUrl.Url +
		jc.config.JiraUrl.GetReportName() +
		jc.config.JiraUrl.GetStartDate() +
		projectInfo.Startdate.GetTimeForUrl() +
		jc.config.JiraUrl.GetEndDate() +
		projectInfo.Enddate.GetTimeForUrl() +
		jc.config.JiraUrl.GetPrjId() +
		strconv.Itoa(projectInfo.Id) +
		jc.config.JiraUrl.GetQuery() +
		projectInfo.Query +
		jc.config.JiraUrl.GetSelectedPrjId() +
		strconv.Itoa(projectInfo.Id) +
		jc.config.JiraUrl.GetReportKey()
}

func (jc *HtmlConnector) getHTMLBodyFromRequest(requ *http.Request) string {
	client := &http.Client{}

	resp, err := client.Do(requ)
	PanicOnError(err)
	defer resp.Body.Close()
	content, errReader := ioutil.ReadAll(resp.Body)
	PanicOnError(errReader)
	return string(content)
}
