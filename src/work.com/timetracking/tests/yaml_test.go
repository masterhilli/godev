package testinterfaces

import (
	"gopkg.in/yaml.v2"
	"path/filepath"
	"io/ioutil"
	"testing"
	. "gopkg.in/check.v1"
)

type YamlTestEngine struct {}

func TestYamlEngine(t *testing.T) {
	Suite(&YamlTestEngine{})
	TestingT(t)
}

type Config struct {
	JiraLogin LoginData
    JiraUrl UrlInformation
}

type LoginData struct {
    Username string
    Password string
}

type UrlInformation struct {
    Url string
}


func (s *HttpTestEngine) TestReadJiraConfig(c *C) {
	filename, _ := filepath.Abs("./jira.yaml")
	yamlInformation, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var config Config
    err = yaml.Unmarshal(yamlInformation, &config)
    if err != nil {
		panic(err)
	}
	c.Assert(config.JiraLogin.Username, Equals, "xyz")
	c.Assert(config.JiraLogin.Password, Equals, "abcdefgh")
	c.Assert(config.JiraUrl.Url, Equals, "http://10.207.121.181/j/secure/")
}

func (s *HttpTestEngine) TestYamlUnmarshaler(c *C) {
	var config Config
	var yamlInformation string = "jiralogin:\n    username: abc\n    password: xyz\njiraurl:\n    url: www.google.at"
    err := yaml.Unmarshal([]byte(yamlInformation), &config)
    if err != nil {
		panic(err)
	}
	c.Assert(config.JiraLogin.Username, Equals, "abc")
	c.Assert(config.JiraLogin.Password, Equals, "xyz")
	c.Assert(config.JiraUrl.Url, Equals, "www.google.at")
}


