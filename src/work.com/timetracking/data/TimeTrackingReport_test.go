package data

import (
	. "gopkg.in/check.v1"
	"testing"
	"time"
	"strings"
)

type DataTestEngine struct{}

func TestDataTestEngine(t *testing.T) {
	Suite(&DataTestEngine{})
	TestingT(t)
}

func (y *DataTestEngine) TestReadPrjDetails(c *C) {
	var projects TimeTrackingReport
	projects.Initialize("./moreofAKind.csv", ',')
	myTime := time.Now()
	c.Assert(len(projects.Settings), Equals, 2)
	if len(projects.Settings) == 2 {
		c.Assert(projects.Settings[strings.ToLower("SOLUT")].Prj, Equals, "SOLUT")
		c.Assert(projects.Settings[strings.ToLower("SOLUT")].Id, Equals, 10941)
		c.Assert(projects.Settings[strings.ToLower("SOLUT")].Query, Equals, "project = SOLUT")
		c.Assert(projects.Settings[strings.ToLower("SOLUT")].Startdate.GetTime().Format("02.01.2006"), Equals, "01.01.2015")
		c.Assert(projects.Settings[strings.ToLower("SOLUT")].Enddate.GetTime().Format("02.01.2006"), Equals, myTime.Format("02.01.2006"))
		c.Assert(projects.Settings[strings.ToLower("SOLUT")].ProductOwner, Equals, "MARTIN")

		c.Assert(projects.Settings[strings.ToLower("NOTSOL")].Prj, Equals, "NOTSOL")
		c.Assert(projects.Settings[strings.ToLower("NOTSOL")].Id, Equals, 10941)
		c.Assert(projects.Settings[strings.ToLower("NOTSOL")].Query, Equals, "project = SOLUT")
		c.Assert(projects.Settings[strings.ToLower("NOTSOL")].Startdate.GetTime().Format("02.01.2006"), Equals, "01.10.2015")
		c.Assert(projects.Settings[strings.ToLower("NOTSOL")].Enddate.GetTime().Format("02.01.2006"), Equals, "31.10.2015")
		c.Assert(projects.Settings[strings.ToLower("NOTSOL")].ProductOwner, Equals, "FRIEDRICH")
	}
}
