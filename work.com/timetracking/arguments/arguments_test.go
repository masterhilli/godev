package arguments

import (
	. "gopkg.in/check.v1"
	"testing"
	"time"
)

type ArgumentTestEngine struct {
	ta TimeTrackingArgs
}

func TestRegisterArgumentTestEngine(t *testing.T) {
	Suite(&ArgumentTestEngine{})
	TestingT(t)
}


// we ignore that one because depending on the IDE we use, the args are different
func (ate *ArgumentTestEngine) IgnoreSettingArgumentsWithOnly1Argument(c *C) {
	var ta TimeTrackingArgs = GetArguments()
	c.Assert(ta.GetCountParsedArgs(), Equals, 1)
}

func (ate *ArgumentTestEngine) TestSetTeammemberFile(c *C) {
	var ta TimeTrackingArgs = GetArguments()
	ta.parseAllArguments([]string{"test.exe", "tm=TestFilename"})
	c.Assert(ta.GetFilePathToTeammembers(), Equals, "TestFilename")
}

func (ate *ArgumentTestEngine) TestSetProjectsFile(c *C) {
	var ta TimeTrackingArgs = GetArguments()
	ta.parseAllArguments([]string{"test.exe", "tm=TestFilename", "prj=myProjectFile.csv"})
	c.Assert(ta.GetFilePathToProjects(), Equals, "myProjectFile.csv")
}

func (ate *ArgumentTestEngine) TestSetProjectsFileNotLowerCase(c *C) {
	var ta TimeTrackingArgs = GetArguments()
	ta.parseAllArguments([]string{"test.exe", "tm=TestFilename", "PRJ=myProjectFile.csv"})
	c.Assert(ta.GetFilePathToProjects(), Equals, "myProjectFile.csv")
}

func (ate *ArgumentTestEngine) TestIsStringArgument(c *C) {
	retVal := ate.ta.isStringArg("bla=SomthingElse")
	c.Assert(retVal, Equals, true)
}

func (ate *ArgumentTestEngine) TestIsNotStringArgument(c *C) {
	retVal := ate.ta.isStringArg("-blaSomthingElse")
	c.Assert(retVal, Equals, false)
}

func (ate *ArgumentTestEngine) TestIsBooleanArgument(c *C) {
	retVal := ate.ta.isBooleanArg("-bla")
	c.Assert(retVal, Equals, true)
}

func (ate *ArgumentTestEngine) TestIsNotBooleanArgument(c *C) {
	retVal := ate.ta.isBooleanArg("b-laSomthingElse")
	c.Assert(retVal, Equals, false)
}

func (ate *ArgumentTestEngine) TestSetFlagForSprintStatistic(c *C) {
	var ta TimeTrackingArgs = GetArguments()
	ta.parseAllArguments([]string{"test.exe", "tm=TestFilename", "prj=myProjectFile.csv", "-sprint"})
	c.Assert(ta.sprintStatistic, Equals, true)
}

func (ate *ArgumentTestEngine) TestNotSetFlagForSprintStatistic(c *C) {
	var ta TimeTrackingArgs = GetArguments()
	ta.parseAllArguments([]string{"test.exe", "tm=TestFilename", "prj=myProjectFile.csv"})
	c.Assert(ta.sprintStatistic, Equals, false)
}

func (ate *ArgumentTestEngine) TestSetStartDate(c *C) {
	var ta TimeTrackingArgs = GetArguments()
	ta.parseAllArguments([]string{"test.exe", "tm=TestFilename", "PRJ=myProjectFile.csv", "start?5.1.2015", "-sprint"})
	t := time.Date(2015, time.January, 5, 0, 0, 0, 0, time.UTC)
	c.Assert(ta.startDate, Equals, t)
}

func (ate *ArgumentTestEngine) TestGetEndDate(c *C) {
	var ta TimeTrackingArgs = GetArguments()
	ta.parseAllArguments([]string{"test.exe", "tm=TestFilename", "PRJ=myProjectFile.csv", "start?25.12.2015", "-sprint"})
	t := time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC)
	c.Assert(ta.GetEndDate(), Equals, t)
}

