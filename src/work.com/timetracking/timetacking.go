package main

import (
    "bufio"
    "fmt"
    "os"
    "path/filepath"
    "strings"
    . "work.com/timetracking/helper"
    jiraConnection "work.com/timetracking/jiraConnector"
    parsehtml "work.com/timetracking/parsehtml"
    pt "work.com/timetracking/personaltime"
    prjinfo "work.com/timetracking/prjinfo"
)

var testing bool = false

type NameTimePair struct {
    // PrjName                string
    NameValues, TimeValues []string
}

func main() {
    var jc jiraConnection.JiraConnector
    var pi prjinfo.Projects

    if len(os.Args) >= 2 {
        if os.Args[1] == "-t" {
            testing = true
        } else if os.Args[1] == "-r" {
            testing = false
        } else if os.Args[1] == "--help" {
            content := ReadInFile("./timetracking_help.txt")
            fmt.Printf("%s\n", string(content))
            return
        } else {
            fmt.Printf("If you do not know how to use this program please call with \"--help\"\n")
            return
        }
    } else {
        fmt.Printf("If you do not know how to use this program please call with \"--help\"\n")
        return
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

    if testing {
        for k := range tm {
            fmt.Printf("Read in:\t%s\n", k)
        }
    }

    var nameTimePairs map[string]NameTimePair = make(map[string]NameTimePair)
    var content string
    for i := range pi.Data {
        fmt.Printf(".")
        if testing {
            content = string(ReadInFile("./testdata/Report-Jira.html"))
        } else {
            content = jc.GetReportContentForProjectInTimeframe(pi.Data[i]) // fix point to retrieve
        }
        var retValues NameTimePair
        var nameValues, timeValues []string
        nameValues, timeValues = ParseHTMLContent(content)
        retValues.NameValues = nameValues
        retValues.TimeValues = timeValues
        nameTimePairs[pi.Data[i].Prj] = retValues
    }
    fmt.Printf(".\n")

    PrintValuesForProject(nameTimePairs, tm)

}

func ParseHTMLContent(data string) ([]string, []string) {
    var generateStatitics parsehtml.ParseHTML
    tableWithNames := generateStatitics.ParseInputForHTMLTableFittingRegexp(generateStatitics.GetRegExpForTableRowToFindEmployeeNames(), data)
    nameValues := generateStatitics.ParseForTableRowsInHTMLTable("[A-Za-z]*\\.[A-Za-z]*", "td", " class=\"main\"", tableWithNames)

    tableWithTimes := generateStatitics.ParseInputForHTMLTableFittingRegexp(generateStatitics.GetRegExpForTableRowToFindTotalTimes(), data)
    timeValues := generateStatitics.ParseForTableRowsInHTMLTable("([0-9]*[wdhms]{1})+", "b", "", tableWithTimes)
    return nameValues, timeValues
}

func PrintValuesForProject(nameTimePairs map[string]NameTimePair, teammembers map[string]bool) {
    var totalPrjs map[string]pt.PersonalTime = make(map[string]pt.PersonalTime)
    var sumOfAllPrj float64 = 0

    for i := range nameTimePairs {
        var retTotalTime pt.PersonalTime
        retTotalTime = CreateTotalOfPrj(i, nameTimePairs[i], teammembers)
        sumOfAllPrj = sumOfAllPrj + retTotalTime.ToFloat64InHours()
        totalPrjs[i] = retTotalTime
    }

    for i := range totalPrjs {
        var prjTime pt.PersonalTime = totalPrjs[i]
        prjTime.SetOverallTime(sumOfAllPrj)
        if prjTime.ToFloat64InHours() > 0.0 {
            fmt.Printf("%s \n", prjTime.ToCsvFormat())
        }
        totalPrjs[i] = prjTime
    }
    fmt.Printf("Sum of TIMES: %fh\n", sumOfAllPrj)

}

func CreateTotalOfPrj(prjName string, nameTimePair NameTimePair, teammembers map[string]bool) pt.PersonalTime {
    var total pt.PersonalTime
    var sumOfTimes float64 = 0.0

    var i int = 0
    var personsTimes []pt.PersonalTime = make([]pt.PersonalTime, len(nameTimePair.NameValues)+1)

    for i = 0; i < len(nameTimePair.NameValues); i++ {
        var person pt.PersonalTime
        person.InitializeFromString(nameTimePair.NameValues[i], nameTimePair.TimeValues[i])
        personsTimes[i] = person
    }

    for key := range nameTimePair.NameValues {
        if teammembers[strings.ToLower(personsTimes[key].GetName())] == true {
            sumOfTimes = sumOfTimes + personsTimes[key].ToFloat64InHours()
            if testing {
                fmt.Printf("Team Member: %s : %f\n", personsTimes[key].ToCsvFormat(), personsTimes[key].ToFloat64InHours())
            }
        } else if testing {
            fmt.Printf("Non-Team Member: %s : %f\n", personsTimes[key].ToCsvFormat(), personsTimes[key].ToFloat64InHours())
        }
    }

    total.InitializeFromFloat(prjName, sumOfTimes)
    return total
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
            teammembers[strings.ToLower(line)] = true
        }
        line, errLine = reader.ReadString('\n')
    }
    return teammembers

}
