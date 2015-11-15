package testinterfaces

import (
	"testing"
	. "gopkg.in/check.v1"
	"time"
)

type TestDateEngine struct {}

func TestDateConversionEngine(t *testing.T) {
	Suite(&TestDateEngine{})
	TestingT(t)
}

//

func (d *TestDateEngine) TestParseDDMMYYYYFormat(c *C) {
	test := d.ParseDateString("31.12.2015")
	t := time.Date(2015, time.December, 31, 0, 0, 0, 0, time.UTC)
    c.Assert(test, Equals, t)
}

func (d *TestDateEngine) TestFormatToUrlFormat(c *C) {
	t := time.Date(2015, time.December, 31, 0, 0, 0, 0, time.UTC)
    c.Assert("31%2FDec%2F15", Equals, d.FormatToUrlDate(t))
}

func (d *TestDateEngine) FormatToUrlDate(t time.Time) string {
	urldate := t.Format("2") + "%2F" + t.Format("Jan")+"%2F" + t.Format("06")
	return urldate
}

func (d *TestDateEngine) ParseDateString(date string) time.Time  {
	test, err := time.Parse("02.01.2006", date)
    if err != nil {
        panic(err)
    }
    return test
}