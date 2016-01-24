package testinterfaces

import (
	"fmt"
	. "gopkg.in/check.v1"
	"regexp"
	"strconv"
	"testing"
)

type RegexpTestEngine struct{}

func TestRegexpTestEngine(t *testing.T) {
	Suite(&RegexpTestEngine{})
	TestingT(t)
}

func (s *RegexpTestEngine) TestParsingForTime(c *C) {
	var endTag string = "h"
	regexpForValue := regexp.MustCompile("(?is)[0-9]+" + endTag)
	valueFromRegExp := regexpForValue.FindStringSubmatch("23w37h30m32112h")
	if valueFromRegExp == nil {
		fmt.Printf("*** DEBUG: valueFromRegExp returned to nil\n")
	}

	match := valueFromRegExp[0]
	match = match[0 : len(match)-1]
	value, err := strconv.Atoi(match)
	if err != nil {
		panic(err)
	}
	c.Assert(value, Equals, 37)
}

func (s *RegexpTestEngine) Test1TimeInDataString(c *C) {
	s.AssertOnSubmatch(c, "[0-9]*[wdhms]{1}", 1, "<>23h<<")
}

func (s *RegexpTestEngine) Test3TimesInDataString(c *C) {
	s.AssertOnSubmatch(c, "[[0-9]*[wdhms]{1}]*", 3, "<>23h<<22m oqp89w<")
}

func (s *RegexpTestEngine) Test3TimesTimeAndTotalAndIssueInDataString(c *C) {
	s.AssertOnSubmatch(c, "(?is)[[0-9]*[wdhms]{1}|total|Issue]*", 5, "<>23h<total<22m oissueqp89w<")
}

func (s *RegexpTestEngine) TestChangableWithinAndSurroundings(c *C) {
	s.AssertOnSubmatch(c, "(?is)<t>([0-9]*[wdhms]{1}|total|Issue)</t>", 1, "<t>23h</t><total<22m oissueqp89w<")
}

func (s *RegexpTestEngine) TestParseForTagSurroundingChangableValueShouldReturn5hits(c *C) {
	s.AssertOnSubmatch(c, "(?is)(<t>([0-9]*[wdhms]{1}|total|Issue)</t>)+", 5, "<t>23h</t><lkslakl<t>total</t><daklskdlfak<t>22m</t> o<t>issue</t>qp<t>89w</t>")
}

//d class=\"total\"><b>
func (s *RegexpTestEngine) TestParseForTagSimilarToJiraSurroundingChangableValueShouldReturn5hits(c *C) {
	s.AssertOnSubmatch(c, "(?is)(<td class=\"total\"><b>([0-9]*[wdhms]{1}|total|Issue)</b></td>)+", 5, "<td class=\"total\"><b>23h</b></td><lkslakl<td class=\"total\"><b>total</b></td><daklskdlfak<td class=\"total\"><b>22m</b></td> o<td class=\"total\"><b>issue</b></td>qp<td class=\"total\"><b>89w</b></td>")
}

var testStringForTotalTimes string = "daskjfdslkjfaldjfladklfalkdsfjlksafdsajf" +
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
	"</tr>   " +
																			"      sakjdlfajsdlkfjdsalfjdsafalkdsjflkas"
var regexpToFindTotalTimesWorking string = "(?is)<tr>([[:space:]]*<td class=\"total\"><b>(([0-9]*[wdhms]{1})+|total|Issue)</b></td>)+[[:space:]]*</tr>" //"(?is)([[:space:]]*<td class=\"total\"><b>(([0-9]*[wdhms]{1})+|total|Issue)</b></td>)+"
func (s *RegexpTestEngine) TestParseForJiraTotalTimeStream(c *C) {
	s.AssertOnSubmatch(c, regexpToFindTotalTimesWorking, 1, testStringForTotalTimes)
}

func (s *RegexpTestEngine) AssertOnSubmatch(c *C, regexpToFind string, countToAssertOn int, dataString string) {
	regexpForMatchingNames := regexp.MustCompile(regexpToFind) //<td class=\"main\">Total</td>"
	indexArray := regexpForMatchingNames.FindAllStringSubmatchIndex(dataString, -1)
	var countOfFoundSubmatches int = 0
	if indexArray != nil {
		countOfFoundSubmatches = len(indexArray)
	}
	c.Assert(countOfFoundSubmatches, Equals, countToAssertOn)
}
