package arguments

import (
	"os"
	"strconv"
	"strings"
	"time"
)

// constant values
const defaultConfigFilepath string = "./__configFiles/jira.yaml"
const testConfigFilepath string = "./__testdata/jira.yaml"

const argConfig string = "config"
const argBSprint string = "-sprint"
const argBTesting string = "-t"
const argBRunning string = "-r"
const argBHelp	string = "--help"
const argDStart string = "start"
const argIReport string = "report"
const noArgumentsMessage string = "If you do not know how to use this program please call with \"--help\""

const wrongArgMessage string = "Unknown argument: "


var args TimeTrackingArgs
var isInitialized bool

type TimeTrackingArgs struct {
	countParsedArgs  int
	reportId         int
	filePathToConfig string
	startDate        time.Time
	sprintStatistic  bool
	testing          bool
	run              bool
	help             bool
	out 			 ArgOut
}

var uiConsole *Console = new(Console)

// Singleton creator
func GetArguments() TimeTrackingArgs {
	if isInitialized == false {
		var timeTrackingArgs TimeTrackingArgs
		timeTrackingArgs.Initialize(true, uiConsole)
		args = timeTrackingArgs
		isInitialized = true
	}
	return args
}

// Initialize TimetrackingArgs
func (this *TimeTrackingArgs) Initialize(parseOSArgs bool, uiOut ArgOut) {
	this.resetArguments()
	this.out = uiOut
	if parseOSArgs {
		this.parseAllArguments(os.Args)
	}
}

// Getter and setter
func (this *TimeTrackingArgs) GetReporterId() int {
	return this.reportId
}

func (this *TimeTrackingArgs) SetReporterId(reportid int) {
	this.reportId = reportid
}

func (this *TimeTrackingArgs) GetCountParsedArgs() int {
	return this.countParsedArgs
}

func (this *TimeTrackingArgs) GetFilePathConfig() string {
	if this.testing {
		return testConfigFilepath
	}
	return this.filePathToConfig
}

func (this *TimeTrackingArgs) GetEndDate() time.Time {
	if this.sprintStatistic {
		duration := time.Hour * 24 * 7
		endDate := this.startDate.Add(duration)
		return endDate
	}
	return time.Now()
}

func (this *TimeTrackingArgs) IsTesting() bool {
	return this.testing
}

func (this *TimeTrackingArgs) IsRunning() bool {
	return this.run
}

func (this *TimeTrackingArgs) IsHelpCall() bool {
	return this.help
}

func (this *TimeTrackingArgs) HasNoRunArgs() bool {
	return !this.IsHelpCall() && !this.IsRunning() && !this.IsTesting()
}

func (this *TimeTrackingArgs) resetArguments() {
	this.countParsedArgs = 0
	this.SetReporterId(0)
	this.filePathToConfig = defaultConfigFilepath
	this.startDate = time.Date(0, time.January, 0, 0, 0, 0, 0, time.UTC)
	this.sprintStatistic = false
	this.testing = false
	this.run = false
	this.help = false
}

func (this *TimeTrackingArgs) parseAllArguments(args []string) {
	this.countParsedArgs = 0
	for i := 1; i < len(args); i++ {
		arg := args[i]
		this.countParsedArgs++
		if this.isIntArg(arg) {
			this.parseIntArg(arg)
		} else if this.isStringArg(arg) {
			this.parseStringArg(arg)
		} else if this.isBooleanArg(arg) {
			this.parseBooleanArg(arg)
		} else if this.isDateArg(arg) {
			this.parseDateArg(arg)
		} else {
			this.printWrongArgMessageToUI(arg)
		}
	}
	if this.HasNoRunArgs() {
		this.printMessageToUI(noArgumentsMessage)
	}
}

func (this *TimeTrackingArgs) isStringArg(arg string) bool {
	return (strings.IndexRune(arg, '=') > 0)
}

func (this *TimeTrackingArgs) isBooleanArg(arg string) bool {
	return (strings.IndexRune(arg, '-') == 0)
}

func (this *TimeTrackingArgs) isDateArg(arg string) bool {
	return (strings.IndexRune(arg, '?') > 0)
}

func (this *TimeTrackingArgs) isIntArg(arg string) bool {
	return (strings.IndexRune(arg, '#') > 0)
}

func (this *TimeTrackingArgs) parseStringArg(stringArg string) {
	index := strings.IndexRune(stringArg, '=')
	if index <= 0 {
		this.printFailedParsedArg("String", stringArg)
		return // this is not a string arg
	}
	this.setStringVariable(strings.ToLower(stringArg[0:index]), stringArg[index+1:])
}

