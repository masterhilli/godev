package prjinfo

import (
	. "gopkg.in/check.v1"
	"testing"
	"time"
)

type DataTestEngine struct{}

func TestDataTestEngine(t *testing.T) {
	Suite(&DataTestEngine{})
	TestingT(t)
}

func (y *DataTestEngine) TestReadPrjDetails(c *C) {
	var projects Projects
	projects.Initialize("./moreofAKind.csv", ',')
	myTime := time.Now()
	c.Assert(len(projects.Data), Equals, 2)
	if len(projects.Data) == 2 {
		c.Assert(projects.Data[0].Prj, Equals, "SOLUT")
		c.Assert(projects.Data[0].Id, Equals, 10941)
		c.Assert(projects.Data[0].Query, Equals, "project = SOLUT")
		c.Assert(projects.Data[0].Startdate.t.Format("02.01.2006"), Equals, "01.01.2015")
		c.Assert(projects.Data[0].Enddate.t.Format("02.01.2006"), Equals, myTime.Format("02.01.2006"))
		c.Assert(projects.Data[1].Prj, Equals, "SOLUT")
		c.Assert(projects.Data[1].Id, Equals, 10941)
		c.Assert(projects.Data[1].Query, Equals, "project = SOLUT")
		c.Assert(projects.Data[1].Startdate.t.Format("02.01.2006"), Equals, "01.10.2015")
		c.Assert(projects.Data[1].Enddate.t.Format("02.01.2006"), Equals, "31.10.2015")
	}
}
