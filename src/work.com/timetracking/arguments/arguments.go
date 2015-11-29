package arguments

import (
	"strings"
)

const defaultTeammemberFilepath string = "./teammembers.txt"
const defaultProjectsFilepath string = "./projects.csv"

type TimetrackingArgs struct {
	countParsedArgs       int
	filePathToTeammembers string
	filePathToProjects    string
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

func (t *TimetrackingArgs) ParseArguments(args []string) {
	t.parseAllArguments(args)
}

func (t *TimetrackingArgs) parseAllArguments(args []string) {
	t.countParsedArgs = 0
	for i := 1; i < len(args); i++ {
		t.countParsedArgs++
		if t.isStringArg(args[i]) {
			t.parseStringArg(args[i])
		} else if t.isBooleanArg(args[i]) {
			t.parseBooleanArg(args[i])
		}

	}
}

func (t *TimetrackingArgs) isStringArg(arg string) bool {
	return (strings.IndexRune(arg, '=') >= 0)
}

func (t *TimetrackingArgs) isBooleanArg(arg string) bool {
	return (strings.IndexRune(arg, '-') == 0)
}

func (t *TimetrackingArgs) parseStringArg(stringArg string) {
	index := strings.IndexRune(stringArg, '=')
	if index < 0 {
		return // this is not a string arg
	}
	t.setStringVariable(stringArg[0:index], stringArg[index+1:])
}

func (t *TimetrackingArgs) setStringVariable(prefix string, value string) {
	switch prefix {
	case "tm":
		t.filePathToTeammembers = value
	case "prj":
		t.filePathToProjects = value
	default:
		// nothing really todo
	}
}

func (t *TimetrackingArgs) parseBooleanArg(boolArg string) {
	index := strings.IndexRune(boolArg, '-')
	if index != 0 {
		return // this is not a string arg
	}
	t.setBooleanVariable(boolArg)
}

func (t *TimetrackingArgs) setBooleanVariable(boolArg string) {
	switch boolArg {
	case "-sprint":
		t.sprintStatistic = true
	default:
		// nothing really todo
	}
}
