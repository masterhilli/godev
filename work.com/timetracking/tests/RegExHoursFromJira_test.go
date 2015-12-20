package testinterfaces

import (
	. "gopkg.in/check.v1"
	"regexp"
	"testing"
)

// Hook up gocheck into the "go test" runner.
type RegexpJiraTestEngine struct {
	testString         string
	specialChars       string
	tdtdTestString     string
	tdtdJiraTestString string
	jiraPageTestString string
}

func NewRegexpJiraTestEngine() *RegexpJiraTestEngine {
	mytestString := "<tr>---</tr>"
	myspecialChars := "^!&%()=?`'#+*_-:.;,><"
	mytdtdTestString := "<td class=\"main\"><b>34h</b></td>\n                    		  <td class=\"main\"><b>13h40m</b></td>"
	mytdtdJiraTestString := "<td class=\"main\" colspan=\"2\">\n" +
		"                1/Sep - 2/Nov\n" +
		"            </---Problemheretd>\n" +
		"                            <td class=\"main\">anton.jessner</td>" +
		"                            <td class=\"main\">david.hangl</td>" +
		"                            <td class=\"main\">leonardo.fisic</td>" +
		"                            <td class=\"main\">marc.dopplinger</td>" +
		"                            <td class=\"main\">martin.hillbrand</td>n" +
		"                            <td class=\"main\">richard.nusser</td>" +
		"                            <td class=\"main\">serhat.ekinci</td>" +
		"                            <td class=\"main\">thomas.pinetz</td>" +
		"                            <td class=\"main\">thomas.rauscher</td>" +
		"                        <td class=\"main\"><b>Total</b></td>"

	myJiraPageTestString := "<div style=\"width: 100%; overflow-x: auto\">\n" +
		"    <table border=\"0\" cellpadding=\"3\" cellspacing=\"1\" class=\"main\">\n" +
		"        <tbody><tr>\n" +
		"            <td class=\"main\" colspan=\"2\">\n" +
		"                1/Sep - 2/Nov\n" +
		"            </---Problemheretd>\n" +
		"                            <td class=\"main\">anton.jessner</td>\n" +
		"                            <td class=\"main\">david.hangl</td>\n" +
		"                            <td class=\"main\">leonardo.fisic</td>\n" +
		"                            <td class=\"main\">marc.dopplinger</td>\n" +
		"                            <td class=\"main\">martin.hillbrand</td>\n" +
		"                            <td class=\"main\">richard.nusser</td>\n" +
		"                            <td class=\"main\">serhat.ekinci</td>\n" +
		"                            <td class=\"main\">thomas.pinetz</td>\n" +
		"                            <td class=\"main\">thomas.rauscher</td>\n" +
		"                        <td class=\"main\"><b>Total</b></td>\n" +
		"        </tr>\n" +
		"        \n" +
		"\n" +
		"<tr>\n" +
		"    <td class=\"total\"><b>Issue</b></td>\n" +
		"    <td class=\"total\"><b>Total</b></td>\n" +
		"                <td class=\"total\"><b>1h</b></td>\n" +
		"                    <td class=\"total\"><b>34h</b></td>\n" +
		"                    <td class=\"total\"><b>13h40m</b></td>\n" +
		"                    <td class=\"total\"><b>44h</b></td>\n" +
		"                    <td class=\"total\"><b>23h30m</b></td>\n" +
		"                    <td class=\"total\"><b>10h50m</b></td>\n" +
		"                    <td class=\"total\"><b>6h</b></td>\n" +
		"                    <td class=\"total\"><b>37h30m</b></td>\n" +
		"                    <td class=\"total\"><b>1h35m</b></td>\n" +
		"                <td class=\"total\"><b>172h5m</b></td>\n" +
		"</tr>" +
		"<tr>\n" +
		"    <td class=\"total\"><b>Issue</b></td>\n" +
		"    <td class=\"total\"><b>Total</b></td>\n" +
		"                <td class=\"total\"><b>1h</b></td>\n" +
		"                    <td class=\"total\"><b>34h</b></td>\n" +
		"                    <td class=\"total\"><b>13h40m</b></td>\n" +
		"                    <td class=\"total\"><b>44h</b></td>\n" +
		"                    <td class=\"total\"><b>23h30m</b></td>\n" +
		"                    <td class=\"total\"><b>10h50m</b></td>\n" +
		"                    <td class=\"total\"><b>6h</b></td>\n" +
		"                    <td class=\"total\"><b>37h30m</b></td>\n" +
		"                    <td class=\"total\"><b>1h35m</b></td>\n" +
		"                <td class=\"total\"><b>172h5m</b></td>\n" +
		"</tr>" +
		"<tr>\n" +
		"    <td class=\"total\"><b>Issue</b></td>\n" +
		"    <td class=\"total\"><b>Total</b></td>\n" +
		"                <td class=\"total\"><b>1h</b></td>\n" +
		"                    <td class=\"total\"><b>34h</b></td>\n" +
		"                    <td class=\"total\"><b>13h40m</b></td>\n" +
		"                    <td class=\"total\"><b>44h</b></td>\n" +
		"                    <td class=\"total\"><b>23h30m</b></td>\n" +
		"                    <td class=\"total\"><b>10h50m</b></td>\n" +
		"                    <td class=\"total\"><b>6h</b></td>\n" +
		"                    <td class=\"total\"><b>37h30m</b></td>\n" +
		"                    <td class=\"total\"><b>1h35m</b></td>\n" +
		"                <td class=\"total\"><b>172h5m</b></td>\n" +
		"</tr>"

	return &RegexpJiraTestEngine{testString: mytestString, specialChars: myspecialChars, tdtdTestString: mytdtdTestString, tdtdJiraTestString: mytdtdJiraTestString, jiraPageTestString: myJiraPageTestString}
}

