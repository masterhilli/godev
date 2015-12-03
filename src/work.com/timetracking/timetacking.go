package main

import (
    "bufio"
    "fmt"
    "os"
    "path/filepath"
    "strconv"
    "strings"
    "time"
    arguments "work.com/timetracking/arguments"
    . "work.com/timetracking/helper"
    jiraConnection "work.com/timetracking/jiraConnector"
    parsehtml "work.com/timetracking/parsehtml"
    pt "work.com/timetracking/personaltime"
    prjinfo "work.com/timetracking/prjinfo"
)

var args arguments.TimetrackingArgs

type NameTimePair struct {
    // PrjName                string
    NameValues, TimeValues []string
}

type ChanelReturnValue struct {
    Prj     string
    Content NameTimePair
}

func main() {
    var jc jiraConnection.JiraConnector
    var pi prjinfo.Projects
    args = arguments.ParseArguments(os.Args)

    if args.IsHelpCall() {
        return
    }

    if !args.IsTesting() && !args.IsRunning() {
        fmt.Println("If you do not know how to use this program please call with \"--help\"")
        return
    }

    if args.IsTesting() {
        pi.Initialize("./projects_test.csv", ',')
    } else {
        pi.Initialize(args.GetFilePathToProjects(), ',')
    }
    jc.Initialize(args.GetFilePathConfig())

    var tm map[string]bool
    if args.IsTesting() {
        tm = ReadTeammembers("./teammembers_test.txt")
    } else {
        tm = ReadTeammembers(args.GetFilePathToTeammembers())
    }

    if args.IsTesting() {
        for k := range tm {
            fmt.Printf("Read in:\t%s\n", k)
        }
    }

    timeStart := time.Now()
    var myRetValueChannel chan ChanelReturnValue = make(chan ChanelReturnValue)
    var nameTimePairs map[string]NameTimePair = make(map[string]NameTimePair)
    for i := range pi.Data {
        go RunRetrieveContent(myRetValueChannel, pi.Data[i], jc)
    }

    for j := 0; j < len(pi.Data); j++ {
        retValue := <-myRetValueChannel
        nameTimePairs[retValue.Prj] = retValue.Content
    }
    close(myRetValueChannel)

    timeStop := time.Now()
    fmt.Printf("-->All projects retrieved in %v\n", timeStop.Sub(timeStart))
    PrintValuesForProject(nameTimePairs, tm)

}

func RunRetrieveContent(returnChannel chan ChanelReturnValue, prjInfo prjinfo.Prjinfo, jc jiraConnection.JiraConnector) {
    timeStart := time.Now()
    var content string
    var retVal ChanelReturnValue

    if args.IsTesting() {
        content = string(ReadInFile("./testdata/Report-Jira.html"))
    } else {
        content = jc.GetReportContentForProjectInTimeframe(prjInfo) // fix point to retrieve
    }
    var retValues NameTimePair
    var nameValues, timeValues []string
    nameValues, timeValues = ParseHTMLContent(content)
    retValues.NameValues = nameValues
    retValues.TimeValues = timeValues
    retVal.Prj = prjInfo.Prj
    retVal.Content = retValues
    timeStop := time.Now()
    fmt.Printf("-->%s DONE in %v\n", retVal.Prj, timeStop.Sub(timeStart))
    returnChannel <- retVal
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

    PrintValuesInCSVFormatSSS("Projectname", "Hours", "Percent")
    for i := range totalPrjs {
        var prjTime pt.PersonalTime = totalPrjs[i]
        prjTime.SetOverallTime(sumOfAllPrj)
        if prjTime.ToFloat64InHours() > 0.0 {
            PrintValuesInCSVFormatPersTime(prjTime)
        }
        totalPrjs[i] = prjTime
    }

    PrintValuesInCSVFormatSFF("OVERALLTIME", sumOfAllPrj, 100.0)
}

var seperator rune = ';'

func PrintValuesInCSVFormatSFF(projectname string, hours float64, percent float64) {
    PrintValuesInCSVFormatSSS(projectname, strconv.FormatFloat(hours, 'f', 2, 64), strconv.FormatFloat(percent, 'f', 1, 64)+"%")
}
func PrintValuesInCSVFormatSSS(projectname string, hours string, percent string) {
    fmt.Printf("%s%c%s%c%s%c\n", projectname, seperator, hours, seperator, percent, seperator)
}

func PrintValuesInCSVFormatPersTime(prjTime pt.PersonalTime) {
    fmt.Printf("%s\n", prjTime.ToCsvFormat(seperator))
}

func CreateTotalOfPrj(prjName string, nameTimePair NameTimePair, teammembers map[string]bool) pt.PersonalTime {
    var total pt.PersonalTime
    var sumOfTimes float64 = 0.0

    var i int = 0
    var personsTimes []pt.PersonalTime = make([]pt.PersonalTime, len(nameTimePair.NameValues)+1)
    var personsWithTime []string = make([]string, 0, len(nameTimePair.NameValues)+1)

    for i = 0; i < len(nameTimePair.NameValues); i++ {
        var person pt.PersonalTime
        person.InitializeFromString(nameTimePair.NameValues[i], nameTimePair.TimeValues[i])
        personsTimes[i] = person
    }

    for key := range nameTimePair.NameValues {
        if teammembers[strings.ToLower(personsTimes[key].GetName())] == true {
            sumOfTimes = sumOfTimes + personsTimes[key].ToFloat64InHours()
            if personsTimes[key].ToFloat64InHours() > 0.0 {
                lastname := personsTimes[key].GetName()
                lastname = lastname[strings.IndexRune(lastname, '.')+1:]
                personsWithTime = append(personsWithTime, lastname)
            }
            if args.IsTesting() {
                fmt.Printf("Team Member: %s : %f\n", personsTimes[key].ToCsvFormat(seperator), personsTimes[key].ToFloat64InHours())
            }
        } else if args.IsTesting() {
            fmt.Printf("Non-Team Member: %s : %f\n", personsTimes[key].ToCsvFormat(seperator), personsTimes[key].ToFloat64InHours())
        }
    }

    total.InitializeFromFloat(prjName, sumOfTimes, personsWithTime)
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
