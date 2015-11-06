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
var expected  string = "<tr>---</tr>"

// Test
func TestTRTRRegExParser(t *testing.T) {
	actual := ParseForTRTR(testString)
	AssertEquals (t, actual, expected)
}