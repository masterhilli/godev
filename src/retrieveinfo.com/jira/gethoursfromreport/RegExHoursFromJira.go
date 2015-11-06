package jiraRegEx
import ( 
//	"fmt"
	"regexp"
	)

var regexpForTRStart string = "<tr>"
var regexpForTRStop string = "</tr>"

type IntPair struct {
    start, stop int
}

func returnIndexArray(stringToParse string, searchRegExp string) []int {
	regexpIndexFinder := regexp.MustCompile(searchRegExp)
	indexArray := regexpIndexFinder.FindAllStringSubmatchIndex(stringToParse, 2000)
//	fmt.Printf("**********%d %d**********", len(indexArray), len(indexArray[0]))
	if  indexArray == nil { 
		return nil 
	} else {
		return indexArray[0]
	}
}

func returnIndexPairs(stringToParse string) []IntPair {
	startIndexArray := returnIndexArray(stringToParse, regexpForTRStart)
	stopIndexArray  := returnIndexArray(stringToParse, regexpForTRStop)
	if len(startIndexArray) != len(stopIndexArray) {
//		fmt.Printf("Length is not the same %s\n", stringToParse)
		return nil
	}

	var indexPairs []IntPair = make([]IntPair, len(startIndexArray))
	for i := 0; i < len(startIndexArray); i++ {
		indexPairs[i].start = startIndexArray[i]
		indexPairs[i].stop  = stopIndexArray[i]
	}
	return indexPairs
}

func ParseForTRTR(stringToParse string) []string {
	indexArray := returnIndexPairs(stringToParse)
	if (indexArray == nil) {
		return []string{""}
	}

	parsedSubMatchedTexts := make([]string, len(indexArray))
	for i := 0; i < len(indexArray); i++ {
//		fmt.Printf("start: %d / stop: %d \"%s\"\n", indexArray[i].start, indexArray[i].stop, stringToParse)
		if (indexArray[i].stop - indexArray[i].start) >= 4 {
			parsedSubMatchedTexts[i] = stringToParse[indexArray[i].start+4:indexArray[i].stop]
		} else {
			parsedSubMatchedTexts[i] = ""
		}
	}	
	return parsedSubMatchedTexts
}

/*
func HasTablesOnHtml(){
	Mat
}*/
