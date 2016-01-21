package arguments

import (
	. "gopkg.in/check.v1"
	"testing"
	"time"
    "strings"
)

// 30 test cases for a test coverage of 92.3% :( I am to tired to work on ;)

var _ = time.ANSIC

const executableArg string = "test.exe"
const fileNameToCompare string = "myTestFilename"
//working arguments:
const stringConfigArg string = "config="+fileNameToCompare

//not found arguments:
const stringArgument string = "tm=TestFilename"
const boolArgSprint string = "-sprint"
const boolArgTest string = "-t"
const boolArgRun string = "-r"
const dateArgStart string =  "start?5.1.2015"
const intArgReport string = "report#7"
const unidentifyableStr string = "someText"
const argBAny string = "-any"

type ArgumentTestEngine struct {
	ta TimeTrackingArgs
}

func TestRegisterArgumentTestEngine(t *testing.T) {
    var argTester ArgumentTestEngine
    argTester.ta.Initialize(false)
    ArgOutGetter.EnableTestMockUp()
	Suite(&ArgumentTestEngine{})
	TestingT(t)
}

// we ignore that one because depending on the IDE we use, the args are different
/*func (ate *ArgumentTestEngine) IgnoreSettingArgumentsWithOnly1Argument(c *C) {
	var ta TimeTrackingArgs = GetArguments()
	c.Assert(ta.GetCountParsedArgs(), Equals, 1)
}*/

// test initialization of the object
func (this *ArgumentTestEngine) TestInitialValueForStartDateResultsInZeroDate(c *C) {
    this.ta.Initialize(false)
    this.ta.parseAllArguments([]string{executableArg})
    t := time.Date(0, time.January, 0, 0, 0, 0, 0, time.UTC)
    c.Assert(this.ta.startDate, Equals, t)
}

func (this *ArgumentTestEngine) TestInitialEndDateResultsInZeroDate(c *C) {
    this.ta.Initialize(false)
    this.ta.parseAllArguments([]string{executableArg})
    c.Assert(this.ta.GetEndDate(), Equals, time.Now())
}


// no real argument type:
func (this *ArgumentTestEngine) TestParseStringThatIsNotInAnyFormat(c *C) {
    this.ta.Initialize(false)
    this.ta.parseAllArguments([]string{executableArg, boolArgRun, unidentifyableStr})
    c.Assert(ArgOutGetter.GetLastArgument(), Equals, unidentifyableStr)
}

func (this *ArgumentTestEngine) TestParseWithHelpArgResultsInHelpString(c *C) {
    this.ta.Initialize(false)
    this.ta.parseAllArguments([]string{executableArg, "--help"})
    c.Assert(ArgOutGetter.GetLastArgument(), Equals, helpContent)
}

// no argument!!!

func (this *ArgumentTestEngine) TestNoArgmuentResultsInInformationToUI(c *C) {
    this.ta.Initialize(false)
    this.ta.parseAllArguments([]string{})
    c.Assert(ArgOutGetter.GetLastArgument(), Equals, noArgumentsMessage)
}

// test strings:
func (ate *ArgumentTestEngine) TestIsStringArgument(c *C) {
    retVal := ate.ta.isStringArg("config=SomthingElse")
    c.Assert(retVal, Equals, true)
}

func (ate *ArgumentTestEngine) TestIsNotStringArgument(c *C) {
    retVal := ate.ta.isStringArg("-blaSomthingElse")
    c.Assert(retVal, Equals, false)
}

func (this *ArgumentTestEngine) TestParseStringArgumentThatIsConfigFile(c *C) {
    this.ta.Initialize(false)
    this.ta.parseAllArguments([]string{executableArg, stringConfigArg})
	c.Assert(this.ta.GetFilePathConfig(), Equals, fileNameToCompare)
}

func (this *ArgumentTestEngine) TestParseStringArgumentThatIsNotPartOfArguments(c *C) {
    this.ta.Initialize(false)
    this.ta.parseAllArguments([]string{executableArg, boolArgRun, stringArgument})
	c.Assert(ArgOutGetter.GetLastArgument(), Equals, "tm") //retrieves a fallback path to a config
}

//this gives no more test coverage, but at least it secures that argument names are not parsed case sensitive
func (this *ArgumentTestEngine) TestParseConfigStringAsLowerCaseMustReturnConfigFilePath(c *C) {
	this.ta.Initialize(false)
	this.ta.parseAllArguments([]string{executableArg, strings.ToLower(stringConfigArg)})
	c.Assert(this.ta.GetFilePathConfig(), Equals, strings.ToLower(fileNameToCompare))
}


//*****************************************
//**** Testing boolean arguments ******
//*****************************************
func (this *ArgumentTestEngine) TestParseBooleanArgSprintResultsInSprintIsSet(c *C) {
	this.ta.Initialize(false)
	this.ta.parseAllArguments([]string{executableArg, stringConfigArg, boolArgSprint})
	c.Assert(this.ta.sprintStatistic, Equals, true)
}

func (this *ArgumentTestEngine) TestParseBooleanArgRunResultsInRunIsSet(c *C) {
    this.ta.Initialize(false)
    this.ta.parseAllArguments([]string{executableArg, stringConfigArg, boolArgRun})
    c.Assert(this.ta.run, Equals, true)
}

func (this *ArgumentTestEngine) TestParseBooleanArgTestResultsInTestIsSet(c *C) {
    this.ta.Initialize(false)
    this.ta.parseAllArguments([]string{executableArg, stringConfigArg, boolArgTest})
    c.Assert(this.ta.IsTesting(), Equals, true)
}

