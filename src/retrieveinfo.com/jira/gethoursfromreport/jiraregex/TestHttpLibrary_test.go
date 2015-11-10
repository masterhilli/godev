package jiraRegEx

import (
	"fmt"
	"net/http"
	"net/url"
	//"regexp"
	"io/ioutil"
	"testing"
	. "gopkg.in/check.v1"
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
		//fmt.Printf("%s\n", htmlBody)
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
	fmt.Printf("***** -- Read in HTML response %s\n", "test")
	if (err == nil) {
		respBody, errReader := ioutil.ReadAll(resp.Body)
		c.Assert(errReader, IsNil)
		s.WriteOutToFile([]byte(respBody), "jira-report.html")
		fmt.Printf("***** -- Read in HTML response %s\n", "test")
		resp.Body.Close()
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
		//fmt.Printf("%s\n", htmlBody)
		resp.Body.Close()
	} else {
		c.Assert(err, NotNil)
	}
}

func (s *HttpTestEngine) TestHttpsPostFormBodyOfJiraReport(c *C) {
	//http://10.207.121.181/j/secure/ConfigureReport.jspa?
	var urlToJiraReport string = "http://10.207.121.181/j/secure/ConfigureReport.jspa"
	//startDateId=1%2FSep%2F15&endDateId=11%2FNov%2F15&projectId=10941&jqlQueryId=&selectedProjectId=10941&reportKey=com.synergyapps.plugins.jira.timepo-timesheet-plugin%3Aissues-report&Next=Next
	urlValues := url.Values{
							"startDateId": {"1/Sep/15"},
							"endDateId": {"11/Nov/15"},
							"projectId": {"10941"},
							"jqlQueryId": {""},
							"selectedProjectId": {"10941"},
							"reportKey": {"com.synergyapps.plugins.jira.timepo-timesheet-plugin:issues-report"},
							"Next": {"Next"}}
	resp, err := http.PostForm(urlToJiraReport, urlValues)
	fmt.Printf("***** -- Read in HTML response %s\n", "test")
	if (err == nil) {
		respBody, errReader := ioutil.ReadAll(resp.Body)
		c.Assert(errReader, IsNil)
		s.WriteOutToFile([]byte(respBody), "jira-report1.html")
		fmt.Printf("***** -- Read in HTML response %s\n", "test")
		resp.Body.Close()
	} else {
		c.Assert(err, IsNil)
	}
}