package personaltime

import (
	"regexp"
	"strconv"
)

type PersonalTime struct {
    name string
    weeks, days, hours, mins, secs int
}

func (p *PersonalTime) Initialize(name string, time string) {
    p.name = name
    p.InitializeTime(time)
}

func (p *PersonalTime) InitializeTime(time string) {
    p.weeks = ParseForInteger(time, "w")
    p.days = ParseForInteger(time, "d")
    p.hours = ParseForInteger(time, "h")
    p.mins = ParseForInteger(time, "m")
    p.secs = ParseForInteger(time, "s")
}

func ParseForInteger(time string, timeIdentifier string) int {
    regexpForValue := regexp.MustCompile("(?is)[0-9]+"+ timeIdentifier)
    valueFromRegExp := regexpForValue.FindStringSubmatch(time)
    if valueFromRegExp == nil {
        //fmt.Printf("*** DEBUG: valueFromRegExp returned to nil(%s/%s)\n", time, timeIdentifier)
        return 0
    }
    
    match := valueFromRegExp[0]
    match = match[0:len(match)-1]
    value, err := strconv.Atoi(match)
    
    if err != nil {
        return 0
    } else {
        return value
    }
}

func (p *PersonalTime) ToString() string {  
    //fmt.Printf("%s: %dw/%dd/%dh/%dm/%ds\n", p.name, p.weeks, p.days, p.hours, p.mins, p.secs)
    return "***" + p.name   + " : "+p.toStringTimes(p.weeks, "week(s)") +
                              " : "+p.toStringTimes(p.days, "day(s)") +
                              " : "+p.toStringTimes(p.hours, "hour(s)")+ 
                              " : "+p.toStringTimes(p.mins, "min(s)")+
                              " : "+p.toStringTimes(p.secs, "sec(s)")
}

func (p *PersonalTime) toStringTimes(time int, name string) string {
    var timeStringBuffer string = ""
    if (time > 0) {
        timeStringBuffer = timeStringBuffer + strconv.Itoa(time) + " " + name
    }
    return timeStringBuffer
}