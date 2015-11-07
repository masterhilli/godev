package jiraRegEx
import ( 
//	"fmt"
	"regexp"
	)

type IntPair struct {
    start, stop int
}

func ReturnTRValues(stringToParse string) [] string {
	indexArray := parseForTagStartEnd(stringToParse, "tr")
	return trimTagsFromArray(indexArray, stringToParse, len("tr"))
}

func parseForTagStartEnd(stringToParse string, tag string) []IntPair {
	indexArray := returnIndexPairs(stringToParse, tag)
	if (indexArray == nil) {
		return []IntPair{{0,0}}
	}
	return indexArray
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

func returnIndexPairs(stringToParse string, tag string) []IntPair {
	startIndexArray := returnIndexArray(stringToParse, "<"+tag+">")
	stopIndexArray  := returnIndexArray(stringToParse, "</" + tag + ">")
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

func trimTagsFromArray(indexArray []IntPair, stringToTrim string, tagNameLen int) []string{
	tagLen := tagNameLen+2

	parsedSubMatchedTexts := make([]string, len(indexArray))
	for i := 0; i < len(indexArray); i++ {
		if (indexArray[i].stop - indexArray[i].start) >= tagLen {
			parsedSubMatchedTexts[i] = stringToTrim[indexArray[i].start+tagLen:indexArray[i].stop]
		} else {
			parsedSubMatchedTexts[i] = ""
		}
	}	
	return parsedSubMatchedTexts
}