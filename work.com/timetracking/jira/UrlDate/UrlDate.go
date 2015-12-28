package jira

import (
	. "../../helper"
	"time"
)

type UrlDate struct {
	t time.Time
}

func (this UrlDate) GetTime() time.Time {
	return this.t
}

func (this *UrlDate) Initialize(date string) {
	if len(date) == 0 {
		this.t = time.Now()
		return
	}
	tmp, err := time.Parse("02.01.2006", date)
	PanicOnError(err)
	this.t = tmp
}

func (this *UrlDate) GetTimeForUrl() string {
	urldate := this.t.Format("2") + "%2F" + this.t.Format("Jan") + "%2F" + this.t.Format("06")
	return urldate
}
