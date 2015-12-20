package jira

import (
	"gopkg.in/yaml.v2"
	. "gopkg.in/check.v1"
	"testing"
	. "../../helper"
)

const pathToTestJiraConfigYaml string = "../../__testdata/jira_config.yaml"
type YamlTestEngine struct{}

func TestYamlEngine(t *testing.T) {
	Suite(&YamlTestEngine{})
	TestingT(t)
}

type YamlWithList struct {
	Mylist []string
	Mypairlist [][]string
}

func (y *YamlTestEngine) TestReadJiraConfig(c *C) {
	config := Reader.Read(pathToTestJiraConfigYaml)
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

func (this *YamlTestEngine) TestReadingInStringList(c *C) {
	var content string = `
mylist: [test1, test2]
mypairlist: [
    [abc, def],
    [ghi, jkl],
    [mno, pqr, stv]
]
`
	yamlList := this.unmarshalList([]byte(content))

	c.Assert(len(yamlList.Mylist), Equals, 2)
	c.Assert(yamlList.Mylist[0], Equals, "test1")
	c.Assert(yamlList.Mylist[1], Equals, "test2")
	c.Assert(len(yamlList.Mypairlist), Equals, 3)
	c.Assert(len(yamlList.Mypairlist[0]), Equals, 2)
	c.Assert(len(yamlList.Mypairlist[1]), Equals, 2)
	c.Assert(len(yamlList.Mypairlist[2]), Equals, 3)
}

func (this *YamlTestEngine) unmarshalList(content []byte) YamlWithList {
	var yamlWithList YamlWithList
	err := yaml.Unmarshal(content, &yamlWithList)
	PanicOnError(err)
	return yamlWithList
}
