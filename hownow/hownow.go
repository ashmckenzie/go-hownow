package hownow

import (
	"strconv"
	"time"

	"github.com/jinzhu/now"
)

// DefaultTimeFormat ...
const DefaultTimeFormat = "2006-01-02 15:04:05.999999999 -0700 MST"

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
func (p PointInTime) Process(command string) time.Time {
	var t time.Time

	switch command {
	case "bod":
		t = p.Now.BeginningOfDay()
	case "eod":
		t = p.Now.EndOfDay()
	case "bow":
		t = p.Now.BeginningOfWeek()
	case "eow":
		t = p.Now.EndOfWeek()
	case "bom":
		t = p.Now.BeginningOfMonth()
	case "eom":
		t = p.Now.EndOfMonth()
	}

	return t
}

// Format ...
func (p PointInTime) Format(in time.Time, format string, epoch bool) string {
	var out string

	if epoch {
		out = strconv.FormatInt(in.Unix(), 10)
	} else {
		out = in.Format(format)
	}

	return out
}
