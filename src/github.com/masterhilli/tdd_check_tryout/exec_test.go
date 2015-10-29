package executeCmds

import (
	"fmt"
	"testing"
	"os/exec"
	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }
type MySuite struct{}
var _ = Suite(&MySuite{})


var teststring = ".\n..\n exec_test.go\n"

func (s *MySuite) TestSimpleExecutionOfLSCmdWithArgument(c *C) {
	out := s.executeCmd(c, "lsa", "-a")
	c.Assert(teststring, Equals, out)
}


func (s *MySuite)  executeCmd(c *C, name string, args ...string) string {
	out, err := exec.Command(name, args...).Output()
	if err != nil {
		c.Fatal(err)
	}
	return fmt.Sprintf("%s", out)
}