func (ate *ArgumentTestEngine) TestNotSetAStartDate(c *C) {
	var ta TimeTrackingArgs = GetArguments()
	ta.parseAllArguments([]string{"test.exe", "tm=TestFilename", "PRJ=myProjectFile.csv"})
	t := time.Date(0, time.January, 0, 0, 0, 0, 0, time.UTC)
	c.Assert(ta.startDate, Equals, t)
}

func (ate *ArgumentTestEngine) TestCreateTimeLayoutAll2digits(c *C) {
	layout := ate.ta.createTimeLayout("15.01.2015")
	c.Assert(layout, Equals, "02.01.2006")
}

func (ate *ArgumentTestEngine) TestCreateTimeLayoutDay1digit(c *C) {
	layout := ate.ta.createTimeLayout("1.01.2015")
	c.Assert(layout, Equals, "2.01.2006")
}

func (ate *ArgumentTestEngine) TestCreateTimeLayoutMonth1digit(c *C) {
	layout := ate.ta.createTimeLayout("15.1.2015")
	c.Assert(layout, Equals, "02.1.2006")
}

func (ate *ArgumentTestEngine) TestCreateTimeLayoutAll1digit(c *C) {
	layout := ate.ta.createTimeLayout("1.1.2015")
	c.Assert(layout, Equals, "2.1.2006")
}

func (ate *ArgumentTestEngine) TestParseIntoTimeObjAll2digits(c *C) {
	parsedTime := ate.ta.parseIntoTimeObj("15.01.2015")
	t := time.Date(2015, time.January, 15, 0, 0, 0, 0, time.UTC)
	c.Assert(parsedTime, Equals, t)
}

func (ate *ArgumentTestEngine) TestParseIntoTimeObjDay1digit(c *C) {
	parsedTime := ate.ta.parseIntoTimeObj("1.11.2015")
	t := time.Date(2015, time.November, 1, 0, 0, 0, 0, time.UTC)
	c.Assert(parsedTime, Equals, t)
}

func (ate *ArgumentTestEngine) TestParseIntoTimeObjMonth1digit(c *C) {
	parsedTime := ate.ta.parseIntoTimeObj("15.1.2015")
	t := time.Date(2015, time.January, 15, 0, 0, 0, 0, time.UTC)
	c.Assert(parsedTime, Equals, t)
}

func (ate *ArgumentTestEngine) TestParseIntoTimeObjAll1digit(c *C) {
	parsedTime := ate.ta.parseIntoTimeObj("1.1.2015")
	t := time.Date(2015, time.January, 1, 0, 0, 0, 0, time.UTC)
	c.Assert(parsedTime, Equals, t)
}


func (ate *ArgumentTestEngine) TestParseAllArgumentsIntIntoReportId(c *C) {
	var arg TimeTrackingArgs
	arg.parseAllArguments([]string{"bla.ignore", "report#7", "-t"})
	c.Assert(arg.reportId, Equals, 7)
}

func (ate *ArgumentTestEngine) TestIsIntArgReturnsTrue(c *C) {
	var arg TimeTrackingArgs
	isInt := arg.isIntArg("report#2")
	c.Assert(isInt, Equals, true)
}

func (ate *ArgumentTestEngine) TestSetIntVariableReturnsReportIdIs5(c *C) {
	var arg TimeTrackingArgs
	arg.setIntVariable("report", "3")
	c.Assert(arg.reportId, Equals, 3)
}

func (ate *ArgumentTestEngine) TestParseIntArgReturns5asReportId(c *C) {
	var arg TimeTrackingArgs
	arg.parseIntArg("report#5")
	c.Assert(arg.reportId, Equals, 5)
}


func (ate *ArgumentTestEngine) TestParseAllArgsReturnsManyAsserts(c *C) {
	var arg TimeTrackingArgs = GetArguments()

	arg.parseAllArguments([]string{"go", "run", "timetracking.go", "-t", "report#4"})

	c.Assert(arg.testing, Equals, true)
	c.Assert(arg.reportId, Equals, 4)
}