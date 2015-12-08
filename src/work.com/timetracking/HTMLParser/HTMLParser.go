package htmlparser

import (
    "fmt"
    "regexp"
    "time"
    . "work.com/timetracking/data"
    . "work.com/timetracking/helper"
    jiraConnection "work.com/timetracking/jira/HtmlConnection"
)

func RetrieveNameTimePairPerProject(retChan chan HTMLParser, prjInfo *Prjinfo, jc jiraConnection.HtmlConnector) {
    timeStart := time.Now()
    var htmlParser HTMLParser

    content := string(ReadInFile("./testdata/Report-Jira.html"))
    /*content := jc.GetReportContentForProjectInTimeframe(prjInfo)
      if args.IsTesting() {

      } else {
          // fix point to retrieve
      }*/

    htmlParser.ParseHTMLContent(content)
    htmlParser.prjInfo = *prjInfo
    prjInfo.SetNames(htmlParser.GetNames())
    prjInfo.SetTimes(htmlParser.GetTimes())
    timeStop := time.Now()
    fmt.Printf("-->%s DONE in %v\n", prjInfo.Prj, timeStop.Sub(timeStart))
    retChan <- htmlParser
}

type HTMLParser struct {
    names, times []string
    prjInfo      Prjinfo
}

func (ghpp *HTMLParser) GetNames() []string {
    return ghpp.names
}

func (ghpp *HTMLParser) GetTimes() []string {
    return ghpp.times
}

func (ghpp *HTMLParser) GetPrjInfo() Prjinfo {
    return ghpp.prjInfo
}

func (ghpp *HTMLParser) ParseHTMLContent(data string) {
    tableWithNames := ghpp.parseInputForHTMLTableFittingRegexp(ghpp.getRegExpForTableRowToFindEmployeeNames(), data)
    ghpp.names = ghpp.parseForTableRowsInHTMLTable("[A-Za-z]*\\.[A-Za-z]*", "td", " class=\"main\"", tableWithNames)

    tableWithTimes := ghpp.parseInputForHTMLTableFittingRegexp(ghpp.getRegExpForTableRowToFindTotalTimes(), data)
    ghpp.times = ghpp.parseForTableRowsInHTMLTable("([0-9]*[wdhms]{1})+", "b", "", tableWithTimes)
}

func (ghpp *HTMLParser) getRegExpForTableRowToFindEmployeeNames() string {
    return "(?is)<tr>.*[<td [[:ascii:]]*>[[:alpha:]]*\\.[[:alpha:]]*</td>]*[[:space:]]*<td class=\"main\"><b>Total</b></td>[[:space:]]*</tr>"
}

func (ghpp *HTMLParser) getRegExpForTableRowToFindTotalTimes() string {
    return "(?is)<tr>([[:space:]]*<td class=\"total\"><b>(([0-9]*[wdhms]{1})+|total|Issue)</b></td>)+[[:space:]]*</tr>"
}

func (ghpp *HTMLParser) parseInputForHTMLTableFittingRegexp(regexpForSearch string, data string) string {
    regexpForMatchingNames := regexp.MustCompile(regexpForSearch)
    indexArray := regexpForMatchingNames.FindStringSubmatchIndex(data)
    if indexArray == nil {
        return string("")
    }
    return data[indexArray[0]:indexArray[1]]
}

func (ghpp *HTMLParser) parseForTableRowsInHTMLTable(regexpToSearchfor string, tag string, attributes string, data string) []string {
    startTag := "<" + tag + attributes + ">"
    endTag := "</" + tag + ">"
    regexpForMatchingNames := regexp.MustCompile("(?is)" + startTag + regexpToSearchfor + endTag)
    indexArray := regexpForMatchingNames.FindAllStringSubmatchIndex(data, -1)
    if indexArray == nil {
        return nil
    }
    var values []string = make([]string, len(indexArray))
    for i := 0; i < len(indexArray); i++ {
        if len(indexArray[i]) >= 2 {
            values[i] = data[indexArray[i][0]+len(startTag) : indexArray[i][1]-len(endTag)]
        }
    }
    return values
}
