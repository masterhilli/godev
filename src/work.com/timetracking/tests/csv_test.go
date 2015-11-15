package testinterfaces

import (
	"fmt"
	"path/filepath"
	"encoding/csv"
	"strings"
	"io/ioutil"
	"testing"
	"strconv"
	. "gopkg.in/check.v1"
)

type CSVTestEngine struct {}

func TestCSVEngine(t *testing.T) {
	Suite(&CSVTestEngine{})
	TestingT(t)
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

func (y *CSVTestEngine) TestReadPrjDetails(c *C) {
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

func (y *CSVTestEngine) readInFile(filename string) []byte {
	filename, _ = filepath.Abs(filename)
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return content
}


func (y *CSVTestEngine) parseProjectsFromByteStream(content []byte) *Projects {
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