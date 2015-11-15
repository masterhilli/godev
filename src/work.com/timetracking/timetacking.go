package main

import (
	"fmt"
	parsehtml "work.com/timetracking/parsehtml"
	pt 		  "work.com/timetracking/personaltime"
)

func printStringArrayForTables(name string, values []string){
    fmt.Printf("Values for %s", name)
    if values != nil {
        for value := range values {
            fmt.Printf(": %s", values[value])
        }
        fmt.Printf(": LEN: %d\n", len(values))
    } else {
        fmt.Printf("no values")
    }
}

func main () {
    var generateStatitics parsehtml.ParseHTML
    data := generateStatitics.RetrieveJIRAReportStream()
    fmt.Printf("Data read in from file, len: %d\n", len(data))

    tableWithNames := generateStatitics.ParseInputForHTMLTableFittingRegexp(generateStatitics.GetRegExpForTableRowToFindEmployeeNames(), data)
    //fmt.Printf("Table with names: %s\n", tableWithNames)
    nameValues := generateStatitics.ParseForTableRowsInHTMLTable("[A-Za-z]*\\.[A-Za-z]*", "td", " class=\"main\"", tableWithNames)
    printStringArrayForTables("Names", nameValues)


    tableWithTimes := generateStatitics.ParseInputForHTMLTableFittingRegexp(generateStatitics.GetRegExpForTableRowToFindTotalTimes(), data)
    //fmt.Printf("Table with times: %s\n", tableWithTimes) // 
    timeValues := generateStatitics.ParseForTableRowsInHTMLTable("([0-9]*[wdhms]{1})+", "b", "", tableWithTimes)
    printStringArrayForTables("Times", timeValues)

    if (len(nameValues)+1 == len(timeValues)) {
        var i int = 0
        var personsTimes []pt.PersonalTime = make([]pt.PersonalTime, len(nameValues)+1)
        for i = 0; i <len(nameValues); i++ {
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