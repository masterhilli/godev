package main

import (
	arguments "./arguments"
	. "./data"
	. "./helper"
	jiraConfig "./jira/Config"
	. "./jira/HTMLParser"
	jiraConnection "./jira/HtmlConnection"
	. "./report"
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	args := arguments.GetArguments()
	if args.IsHelpCall() || args.HasNoRunArgs() {
		return
	}

	var config jiraConfig.Config = jiraConfig.Reader.Read(args.GetFilePathConfig())
	config.JiraLogin.Password = SetEmptyPasswordOverConsoleInput(config.JiraLogin.Password)

	var jc jiraConnection.HtmlConnector = jiraConnection.NewHtmlConnector(config)
	var tm map[string]bool = ReadTeammembers(args.GetFilePathToTeammembers())
	var pi TimeTrackingReport

	pi.Initialize(args.GetFilePathToProjects())

	if args.IsTesting() {
		for k := range tm {
			fmt.Printf("Read in:\t%s\n", k)
		}
	}

	timeStart := time.Now()
	var retChannel chan ProjectReportSetting = make(chan ProjectReportSetting)
	//var nameTimePairs map[string]HTMLParser = make(map[string]HTMLParser)
	for i := range pi.GetAllSettings() {
		go RetrieveNameTimePairPerProject(retChannel, pi.GetEntry(i), jc) // I really do not like this one! you can not add the pointer of an element of the map as param to a method :(
	}

	for j := 0; j < pi.GetSettingsLen(); j++ {
		prjSetting := <-retChannel
		pi.SetEntry(prjSetting)
	}
	close(retChannel)

	pi.SetTeamMembers(tm)

	timeStop := time.Now()
	fmt.Printf("-->All projects retrieved in %v\n", timeStop.Sub(timeStart))
	var reporter ReporterInterface
	pi.Finish()
	reporter = GetReporter(';')
	reporter.ExportReport(pi)

}

func SetEmptyPasswordOverConsoleInput(pwd string) string {
	if len(pwd) == 0 {

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter password: ")
		pwd, _ = reader.ReadString('\n')
		pwd = strings.TrimSpace(pwd)
	}
	return pwd
}

func ReadTeammembers(file string) map[string]bool {
	var reader *bufio.Reader
	reader = GetBufferIOReader(file)
	return ReadEachLineAndAddTeamMemberToMap(reader)
}

func GetBufferIOReader(file string) *bufio.Reader {
	filename, errAbs := filepath.Abs(file)
	PanicOnError(errAbs)

	f, errOpen := os.Open(filename)
	PanicOnError(errOpen)

	return bufio.NewReader(f)
}

func ReadEachLineAndAddTeamMemberToMap(reader *bufio.Reader) map[string]bool {
	var teammembers map[string]bool = make(map[string]bool)
	line, errLine := reader.ReadString('\n')

	for errLine == nil {
		line = strings.TrimSpace(line)
		if len(line) > 0 {
			teammembers[line] = true
		}
		line, errLine = reader.ReadString('\n')
	}
	return teammembers
}
