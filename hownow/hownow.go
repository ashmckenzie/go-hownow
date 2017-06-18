package main

import (
	"strconv"
	"time"

	"github.com/jinzhu/now"
)

type PointInTime struct {
	Offset int
	Now    *now.Now
}

func New(t time.Time, offset int) *PointInTime {
	n := t.AddDate(0, 0, offset)
	return &PointInTime{Offset: offset, Now: now.New(n)}
}

func (self *PointInTime) Process(command string, epoch bool) string {
	var theTime time.Time

	switch command {
	case "bod":
		theTime = self.Now.BeginningOfDay()

	case "bow":
		theTime = self.Now.BeginningOfWeek()

	case "eod":
		theTime = self.Now.EndOfDay()

	case "eow":
		theTime = self.Now.EndOfWeek()
	}

	return toString(theTime, epoch)
}

func toString(in time.Time, epoch bool) string {
	var out string

	if epoch {
		out = strconv.FormatInt(in.Unix(), 10)
	} else {
		out = in.String()
	}

	return out
}
