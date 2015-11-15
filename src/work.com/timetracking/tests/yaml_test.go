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

func (y *YamlTestEngine) TestReadJiraConfig(c *C) {
	content := y.readInFile("./jira.yaml")
	config := y.parseConfigFromByteStream(content)
	c.Assert(config.JiraLogin.Username, Equals, "xyz")
	c.Assert(config.JiraLogin.Password, Equals, "abcdefgh")
	c.Assert(config.JiraUrl.Url, Equals, "http://10.207.121.181/j/secure/")
}

func (y *YamlTestEngine) TestYamlUnmarshaler(c *C) {
	var content string = "jiralogin:\n    username: abc\n    password: xyz\njiraurl:\n    url: www.google.at"
	config := y.parseConfigFromByteStream([]byte(content))
	c.Assert(config.JiraLogin.Username, Equals, "abc")
	c.Assert(config.JiraLogin.Password, Equals, "xyz")
	c.Assert(config.JiraUrl.Url, Equals, "www.google.at")
}

func (y *YamlTestEngine) readInFile(filename string) []byte {
	filename, _ = filepath.Abs(filename)
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return content
}

func (y *YamlTestEngine) parseConfigFromByteStream(content []byte) Config {
	var config Config
    err := yaml.Unmarshal(content, &config)
    if err != nil {
		panic(err)
	}
	return config
}




