package jiraRegEx

import (
//	"fmt"
	"regexp"
	"testing"
)

func AssertEqualsString (t *testing.T, actual string, expected string) {
	if actual != expected {
		t.Fatalf("\nexpected: %s\nactual:  %s", expected, actual)
	}
}

func AssertEqualsInt (t *testing.T, actual int, expected int) {
	if actual != expected {
		t.Fatalf("\nexpected: %d\nactual:  %d", expected, actual)
	}
}

var testString string = "<tr>---</tr>"
var sepcialChars string = "^!&%()=?`'#+*_-:.;,><"
// TestTRTRNovaluesInst
func TestTRTRRegExParser(t *testing.T) {
	actual := ReturnTRValues(testString)
	expected := "---"
	AssertEqualsString (t, actual[0], expected)
}

func TestTRTRWithNumAndChars(t *testing.T) {
	actual := ReturnTRValues("<tr>BLA123BLA</tr>")
	expected := "BLA123BLA"
	AssertEqualsString (t, actual[0], expected)
}


func TestTRTRNovaluesIn(t *testing.T) {
	actual := ReturnTRValues("<tr></tr>")
	expected := ""
	AssertEqualsString (t, actual[0], expected)
}

func TestTRTRNoMatch(t *testing.T) {
	actual := ReturnTRValues("<tr</tr>")
	expected := ""
	AssertEqualsString (t, actual[0], expected)
}


func TestTRTRSpecialCharacters(t *testing.T) {
	actual := ReturnTRValues("<tr>"+sepcialChars+"</tr>")
	expected := sepcialChars
	AssertEqualsString (t, actual[0], expected)
}


func TestTRTRSearchInLongerString(t *testing.T) {
	actual := ReturnTRValues("scvjewoqfäiusaölkenskajflaöwkepori<tr>"+sepcialChars+"</tr>´dsfkjhjhsajkhfdahdskhaksdfkjas")
	expected := sepcialChars
	AssertEqualsString (t, actual[0], expected)
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

func TestTRTROwnHTMLStreamShouldReturn2matches(t *testing.T) {
	actual := ReturnTRValues("<tr>...</tr>afkjajfladslkflkasd<tr>llll</tr>")
	expected := 2
	AssertEqualsInt (t, len(actual), expected)	
}

var tdtdTestString string = "<td class=\"main\"><b>34h</b></td>\n"+
"                    		  <td class=\"main\"><b>13h40m</b></td>"

func TestTagTDWithTwoMatchesShouldResult2(t *testing.T) {
	//trTagValuesInArray := ReturnTRValues(myJiraPageTestString)
	actual := ReturnTDClassMainValues(tdtdTestString)
	expected := 2
	AssertEqualsInt (t, len(actual), expected)	
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


func TestTagTDWithJiraStringShouldResult10(t *testing.T) {
	//trTagValuesInArray := ReturnTRValues(myJiraPageTestString)
	actual := ReturnTDClassMainValues(tdtdJiraTestString)
	expected := 10
	AssertEqualsInt (t, len(actual), expected)	
}

func TestParseForTagStartEnd(t *testing.T){
	indexArray := parseForTagStartEnd(tdtdJiraTestString, "td", " class=\"main\"")
	AssertEqualsInt(t, len(indexArray), 10)
}

func TestReturnIndexPairsShouldReturn10(t *testing.T){
	indexArray := returnIndexPairs(tdtdJiraTestString, "td", " class=\"main\"")
	AssertEqualsInt(t, len(indexArray), 10)
}


func TestTRReturnIndexArrayForStartTagShouldReturn4(t *testing.T){
	//var tag string = "td"
	//var attributes string = " class=\"main\""
	indexArray := returnIndexArray(myJiraPageTestString, "(?is)<tr>")
	//indexArray := returnIndexArray(tdtdJiraTestString, "<"+tag+ attributes+">")
	AssertEqualsInt(t, len(indexArray), 4)
}

func TestFindAllStringSubmatchIndex(t *testing.T) {
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
	AssertEqualsInt(t, arrayLen, 4)
}

func TestTRReturnIndexArrayForEndTagShouldReturn4(t *testing.T){
	//var tag string = "td"
	//var attributes string = " class=\"main\""
	indexArray := returnIndexArray(tdtdJiraTestString, "</tr>")
	//indexArray := returnIndexArray(tdtdJiraTestString, "<"+tag+ attributes+">")
	AssertEqualsInt(t, len(indexArray), 0)
}


func TestReturnIndexArrayForStartTagShouldReturn10(t *testing.T){
	//var tag string = "td"
	//var attributes string = " class=\"main\""
	indexArray := returnIndexArray(tdtdJiraTestString, "<td class=\"main\">")
	//indexArray := returnIndexArray(tdtdJiraTestString, "<"+tag+ attributes+">")
	AssertEqualsInt(t, len(indexArray), 10)
}

func TestReturnIndexArrayForEndTagShouldReturn10(t *testing.T){
	//var tag string = "td"
	//var attributes string = " class=\"main\""
	indexArray := returnIndexArray(tdtdJiraTestString, "</td>")
	//indexArray := returnIndexArray(tdtdJiraTestString, "<"+tag+ attributes+">")
	AssertEqualsInt(t, len(indexArray), 10)
}

/*func returnValuesOfTag(stringToParse string, tag string, attributes string) [] string {
	indexArray := parseForTagStartEnd(stringToParse, tag, attributes)
	return trimTagsFromArray(indexArray, stringToParse, len(tag+attributes))
}*/

func TestTagTDJiraHtmlStreamMatches10(t *testing.T) {
	trTagValuesInArray := ReturnTRValues(myJiraPageTestString)
	actual := ReturnTDClassMainValues(trTagValuesInArray[0])
	expected := 10
	AssertEqualsInt (t, len(actual), expected)	
}


func TestReturnIndexShouldReturn2(t *testing.T) {
	index := returnIndexArray("<tr>...</tr>afkjajfladslkflkasd<tr>llll</tr>", "<tr>")

	AssertEqualsInt(t, len(index), 2)
}


func TestJiraStreamShouldReturn2Index(t *testing.T) {
	index := returnIndexArray(myJiraPageTestString, "<tr>")
	AssertEqualsInt(t, len(index), 4)
}
