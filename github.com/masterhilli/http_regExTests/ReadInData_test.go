package testinterfaces

import (
	. "gopkg.in/check.v1"
	"io/ioutil"
	"regexp"
	"testing"
)

const pathToTestReportJiraHtml string = "../__testdata/Report-Jira.html"

// Hook up gocheck into the "go test" runner.
type ReadInDataTestEngine struct {
	regexpToFindNames      string
	regexpToFindTotalTimes string
}

func NewReadInDataTestEngine() *ReadInDataTestEngine {
	return &ReadInDataTestEngine{
		regexpToFindNames:      "(?is)<tr>.*[<td [[:ascii:]]*>[[:alpha:]]*\\.[[:alpha:]]*</td>]*[[:space:]]*<td class=\"main\"><b>Total</b></td>[[:space:]]*</tr>",
		regexpToFindTotalTimes: "(?is)<tr>([[:space:]]*<td class=\"total\"><b>(([0-9]*[wdhms]{1})+|total|Issue)</b></td>)+[[:space:]]*</tr>",
	}
}

func TestReadInData(t *testing.T) {
	Suite(NewReadInDataTestEngine())
	TestingT(t)
}

func (s *ReadInDataTestEngine) checkForError(c *C, e error) {
	c.Assert(e, Equals, nil)
}

func (s *ReadInDataTestEngine) TestReadingInWholeFile(c *C) {
	data, err := ioutil.ReadFile(pathToTestReportJiraHtml)
	s.checkForError(c, err)
	c.Assert(len(data), Equals, 176349)
}

func (s *ReadInDataTestEngine) TestReadInAndSubMatchForNames(c *C) {
	s.ReadInFileAndFindRegExp(c, s.regexpToFindNames, 1)
}

func (s *ReadInDataTestEngine) TestReadInAndSubMatchForTotalTimes(c *C) {
	s.ReadInFileAndFindRegExp(c, s.regexpToFindTotalTimes, 2)
}

func (s *ReadInDataTestEngine) ReadInFileAndFindRegExp(c *C, regexpToFind string, countToAssertOn int) {
	data, err := ioutil.ReadFile(pathToTestReportJiraHtml)
	s.checkForError(c, err)
	s.AssertOnSubmatch(c, regexpToFind, countToAssertOn, string(data))
}

func (s *ReadInDataTestEngine) AssertOnSubmatch(c *C, regexpToFind string, countToAssertOn int, dataString string) {
	regexpForMatchingNames := regexp.MustCompile(regexpToFind) //<td class=\"main\">Total</td>"
	indexArray := regexpForMatchingNames.FindAllStringSubmatchIndex(dataString, -1)
	var countOfFoundSubmatches int = 0
	if indexArray != nil {
		countOfFoundSubmatches = len(indexArray)
	}
	c.Assert(countOfFoundSubmatches, Equals, countToAssertOn)
}