func (this *ArgumentTestEngine) TestParseBooleanArgAnyResultsInWrongArgMessage(c *C) {
    this.ta.Initialize(false)
    this.ta.parseAllArguments([]string{executableArg, boolArgRun, stringConfigArg, argBAny})
    c.Assert(ArgOutGetter.GetLastArgument(), Equals, argBAny)
}

//*************************************
//******* TESTING DATE ARGUMENTS
//*************************************
func (this *ArgumentTestEngine) TestParseDateArgStartResultsInStarDateIsSet(c *C) {
	this.ta.Initialize(false)
	this.ta.parseAllArguments([]string{executableArg, stringConfigArg, stringConfigArg,dateArgStart, boolArgSprint})
	t := time.Date(2015, time.January, 5, 0, 0, 0, 0, time.UTC)
	c.Assert(this.ta.startDate, Equals, t)
}

func (this *ArgumentTestEngine) TestParseDateArgAnyResultsInStarDateIsSet(c *C) {
    this.ta.Initialize(false)
    this.ta.parseAllArguments([]string{executableArg, boolArgRun, stringConfigArg, "any?1.1.2015", boolArgSprint})
    c.Assert(ArgOutGetter.GetLastArgument(), Equals, "any")
}


func (this *ArgumentTestEngine) TestParseStartAndSprintArgResultsInEndDateCalculated(c *C) {
	this.ta.parseAllArguments([]string{executableArg, stringConfigArg, dateArgStart, boolArgSprint})
	t := time.Date(2015, time.January, 12, 0, 0, 0, 0, time.UTC)
	c.Assert(this.ta.GetEndDate(), Equals, t)
}

func (this *ArgumentTestEngine) TestParseIntArgReportIdResultsReportIdIsSetTo7(c *C) {
	this.ta.Initialize(false)
	this.ta.parseAllArguments([]string{executableArg, intArgReport})
	c.Assert(this.ta.GetReporterId(), Equals, 7)
}

func (this *ArgumentTestEngine) TestParseIntArgAnyResultsInWrongArgumentMessage(c *C) {
    this.ta.Initialize(false)
    this.ta.parseAllArguments([]string{executableArg, boolArgRun, "any#1"})
    c.Assert(ArgOutGetter.GetLastArgument(), Equals, "any")
}


func (this *ArgumentTestEngine) TestParseAllExistingArguments(c *C) {
	this.ta = GetArguments()
    ArgOutGetter.DisableTestMockUp()
	this.ta.parseAllArguments([]string{ executableArg,  stringConfigArg,
                                        boolArgRun,     boolArgSprint,  boolArgTest,
                                        dateArgStart,   intArgReport})

	c.Assert(this.ta.GetCountParsedArgs(), Equals, 6)
    ArgOutGetter.EnableTestMockUp()
}



// *****************************************************
// Testing helper functions that are internally needed
// Those tests help on debugging but not on tst coverage
// *****************************************************
func (this *ArgumentTestEngine) TestCreateTimeLayoutAll2digits(c *C) {
    layout := this.ta.createTimeLayout("15.01.2015")
    c.Assert(layout, Equals, "02.01.2006")
}

func (this *ArgumentTestEngine) TestCreateTimeLayoutDay1digit(c *C) {
    layout := this.ta.createTimeLayout("1.01.2015")
    c.Assert(layout, Equals, "2.01.2006")
}

func (this *ArgumentTestEngine) TestCreateTimeLayoutMonth1digit(c *C) {
    layout := this.ta.createTimeLayout("15.1.2015")
    c.Assert(layout, Equals, "02.1.2006")
}

func (this *ArgumentTestEngine) TestCreateTimeLayoutAll1digit(c *C) {
    layout := this.ta.createTimeLayout("1.1.2015")
    c.Assert(layout, Equals, "2.1.2006")
}


func (ate *ArgumentTestEngine) TestParseIntoTimeObjAll2digits(c *C) {
    parsedTime := ate.ta.parseIntoTimeObj("15.01.2015")
    t := time.Date(2015, time.January, 15, 0, 0, 0, 0, time.UTC)
    c.Assert(parsedTime, Equals, t)
}

func (ate *ArgumentTestEngine) TestParseIntoTimeObjDay1digit(c *C) {
    parsedTime := ate.ta.parseIntoTimeObj("1.11.2015")
    t := time.Date(2015, time.November, 1, 0, 0, 0, 0, time.UTC)
    c.Assert(parsedTime, Equals, t)
}

func (ate *ArgumentTestEngine) TestParseIntoTimeObjMonth1digit(c *C) {
    parsedTime := ate.ta.parseIntoTimeObj("15.1.2015")
    t := time.Date(2015, time.January, 15, 0, 0, 0, 0, time.UTC)
    c.Assert(parsedTime, Equals, t)
}

func (ate *ArgumentTestEngine) TestParseIntoTimeObjAll1digit(c *C) {
    parsedTime := ate.ta.parseIntoTimeObj("1.1.2015")
    t := time.Date(2015, time.January, 1, 0, 0, 0, 0, time.UTC)
    c.Assert(parsedTime, Equals, t)
}

func (this *ArgumentTestEngine) TestIsBooleanArgument(c *C) {
    retVal := this.ta.isBooleanArg("-bla")
    c.Assert(retVal, Equals, true)
}

func (this *ArgumentTestEngine) TestIsNotBooleanArgument(c *C) {
    retVal := this.ta.isBooleanArg("b-laSomthingElse")
    c.Assert(retVal, Equals, false)
}