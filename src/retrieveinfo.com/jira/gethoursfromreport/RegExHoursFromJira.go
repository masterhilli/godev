package jiraRegEx
import "regexp"
func ParseForTRTR(stringToParse string) string {
	searchText := regexp.MustCompile("<tr>[[:ascii:]]*</tr>")
	if searchText.MatchString(stringToParse) == true {
		submatchOfFoundString := searchText.FindSubmatch([]byte(stringToParse))
		substring :=  string(submatchOfFoundString[0])
		return substring[4:len(substring)-5]
	} else {
		return "";
	}
}

/*
func HasTablesOnHtml(){
	Mat
}*/
