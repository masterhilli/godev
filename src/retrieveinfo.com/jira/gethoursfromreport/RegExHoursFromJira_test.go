package jiraRegEx

import (
	//"fmt"
	"testing"
)

func AssertEquals (t *testing.T, actual string, expected string) {
	if actual != expected {
		t.Fatalf("\nexpected: %s\nactual:  %s", expected, actual)
	}
}

var testString string = "<tr>---</tr>"
// Test
func TestTRTRRegExParser(t *testing.T) {
	actual := ParseForTRTR(testString)
	expected := "---"
	AssertEquals (t, actual, expected)
}

func TestTRTRWithNumAndChars(t *testing.T) {
	actual := ParseForTRTR("<tr>BLA123BLA</tr>")
	expected := "BLA123BLA"
	AssertEquals (t, actual, expected)
}


func TestTRTRNovaluesIn(t *testing.T) {
	actual := ParseForTRTR("<tr></tr>")
	expected := ""
	AssertEquals (t, actual, expected)
}

func TestTRTRNoMatch(t *testing.T) {
	actual := ParseForTRTR("<tr</tr>")
	expected := ""
	AssertEquals (t, actual, expected)
}