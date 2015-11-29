package arguments

import (
	"fmt"
	"strings"
	"time"
)

const defaultTeammemberFilepath string = "./teammembers.txt"
const defaultProjectsFilepath string = "./projects.csv"

type TimetrackingArgs struct {
	countParsedArgs       int
	filePathToTeammembers string
	filePathToProjects    string
	startDate             time.Time
	sprintStatistic       bool
}

func (t *TimetrackingArgs) GetCountParsedArgs() int {
	return t.countParsedArgs
}

func (t *TimetrackingArgs) GetFilePathToTeammembers() string {
	retVal := t.filePathToTeammembers
	if len(retVal) == 0 {
		retVal = defaultTeammemberFilepath
	}
	return retVal
}

func (t *TimetrackingArgs) GetFilePathToProjects() string {
	retVal := t.filePathToProjects
	if len(retVal) == 0 {
		retVal = defaultProjectsFilepath
	}
	return retVal
}

func (t *TimetrackingArgs) GetEndDate() time.Time {
	if t.sprintStatistic {
		duration := time.Hour * 24 * 7
		endDate := t.startDate.Add(duration)
		return endDate
	}
	return time.Now()

}

func (t *TimetrackingArgs) clearArguments() {
	t.countParsedArgs = 0
	t.filePathToProjects = ""
	t.filePathToTeammembers = ""
	t.startDate = time.Date(0, time.January, 0, 0, 0, 0, 0, time.UTC)
	t.sprintStatistic = false
}

func ParseArguments(args []string) TimetrackingArgs {
	var timeTrackingArgs TimetrackingArgs
	timeTrackingArgs.clearArguments()
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
