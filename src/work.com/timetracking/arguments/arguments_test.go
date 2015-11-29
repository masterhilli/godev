package arguments

import (
	. "gopkg.in/check.v1"
	"testing"
)

type ArgumentTestEngine struct {
	ta TimetrackingArgs
}

func TestRegisterArgumentTestEngine(t *testing.T) {
	Suite(&ArgumentTestEngine{})
	TestingT(t)
}

func (ate *ArgumentTestEngine) TestSettingArgumentsWithOnly1Argument(c *C) {
	ate.ta.ParseArguments([]string{"test.exe"})
	c.Assert(ate.ta.GetCountParsedArgs(), Equals, 0)
}

func (ate *ArgumentTestEngine) TestSetTeammemberFile(c *C) {
	ate.ta.ParseArguments([]string{"test.exe", "tm=TestFilename"})
	c.Assert(ate.ta.GetFilePathToTeammembers(), Equals, "TestFilename")
}

func (ate *ArgumentTestEngine) TestSetProjectsFile(c *C) {
	ate.ta.ParseArguments([]string{"test.exe", "tm=TestFilename", "prj=myProjectFile.csv"})
	c.Assert(ate.ta.GetFilePathToProjects(), Equals, "myProjectFile.csv")
}

func (ate *ArgumentTestEngine) TestSetProjectsFileNotLowerCase(c *C) {
	ate.ta.ParseArguments([]string{"test.exe", "tm=TestFilename", "PRJ=myProjectFile.csv"})
	c.Assert(ate.ta.GetFilePathToProjects(), Equals, "myProjectFile.csv")
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
	ate.ta.ParseArguments([]string{"test.exe", "tm=TestFilename", "PRJ=myProjectFile.csv", "-sprint"})
	c.Assert(ate.ta.sprintStatistic, Equals, true)
}

func (ate *ArgumentTestEngine) TestNotSetFlagForSprintStatistic(c *C) {
	ate.ta.ParseArguments([]string{"test.exe", "tm=TestFilename", "PRJ=myProjectFile.csv"})
	c.Assert(ate.ta.sprintStatistic, Equals, false)
}

func (ate *ArgumentTestEngine) TestSetStartDate(c *C) {
	ate.ta.ParseArguments([]string{"test.exe", "tm=TestFilename", "PRJ=myProjectFile.csv", "start?5.1.2015", "-sprint"})
	c.Assert(ate.ta.startDate, Equals, "5.1.2015")
}

func (ate *ArgumentTestEngine) TestNotSetAStartDate(c *C) {
	ate.ta.ParseArguments([]string{"test.exe", "tm=TestFilename", "PRJ=myProjectFile.csv"})
	c.Assert(ate.ta.startDate, Equals, "")
}
