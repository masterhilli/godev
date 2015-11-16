package main

import (
    "fmt"
    . "work.com/timetracking/helper"
    //    jiraConnect "work.com/timetracking/jiraConnector"
    parsehtml "work.com/timetracking/parsehtml"
    pt "work.com/timetracking/personaltime"
)

/* TODO:
* Create a struct holding all jira log in information & connection strings
* Create a struct that holds the projects to retrieve (time and so on)
* Create a struct that parses the data and creates the time overview
* Create a struct that creates a csv file with the necessary information (in future this might be possible to create in confluence itself)
 */

func main() {
    var generateStatitics parsehtml.ParseHTML
    data := string(ReadInFile("./testdata/report-jira.html"))
    fmt.Printf("Data read in from file, len: %d\n", len(data))

    tableWithNames := generateStatitics.ParseInputForHTMLTableFittingRegexp(generateStatitics.GetRegExpForTableRowToFindEmployeeNames(), data)
    //fmt.Printf("Table with names: %s\n", tableWithNames)
    nameValues := generateStatitics.ParseForTableRowsInHTMLTable("[A-Za-z]*\\.[A-Za-z]*", "td", " class=\"main\"", tableWithNames)
    PrintStringArrayForTables("Names", nameValues)

    tableWithTimes := generateStatitics.ParseInputForHTMLTableFittingRegexp(generateStatitics.GetRegExpForTableRowToFindTotalTimes(), data)
    //fmt.Printf("Table with times: %s\n", tableWithTimes) //
    timeValues := generateStatitics.ParseForTableRowsInHTMLTable("([0-9]*[wdhms]{1})+", "b", "", tableWithTimes)
    PrintStringArrayForTables("Times", timeValues)

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
            fmt.Printf("PERSON: %s\n", personsTimes[i].ToString())
        }
    }

}
