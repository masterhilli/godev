package main

import (
    "fmt"
    //. "work.com/timetracking/helper"
    jiraConnection "work.com/timetracking/jiraConnector"
    parsehtml "work.com/timetracking/parsehtml"
    pt "work.com/timetracking/personaltime"
    prjinfo "work.com/timetracking/prjinfo"
)

/* TODO:
* Create a struct holding all jira log in information & connection strings
* Create a struct that holds the projects to retrieve (time and so on)
* Create a struct that parses the data and creates the time overview
* Create a struct that creates a csv file with the necessary information (in future this might be possible to create in confluence itself)
 */

func main() {
    var jc jiraConnection.JiraConnector
    var pi prjinfo.Projects
    pi.Initialize("./projects.csv", ',')
    jc.Initialize("./jira.yaml")

    for i := range pi.Data {
        //for i := 0; i < len(pi.Data); i++ {
        content := jc.GetReportContentForProjectInTimeframe(pi.Data[i]) // fix point to retrieve
        //content := string(ReadInFile("./testdata/Report-Jira.html"))
        nameValues, timeValues := ParseHTMLContent(content)
        //PrintValuesPerItem(pi.Data[i].Prj, nameValues, timeValues)
        PrintValuesForProject(pi.Data[i].Prj, nameValues, timeValues)
    }

}

func ParseHTMLContent(data string) ([]string, []string) {
    var generateStatitics parsehtml.ParseHTML
    tableWithNames := generateStatitics.ParseInputForHTMLTableFittingRegexp(generateStatitics.GetRegExpForTableRowToFindEmployeeNames(), data)
    //fmt.Printf("Table with names: %s\n", tableWithNames)
    nameValues := generateStatitics.ParseForTableRowsInHTMLTable("[A-Za-z]*\\.[A-Za-z]*", "td", " class=\"main\"", tableWithNames)

    tableWithTimes := generateStatitics.ParseInputForHTMLTableFittingRegexp(generateStatitics.GetRegExpForTableRowToFindTotalTimes(), data)
    //fmt.Printf("Table with times: %s\n", tableWithTimes) //
    timeValues := generateStatitics.ParseForTableRowsInHTMLTable("([0-9]*[wdhms]{1})+", "b", "", tableWithTimes)
    return nameValues, timeValues
}

func PrintValuesPerItem(nameValues, timeValues []string) {
    if len(nameValues)+1 == len(timeValues) {
        var i int = 0
        var personsTimes []pt.PersonalTime = make([]pt.PersonalTime, len(nameValues)+1)
        for i = 0; i < len(nameValues); i++ {
            var person pt.PersonalTime
            person.Initialize(nameValues[i], timeValues[i])
            personsTimes[i] = person
        }
        var total pt.PersonalTime
        total.Initialize("TOTAL", timeValues[i])
        personsTimes[i] = total
        for i := range personsTimes {
            fmt.Printf("PERSON: %s\n", personsTimes[i].ToCsvFormat())
        }
    }
}

func PrintValuesForProject(prjName string, nameValues, timeValues []string) {
    var total pt.PersonalTime
    if len(timeValues) > 0 {
        total.Initialize(prjName, timeValues[len(timeValues)-1])
    } else {
        total.Initialize(prjName, "0h")
    }
    fmt.Printf("%s\n", total.ToCsvFormat())
}
