package jiraRegEx

	
import (
//    "bufio"
    "fmt"
//    "io"
    "io/ioutil"
    "strconv"
    "regexp"
    "testing"
    . "gopkg.in/check.v1"
)

func (s *ReadInDataTestEngine) TestParsingForTime(c *C) {
    var endTag string = "h"
	regexpForValue := regexp.MustCompile("(?is)[0-9]+"+endTag)
    valueFromRegExp := regexpForValue.FindStringSubmatch("23w37h30m32112h")
    if valueFromRegExp == nil {
        fmt.Printf("*** DEBUG: valueFromRegExp returned to nil\n")
    }
    
    match := valueFromRegExp[0]
    match = match[0:len(match)-1]
    value, err := strconv.Atoi(match)
    if err != nil {
        panic(err)
    }
    c.Assert(value, Equals, 37)
 }

// Hook up gocheck into the "go test" runner.
type ReadInDataTestEngine struct{}
func TestReadInData(t *testing.T) { 
	Suite(&ReadInDataTestEngine{})
	TestingT(t) 
}
//var _ = 

func (s *ReadInDataTestEngine) checkForError(c *C, e error) {
	c.Assert(e, Equals, nil)
}

func (s *ReadInDataTestEngine) TestReadingInWholeFile(c *C) {
	data, err := ioutil.ReadFile("./testdata/Report-Jira.html")
    s.checkForError(c, err)
    c.Assert(len(data), Equals, 172252)
}

var regexpToFindNames string = "(?is)<tr>.*[<td [[:ascii:]]*>[[:alpha:]]*\\.[[:alpha:]]*</td>]*[[:space:]]*<td class=\"main\"><b>Total</b></td>[[:space:]]*</tr>"
var regexpToFindTotalTimes string = "(?is)<tr>([[:space:]]*<td class=\"total\"><b>(([0-9]*[wdhms]{1})+|total|Issue)</b></td>)+[[:space:]]*</tr>"

func (s *ReadInDataTestEngine) TestReadInAndSubMatchForNames(c *C) {
	s.ReadInFileAndFindRegExp(c, regexpToFindNames, 1)
}

func (s *ReadInDataTestEngine) TestReadInAndSubMatchForTotalTimes(c *C) {
	s.ReadInFileAndFindRegExp(c, regexpToFindTotalTimes, 2) //occures twice because also at the end of the table!
}

func (s *ReadInDataTestEngine) ReadInFileAndFindRegExp(c *C, regexpToFind string, countToAssertOn int) {
	data, err := ioutil.ReadFile("./testdata/Report-Jira.html")
    s.checkForError(c, err)
    s.AssertOnSubmatch(c, regexpToFind, countToAssertOn, string(data))
}


func (s *ReadInDataTestEngine) Test1TimeInDataString(c *C) {
	s.AssertOnSubmatch(c, "[0-9]*[wdhms]{1}", 1, "<>23h<<")
}


func (s *ReadInDataTestEngine) Test3TimesInDataString(c *C) {
	s.AssertOnSubmatch(c, "[[0-9]*[wdhms]{1}]*", 3, "<>23h<<22m oqp89w<")
}


func (s *ReadInDataTestEngine) Test3TimesTimeAndTotalAndIssueInDataString(c *C) {
	s.AssertOnSubmatch(c, "(?is)[[0-9]*[wdhms]{1}|total|Issue]*", 5, "<>23h<total<22m oissueqp89w<")
}


func (s *ReadInDataTestEngine) TestChangableWithinAndSurroundings(c *C) {
	s.AssertOnSubmatch(c, "(?is)<t>([0-9]*[wdhms]{1}|total|Issue)</t>", 1, "<t>23h</t><total<22m oissueqp89w<")
}

func (s *ReadInDataTestEngine) TestParseForTagSurroundingChangableValueShouldReturn5hits(c *C) {
	s.AssertOnSubmatch(c, "(?is)(<t>([0-9]*[wdhms]{1}|total|Issue)</t>)+", 5, "<t>23h</t><lkslakl<t>total</t><daklskdlfak<t>22m</t> o<t>issue</t>qp<t>89w</t>")
}