func TestRegexp(t *testing.T) {
	Suite(NewRegexpJiraTestEngine())
	TestingT(t)
}

func (s *RegexpJiraTestEngine) TestTRTRRegExParser(c *C) {
	actual := ReturnTRValues(s.testString)
	expected := "---"
	c.Assert(actual[0], Equals, expected)
}

func (s *RegexpJiraTestEngine) TestTRTRWithNumAndChars(c *C) {
	actual := ReturnTRValues("<tr>BLA123BLA</tr>")
	expected := "BLA123BLA"
	c.Assert(actual[0], Equals, expected)
}

func (s *RegexpJiraTestEngine) TestTRTRNovaluesIn(c *C) {
	actual := ReturnTRValues("<tr></tr>")
	expected := ""
	c.Assert(actual[0], Equals, expected)
}

func (s *RegexpJiraTestEngine) TestTRTRNoMatch(c *C) {
	actual := ReturnTRValues("<tr</tr>")
	expected := ""
	c.Assert(actual[0], Equals, expected)
}

func (s *RegexpJiraTestEngine) TestTRTRSpecialCharacters(c *C) {
	actual := ReturnTRValues("<tr>" + s.specialChars + "</tr>")
	expected := s.specialChars
	c.Assert(actual[0], Equals, expected)
}

func (s *RegexpJiraTestEngine) TestTRTRSearchInLongerString(c *C) {
	actual := ReturnTRValues("scvjewoqfäiusaölkenskajflaöwkepori<tr>" + s.specialChars + "</tr>´dsfkjhjhsajkhfdahdskhaksdfkjas")
	expected := s.specialChars
	c.Assert(actual[0], Equals, expected)
}

func (s *RegexpJiraTestEngine) TestTRTROwnHTMLStreamShouldReturn2matches(c *C) {
	actual := ReturnTRValues("<tr>...</tr>afkjajfladslkflkasd<tr>llll</tr>")
	expected := 2
	c.Assert(len(actual), Equals, expected)
}

func (s *RegexpJiraTestEngine) TestTagTDWithTwoMatchesShouldResult2(c *C) {
	actual := ReturnTDClassMainValues(s.tdtdTestString)
	expected := 2
	c.Assert(len(actual), Equals, expected)
}

func (s *RegexpJiraTestEngine) TestTagTDWithJiraStringShouldResult10(c *C) {
	actual := ReturnTDClassMainValues(s.tdtdJiraTestString)
	expected := 10
	c.Assert(len(actual), Equals, expected)
}

func (s *RegexpJiraTestEngine) TestParseForTagStartEnd(c *C) {
	indexArray := parseForTagStartEnd(s.tdtdJiraTestString, "td", " class=\"main\"")
	c.Assert(len(indexArray), Equals, 10)
}

func (s *RegexpJiraTestEngine) TestReturnIndexPairsShouldReturn10(c *C) {
	indexArray := returnIndexPairs(s.tdtdJiraTestString, "td", " class=\"main\"")
	c.Assert(len(indexArray), Equals, 10)
}

func (s *RegexpJiraTestEngine) TestTRReturnIndexArrayForStartTagShouldReturn4(c *C) {
	indexArray := returnIndexArray(s.jiraPageTestString, "(?is)<tr>")
	c.Assert(len(indexArray), Equals, 4)
}

func (s *RegexpJiraTestEngine) TestFindAllStringSubmatchIndex(c *C) {
	var stringToParse string = "<tr><tr>dajfskjhfakdsjfhkasfhkjashfd<tr>\n\nsdadsasd<tr>"
	regexpIndexFinder := regexp.MustCompile("(?is)<tr>")
	indexArray := regexpIndexFinder.FindAllStringSubmatchIndex(stringToParse, -1)
	var arrayLen int = 0
	if indexArray != nil {
		arrayLen = len(indexArray)
	}
	c.Assert(arrayLen, Equals, 4)
}

func (s *RegexpJiraTestEngine) TestTRReturnIndexArrayForEndTagShouldReturn4(c *C) {
	indexArray := returnIndexArray(s.tdtdJiraTestString, "</tr>")
	c.Assert(len(indexArray), Equals, 0)
}

func (s *RegexpJiraTestEngine) TestReturnIndexArrayForStartTagShouldReturn10(c *C) {
	indexArray := returnIndexArray(s.tdtdJiraTestString, "<td class=\"main\">")
	c.Assert(len(indexArray), Equals, 10)
}

func (s *RegexpJiraTestEngine) TestReturnIndexArrayForEndTagShouldReturn10(c *C) {
	indexArray := returnIndexArray(s.tdtdJiraTestString, "</td>")
	c.Assert(len(indexArray), Equals, 10)
}

func (s *RegexpJiraTestEngine) TestTagTDJiraHtmlStreamMatches10(c *C) {
	trTagValuesInArray := ReturnTRValues(s.jiraPageTestString)
	actual := ReturnTDClassMainValues(trTagValuesInArray[0])
	expected := 10
	c.Assert(len(actual), Equals, expected)
}

func (s *RegexpJiraTestEngine) TestReturnIndexShouldReturn2(c *C) {
	index := returnIndexArray("<tr>...</tr>afkjajfladslkflkasd<tr>llll</tr>", "<tr>")
	c.Assert(len(index), Equals, 2)
}

func (s *RegexpJiraTestEngine) TestJiraStreamShouldReturn2Index(c *C) {
	index := returnIndexArray(s.jiraPageTestString, "<tr>")
	c.Assert(len(index), Equals, 4)
}
