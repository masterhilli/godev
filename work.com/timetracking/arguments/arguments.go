package arguments

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const defaultConfigFilepath string = "./__configFiles/jira.yaml"
const testConfigFilepath string = "./__testdata/jira.yaml"

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
}

func GetArguments() TimeTrackingArgs {
	if isInitialized == false {
		var timeTrackingArgs TimeTrackingArgs
		timeTrackingArgs.resetArguments()
		timeTrackingArgs.parseAllArguments(os.Args)
		if timeTrackingArgs.HasNoRunArgs() {
			fmt.Println("If you do not know how to use this program please call with \"--help\"")
		}
		args = timeTrackingArgs
		isInitialized = true
	}
	return args
}

func (this *TimeTrackingArgs) GetReporterId() int {
	return this.reportId
}

func (this *TimeTrackingArgs) SetReporterId(reportid int) {
	this.reportId = reportid
}

func (t *TimeTrackingArgs) GetCountParsedArgs() int {
	return t.countParsedArgs
}

func (t *TimeTrackingArgs) GetFilePathConfig() string {
	if t.testing {
		return testConfigFilepath
	}
	return t.filePathToConfig
}

func (t *TimeTrackingArgs) GetEndDate() time.Time {
	if t.sprintStatistic {
		duration := time.Hour * 24 * 7
		endDate := t.startDate.Add(duration)
		return endDate
	}
	return time.Now()
}

func (t *TimeTrackingArgs) IsTesting() bool {
	return t.testing
}

func (t *TimeTrackingArgs) IsRunning() bool {
	return t.run
}

func (t *TimeTrackingArgs) IsHelpCall() bool {
	return t.help
}

func (t *TimeTrackingArgs) HasNoRunArgs() bool {
	return !t.IsHelpCall() && !t.IsRunning() && !t.IsTesting()
}

func (t *TimeTrackingArgs) resetArguments() {
	t.countParsedArgs = 0
	t.SetReporterId(0)
	t.filePathToConfig = defaultConfigFilepath
	t.startDate = time.Date(0, time.January, 0, 0, 0, 0, 0, time.UTC)
	t.sprintStatistic = false
	t.testing = false
	t.run = false
	t.help = false
}

func (t *TimeTrackingArgs) parseAllArguments(args []string) {
	t.countParsedArgs = 0

	for i := 1; i < len(args); i++ {
		arg := args[i]
		t.countParsedArgs++
		if t.isIntArg(arg) {
			t.parseIntArg(arg)
		} else if t.isStringArg(arg) {
			t.parseStringArg(arg)
		} else if t.isBooleanArg(arg) {
			t.parseBooleanArg(arg)
		} else if t.isDateArg(arg) {
			t.parseDateArg(arg)
		} else {
			fmt.Printf("Unknown argument: %d", arg)
		}

	}
}

func (t *TimeTrackingArgs) isStringArg(arg string) bool {
	return (strings.IndexRune(arg, '=') >= 0)
}

func (t *TimeTrackingArgs) isBooleanArg(arg string) bool {
	return (strings.IndexRune(arg, '-') == 0)
}

func (t *TimeTrackingArgs) isDateArg(arg string) bool {
	return (strings.IndexRune(arg, '?') >= 0)
}

func (this *TimeTrackingArgs) isIntArg(arg string) bool {
	return (strings.IndexRune(arg, '#') >= 0)
}

func (t *TimeTrackingArgs) parseStringArg(stringArg string) {
	index := strings.IndexRune(stringArg, '=')
	if index < 0 {
		return // this is not a string arg
	}
	t.setStringVariable(strings.ToLower(stringArg[0:index]), stringArg[index+1:])
}

func (t *TimeTrackingArgs) setStringVariable(prefix string, value string) {
	switch prefix {
	case "config":
		t.filePathToConfig = value
	default:
		fmt.Printf("Unknown String argument: %s\n", prefix)
	}
}

func (t *TimeTrackingArgs) parseBooleanArg(boolArg string) {
	index := strings.IndexRune(boolArg, '-')
	if index != 0 {
		return // this is not a string arg
	}
	t.setBooleanVariable(strings.ToLower(boolArg))
}

func (t *TimeTrackingArgs) setBooleanVariable(boolArg string) {
	switch boolArg {
	case "-sprint":
		t.sprintStatistic = true
	case "-t":
		t.testing = true
	case "-r":
		t.run = true
	case "--help":
		fmt.Printf("%s\n", helpContent)
		t.help = true
	default:
		fmt.Printf("Unknown Boolean argument: %s\n", boolArg)
	}
}

func (t *TimeTrackingArgs) parseDateArg(dateArg string) {
	index := strings.IndexRune(dateArg, '?')
	if index <= 0 {
		return // this is not a string arg
	}
	t.setDateVariable(strings.ToLower(dateArg[0:index]), dateArg[index+1:])
}

func (t *TimeTrackingArgs) setDateVariable(prefix, dateArg string) {
	switch prefix {
	case "start":
		t.startDate = t.parseIntoTimeObj(dateArg)
	default:
		fmt.Printf("Unknown Date argument: %s\n", prefix)
	}
}

func (t *TimeTrackingArgs) parseIntArg(intArg string) {
	index := strings.IndexRune(intArg, '#')
	if index <= 0 {
		return // this is not a int arg
	}
	t.setIntVariable(strings.ToLower(intArg[0:index]), intArg[index+1:])
}

func (this *TimeTrackingArgs) setIntVariable(prefix, intArg string) {
	switch prefix {
	case "report":
		report, err := strconv.Atoi(intArg)
		if err != nil {
			panic(err)
		}
		this.SetReporterId(report)
	default:
		fmt.Printf("Unknown Int argument: %s\n", prefix)
	}
}

func (t *TimeTrackingArgs) parseIntoTimeObj(date string) time.Time {
	layout := t.createTimeLayout(date)
	var myTime time.Time
	myTime, err := time.Parse(layout, date)
	if err != nil {
		panic(err)
	}
	return myTime
}

func (t *TimeTrackingArgs) createTimeLayout(date string) string {
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
        excludeothers: true
    SOLUT:
        project: SOLUT
        productowner: Hillbrand
    TAIR:
        project: TAIR
        productowner: Priesching
    RMA:
        project: RMA
        productowner: HILLBRAND
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
