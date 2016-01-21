package report

import (
	"fmt"
)



type Writer interface {
	Initialize(values []string, reportname string)
	PrintLine(teamName string, teamMembers string, projectname string, hours string, percent string)
	Close()
}

type CmdLineWriter struct {
	separator rune
}

var cmdLineWriter CmdLineWriter

func NewCmdLineWriter() *CmdLineWriter {
    return &cmdLineWriter
}

func (this *CmdLineWriter) Initialize(values []string, reportname string) {
	valLen := len(values)
	if valLen > 0 {
		this.separator = ([]rune(values[0]))[0]
	}
}

func (this *CmdLineWriter) PrintLine(teamName string, teamMembers string, projectname string, hours string, percent string) {
	fmt.Printf("%s%c%s%c%s%c%s%c%s%c\n", teamName, this.separator, teamMembers, this.separator, projectname, this.separator, hours, this.separator, percent, this.separator)
}

func (this *CmdLineWriter) Close() {

}
