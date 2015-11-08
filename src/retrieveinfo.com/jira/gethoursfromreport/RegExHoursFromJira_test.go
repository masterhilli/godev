package jiraRegEx

import (
//	"fmt"
	"regexp"
	"testing"
	. "gopkg.in/check.v1"
)


// Hook up gocheck into the "go test" runner.
type RegexpTestEngine struct{}
func TestRegexp(t *testing.T) { 
	Suite(&RegexpTestEngine{})
	TestingT(t) 
}

var testString string = "<tr>---</tr>"
var sepcialChars string = "^!&%()=?`'#+*_-:.;,><"
// TestTRTRNovaluesInst
func (s *RegexpTestEngine) TestTRTRRegExParser(c *C) {
	actual := ReturnTRValues(testString)
	expected := "---"
	c.Assert (actual[0], Equals, expected)
}

func (s *RegexpTestEngine) TestTRTRWithNumAndChars(c *C) {
	actual := ReturnTRValues("<tr>BLA123BLA</tr>")
	expected := "BLA123BLA"
	c.Assert ( actual[0], Equals,  expected)
}


func (s *RegexpTestEngine) TestTRTRNovaluesIn(c *C) {
	actual := ReturnTRValues("<tr></tr>")
	expected := ""
	c.Assert ( actual[0], Equals,  expected)
}

func (s *RegexpTestEngine) TestTRTRNoMatch(c *C) {
	actual := ReturnTRValues("<tr</tr>")
	expected := ""
	c.Assert ( actual[0], Equals,  expected)
}


func (s *RegexpTestEngine) TestTRTRSpecialCharacters(c *C) {
	actual := ReturnTRValues("<tr>"+sepcialChars+"</tr>")
	expected := sepcialChars
	c.Assert ( actual[0], Equals,  expected)
}


func (s *RegexpTestEngine) TestTRTRSearchInLongerString(c *C) {
	actual := ReturnTRValues("scvjewoqfäiusaölkenskajflaöwkepori<tr>"+sepcialChars+"</tr>´dsfkjhjhsajkhfdahdskhaksdfkjas")
	expected := sepcialChars
	c.Assert ( actual[0], Equals,  expected)
}

