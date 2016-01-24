package main

import (
	arguments "./arguments"
	. "./data"
	reader "./jira/Config/Reader"
	. "./jira/HTMLParser"
	jiraConnection "./jira/HtmlConnection"
	. "./report"
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// TODO: have to split up to more than 1 method
	args := arguments.GetArguments()
	if args.IsHelpCall() || args.HasNoRunArgs() {
		return
	}

	config := reader.GetReader().Read(args.GetFilePathConfig())
	if (config == nil) {
		return
	}
	config.Jiradata.Password = SetEmptyPasswordOverConsoleInput(config.Jiradata.Password)

	var jc jiraConnection.HtmlConnector = jiraConnection.NewHtmlConnector(config)
	var pi TimeTrackingReport = config.GetTimeTrackingReportData()

	timeStart := time.Now()
	var retChannel chan ProjectReportSetting = make(chan ProjectReportSetting)
	for i := range pi.GetAllSettings() {
		go RetrieveNameTimePairPerProject(retChannel, pi.GetEntry(i), jc) // I really do not like this one! you can not add the pointer of an element of the map as param to a method :(
	}

	for j := 0; j < pi.GetSettingsLen(); j++ {
		prjSetting := <-retChannel
		pi.SetEntry(prjSetting)
	}
	close(retChannel)
	timeStop := time.Now()
	fmt.Printf("-->All projects retrieved in %v\n", timeStop.Sub(timeStart))
	var reporter ReporterInterface
	pi.Finish()
	reporter = GetReporter(';', config)
	reporter.ExportReport(pi)

}

func SetEmptyPasswordOverConsoleInput(pwd string) string {
	if len(pwd) == 0 {
		//TODO: need to read it in without echo of input, by 28.12.2015 there is no possibility to turn off echoing
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter password: ")
		pwd, _ = reader.ReadString('\n')
		pwd = strings.TrimSpace(pwd)
	}
	return pwd
}
