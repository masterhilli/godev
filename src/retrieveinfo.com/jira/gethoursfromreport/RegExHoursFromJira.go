package jiraRegEx
import ( 
	"fmt"
	"regexp"
	)

type IntPair struct {
    start, stop int
}

func ReturnTRValues(stringToParse string) [] string {
	return returnValuesOfTag(stringToParse, "tr", "")
}

func ReturnTDClassMainValues(stringToParse string) [] string {
	fmt.Printf("Input for TD Search: %s", stringToParse)
	return returnValuesOfTag(stringToParse, "td",  " class=\"main\"")
}

func returnValuesOfTag(stringToParse string, tag string, attributes string) [] string {
	indexArray := parseForTagStartEnd(stringToParse, tag, attributes)
	return trimTagsFromArray(indexArray, stringToParse, len(tag+attributes))
}

func parseForTagStartEnd(stringToParse string, tag string, attributes string) []IntPair {
	indexArray := returnIndexPairs(stringToParse, tag, attributes)
	if (indexArray == nil) {
		return []IntPair{{0,0}}
	}
	return indexArray
}

func returnIndexArray(stringToParse string, searchRegExp string) []int {
	regexpIndexFinder := regexp.MustCompile(searchRegExp)
	indexArray := regexpIndexFinder.FindAllStringSubmatchIndex(stringToParse, 2000)
	
	if  indexArray == nil { 
//		fmt.Printf("**********returnIndexArray nil**********\n")
		return nil 
	} else {
//		fmt.Printf("**********returnIndexArray:%d %d**********\n", len(indexArray), len(indexArray[0]))
		return indexArray[0]
	}
}

func returnIndexPairs(stringToParse string, tag string, attributes string) []IntPair {
	startIndexArray := returnIndexArray(stringToParse, "<"+tag+ attributes+">")
	stopIndexArray  := returnIndexArray(stringToParse, "</" + tag + ">")
	if len(startIndexArray) != len(stopIndexArray) {
		return nil
	}

	//fmt.Printf("Length is not the same %s\n", stringToParse)

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
		fmt.Printf("*****************ParsedSubMatch: %s\n", parsedSubMatchedTexts[i])
	}	
	return parsedSubMatchedTexts
}