package main

import (
    "bufio"
    "fmt"
    "os"
    "path/filepath"
    "strconv"
    "strings"
    . "work.com/timetracking/helper"
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
    var testing bool = false

    if len(os.Args) >= 2 && os.Args[1] == "-t" {
        testing = true
    }

    if testing {
        pi.Initialize("./projects_test.csv", ',')
    } else {
        pi.Initialize("./projects.csv", ',')
    }
    jc.Initialize("./jira.yaml")

    var tm map[string]bool
    if testing {
        tm = ReadTeammembers("./teammembers_test.txt")
    } else {
        tm = ReadTeammembers("./teammembers.txt")
    }

    for k := range tm {
        fmt.Printf("Read in:\t%s\n", k)
    }

    var content string
    for i := range pi.Data {
        if testing {
            content = string(ReadInFile("./testdata/Report-Jira.html"))
        } else {
            content = jc.GetReportContentForProjectInTimeframe(pi.Data[i]) // fix point to retrieve
        }
        nameValues, timeValues := ParseHTMLContent(content)
        PrintValuesForProject(pi.Data[i].Prj, nameValues, timeValues, tm)
    }

}

func ParseHTMLContent(data string) ([]string, []string) {
    var generateStatitics parsehtml.ParseHTML
    tableWithNames := generateStatitics.ParseInputForHTMLTableFittingRegexp(generateStatitics.GetRegExpForTableRowToFindEmployeeNames(), data)
    nameValues := generateStatitics.ParseForTableRowsInHTMLTable("[A-Za-z]*\\.[A-Za-z]*", "td", " class=\"main\"", tableWithNames)

    tableWithTimes := generateStatitics.ParseInputForHTMLTableFittingRegexp(generateStatitics.GetRegExpForTableRowToFindTotalTimes(), data)
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

func PrintValuesForProject(prjName string, nameValues, timeValues []string, teammembers map[string]bool) {
    var total pt.PersonalTime
    var sumOfTimes float64 = 0.0

    var i int = 0
    var personsTimes []pt.PersonalTime = make([]pt.PersonalTime, len(nameValues)+1)

    for i = 0; i < len(nameValues); i++ {
        var person pt.PersonalTime
        person.Initialize(nameValues[i], timeValues[i])
        personsTimes[i] = person
    }

    for key := range nameValues {
        if teammembers[personsTimes[key].GetName()] == true {
            sumOfTimes = sumOfTimes + personsTimes[key].ToFloat64InHours()
            fmt.Printf("Team Member: %s : %f\n", personsTimes[key].ToCsvFormat(), personsTimes[key].ToFloat64InHours())
        } else {
            fmt.Printf("Non-Team Member: %s : %f\n", personsTimes[key].ToCsvFormat(), personsTimes[key].ToFloat64InHours())
        }
    }

    if len(timeValues) > 0 {
        total.Initialize(prjName, timeValues[len(timeValues)-1])
    } else {
        total.Initialize(prjName, "0h")
    }

    fmt.Printf("%s : %s\n", total.ToCsvFormat(), strconv.FormatFloat(sumOfTimes, 'f', 2, 64))
}

func ReadTeammembers(file string) map[string]bool {
    var reader *bufio.Reader
    var teammembers map[string]bool = make(map[string]bool)
    filename, errAbs := filepath.Abs(file)
    PanicOnError(errAbs)

    f, errOpen := os.Open(filename)
    PanicOnError(errOpen)
    reader = bufio.NewReader(f)
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