var myJiraPageTestString string = "<div style=\"width: 100%; overflow-x: auto\">\n"+
"    <table border=\"0\" cellpadding=\"3\" cellspacing=\"1\" class=\"main\">\n"+
"        <tbody><tr>\n"+
"            <td class=\"main\" colspan=\"2\">\n"+
"                1/Sep - 2/Nov\n"+
"            </---Problemheretd>\n"+
"                            <td class=\"main\">anton.jessner</td>\n"+
"                            <td class=\"main\">david.hangl</td>\n"+
"                            <td class=\"main\">leonardo.fisic</td>\n"+
"                            <td class=\"main\">marc.dopplinger</td>\n"+
"                            <td class=\"main\">martin.hillbrand</td>\n"+
"                            <td class=\"main\">richard.nusser</td>\n"+
"                            <td class=\"main\">serhat.ekinci</td>\n"+
"                            <td class=\"main\">thomas.pinetz</td>\n"+
"                            <td class=\"main\">thomas.rauscher</td>\n"+
"                        <td class=\"main\"><b>Total</b></td>\n"+
"        </tr>\n"+
"        \n"+
"\n"+
"<tr>\n"+
"    <td class=\"total\"><b>Issue</b></td>\n"+
"    <td class=\"total\"><b>Total</b></td>\n"+
"                <td class=\"total\"><b>1h</b></td>\n"+
"                    <td class=\"total\"><b>34h</b></td>\n"+
"                    <td class=\"total\"><b>13h40m</b></td>\n"+
"                    <td class=\"total\"><b>44h</b></td>\n"+
"                    <td class=\"total\"><b>23h30m</b></td>\n"+
"                    <td class=\"total\"><b>10h50m</b></td>\n"+
"                    <td class=\"total\"><b>6h</b></td>\n"+
"                    <td class=\"total\"><b>37h30m</b></td>\n"+
"                    <td class=\"total\"><b>1h35m</b></td>\n"+
"                <td class=\"total\"><b>172h5m</b></td>\n"+
"</tr>"+
"<tr>\n"+
"    <td class=\"total\"><b>Issue</b></td>\n"+
"    <td class=\"total\"><b>Total</b></td>\n"+
"                <td class=\"total\"><b>1h</b></td>\n"+
"                    <td class=\"total\"><b>34h</b></td>\n"+
"                    <td class=\"total\"><b>13h40m</b></td>\n"+
"                    <td class=\"total\"><b>44h</b></td>\n"+
"                    <td class=\"total\"><b>23h30m</b></td>\n"+
"                    <td class=\"total\"><b>10h50m</b></td>\n"+
"                    <td class=\"total\"><b>6h</b></td>\n"+
"                    <td class=\"total\"><b>37h30m</b></td>\n"+
"                    <td class=\"total\"><b>1h35m</b></td>\n"+
"                <td class=\"total\"><b>172h5m</b></td>\n"+
"</tr>"+
"<tr>\n"+
"    <td class=\"total\"><b>Issue</b></td>\n"+
"    <td class=\"total\"><b>Total</b></td>\n"+
"                <td class=\"total\"><b>1h</b></td>\n"+
"                    <td class=\"total\"><b>34h</b></td>\n"+
"                    <td class=\"total\"><b>13h40m</b></td>\n"+
"                    <td class=\"total\"><b>44h</b></td>\n"+
"                    <td class=\"total\"><b>23h30m</b></td>\n"+
"                    <td class=\"total\"><b>10h50m</b></td>\n"+
"                    <td class=\"total\"><b>6h</b></td>\n"+
"                    <td class=\"total\"><b>37h30m</b></td>\n"+
"                    <td class=\"total\"><b>1h35m</b></td>\n"+
"                <td class=\"total\"><b>172h5m</b></td>\n"+
"</tr>"

func (s *RegexpTestEngine) TestTRTROwnHTMLStreamShouldReturn2matches(c *C) {
	actual := ReturnTRValues("<tr>...</tr>afkjajfladslkflkasd<tr>llll</tr>")
	expected := 2
	c.Assert ( len(actual), Equals,  expected)	
}

var tdtdTestString string = "<td class=\"main\"><b>34h</b></td>\n"+
"                    		  <td class=\"main\"><b>13h40m</b></td>"

func (s *RegexpTestEngine) TestTagTDWithTwoMatchesShouldResult2(c *C) {
	//trTagValuesInArray := ReturnTRValues(myJiraPageTestString)
	actual := ReturnTDClassMainValues(tdtdTestString)
	expected := 2
	c.Assert ( len(actual), Equals,  expected)	
}

var tdtdJiraTestString string = "<td class=\"main\" colspan=\"2\">\n"+
"                1/Sep - 2/Nov\n"+
"            </---Problemheretd>\n"+
"                            <td class=\"main\">anton.jessner</td>"+
"                            <td class=\"main\">david.hangl</td>"+
"                            <td class=\"main\">leonardo.fisic</td>"+
"                            <td class=\"main\">marc.dopplinger</td>"+
"                            <td class=\"main\">martin.hillbrand</td>n"+
"                            <td class=\"main\">richard.nusser</td>"+
"                            <td class=\"main\">serhat.ekinci</td>"+
"                            <td class=\"main\">thomas.pinetz</td>"+
"                            <td class=\"main\">thomas.rauscher</td>"+
"                        <td class=\"main\"><b>Total</b></td>"


func (s *RegexpTestEngine) TestTagTDWithJiraStringShouldResult10(c *C) {
	//trTagValuesInArray := ReturnTRValues(myJiraPageTestString)
	actual := ReturnTDClassMainValues(tdtdJiraTestString)
	expected := 10
	c.Assert ( len(actual), Equals,  expected)	
}

func (s *RegexpTestEngine) TestParseForTagStartEnd(c *C){
	indexArray := parseForTagStartEnd(tdtdJiraTestString, "td", " class=\"main\"")
	c.Assert( len(indexArray), Equals, 10)
}

