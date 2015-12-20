package TimeEntry

import (
	"regexp"
	"strconv"
	"strings"
)

type TimeEntry struct {
	name                           string
	weeks, days, hours, mins, secs int
	overallTimeOfAllProjects       float64
	participants                   []string
}

func (p *TimeEntry) GetName() string {
	return p.name
}

func (p *TimeEntry) SetOverallTime(overallTime float64) {
	p.overallTimeOfAllProjects = overallTime
}

func (p *TimeEntry) GetInPercent() float64 {
	return (100.0 / p.overallTimeOfAllProjects) * p.ToFloat64InHours()
}

func (p *TimeEntry) InitializeFromString(name string, time string) {
	p.name = name
	p.InitializeTime(time)
}

func (p *TimeEntry) InitializeFromFloat(name string, timeInHours float64, participants []string) {
	p.name = name
	p.hours = int(timeInHours)
	p.mins = int((timeInHours - float64(p.hours)) * 60)
	p.participants = participants
}

func (p *TimeEntry) InitializeTime(time string) {
	p.weeks = ParseForInteger(time, "w")
	p.days = ParseForInteger(time, "d")
	p.hours = ParseForInteger(time, "h")
	p.mins = ParseForInteger(time, "m")
	p.secs = ParseForInteger(time, "s")
}

func (this *TimeEntry) GetTeamMembersCommaSeperated(poToSkip string) string {
	var teamMembers string
	participantCount := len(this.participants)
	for i := range this.participants {
		nextTeamMember := this.participants[i]
		if strings.ToLower(nextTeamMember) != strings.ToLower(poToSkip) {
			teamMembers = teamMembers + nextTeamMember
			if i < (participantCount - 1) {
				teamMembers = teamMembers + ","
			}
		}
	}
	return teamMembers
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

func (p *TimeEntry) ToString() string {
	//fmt.Printf("%s: %dw/%dd/%dh/%dm/%ds\n", p.name, p.weeks, p.days, p.hours, p.mins, p.secs)
	return "***" + p.name + " : " + p.toStringTimes(p.weeks, "week(s)") +
		" : " + p.toStringTimes(p.days, "day(s)") +
		" : " + p.toStringTimes(p.hours, "hour(s)") +
		" : " + p.toStringTimes(p.mins, "min(s)") +
		" : " + p.toStringTimes(p.secs, "sec(s)")
}

func (p *TimeEntry) toStringTimes(time int, name string) string {
	var timeStringBuffer string = ""
	if time > 0 {
		timeStringBuffer = timeStringBuffer + strconv.Itoa(time) + " " + name
	}
	return timeStringBuffer
}

func (p *TimeEntry) ToFloat64InHours() float64 {
	dmins := float64(p.mins)
	mins := dmins / 60.0
	return float64(p.hours) + mins
}

func (p *TimeEntry) ToCsvFormat(seperator rune) string {
	time := strconv.FormatFloat(p.ToFloat64InHours(), 'f', 2, 64)
	var sumParticipants string
	for k := range p.participants {
		sumParticipants = sumParticipants + p.participants[k]
		if k < (len(p.participants) - 1) {
			sumParticipants = sumParticipants + ","
		}

	}
	var retVal string = sumParticipants + string(seperator) + p.name
	if p.overallTimeOfAllProjects > 0.0 {
		percentOfOverallPrj := strconv.FormatFloat(p.GetInPercent(), 'f', 1, 64)
		retVal = retVal + string(seperator) + percentOfOverallPrj + "%"
	}
	retVal = retVal + string(seperator) + time
	return retVal
}
