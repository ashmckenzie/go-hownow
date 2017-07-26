package hownow

import (
	"strconv"
	"time"

	"github.com/jinzhu/now"
)

// PointInTime ...
type PointInTime struct {
	Offset int
	Now    *now.Now
}

// New ...
func New(t time.Time, offset int) PointInTime {
	n := t.AddDate(0, 0, offset)
	now.FirstDayMonday = true
	return PointInTime{Offset: offset, Now: now.New(n)}
}

// Process ...
func (p PointInTime) Process(command string, epoch bool) string {
	var theTime time.Time

	switch command {
	case "bod":
		theTime = p.Now.BeginningOfDay()
	case "eod":
		theTime = p.Now.EndOfDay()
	case "bow":
		theTime = p.Now.BeginningOfWeek()
	case "eow":
		theTime = p.Now.EndOfWeek()
	case "bom":
		theTime = p.Now.BeginningOfMonth()
	case "eom":
		theTime = p.Now.EndOfMonth()
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
