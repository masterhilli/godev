package main


import	(
	"fmt"
    "io/ioutil"
)


func main () {
    var generateStatitics GenerateHoursPerProject
    data := generateStatitics.RetrieveJIRAReportStream()
    fmt.Printf("Data read in from file, len: %d\n", len(data))
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
    return "(?is)<tr>.*[<td [[:ascii:]]*>[[:alpha:]]*\\.[[:alpha:]]*</td>]*[[:space:]]*<td class=\"main\"><b>Total</b></td>[[:space:]]*</tr>"
}