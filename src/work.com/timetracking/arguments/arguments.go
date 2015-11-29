package arguments

import (
	"fmt"
	"strings"
	"time"
)

const defaultTeammemberFilepath string = "./teammembers.txt"
const defaultProjectsFilepath string = "./projects.csv"
const defaultConfigFilepath string = "./jira.yaml"

type TimetrackingArgs struct {
	countParsedArgs       int
	filePathToTeammembers string
	filePathToProjects    string
	filePathToConfig      string
	startDate             time.Time
	sprintStatistic       bool
	testing               bool
	run                   bool
	help                  bool
}

func (t *TimetrackingArgs) GetCountParsedArgs() int {
	return t.countParsedArgs
}

func (t *TimetrackingArgs) GetFilePathToTeammembers() string {
	return t.filePathToTeammembers
}

func (t *TimetrackingArgs) GetFilePathToProjects() string {
	return t.filePathToProjects
}

func (t *TimetrackingArgs) GetFilePathConfig() string {
	return t.filePathToConfig
}

func (t *TimetrackingArgs) GetEndDate() time.Time {
	if t.sprintStatistic {
		duration := time.Hour * 24 * 7
		endDate := t.startDate.Add(duration)
		return endDate
	}
	return time.Now()

}

func (t *TimetrackingArgs) IsTesting() bool {
	return t.testing
}

func (t *TimetrackingArgs) IsRunning() bool {
	return t.run
}

func (t *TimetrackingArgs) IsHelpCall() bool {
	return t.help
}

func (t *TimetrackingArgs) resetArguments() {
	t.countParsedArgs = 0
	t.filePathToConfig = defaultConfigFilepath
	t.filePathToProjects = defaultProjectsFilepath
	t.filePathToTeammembers = defaultTeammemberFilepath
	t.startDate = time.Date(0, time.January, 0, 0, 0, 0, 0, time.UTC)
	t.sprintStatistic = false
	t.testing = false
	t.run = false
	t.help = false
}

func ParseArguments(args []string) TimetrackingArgs {
	var timeTrackingArgs TimetrackingArgs
	timeTrackingArgs.resetArguments()
	timeTrackingArgs.parseAllArguments(args)
	return timeTrackingArgs
}

func (t *TimetrackingArgs) parseAllArguments(args []string) {
	t.countParsedArgs = 0
	for i := 1; i < len(args); i++ {
		arg := args[i]
		t.countParsedArgs++
		if t.isStringArg(arg) {
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

func (t *TimetrackingArgs) isStringArg(arg string) bool {
	return (strings.IndexRune(arg, '=') >= 0)
}

func (t *TimetrackingArgs) isBooleanArg(arg string) bool {
	return (strings.IndexRune(arg, '-') == 0)
}

func (t *TimetrackingArgs) isDateArg(arg string) bool {
	return (strings.IndexRune(arg, '?') >= 0)
}

func (t *TimetrackingArgs) parseStringArg(stringArg string) {
	index := strings.IndexRune(stringArg, '=')
	if index < 0 {
		return // this is not a string arg
	}
	t.setStringVariable(strings.ToLower(stringArg[0:index]), stringArg[index+1:])
}

func (t *TimetrackingArgs) setStringVariable(prefix string, value string) {
	switch prefix {
	case "tm":
		t.filePathToTeammembers = value
	case "prj":
		t.filePathToProjects = value
	default:
		fmt.Printf("Unknown String argument: %s\n", prefix)
	}
}

func (t *TimetrackingArgs) parseBooleanArg(boolArg string) {
	index := strings.IndexRune(boolArg, '-')
	if index != 0 {
		return // this is not a string arg
	}
	t.setBooleanVariable(strings.ToLower(boolArg))
}

func (t *TimetrackingArgs) setBooleanVariable(boolArg string) {
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

func (t *TimetrackingArgs) parseDateArg(dateArg string) {
	index := strings.IndexRune(dateArg, '?')
	if index <= 0 {
		return // this is not a string arg
	}
	t.setDateVariable(strings.ToLower(dateArg[0:index]), dateArg[index+1:])
}

func (t *TimetrackingArgs) setDateVariable(prefix, dateArg string) {
	switch prefix {
	case "start":
		t.startDate = t.parseIntoTimeObj(dateArg)
	default:
		fmt.Printf("Unknown Date argument: %s\n", prefix)
	}
}

func (t *TimetrackingArgs) parseIntoTimeObj(date string) time.Time {
	layout := t.createTimeLayout(date)
	var myTime time.Time
	myTime, err := time.Parse(layout, date)
	if err != nil {
		panic(err)
	}
	return myTime
}

func (t *TimetrackingArgs) createTimeLayout(date string) string {
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
  
<no other arguments till now>


-----------------------------------------------------------------------
What you need to run the program:
-----------------------------------------------------------------------

A file called projects.csv in the same folder as the executable in the following format:
PROJECT SHORTNAME, PRJ ID, <not yet used, keep empty>, START DATE in the format: DD.MM.YYYY (fill up with zero), END DATE (same format, if ZERO then actual is taken)
pjname, 001,, 16.11.2015,
pjname1, 002,, 16.11.2015,
pjname2, 003,, 16.11.2015,


A file called teammembers.txt, with all teammembers of the team. JIRA username in every line.  Last Line must be empty!
vorname.nachname
vorname.nachname
vorname.nachname
vorname.nachname
vorname.nachname
vorname.nachname
vorname.nachname
vorname.nachname
vorname.nachname
vorname.nachname
vorname.nachname
vorname.nachname

And finally a jira.yaml, with the connection details:
jiralogin:
    username: <yourJIRAUsername>
    password: <yourPWD>
jiraurl:
    url: "http://10.207.121.181/j/secure/"
    reportname: "ConfigureReport.jspa?"
    startdate: "startDateId="
    enddate: "&endDateId="
    prjid: "&projectId="
    query: "&jqlQueryId="
    selectedprjid: "&selectedProjectId="
    prefix: "&reportKey=com.synergyapps.plugins.jira.timepo-timesheet-plugin%3Aissues-report&Next=Next"`
