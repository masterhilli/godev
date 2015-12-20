package data

import (
	. "gopkg.in/check.v1"
	"strings"
	"testing"
	"time"
)

type DataTestEngine struct{}

const pathToTestCSVFileOfProjects string = "../__testdata/moreofAKind.csv"

func TestDataTestEngine(t *testing.T) {
	Suite(&DataTestEngine{})
	TestingT(t)
}

func (y *DataTestEngine) TestReadPrjDetails(c *C) {
	var projects TimeTrackingReport
	projects.Initialize(pathToTestCSVFileOfProjects)
	myTime := time.Now()
	c.Assert(projects.GetSettingsLen(), Equals, 2)
	if projects.GetSettingsLen() == 2 {
		entry := projects.GetEntry(strings.ToLower("SOLUT"))
		c.Assert(entry.prj, Equals, "SOLUT")
		c.Assert(entry.id, Equals, 10941)
		c.Assert(entry.query, Equals, "project = SOLUT")
		c.Assert(entry.startdate.GetTime().Format("02.01.2006"), Equals, "01.01.2015")
		c.Assert(entry.enddate.GetTime().Format("02.01.2006"), Equals, myTime.Format("02.01.2006"))
		c.Assert(entry.GetProductOwner(), Equals, "MARTIN")

		entry = projects.GetEntry(strings.ToLower("NOTSOL"))
		c.Assert(entry.prj, Equals, "NOTSOL")
		c.Assert(entry.id, Equals, 10941)
		c.Assert(entry.query, Equals, "project = SOLUT")
		c.Assert(entry.startdate.GetTime().Format("02.01.2006"), Equals, "01.10.2015")
		c.Assert(entry.enddate.GetTime().Format("02.01.2006"), Equals, "31.10.2015")
		c.Assert(entry.GetProductOwner(), Equals, "FRIEDRICH")
	}
}
