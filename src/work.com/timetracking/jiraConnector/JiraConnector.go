package jiraconnector

import (
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "net/http"
    "strconv"
    . "work.com/timetracking/helper"
    prj "work.com/timetracking/prjinfo"
)

type JiraConnector struct {
    config Config
}

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

func (jc *JiraConnector) NewJiraConnector() {
}

func (jc *JiraConnector) Initialize(pathToConfig string) {
    content := ReadInFile(pathToConfig)
    jc.config = jc.unmarshalToConfig(content)
}

func (jc *JiraConnector) unmarshalToConfig(content []byte) Config {
    var config Config
    err := yaml.Unmarshal(content, &config)
    PanicOnError(err)
    return config
}

func (jc *JiraConnector) GetReportContentForProjectInTimeframe(projectInfo prj.Prjinfo) string {
    requ := jc.generateRequest(projectInfo)
    return jc.getHTMLBodyFromRequest(requ)
}

func (jc *JiraConnector) generateRequest(projectInfo prj.Prjinfo) *http.Request {
    requ, err := http.NewRequest("GET", jc.generateUrlToConnect(projectInfo), nil)
    PanicOnError(err)
    requ.SetBasicAuth(jc.config.JiraLogin.Username, jc.config.JiraLogin.Password)
    return requ
}

func (jc *JiraConnector) generateUrlToConnect(projectInfo prj.Prjinfo) string {
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

func (jc *JiraConnector) getHTMLBodyFromRequest(requ *http.Request) string {
    client := &http.Client{}

    resp, err := client.Do(requ)
    defer resp.Body.Close()
    PanicOnError(err)
    content, errReader := ioutil.ReadAll(resp.Body)
    PanicOnError(errReader)
    return string(content)
}
