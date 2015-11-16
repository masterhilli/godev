package parsehtml

import (
    "regexp"
)

type ParseHTML struct {
}

func (ghpp *ParseHTML) GetRegExpForTableRowToFindEmployeeNames() string {
    return "(?is)<tr>.*[<td [[:ascii:]]*>[[:alpha:]]*\\.[[:alpha:]]*</td>]*[[:space:]]*<td class=\"main\"><b>Total</b></td>[[:space:]]*</tr>"
}

func (ghpp *ParseHTML) GetRegExpForTableRowToFindTotalTimes() string {
    return "(?is)<tr>([[:space:]]*<td class=\"total\"><b>(([0-9]*[wdhms]{1})+|total|Issue)</b></td>)+[[:space:]]*</tr>"
}

func (ghpp *ParseHTML) ParseInputForHTMLTableFittingRegexp(regexpForSearch string, data string) string {
    regexpForMatchingNames := regexp.MustCompile(regexpForSearch)
    indexArray := regexpForMatchingNames.FindStringSubmatchIndex(data)
    if indexArray == nil {
        return string("")
    }
    return data[indexArray[0]:indexArray[1]]
}

func (ghpp *ParseHTML) ParseForTableRowsInHTMLTable(regexpToSearchfor string, tag string, attributes string, data string) []string {
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
