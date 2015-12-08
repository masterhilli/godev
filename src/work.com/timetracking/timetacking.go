package main

import (
    "bufio"
    "fmt"
    "os"
    "path/filepath"
    "strconv"
    "strings"
    "time"
    . "work.com/timetracking/HTMLParser"
    arguments "work.com/timetracking/arguments"
    . "work.com/timetracking/helper"
    jiraConfig "work.com/timetracking/jira/Config"
    jiraConnection "work.com/timetracking/jira/HtmlConnection"
    jiraTime "work.com/timetracking/jira/Timeentry"
    prjinfo "work.com/timetracking/prjinfo"
)

var args arguments.TimetrackingArgs

func main() {
    args = arguments.NewArguments()
    if args.IsHelpCall() || args.HasNoRunArgs() {
        return
    }

    var config jiraConfig.Config = jiraConfig.Reader.Read(args.GetFilePathConfig())
    var jc jiraConnection.HtmlConnector = jiraConnection.NewHtmlConnector(config)
    var tm map[string]bool = ReadTeammembers(args.GetFilePathToTeammembers())
    var pi prjinfo.Projects

    pi.Initialize(args.GetFilePathToProjects(), ',')

    if args.IsTesting() {
        for k := range tm {
            fmt.Printf("Read in:\t%s\n", k)
        }
    }

    timeStart := time.Now()
    var retChannel chan HTMLParser = make(chan HTMLParser)
    var nameTimePairs map[string]HTMLParser = make(map[string]HTMLParser)
    for i := range pi.Data {
        go RetrieveNameTimePairPerProject(retChannel, pi.Data[i], jc)
    }

    for j := 0; j < len(pi.Data); j++ {
        retValue := <-retChannel
        nameTimePairs[retValue.GetPrjInfo().Prj] = retValue
        // we just wait for the threads to end
    }
    close(retChannel)

    timeStop := time.Now()
    fmt.Printf("-->All projects retrieved in %v\n", timeStop.Sub(timeStart))
    PrintValuesForProject(nameTimePairs, tm)
    //PrintValuesForProject(pi, tm)

}

func PrintValuesForProject(nameTimePairs map[string]HTMLParser, teammembers map[string]bool) {
    //func PrintValuesForProject(pi prjinfo.Projects, teammembers map[string]bool) {
    var totalPrjs map[string]jiraTime.TimeEntry = make(map[string]jiraTime.TimeEntry)
    var sumOfAllPrj float64 = 0

    /*for i := range pi.Data {
        var retTotalTime jiraTime.TimeEntry
        retTotalTime = CreateTotalOfPrj(pi.Data[i].Prj, pi.Data[i], teammembers)
        sumOfAllPrj = sumOfAllPrj + retTotalTime.ToFloat64InHours()
        totalPrjs[pi.Data[i].Prj] = retTotalTime
    }*/

    for i := range nameTimePairs {
        var retTotalTime jiraTime.TimeEntry
        retTotalTime = CreateTotalOfPrj(i, nameTimePairs[i], teammembers)
        sumOfAllPrj = sumOfAllPrj + retTotalTime.ToFloat64InHours()
        totalPrjs[i] = retTotalTime
    }

    PrintValuesInCSVFormatSSS("Projectname", "Hours", "Percent")
    for i := range totalPrjs {
        var prjTime jiraTime.TimeEntry = totalPrjs[i]
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

func PrintValuesInCSVFormatPersTime(prjTime jiraTime.TimeEntry) {
    fmt.Printf("%s\n", prjTime.ToCsvFormat(seperator))
}

//func CreateTotalOfPrj(prjName string, nameTimePair prjinfo.Prjinfo, teammembers map[string]bool) jiraTime.TimeEntry {
func CreateTotalOfPrj(prjName string, nameTimePair HTMLParser, teammembers map[string]bool) jiraTime.TimeEntry {
    var total jiraTime.TimeEntry
    var sumOfTimes float64 = 0.0

    var i int = 0
    var personsTimes []jiraTime.TimeEntry = make([]jiraTime.TimeEntry, len(nameTimePair.GetNames())+1)
    var personsWithTime []string = make([]string, 0, len(nameTimePair.GetNames())+1)

    for i = 0; i < len(nameTimePair.GetNames()); i++ {
        var person jiraTime.TimeEntry
        person.InitializeFromString(nameTimePair.GetNames()[i], nameTimePair.GetTimes()[i])
        personsTimes[i] = person
    }

    for key := range nameTimePair.GetNames() {
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
