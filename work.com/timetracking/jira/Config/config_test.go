package jira

import (
	. "../../helper"
	. "gopkg.in/check.v1"
	"gopkg.in/yaml.v2"
	"testing"
)

const pathToTestJiraConfigYaml string = "../../__testdata/jira.yaml"
const typicalProjectTemplate string = `
jiradata:
    username: abc
    password: xyz
    url: www.google.at
projects:
    SOLUT:
        project: SOLUT
        platform: Solution Tool
        excludeothers: true
        productowner: Hillbrand
    RMA:
        project: RMA
        platform: RMA
        productowner: Schiller
    RB:
        project: RBLET
        platform: RasenBall
        productowner: Kater
teammembers: [
    anton.Fressner,
    david.Huggl,
    Leonardo.Vastic,
    Marc.Soriano,
    martin.Hillbrand,
    richard.Friesenbichler,
    serhat.Kalowski,
    thomas.Eisenherz,
    thomas.Fridolino,
]
`
const contentForTestingYaml string = `
mylist: [test1, test2]
mypairlist: [
    [abc, def],
    [ghi, jkl],
    [mno, pqr, stv]
]
mymaplist:
    abc:
        - test
        - test2
    def:
        - again
        - again2
myobjectlist:
    mytest:
        test: my
        another: soon
    myOther:
        test: second
        another: so on
`

type YamlTestEngine struct{}

func TestYamlEngine(t *testing.T) {
	Suite(&YamlTestEngine{})
	TestingT(t)
}

type YamlWithList struct {
	Mylist       []string
	Mypairlist   [][]string
	Mymaplist    map[string][]string
	Myobjectlist map[string]Pair
}

type Pair struct {
	Test    string
	Another string
}

func (y *YamlTestEngine) TestReadJiraConfig(c *C) {
	config := Reader.Read(pathToTestJiraConfigYaml)
    jiraData := config.Jiradata

    c.Assert(jiraData.Username, Equals, "xyz")
	c.Assert(jiraData.Password, Equals, "abcdefgh")

	c.Assert(jiraData.Url, Equals, "http://10.99.11.333/j/secure/")
	c.Assert(jiraData.GetQuery(), Equals, "&jqlQueryId=")
}

func (y *YamlTestEngine) TestYamlUnmarshaler(c *C) {
	config := Reader.unmarshalToConfig([]byte(typicalProjectTemplate))
	c.Assert(config.Jiradata.Username, Equals, "abc")
	c.Assert(config.Jiradata.Password, Equals, "xyz")
	c.Assert(config.Jiradata.Url, Equals, "www.google.at")
	c.Assert(len(config.Projects), Equals, 3)
	c.Assert(config.Projects["RB"].Platform, Equals, "RasenBall")
    c.Assert(config.Projects["RB"].Excludeothers, Equals, false)
    c.Assert(config.Projects["SOLUT"].Excludeothers, Equals, true)
	c.Assert(len(config.Teammembers), Equals, 9)
	c.Assert(config.Teammembers[2], Equals, "Leonardo.Vastic")
}

func (this *YamlTestEngine) TestGetTeamMembersAsMapReturnsCount9(c *C) {
    config := Reader.unmarshalToConfig([]byte(typicalProjectTemplate))
    c.Assert(len(config.GetTeammembersAsMap()), Equals, 9)
}

func (this *YamlTestEngine) TestReadingInStringList(c *C) {
	yamlList := this.unmarshalList([]byte(contentForTestingYaml))
	c.Assert(len(yamlList.Mylist), Equals, 2)
	c.Assert(yamlList.Mylist[0], Equals, "test1")
	c.Assert(yamlList.Mylist[1], Equals, "test2")
	c.Assert(len(yamlList.Mypairlist), Equals, 3)
	c.Assert(len(yamlList.Mypairlist[0]), Equals, 2)
	c.Assert(len(yamlList.Mypairlist[1]), Equals, 2)
	c.Assert(len(yamlList.Mypairlist[2]), Equals, 3)
	c.Assert(len(yamlList.Mymaplist["abc"]), Equals, 2)
	c.Assert(yamlList.Mymaplist["abc"][1], Equals, "test2")
}

func (this *YamlTestEngine) unmarshalList(content []byte) YamlWithList {
	var yamlWithList YamlWithList
	err := yaml.Unmarshal(content, &yamlWithList)
	PanicOnError(err)
	return yamlWithList
}
