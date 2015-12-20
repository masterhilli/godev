package testinterfaces

import (
	"net/http"
	"path/filepath"
	"io/ioutil"
	"testing"
	. "gopkg.in/check.v1"
	"gopkg.in/yaml.v2"
	. "../jira/Config"
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

func (s *HttpTestEngine) TTestHttpsGetJiraUrl(c *C) {
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

func (s *HttpTestEngine) TTestJiraCreateRequestAndLogin(c *C) {
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