//d class=\"total\"><b>
func (s *ReadInDataTestEngine) TestParseForTagSimilarToJiraSurroundingChangableValueShouldReturn5hits(c *C) {
	s.AssertOnSubmatch(c, "(?is)(<td class=\"total\"><b>([0-9]*[wdhms]{1}|total|Issue)</b></td>)+", 5, "<td class=\"total\"><b>23h</b></td><lkslakl<td class=\"total\"><b>total</b></td><daklskdlfak<td class=\"total\"><b>22m</b></td> o<td class=\"total\"><b>issue</b></td>qp<td class=\"total\"><b>89w</b></td>")
}



var testStringForTotalTimes string = "daskjfdslkjfaldjfladklfalkdsfjlksafdsajf"+
"<tr>\n"+
"    <td class=\"total\"><b>Issue</b></td>\n"+
"    <td class=\"total\"><b>Total</b></td>\n"+
"                <td class=\"total\"><b>1h</b></td>\n"+
"                    <td class=\"total\"><b>34h</b></td>\n"+
"                    <td class=\"total\"><b>13h40m</b></td>\n"+
"                    <td class=\"total\"><b>44h</b></td>\n"+
"                    <td class=\"total\"><b>23h30m</b></td>\n"+
"                    <td class=\"total\"><b>10h50m</b></td>\n"+
"                    <td class=\"total\"><b>6h</b></td>\n"+
"                    <td class=\"total\"><b>37h30m</b></td>\n"+
"                    <td class=\"total\"><b>1h35m</b></td>\n"+
"                <td class=\"total\"><b>172h5m</b></td>\n"+
"</tr>   "+
"      sakjdlfajsdlkfjdsalfjdsafalkdsjflkas"
var regexpToFindTotalTimesWorking string = "(?is)<tr>([[:space:]]*<td class=\"total\"><b>(([0-9]*[wdhms]{1})+|total|Issue)</b></td>)+[[:space:]]*</tr>"//"(?is)([[:space:]]*<td class=\"total\"><b>(([0-9]*[wdhms]{1})+|total|Issue)</b></td>)+" 
func (s *ReadInDataTestEngine) TestParseForJiraTotalTimeStream(c *C) {
	s.AssertOnSubmatch(c, regexpToFindTotalTimesWorking, 1, testStringForTotalTimes)	
}


func (s *ReadInDataTestEngine) AssertOnSubmatch(c *C, regexpToFind string, countToAssertOn int, dataString string) {
    regexpForMatchingNames := regexp.MustCompile(regexpToFind)//<td class=\"main\">Total</td>"
    indexArray := regexpForMatchingNames.FindAllStringSubmatchIndex(dataString, -1)
    var countOfFoundSubmatches int = 0
    if (indexArray != nil) {
    	countOfFoundSubmatches = len(indexArray)
    	/*for countOfFoundSubmatches = 0; countOfFoundSubmatches < len(indexArray); countOfFoundSubmatches++ {
    		for k := 0; k < len(indexArray[countOfFoundSubmatches]); k++ {
    			//fmt.Printf("Index Array[%d][%d]: %d\n", countOfFoundSubmatches, k, indexArray[countOfFoundSubmatches][k])
    		}
    		
    		if ((indexArray[countOfFoundSubmatches][1] - indexArray[countOfFoundSubmatches][0]) < 1000) {
    			subMatchLen = indexArray[countOfFoundSubmatches][1] - indexArray[countOfFoundSubmatches][0]
    		} 
    		submatch = dataString[indexArray[countOfFoundSubmatches][0]:indexArray[countOfFoundSubmatches][1]]
    		//submatch := dataString[indexArray[i][0]:IndexArray[countOfFoundSubmatches][0]+countToAssertOn]
    	//	fmt.Printf("Index Array (\"%s\")[%d]: (%s)\n\n", regexpToFind, countOfFoundSubmatches, submatch)
    	}*/
    }	
    c.Assert(countOfFoundSubmatches, Equals, countToAssertOn)
}