package personaltime

import (
    "regexp"
    "strconv"
)

type PersonalTime struct {
    name                           string
    weeks, days, hours, mins, secs int
    overallTimeOfAllProjects       float64
}

func (p *PersonalTime) GetName() string {
    return p.name
}

func (p *PersonalTime) SetOverallTime(overallTime float64) {
    p.overallTimeOfAllProjects = overallTime
}

func (p *PersonalTime) InitializeFromString(name string, time string) {
    p.name = name
    p.InitializeTime(time)
}

func (p *PersonalTime) InitializeFromFloat(name string, timeInHours float64) {
    p.name = name
    p.hours = int(timeInHours)
    p.mins = int((timeInHours - float64(p.hours)) * 60)
}

func (p *PersonalTime) InitializeTime(time string) {
    p.weeks = ParseForInteger(time, "w")
    p.days = ParseForInteger(time, "d")
    p.hours = ParseForInteger(time, "h")
    p.mins = ParseForInteger(time, "m")
    p.secs = ParseForInteger(time, "s")
}

func ParseForInteger(time string, timeIdentifier string) int {
    regexpForValue := regexp.MustCompile("(?is)[0-9]+" + timeIdentifier)
    valueFromRegExp := regexpForValue.FindStringSubmatch(time)
    if valueFromRegExp == nil {
        //fmt.Printf("*** DEBUG: valueFromRegExp returned to nil(%s/%s)\n", time, timeIdentifier)
        return 0
    }

    match := valueFromRegExp[0]
    match = match[0 : len(match)-1]
    value, err := strconv.Atoi(match)

    if err != nil {
        return 0
    } else {
        return value
    }
}

func (p *PersonalTime) ToString() string {
    //fmt.Printf("%s: %dw/%dd/%dh/%dm/%ds\n", p.name, p.weeks, p.days, p.hours, p.mins, p.secs)
    return "***" + p.name + " : " + p.toStringTimes(p.weeks, "week(s)") +
        " : " + p.toStringTimes(p.days, "day(s)") +
        " : " + p.toStringTimes(p.hours, "hour(s)") +
        " : " + p.toStringTimes(p.mins, "min(s)") +
        " : " + p.toStringTimes(p.secs, "sec(s)")
}

func (p *PersonalTime) toStringTimes(time int, name string) string {
    var timeStringBuffer string = ""
    if time > 0 {
        timeStringBuffer = timeStringBuffer + strconv.Itoa(time) + " " + name
    }
    return timeStringBuffer
}

func (p *PersonalTime) ToFloat64InHours() float64 {
    dmins := float64(p.mins)
    mins := dmins / 60.0
    return float64(p.hours) + mins
}

func (p *PersonalTime) ToCsvFormat() string {
    time := strconv.FormatFloat(p.ToFloat64InHours(), 'f', 2, 64)
    var retVal string = p.name + "," + time
    if p.overallTimeOfAllProjects > 0.0 {
        percentOfOverallPrj := strconv.FormatFloat((100.0/p.overallTimeOfAllProjects)*p.ToFloat64InHours(), 'f', 1, 64)
        retVal = retVal + "," + percentOfOverallPrj + "%"
    }
    return retVal
}
