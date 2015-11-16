package jiraconnector

import (
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "path/filepath"
    . "work.com/timetracking/helper"
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
    content := jc.readInFile(pathToConfig)
    jc.config = jc.unmarshalToConfig(content)
}

// this is a test support function, should have nothing todo in here!
func (jc *JiraConnector) RetrieveJIRAReportStream() string {
    data, err := ioutil.ReadFile("./testdata/Report-Jira.html")
    PanicOnError(err)
    return string(data)
}

func (jc *JiraConnector) readInFile(path string) []byte {
    filename, errAbs := filepath.Abs(path)
    PanicOnError(errAbs)
    content, errReadFile := ioutil.ReadFile(filename)
    PanicOnError(errReadFile)
    return content
}

func (jc *JiraConnector) unmarshalToConfig(content []byte) Config {
    var config Config
    err := yaml.Unmarshal(content, &config)
    PanicOnError(err)
    return config
}
