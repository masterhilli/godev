package jiraRegEx
import ( 
//	"fmt"
	"regexp"
	)

type IntPair struct {
    start, stop int
}

func ReturnTRValues(stringToParse string) [] string {
	return returnValuesOfTag(stringToParse, "tr", "")
}

func ReturnTDClassMainValues(stringToParse string) [] string {
//	fmt.Printf("*****************Input for TD Search: %s\n", stringToParse)
	return returnValuesOfTag(stringToParse, "td",  " class=\"main\"")
}

func returnValuesOfTag(stringToParse string, tag string, attributes string) [] string {
	indexArray := parseForTagStartEnd(stringToParse, tag, attributes)
	return trimTagsFromArray(indexArray, stringToParse, len(tag+attributes))
}

func parseForTagStartEnd(stringToParse string, tag string, attributes string) []IntPair {
	indexArray := returnIndexPairs(stringToParse, tag, attributes)
	if (indexArray == nil) {
		return []IntPair{{0,0}} // really return nil!
	}
	return indexArray
}

func returnIndexArray(stringToParse string, searchRegExp string) [][]int {
	regexpIndexFinder := regexp.MustCompile(searchRegExp)
	indexArray := regexpIndexFinder.FindAllStringSubmatchIndex(stringToParse, -1)
	
	if  indexArray == nil { 
//		fmt.Printf("**********returnIndexArray nil**********\n")
		return nil 
	} else {
//		fmt.Printf("**********returnIndexArray:%d %d**********\n", len(indexArray), len(indexArray[0]))
		return indexArray
	}
}

func returnIndexPairs(stringToParse string, tag string, attributes string) []IntPair {
	startIndexArray := returnIndexArray(stringToParse, "(?is)<"+tag+ attributes+">")
	stopIndexArray  := returnIndexArray(stringToParse, "(?is)</" + tag + ">")
	if len(startIndexArray) != len(stopIndexArray) {
		return nil
	}

//	fmt.Printf("******************* returnIndexPairs: %d %s  (Start: %d / End: %d)\n", len(startIndexArray), stringToParse, startIndexArray[0][1], stopIndexArray[0][0])

	var indexPairs []IntPair = make([]IntPair, len(startIndexArray))
	for i := 0; i < len(startIndexArray); i++ {
		indexPairs[i].start = startIndexArray[i][1]
		indexPairs[i].stop  = stopIndexArray[i][0]
	}
	return indexPairs
}

func trimTagsFromArray(indexArray []IntPair, stringToTrim string, tagNameLen int) []string{
	parsedSubMatchedTexts := make([]string, len(indexArray))
	for i := 0; i < len(indexArray); i++ {
		if (indexArray[i].stop - indexArray[i].start) >= 0 {
			parsedSubMatchedTexts[i] = stringToTrim[indexArray[i].start:indexArray[i].stop] //+tagLen
		} else {
			parsedSubMatchedTexts[i] = ""
		}
//		fmt.Printf("*****************ParsedSubMatch: %s\n", parsedSubMatchedTexts[i])
	}	
	return parsedSubMatchedTexts
}