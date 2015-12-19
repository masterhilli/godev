package jira

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
	config := Reader.Read("./jira.yaml")
	c.Assert(config.JiraLogin.Username, Equals, "xyz")
	c.Assert(config.JiraLogin.Password, Equals, "abcdefgh")
	c.Assert(config.JiraUrl.Url, Equals, "http://10.207.121.181/j/secure/")
	c.Assert(config.JiraUrl.Query, Equals, "&jqlQueryId=")
}

func (y *YamlTestEngine) TestYamlUnmarshaler(c *C) {
	var content string = "jiralogin:\n    username: abc\n    password: xyz\njiraurl:\n    url: www.google.at"
	config := Reader.unmarshalToConfig([]byte(content))
	c.Assert(config.JiraLogin.Username, Equals, "abc")
	c.Assert(config.JiraLogin.Password, Equals, "xyz")
	c.Assert(config.JiraUrl.Url, Equals, "www.google.at")
}
