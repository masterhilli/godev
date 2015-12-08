package jira

import (
    "io/ioutil"
    "net/http"
    "strconv"
    . "work.com/timetracking/data"
    . "work.com/timetracking/helper"
    . "work.com/timetracking/jira/Config"
)

type HtmlConnector struct {
    config Config
}

func NewHtmlConnector(config Config) HtmlConnector {
    var htmlConnector HtmlConnector
    htmlConnector.config = config
    return htmlConnector
}

func (jc *HtmlConnector) GetReportContentForProjectInTimeframe(projectInfo Prjinfo) string {
    requ := jc.generateRequest(projectInfo)
    return jc.getHTMLBodyFromRequest(requ)
}

func (jc *HtmlConnector) generateRequest(projectInfo Prjinfo) *http.Request {
    requ, err := http.NewRequest("GET", jc.generateUrlToConnect(projectInfo), nil)
    PanicOnError(err)
    requ.SetBasicAuth(jc.config.JiraLogin.Username, jc.config.JiraLogin.Password)
    return requ
}

func (jc *HtmlConnector) generateUrlToConnect(projectInfo Prjinfo) string {
    return jc.config.JiraUrl.Url +
        jc.config.JiraUrl.Reportname +
        jc.config.JiraUrl.Startdate +
        projectInfo.Startdate.GetTimeForUrl() +
        jc.config.JiraUrl.Enddate +
        projectInfo.Enddate.GetTimeForUrl() +
        jc.config.JiraUrl.Prjid +
        strconv.Itoa(projectInfo.Id) +
        jc.config.JiraUrl.Query +
        projectInfo.Query +
        jc.config.JiraUrl.Selectedprjid +
        strconv.Itoa(projectInfo.Id) +
        jc.config.JiraUrl.Prefix
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
