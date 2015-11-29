package prjinfo

import (
	. "gopkg.in/check.v1"
	"testing"
	"time"
)

type TestDateEngine struct{}

func TestDateConversionEngine(t *testing.T) {
	Suite(&TestDateEngine{})
	TestingT(t)
}

func (d *TestDateEngine) TestParseDDMMYYYYFormat(c *C) {
	var jd JiraDate
	jd.Initialize("31.12.2015")
	t := time.Date(2015, time.December, 31, 0, 0, 0, 0, time.UTC)
	c.Assert(jd.t, Equals, t)
}

func (d *TestDateEngine) TestFormatToUrlFormat(c *C) {
	var jd JiraDate
	jd.Initialize("31.12.2015")
	c.Assert(jd.GetTimeForUrl(), Equals, "31%2FDec%2F15")
}

func (d *TestDateEngine) TestDurationParser(c *C) {
	duration, err := time.ParseDuration("5h30m")
	if err != nil {
		panic(err)
	}

	c.Assert(duration.Hours(), Equals, 5.5)
}
