package jiraRegEx

import (
	"net/http"
	"net/url"
	"path/filepath"
	"io/ioutil"
	"testing"
	. "gopkg.in/check.v1"
	"gopkg.in/yaml.v2"
)


// Hook up gocheck into the "go test" runner.
type HttpTestEngine struct{}
func TestHttpLibrary(t *testing.T) { 
	Suite(&HttpTestEngine{})
	TestingT(t) 
}

func (s *HttpTestEngine) TestHttpGetOnGoogleReturnsNoError(c *C) {
	_, err := http.Get("http://www.google.com")
	c.Assert(err, IsNil)
}

func (s *HttpTestEngine) TestHttpGetOnExampleReturnsError(c *C) {
	_, err := http.Get("abc")
	c.Assert(err, NotNil)
}

func (s *HttpTestEngine) TestHttpGetParseForBodyNoError(c *C) {
	resp, err := http.Get("http://www.google.com/")
	if (err == nil) {
		_, errReader := ioutil.ReadAll(resp.Body)
		c.Assert(errReader, IsNil)
		resp.Body.Close()
	} else {
		c.Assert(err, NotNil)
	}
}

func (s *HttpTestEngine) TestHttpsGetParseForBodyNoError(c *C) {
	resp, err := http.Get("https://www.google.com")
	if (err == nil) {
		_, errReader := ioutil.ReadAll(resp.Body)
		c.Assert(errReader, IsNil)
		resp.Body.Close()
	} else {
		c.Assert(err, NotNil)
	}
}

func (s *HttpTestEngine) TestHttpsGetJiraUrl(c *C) {
	var url string = "http://10.207.121.181/j/"
	resp, err := http.Get(url)
	if (err == nil) {
		respBody, errReader := ioutil.ReadAll(resp.Body)
		c.Assert(errReader, IsNil)
		s.WriteOutToFile([]byte(respBody), "jira-report.html")
		resp.Body.Close()
	} else {
		c.Assert(err, IsNil)
	}
}


var urlToJiraReport string = "http://10.207.121.181/j/secure/ConfigureReport.jspa?startDateId=1%2FSep%2F15&endDateId=11%2FNov%2F15&projectId=10941&jqlQueryId=&selectedProjectId=10941&reportKey=com.synergyapps.plugins.jira.timepo-timesheet-plugin%3Aissues-report&Next=Next"
func (s *HttpTestEngine) TestJiraCreateRequestAndLogin(c *C) {
	filename, _ := filepath.Abs("..\\jira.yaml")
	yamlInformation, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var config Config
    err = yaml.Unmarshal(yamlInformation, &config)
    if err != nil {
		panic(err)
	}
	requ, errReq := http.NewRequest("GET", config.JiraUrl.Url, nil)
	if errReq != nil {
		panic("Error while building jira request")
	}
	requ.SetBasicAuth(config.JiraLogin.Username, config.JiraLogin.Password)
	client := &http.Client{}

	resp, errDo := client.Do(requ)
	defer resp.Body.Close()
	if (errDo == nil) {
		respBody, errReader := ioutil.ReadAll(resp.Body)
		c.Assert(errReader, IsNil)
		s.WriteOutToFile([]byte(respBody), "TestJiraCreateRequestAndLogin.html")
	} else {
		c.Assert(err, IsNil)
	}
}

func (s *HttpTestEngine) WriteOutToFile(data []byte, fileName string)  {
	err := ioutil.WriteFile("./"+fileName, data, 0644)
    if err != nil {
        panic(err)
    }
}

func (s *HttpTestEngine) TestHttpPostFormParseForBodyNoError(c *C) {
	resp, err := http.PostForm("http://www.google.com/", url.Values{"gfe_rd" : {"cr"}, "ei":{"lVhCVreWAsOH8Qe6lq2gDQ"}, "gws_rd" : {"ssl"}})
	if (err == nil) {
		_, errReader := ioutil.ReadAll(resp.Body)
		c.Assert(errReader, IsNil)
		resp.Body.Close()
	} else {
		c.Assert(err, NotNil)
	}
}

func (s *HttpTestEngine) TestHttpsPostFormBodyOfJiraReport(c *C) {
	//
	var urlToJiraReport string = "http://10.207.121.181/j/secure/ConfigureReport.jspa"
	urlValues := url.Values{
							"startDateId": {"1/Sep/15"},
							"endDateId": {"11/Nov/15"},
							"projectId": {"10941"},
							"jqlQueryId": {""},
							"selectedProjectId": {"10941"},
							"reportKey": {"com.synergyapps.plugins.jira.timepo-timesheet-plugin:issues-report"},
							"Next": {"Next"}}
	resp, err := http.PostForm(urlToJiraReport, urlValues)
	if (err == nil) {
		respBody, errReader := ioutil.ReadAll(resp.Body)
		c.Assert(errReader, IsNil)
		s.WriteOutToFile([]byte(respBody), "TestHttpsPostFormBodyOfJiraReport.html")
		resp.Body.Close()
	} else {
		c.Assert(err, IsNil)
	}
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


