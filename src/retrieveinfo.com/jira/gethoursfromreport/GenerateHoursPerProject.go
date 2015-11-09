package main


import	(
	"fmt"
    "regexp"
    "io/ioutil"
    "strconv"
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
    var generateStatitics GenerateHoursPerProject
    data := generateStatitics.RetrieveJIRAReportStream()
    fmt.Printf("Data read in from file, len: %d\n", len(data))

    tableWithNames := generateStatitics.ParseInputForHTMLTableFittingRegexp(generateStatitics.getRegExpForTableRowToFindEmployeeNames(), data)
    //fmt.Printf("Table with names: %s\n", tableWithNames)
    nameValues := generateStatitics.parseForTableRowsInHTMLTable("[A-Za-z]*\\.[A-Za-z]*", "td", " class=\"main\"", tableWithNames)
    printStringArrayForTables("Names", nameValues)


    tableWithTimes := generateStatitics.ParseInputForHTMLTableFittingRegexp(generateStatitics.getRegExpForTableRowToFindTotalTimes(), data)
    //fmt.Printf("Table with times: %s\n", tableWithTimes) // 
    timeValues := generateStatitics.parseForTableRowsInHTMLTable("([0-9]*[wdhms]{1})+", "b", "", tableWithTimes)
    printStringArrayForTables("Times", timeValues)

    if (len(nameValues)+1 == len(timeValues)) {
        var i int = 0
        var personsTimes []PersonalTime = make([]PersonalTime, len(nameValues)+1)
        for i = 0; i <len(nameValues); i++ {
            var person PersonalTime
            person.Initialize(nameValues[i], timeValues[i])
            personsTimes[i] = person
        }
        var total PersonalTime
        total.Initialize("TOTAL", timeValues[i])
        personsTimes[i] = total
        for i := range personsTimes {
            fmt.Printf("PERSON: %s\n", personsTimes[i].ToString())
        }
    }

}

type GenerateHoursPerProject struct {
    
}

func (ghpp *GenerateHoursPerProject) check(e error) {
    if e != nil {
        panic(e)
    }
}

func (ghpp *GenerateHoursPerProject) RetrieveJIRAReportStream() string {
    data, err := ioutil.ReadFile("./testdata/Report-Jira.html")
    ghpp.check(err)
    return string(data)
}

func (ghpp *GenerateHoursPerProject) getRegExpForTableRowToFindEmployeeNames() string {
    return "(?is)<tr>.*[<td [[:ascii:]]*>[[:alpha:]]*\\.[[:alpha:]]*</td>]*[[:space:]]*<td class=\"main\"><b>Total</b></td>[[:space:]]*</tr>"
}

func (ghpp *GenerateHoursPerProject) getRegExpForTableRowToFindTotalTimes() string {
    return "(?is)<tr>([[:space:]]*<td class=\"total\"><b>(([0-9]*[wdhms]{1})+|total|Issue)</b></td>)+[[:space:]]*</tr>"
}

func (ghpp *GenerateHoursPerProject) ParseInputForHTMLTableFittingRegexp(regexpForSearch string, data string) string {
    regexpForMatchingNames := regexp.MustCompile(regexpForSearch)
    indexArray := regexpForMatchingNames.FindStringSubmatchIndex(data)
    if (indexArray == nil) {
        return string("")
    }
    return data[indexArray[0]:indexArray[1]]
}

func (ghpp *GenerateHoursPerProject)  parseForTableRowsInHTMLTable(regexpToSearchfor string, tag string, attributes string, data string) []string {
    startTag := "<" + tag + attributes +">"
    endTag := "</"+tag+">"
    regexpForMatchingNames := regexp.MustCompile("(?is)"+ startTag+regexpToSearchfor + endTag)
    indexArray := regexpForMatchingNames.FindAllStringSubmatchIndex(data, -1)
    if indexArray == nil {
        return nil
    }
    var values []string = make([]string, len(indexArray))
    for i := 0 ; i < len(indexArray); i++ {
        if (len(indexArray[i]) >= 2) {
            values[i] = data[indexArray[i][0]+len(startTag): indexArray[i][1]-len(endTag)]
        }
    }
    return values
}

type PersonalTime struct {
    name string
    weeks, days, hours, mins, secs int
}

func (p *PersonalTime) Initialize(name string, time string) {
    p.name = name
    p.InitializeTime(time)
}

func (p *PersonalTime) InitializeTime(time string) {
    p.weeks = ParseForInteger(time, "w")
    p.days = ParseForInteger(time, "d")
    p.hours = ParseForInteger(time, "h")
    p.mins = ParseForInteger(time, "m")
    p.secs = ParseForInteger(time, "s")
}

func ParseForInteger(time string, timeIdentifier string) int {
    regexpForValue := regexp.MustCompile("(?is)[0-9]+"+ timeIdentifier)
    valueFromRegExp := regexpForValue.FindStringSubmatch(time)
    if valueFromRegExp == nil {
        //fmt.Printf("*** DEBUG: valueFromRegExp returned to nil(%s/%s)\n", time, timeIdentifier)
        return 0
    }
    
    match := valueFromRegExp[0]
    match = match[0:len(match)-1]
    value, err := strconv.Atoi(match)
    
    if err != nil {
        return 0
    } else {
        return value
    }
}

func (p *PersonalTime) ToString() string {  
    //fmt.Printf("%s: %dw/%dd/%dh/%dm/%ds\n", p.name, p.weeks, p.days, p.hours, p.mins, p.secs)
    return "***" + p.name   + " : "+p.toStringTimes(p.weeks, "week(s)") +
                              " : "+p.toStringTimes(p.days, "day(s)") +
                              " : "+p.toStringTimes(p.hours, "hour(s)")+ 
                              " : "+p.toStringTimes(p.mins, "min(s)")+
                              " : "+p.toStringTimes(p.secs, "sec(s)")
}

func (p *PersonalTime) toStringTimes(time int, name string) string {
    var timeStringBuffer string = ""
    if (time > 0) {
        timeStringBuffer = timeStringBuffer + strconv.Itoa(time) + " " + name
    }
    return timeStringBuffer
}