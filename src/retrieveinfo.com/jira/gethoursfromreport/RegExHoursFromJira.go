package jiraRegEx
import ( 
	"fmt"
	"regexp"
	)

func ParseForTRTR(stringToParse string) []string {
	searchText := regexp.MustCompile("<tr>.*</tr>+")
	if searchText.MatchString(stringToParse) == true {
		submatchOfFoundStrings := searchText.FindAllString(stringToParse, 2000)
		countFoundMatches := len(submatchOfFoundStrings)
		allFoundMatchesAsStrings := make([]string, countFoundMatches);
		fmt.Printf("first: %d %s \n", len(submatchOfFoundStrings), submatchOfFoundStrings)
		for i := 0; i < len(submatchOfFoundStrings); i++{
				manipulate := submatchOfFoundStrings[i];
				if (len(manipulate) >=9){
					allFoundMatchesAsStrings[i] = manipulate[4:len(manipulate)-5]
					fmt.Printf("Foundstream: %s", allFoundMatchesAsStrings[i])
				} else {
					allFoundMatchesAsStrings[i] = ""
				}
		}
		
		return allFoundMatchesAsStrings
	} 
	return []string{""}
}

/*
func HasTablesOnHtml(){
	Mat
}*/