func (s *RegexpTestEngine) TestReturnIndexPairsShouldReturn10(c *C){
	indexArray := returnIndexPairs(tdtdJiraTestString, "td", " class=\"main\"")
	c.Assert( len(indexArray),  Equals,10)
}


func (s *RegexpTestEngine) TestTRReturnIndexArrayForStartTagShouldReturn4(c *C){
	//var tag string = "td"
	//var attributes string = " class=\"main\""
	indexArray := returnIndexArray(myJiraPageTestString, "(?is)<tr>")
	//indexArray := returnIndexArray(tdtdJiraTestString, "<"+tag+ attributes+">")
	c.Assert( len(indexArray),  Equals,4)
}

func (s *RegexpTestEngine) TestFindAllStringSubmatchIndex(c *C) {
	var stringToParse string = "<tr><tr>dajfskjhfakdsjfhkasfhkjashfd<tr>\n\nsdadsasd<tr>"
	regexpIndexFinder := regexp.MustCompile("(?is)<tr>")
	indexArray := regexpIndexFinder.FindAllStringSubmatchIndex(stringToParse, -1)
	var arrayLen int = 0
	if indexArray != nil {
		arrayLen = len(indexArray)
//		fmt.Printf("------- indexArray: %d / %d", arrayLen, len(indexArray[0]))
		for i := 0; i < arrayLen; i++ {
			for k := 0; k < len(indexArray[i]); k++ {
//				fmt.Printf("------- indexArray[%d][%d]: %d : %s\n",i, k, indexArray[i][k], stringToParse[indexArray[i][k]: len(stringToParse)])
			}
		}
	}
	c.Assert( arrayLen,  Equals,4)
}

func (s *RegexpTestEngine) TestTRReturnIndexArrayForEndTagShouldReturn4(c *C){
	//var tag string = "td"
	//var attributes string = " class=\"main\""
	indexArray := returnIndexArray(tdtdJiraTestString, "</tr>")
	//indexArray := returnIndexArray(tdtdJiraTestString, "<"+tag+ attributes+">")
	c.Assert( len(indexArray),  Equals,0)
}


func (s *RegexpTestEngine) TestReturnIndexArrayForStartTagShouldReturn10(c *C){
	//var tag string = "td"
	//var attributes string = " class=\"main\""
	indexArray := returnIndexArray(tdtdJiraTestString, "<td class=\"main\">")
	//indexArray := returnIndexArray(tdtdJiraTestString, "<"+tag+ attributes+">")
	c.Assert( len(indexArray),  Equals,10)
}

func (s *RegexpTestEngine) TestReturnIndexArrayForEndTagShouldReturn10(c *C){
	//var tag string = "td"
	//var attributes string = " class=\"main\""
	indexArray := returnIndexArray(tdtdJiraTestString, "</td>")
	//indexArray := returnIndexArray(tdtdJiraTestString, "<"+tag+ attributes+">")
	c.Assert( len(indexArray),  Equals,10)
}

/*func returnValuesOfTag(stringToParse string, tag string, attributes string) [] string {
	indexArray := parseForTagStartEnd(stringToParse, tag, attributes)
	return trimTagsFromArray(indexArray, stringToParse, len(tag+attributes))
}*/

func (s *RegexpTestEngine) TestTagTDJiraHtmlStreamMatches10(c *C) {
	trTagValuesInArray := ReturnTRValues(myJiraPageTestString)
	actual := ReturnTDClassMainValues(trTagValuesInArray[0])
	expected := 10
	c.Assert ( len(actual), Equals,  expected)	
}


func (s *RegexpTestEngine) TestReturnIndexShouldReturn2(c *C) {
	index := returnIndexArray("<tr>...</tr>afkjajfladslkflkasd<tr>llll</tr>", "<tr>")

	c.Assert( len(index),  Equals,2)
}


func (s *RegexpTestEngine) TestJiraStreamShouldReturn2Index(c *C) {
	index := returnIndexArray(myJiraPageTestString, "<tr>")
	c.Assert( len(index),  Equals,4)
}
