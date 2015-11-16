package prjinfo

import (
	. "gopkg.in/check.v1"
	"testing"
	"time"
)

type PrjInfoTestEngine struct{}

func TestPrjInfoTestEngine(t *testing.T) {
	Suite(&PrjInfoTestEngine{})
	TestingT(t)
}

func (y *PrjInfoTestEngine) TestReadPrjDetails(c *C) {
	var projects Projects
	projects.Initialize("./moreofAKind.csv", ',')
	myTime := time.Now()
	c.Assert(len(projects.ParseData), Equals, 2)
	if len(projects.ParseData) == 2 {
		c.Assert(projects.ParseData[0].Prj, Equals, "SOLUT")
		c.Assert(projects.ParseData[0].Id, Equals, 10941)
		c.Assert(projects.ParseData[0].Query, Equals, "project = SOLUT")
		c.Assert(projects.ParseData[0].Startdate.t.Format("02.01.2006"), Equals, "01.01.2015")
		c.Assert(projects.ParseData[0].Enddate.t.Format("02.01.2006"), Equals, myTime.Format("02.01.2006"))
		c.Assert(projects.ParseData[1].Prj, Equals, "SOLUT")
		c.Assert(projects.ParseData[1].Id, Equals, 10941)
		c.Assert(projects.ParseData[1].Query, Equals, "project = SOLUT")
		c.Assert(projects.ParseData[1].Startdate.t.Format("02.01.2006"), Equals, "01.10.2015")
		c.Assert(projects.ParseData[1].Enddate.t.Format("02.01.2006"), Equals, "31.10.2015")
	}
}
