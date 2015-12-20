package jira

import (
    "fmt"
    "regexp"
    "time"
    . "../../data"
    . "../../helper"
    jiraConnection "../HtmlConnection"
    . "../../arguments"
)

const pathToTemplateReportJiraHTML string = "./__testdata/Report-Jira.html"
func RetrieveNameTimePairPerProject(retChan chan ProjectReportSetting, prjInfo ProjectReportSetting, jc jiraConnection.HtmlConnector) {
    timeStart := time.Now()
    var htmlParser HTMLParser
    var content string
    args := GetArguments()
    if args.IsTesting() {
        content = string(ReadInFile(pathToTemplateReportJiraHTML))
    } else {
      content = jc.GetReportContentForProjectInTimeframe(prjInfo)
    }

    htmlParser.ParseHTMLContent(content)
    htmlParser.prjInfo = prjInfo
    prjInfo.SetNames(htmlParser.GetNames())
    prjInfo.SetTimes(htmlParser.GetTimes())
    timeStop := time.Now()
    fmt.Printf("-->%s DONE in %v\n", prjInfo.Prj, timeStop.Sub(timeStart))
    retChan <- prjInfo
}

type HTMLParser struct {
    names, times []string
    prjInfo      ProjectReportSetting
}

func (ghpp *HTMLParser) GetNames() []string {
    return ghpp.names
}

func (ghpp *HTMLParser) GetTimes() []string {
    return ghpp.times
}

func (ghpp *HTMLParser) GetPrjInfo() ProjectReportSetting {
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
