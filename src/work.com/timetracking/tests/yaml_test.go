package testinterfaces

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"path/filepath"
	"encoding/csv"
	"strings"
	"io/ioutil"
	"testing"
	"strconv"
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
/*
prjinfo:
   	    prj: SOLUT
        id: 10941
        query: "project = SOLUT"
        startdate: 01.10.2015
        enddate: 31.10.2015
*/
type Projects struct {
	ParseData []Prjinfo
}

type Prjinfo struct {
	Prj string
	Id int
	Query string
	Startdate string
	Enddate string
}

func (y *YamlTestEngine) TestReadPrjDetails(c *C) {
	content := y.readInFile("./moreofAKind.csv")
	projects := y.parseProjectsFromByteStream(content)
	c.Assert(len(projects.ParseData), Equals, 2)
	if len(projects.ParseData) == 2 {
		c.Assert(projects.ParseData[0].Prj, Equals, "SOLUT")
		c.Assert(projects.ParseData[0].Id, Equals, 10941)
		c.Assert(projects.ParseData[0].Query, Equals, "project = SOLUT")
		c.Assert(projects.ParseData[0].Startdate, Equals, "01.01.2015")
		c.Assert(projects.ParseData[0].Enddate, Equals, "")
		c.Assert(projects.ParseData[1].Prj, Equals, "SOLUT")
		c.Assert(projects.ParseData[1].Id, Equals, 10941)
		c.Assert(projects.ParseData[1].Query, Equals, "project = SOLUT")
		c.Assert(projects.ParseData[1].Startdate, Equals, "01.10.2015")
		c.Assert(projects.ParseData[1].Enddate, Equals, "31.10.2015")
	}
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

func (y *YamlTestEngine) parseProjectsFromByteStream(content []byte) *Projects {
	r := csv.NewReader(strings.NewReader(string(content)))
	r.Comma = ','
	r.Comment = '#'

	records, err := r.ReadAll()
	if err != nil {
		panic(err)
		return nil
	}

	var projects Projects
	projects.ParseData = make([]Prjinfo, len(records))

	for i := 0; i < len(records); i++ {
		projects.ParseData[i].Prj       = strings.TrimSpace(records[i][0])
		k, parseErr        			   := strconv.Atoi(strings.TrimSpace(records[i][1]))
		if parseErr != nil {
			fmt.Printf("\"%s\"", records[i][1])
			projects.ParseData[i].Id = -1
		} else {
			projects.ParseData[i].Id = k
		}
		projects.ParseData[i].Query     = strings.TrimSpace(records[i][2])
		projects.ParseData[i].Startdate = strings.TrimSpace(records[i][3])
		projects.ParseData[i].Enddate   = strings.TrimSpace(records[i][4])
	}
	return &projects
}


