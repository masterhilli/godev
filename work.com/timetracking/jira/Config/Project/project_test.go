package jira

import (
	. "gopkg.in/check.v1"
	"testing"
)

type ProjectTestEngine struct {
	project Project
}

func TestProjectEngine(t *testing.T) {
	var tstEngine ProjectTestEngine
	var prj Project
	prj.Excludeothers = true
	prj.Platform = ""
	prj.Productowner = "Me"
	prj.Project = "DAILCS"
	tstEngine.project = prj
	Suite(&tstEngine)
	TestingT(t)
}

func (this *ProjectTestEngine) TestGetQuery(c *C) {
	platforms := []string{"Test1", "Test2", "Test3"}

	query := this.project.GetQuery(platforms)

	c.Assert(query, Equals, "%28project+%3D+%22DAILCS%22%29+AND+Platform+not+in+%28%22Test1%22%2C%22Test2%22%2C%22Test3%22%29")

}
