package jira

import (
	. "gopkg.in/check.v1"
	"testing"
)

type ConfigTestEngine struct {
	config Config
}

func TestYamlEngine(t *testing.T) {
	var testEngine ConfigTestEngine

	Suite(&testEngine)
	TestingT(t)
}

func (this *ConfigTestEngine) TestSmthg(c *C) {
	//TODO: add tests for Config Test Engine here
}


