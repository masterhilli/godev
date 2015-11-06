package jiraRegEx

import (
	//"fmt"
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
// Test
func TestTRTRRegExParser(t *testing.T) {
	actual := ParseForTRTR(testString)
	expected := "---"
	AssertEqualsString (t, actual[0], expected)
}

func TestTRTRWithNumAndChars(t *testing.T) {
	actual := ParseForTRTR("<tr>BLA123BLA</tr>")
	expected := "BLA123BLA"
	AssertEqualsString (t, actual[0], expected)
}


func TestTRTRNovaluesIn(t *testing.T) {
	actual := ParseForTRTR("<tr></tr>")
	expected := ""
	AssertEqualsString (t, actual[0], expected)
}

func TestTRTRNoMatch(t *testing.T) {
	actual := ParseForTRTR("<tr</tr>")
	expected := ""
	AssertEqualsString (t, actual[0], expected)
}


func TestTRTRSpecialCharacters(t *testing.T) {
	actual := ParseForTRTR("<tr>"+sepcialChars+"</tr>")
	expected := sepcialChars
	AssertEqualsString (t, actual[0], expected)
}


func TestTRTRSearchInLongerString(t *testing.T) {
	actual := ParseForTRTR("scvjewoqfäiusaölkenskajflaöwkepori<tr>"+sepcialChars+"</tr>´dsfkjhjhsajkhfdahdskhaksdfkjas")
	expected := sepcialChars
	AssertEqualsString (t, actual[0], expected)
}

var myJiraPageTestString string = "<div style=\"width: 100%; overflow-x: auto\">\n"+
"    <table border=\"0\" cellpadding=\"3\" cellspacing=\"1\" class=\"main\">\n"+
"        <tbody><tr>\n"+
"            <td class=\"main\" colspan=\"2\">\n"+
"                1/Sep - 2/Nov\n"+
"            </td>\n"+
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
"</tr>"
/*

func TestTRTRCorrectJiraHTMLStreamReturns2Matches(t *testing.T) {
	actual := ParseForTRTR(myJiraPageTestString)
	expected := 2
	AssertEqualsInt (t, len(actual), expected)	
}*/



func TestTRTROwnHTMLStreamShouldReturn2matches(t *testing.T) {
	actual := ParseForTRTR("<tr>...</tr>afkjajfladslkflkasd<tr>llll</tr>")
	expected := 2
	AssertEqualsInt (t, len(actual), expected)	
}