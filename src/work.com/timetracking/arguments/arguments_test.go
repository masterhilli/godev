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
