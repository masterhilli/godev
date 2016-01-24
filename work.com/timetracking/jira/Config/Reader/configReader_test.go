package reader

import (
    . "../../Config"
    . "gopkg.in/check.v1"
    "testing"
)

const pathToTestJiraConfigYaml string = "../../../__testdata/jira.yaml"
const pathToMorallyWrongTestFile string = "../../../__testdata/morallyNotOkFile.yaml"

type ConfigReaderTestEngine struct {
    config Config
}
func TestConfigReaderEngine(t *testing.T) {
    testEngine := ConfigReaderTestEngine{}
    testEngine.config = GetReader().unmarshalToConfig([]byte(typicalProjectTemplate))
    GetReader().setErrorChannel(mockForPanic)
    Suite(&testEngine)
    TestingT(t)
}

func (this *ConfigReaderTestEngine) TestReadJiraConfig(c *C) {
    config := GetReader().Read(pathToTestJiraConfigYaml)
    jiraData := config.Jiradata

    c.Assert(jiraData.Username, Equals, "xyz")
    c.Assert(jiraData.Password, Equals, "abcdefgh")
    c.Assert(jiraData.Url, Equals, "http://10.99.11.333/j/secure/")
}

func (this *ConfigReaderTestEngine) TestGetTeamMembersAsMapReturnsCount9(c *C) {
    config := GetReader().unmarshalToConfig([]byte(typicalProjectTemplate))
    c.Assert(len(config.GetTeammembersAsMap()), Equals, 9)
}

func (this *ConfigReaderTestEngine) TestUnmarshalFileWithOneTeamMemberResultsInPanic(c *C) {
    GetReader().Read(pathToMorallyWrongTestFile)
    c.Assert(mockForPanic.wasPanicCalled, Equals, true)
}


func (y *ConfigReaderTestEngine) TestYamlUnmarshaler(c *C) {
    c.Assert(y.config.Jiradata.Username, Equals, "abc")
    c.Assert(y.config.Jiradata.Password, Equals, "xyz")
    c.Assert(y.config.Jiradata.Url, Equals, "www.google.at")
    c.Assert(len(y.config.Projects), Equals, 3)
    c.Assert(y.config.Projects["RB"].Platform, Equals, "RasenBall")
    c.Assert(y.config.Projects["RB"].Excludeothers, Equals, false)
    c.Assert(y.config.Projects["SOLUT"].Excludeothers, Equals, true)
    c.Assert(len(y.config.Teammembers), Equals, 9)
    c.Assert(y.config.Teammembers[2], Equals, "Leonardo.Vastic")
}

func (this *ConfigReaderTestEngine) TestGetQueryForPlatformAndProject(c *C) {
    c.Assert(this.config.Projects["SOLUT"].GetQuery(nil), Equals, "%28Platform+%3D+%22Solution+Tool%22+OR+project+%3D+%22SOLUT%22%29")
}

func (this *ConfigReaderTestEngine) TestGetQueryForNoPlatformButProject(c *C) {
    c.Assert(this.config.Projects["RMA"].GetQuery(nil), Equals, "%28project+%3D+%22RMA%22%29")
}

func (this *ConfigReaderTestEngine) TestGetQueryForPlatformAndNoProject(c *C) {
    c.Assert(this.config.Projects["RB"].GetQuery(nil), Equals, "%28Platform+%3D+%22RasenBall%22%29")
}




/*******************************************************************
Panic MOCK
********************************************************************/
type MockPanicCalled struct {
    wasPanicCalled bool
}
func (this* MockPanicCalled) setPanicCalled(val bool) {
    this.wasPanicCalled = val
}
func (this* MockPanicCalled) Error (msg string) {
    this.wasPanicCalled = true
}
var mockForPanic *MockPanicCalled = new(MockPanicCalled)


/******************************************************************
Fixed paths that are needed for testing
*******************************************************************/
const typicalProjectTemplate string = `
jiradata:
    username: abc
    password: xyz
    url: www.google.at
projects:
    SOLUT:
        project: SOLUT
        platform: "Solution Tool"
        excludeothers: true
        productowner: Hillbrand
    RMA:
        project: RMA
        productowner: Schiller
    RB:
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