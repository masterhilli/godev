package jiraconnector

import (
	. "gopkg.in/check.v1"
	"testing"
)

type YamlTestEngine struct{}

func TestYamlEngine(t *testing.T) {
	Suite(&YamlTestEngine{})
	TestingT(t)
}

func (y *YamlTestEngine) TestReadJiraConfig(c *C) {
	var jc JiraConnector
	jc.Initialize("./jira.yaml")
	config := jc.config
	c.Assert(config.JiraLogin.Username, Equals, "xyz")
	c.Assert(config.JiraLogin.Password, Equals, "abcdefgh")
	c.Assert(config.JiraUrl.Url, Equals, "http://10.207.121.181/j/secure/")
}

func (y *YamlTestEngine) TestYamlUnmarshaler(c *C) {
	var content string = "jiralogin:\n    username: abc\n    password: xyz\njiraurl:\n    url: www.google.at"
	var jc JiraConnector
	config := jc.unmarshalToConfig([]byte(content))
	c.Assert(config.JiraLogin.Username, Equals, "abc")
	c.Assert(config.JiraLogin.Password, Equals, "xyz")
	c.Assert(config.JiraUrl.Url, Equals, "www.google.at")
}
