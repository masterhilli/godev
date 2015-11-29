package arguments

import (
	"strings"
)

type TimetrackingArgs struct {
	countParsedArgs       int
	filePathToTeammembers string
}

func (t *TimetrackingArgs) GetCountParsedArgs() int {
	return t.countParsedArgs
}

func (t *TimetrackingArgs) GetFilePathToTeammembers() string {
	return t.filePathToTeammembers
}

func (t *TimetrackingArgs) ParseArguments(args []string) {
	t.parseAllArguments(args)
}

func (t *TimetrackingArgs) parseAllArguments(args []string) {
	t.countParsedArgs = 0
	for i := 1; i < len(args); i++ {
		t.countParsedArgs++
		t.parseStringArg(args[i])
	}
}

func (t *TimetrackingArgs) parseStringArg(stringArg string) {
	index := strings.IndexRune(stringArg, '=')
	if index < 0 {
		return // this is not a string arg
	}
	t.filePathToTeammembers = stringArg[index+1:]
}