func (this *TimeTrackingArgs) setStringVariable(prefix string, value string) {
	switch prefix {
	case argConfig:
		this.filePathToConfig = value
	default:
		this.printWrongArgMessageToUI(prefix)
	}
}

func (this *TimeTrackingArgs) parseBooleanArg(boolArg string) {
	index := strings.IndexRune(boolArg, '-')
	if index != 0 {
		this.printFailedParsedArg("Bool", boolArg)
		return
	}
	this.setBooleanVariable(strings.ToLower(boolArg))
}




func (this *TimeTrackingArgs) setBooleanVariable(boolArg string) {
	switch boolArg {
	case argBSprint:
		this.sprintStatistic = true
	case argBTesting:
		this.testing = true
	case argBRunning:
		this.run = true
	case argBHelp:
		this.printMessageToUI(helpContent)
		this.help = true
	default:
		this.printWrongArgMessageToUI(boolArg)
	}
}

func (this *TimeTrackingArgs) parseDateArg(dateArg string) {
	index := strings.IndexRune(dateArg, '?')
	if index <= 0 {
		this.printFailedParsedArg("Date", dateArg)
		return
	}
	this.setDateVariable(strings.ToLower(dateArg[0:index]), dateArg[index+1:])
}

func (this *TimeTrackingArgs) setDateVariable(prefix, dateArg string) {
	switch prefix {
	case argDStart:
		this.startDate = this.parseIntoTimeObj(dateArg)
	default:
		this.printWrongArgMessageToUI( prefix)
	}
}

func (this *TimeTrackingArgs) parseIntArg(intArg string) {
	index := strings.IndexRune(intArg, '#')
	if index <= 0 {
		this.printFailedParsedArg("Number", intArg)
		return
	}
	this.setIntVariable(strings.ToLower(intArg[0:index]), intArg[index+1:])
}

func (this *TimeTrackingArgs) setIntVariable(prefix, intArg string) {
	switch prefix {
	case argIReport:
		report, err := strconv.Atoi(intArg)
		if err != nil {
			this.printParsingError(intArg, err)
			return
		}
		this.SetReporterId(report)
	default:
		this.printWrongArgMessageToUI( prefix)
	}
}

func (this *TimeTrackingArgs) parseIntoTimeObj(date string) time.Time {
	layout := this.createTimeLayout(date)
	var myTime time.Time
	myTime, err := time.Parse(layout, date)
	if err != nil {
		this.printParsingError(date, err)
		return time.Now()  // is this good? we do not know the outcome here!
	}
	return myTime
}

func (this *TimeTrackingArgs) createTimeLayout(date string) string {
	index := strings.IndexRune(date, '.')
	var layout string = ""
	if index == 1 {
		layout = layout + "2."
	} else {
		layout = layout + "02."
	}

	index = strings.IndexRune(date[index+1:], '.')
	if index == 1 {
		layout = layout + "1."
	} else {
		layout = layout + "01."
	}

	layout = layout + "2006"
	return layout

}

// printParsingError
func (this TimeTrackingArgs) printParsingError(argument string, err error) {
	this.out.Println("Error parsing argument:  %s \n paniced error: \n%s", argument, err.Error())
}


func (this TimeTrackingArgs) printFailedParsedArg(argType string, argument string) {
	this.out.Println("Unknown Argument: %s (%s)", argument, argType)
}

func (this TimeTrackingArgs) printWrongArgMessageToUI(argument string) {
	this.out.Println("%s", argument, "")
}

func (this TimeTrackingArgs) printMessageToUI(argument string) {
	this.out.Println("%s", argument, "")
}

const helpContent string = `Possible Arguments: 
--help calls this help
-r    runs the program
options:
#report=0
  
<no other arguments till now>


-----------------------------------------------------------------------
What you need to run the program:
-----------------------------------------------------------------------

A jira.yaml, with the connection details, Teammembers and the project you need: (for further details ask me :) )
jiradata:
    username: <yourJIRAUsername>
    password: <yourPWD>
    url: "http://10.207.121.181/j/secure/"
projects:
    DAILCS:
        project: DAILCS
        productowner: Priesching
        excludeothers: true ----------> This can only set once, to exclude all others
IMCD: --------------------------------> Projektname wie er dann im EXCEL steht
    project: IMINT -------------------> Shortcut zur Identifizierung des JIRA Projektes
    platform: IMCD -------------------> Wert der unter Plattform steht.
    productowner: Priesching ---------> ProductOwner für diese Plattform
teammembers: [
    anton.fressner,
    david.huggl,
    leonardo.vastic,
    marc.Soriano,
    martin.hillbrand,
    richard.Friesenbichler,
    serhat.Kalowski,
    thomas.Eisenherz,
    thomas.Fridolino
]
`
