package jiraconnector

import (
    "gopkg.in/yaml.v2"
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
    Url string
}

func (jc *JiraConnector) NewJiraConnector() {
}

func (jc *JiraConnector) Initialize(pathToConfig string) {
    content := ReadInFile(pathToConfig)
    jc.config = jc.unmarshalToConfig(content)
}

func (jc *JiraConnector) GetReportContentForProjectInTimeframe(projectInfo prj.Prjinfo) string {

    return ""
}

func (jc *JiraConnector) unmarshalToConfig(content []byte) Config {
    var config Config
    err := yaml.Unmarshal(content, &config)
    PanicOnError(err)
    return config
}
