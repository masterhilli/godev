package jiraconnector

import (
    "io/ioutil"
    . "work.com/timetracking/helper"
)

type JiraConnector struct {
}

// this is a test support function, should have nothing todo in here!
func (jc *JiraConnector) RetrieveJIRAReportStream() string {
    data, err := ioutil.ReadFile("./testdata/Report-Jira.html")
    PanicOnError(err)
    return string(data)
}
