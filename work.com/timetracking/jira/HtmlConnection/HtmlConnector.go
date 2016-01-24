package jira

import (
	. "../../data"
	. "../../helper"
	. "../Config"
	"io/ioutil"
	"net/http"
)

type HtmlConnector struct {
	config *Config
}

func NewHtmlConnector(config *Config) HtmlConnector {
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
	requ.SetBasicAuth(jc.config.Jiradata.Username, jc.config.Jiradata.Password)
	return requ
}

func (jc *HtmlConnector) generateUrlToConnect(projectInfo ProjectReportSetting) string {
	startDate := projectInfo.GetStartDate()
	endDate := projectInfo.GetEndDate()
	return jc.config.Jiradata.Url +
		jc.config.Jiradata.GetReportName() +
		jc.config.Jiradata.GetStartDate() +
		startDate.GetTimeForUrl() +
		jc.config.Jiradata.GetEndDate() +
		endDate.GetTimeForUrl() +
		jc.config.Jiradata.GetPrjId() +
		projectInfo.GetIdAsString() +
		jc.config.Jiradata.GetQuery() +
		projectInfo.GetQuery() +
		jc.config.Jiradata.GetSelectedPrjId() +
		projectInfo.GetIdAsString() +
		jc.config.Jiradata.GetReportKey()
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
