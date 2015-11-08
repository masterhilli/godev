package jiraRegEx

	
import (
//    "bufio"
//    "fmt"
//    "io"
    "io/ioutil"
//    "os"
    "testing"
    . "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
type ReadInDataTestEngine struct{}
func TestReadInData(t *testing.T) { 
	Suite(&ReadInDataTestEngine{})
	TestingT(t) 
}
//var _ = 

func (s *ReadInDataTestEngine) checkForError(c *C, e error) {
	c.Assert(e, Equals, nil)
}

func (s *ReadInDataTestEngine) TestReadingInWholeFile(c *C) {
	data, err := ioutil.ReadFile("./testdata/Report-Jira.html")
    s.checkForError(c, err)
    c.Assert(173345, Equals, len(data))
}