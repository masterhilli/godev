package arguments

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const defaultTeamMemberFilepath string = "./configFiles/teammembers.txt"
const defaultProjectsFilepath string = "./configFiles/projects.csv"
const defaultConfigFilepath string = "./configFiles/jira.yaml"
const testTeamMemberFilepath string = "./configFiles/teammembers_test.txt"
const testProjectsFilepath string = "./configFiles/projects_test.csv"

type TimeTrackingArgs struct {
	countParsedArgs       int
	filePathToTeamMembers string
	filePathToProjects    string
	filePathToConfig      string
	startDate             time.Time
	sprintStatistic       bool
	testing               bool
	run                   bool
	help                  bool
}

func (t *TimeTrackingArgs) GetCountParsedArgs() int {
	return t.countParsedArgs
}

func (t *TimeTrackingArgs) GetFilePathToTeammembers() string {
	if t.testing {
		return testTeamMemberFilepath
	}
	return t.filePathToTeamMembers
}

func (t *TimeTrackingArgs) GetFilePathToProjects() string {
	if t.testing {
		return testProjectsFilepath
	}
	return t.filePathToProjects
}

func (t *TimeTrackingArgs) GetFilePathConfig() string {
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
	t.filePathToConfig = defaultConfigFilepath
	t.filePathToProjects = defaultProjectsFilepath
	t.filePathToTeamMembers = defaultTeamMemberFilepath
	t.startDate = time.Date(0, time.January, 0, 0, 0, 0, 0, time.UTC)
	t.sprintStatistic = false
	t.testing = false
	t.run = false
	t.help = false
}

func NewArguments() TimeTrackingArgs {
	var timeTrackingArgs TimeTrackingArgs
	timeTrackingArgs.resetArguments()
	timeTrackingArgs.parseAllArguments(os.Args)
	if timeTrackingArgs.HasNoRunArgs() {
		fmt.Println("If you do not know how to use this program please call with \"--help\"")
	}
	return timeTrackingArgs
}

func (t *TimeTrackingArgs) parseAllArguments(args []string) {
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

func (t *TimeTrackingArgs) isStringArg(arg string) bool {
	return (strings.IndexRune(arg, '=') >= 0)
}

func (t *TimeTrackingArgs) isBooleanArg(arg string) bool {
	return (strings.IndexRune(arg, '-') == 0)
}

func (t *TimeTrackingArgs) isDateArg(arg string) bool {
	return (strings.IndexRune(arg, '?') >= 0)
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
	case "tm":
		t.filePathToTeamMembers = value
	case "prj":
		t.filePathToProjects = value
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
