package hownow_test

import (
	"testing"
	"time"

	hownow "github.com/ashmckenzie/go-hownow/hownow"
	"github.com/stretchr/testify/assert"
)

var pastDateTime = time.Date(2017, 6, 18, 17, 51, 49, 123456789, time.UTC)
var pointInTimeInPast = hownow.New(pastDateTime, -3)

var currentDateTime = time.Date(2017, 6, 17, 17, 51, 49, 12345678, time.UTC)
var pointInTime = hownow.New(currentDateTime, 0)

var futureDateTime = time.Date(2017, 6, 18, 17, 51, 49, 123456789, time.UTC)
var pointInTimeInFuture = hownow.New(futureDateTime, 3)

func format(p hownow.PointInTime, t time.Time) string {
	return p.Format(t, hownow.DefaultTimeFormat, false)
}

func formatEpoch(p hownow.PointInTime, t time.Time) string {
	return p.Format(t, "", true)
}

// func TestFormatFutureOffsetBOD(t *testing.T) {
// 	assert.Equal(t, "2017-06-21 00:00:00 +0000 UTC", pointInTimeInFuture.Process("bod"))
// 	assert.Equal(t, "1498003200", pointInTimeInFuture.Process("bod"))
// }

// func TestFormatFutureOffsetEOD(t *testing.T) {
// 	assert.Equal(t, "2017-06-21 23:59:59.999999999 +0000 UTC", pointInTimeInFuture.Process("eod"))
// 	assert.Equal(t, "1498089599", pointInTimeInFuture.Process("eod"))
// }

func TestFormatFutureOffsetBOW(t *testing.T) {
	ti := pointInTimeInFuture.Process("bow")
	assert.Equal(t, "2017-06-19 00:00:00 +0000 UTC", format(pointInTimeInFuture, ti))
	assert.Equal(t, "1497830400", formatEpoch(pointInTimeInFuture, ti))
}

func TestFormatFutureOffsetEOW(t *testing.T) {
	ti := pointInTimeInFuture.Process("eow")
	assert.Equal(t, "2017-06-25 23:59:59.999999999 +0000 UTC", format(pointInTimeInFuture, ti))
	assert.Equal(t, "1498435199", formatEpoch(pointInTimeInFuture, ti))
}

func TestFormatFutureOffsetBOM(t *testing.T) {
	assert.Equal(t, "2017-06-01 00:00:00 +0000 UTC", pointInTimeInFuture.Process("bom"))
	assert.Equal(t, "1496275200", pointInTimeInFuture.Process("bom"))
}

// func TestFormatFutureOffsetEOM(t *testing.T) {
// 	assert.Equal(t, "2017-06-30 23:59:59.999999999 +0000 UTC", pointInTimeInFuture.Process("eom"))
// 	assert.Equal(t, "1498867199", pointInTimeInFuture.Process("eom"))
// }

// func TestFormatBOD(t *testing.T) {
// 	assert.Equal(t, "2017-06-17 00:00:00 +0000 UTC", pointInTime.Process("bod"))
// 	assert.Equal(t, "1497657600", pointInTime.Process("bod"))
// }

// func TestFormatEOD(t *testing.T) {
// 	assert.Equal(t, "2017-06-17 23:59:59.999999999 +0000 UTC", pointInTime.Process("eod"))
// 	assert.Equal(t, "1497743999", pointInTime.Process("eod"))
// }

// func TestFormatBOW(t *testing.T) {
// 	assert.Equal(t, "2017-06-12 00:00:00 +0000 UTC", pointInTime.Process("bow"))
// 	assert.Equal(t, "1497225600", pointInTime.Process("bow"))
// }

// func TestFormatEOW(t *testing.T) {
// 	assert.Equal(t, "2017-06-18 23:59:59.999999999 +0000 UTC", pointInTime.Process("eow"))
// 	assert.Equal(t, "1497830399", pointInTime.Process("eow"))
// }

// func TestFormatBOM(t *testing.T) {
// 	assert.Equal(t, "2017-06-01 00:00:00 +0000 UTC", pointInTime.Process("bom"))
// 	assert.Equal(t, "1496275200", pointInTime.Process("bom"))
// }

// func TestFormatEOM(t *testing.T) {
// 	assert.Equal(t, "2017-06-30 23:59:59.999999999 +0000 UTC", pointInTime.Process("eom"))
// 	assert.Equal(t, "1498867199", pointInTime.Process("eom"))
// }

// func TestFormatPastOffsetBOD(t *testing.T) {
// 	assert.Equal(t, "2017-06-15 00:00:00 +0000 UTC", pointInTimeInPast.Process("bod"))
// 	assert.Equal(t, "1497484800", pointInTimeInPast.Process("bod"))
// }

// func TestFormatPastOffsetEOD(t *testing.T) {
// 	assert.Equal(t, "2017-06-15 23:59:59.999999999 +0000 UTC", pointInTimeInPast.Process("eod"))
// 	assert.Equal(t, "1497571199", pointInTimeInPast.Process("eod"))
// }

// func TestFormatPastOffsetBOW(t *testing.T) {
// 	assert.Equal(t, "2017-06-12 00:00:00 +0000 UTC", pointInTimeInPast.Process("bow"))
// 	assert.Equal(t, "1497225600", pointInTimeInPast.Process("bow"))
// }

// func TestFormatPastOffsetEOW(t *testing.T) {
// 	assert.Equal(t, "2017-06-18 23:59:59.999999999 +0000 UTC", pointInTimeInPast.Process("eow"))
// 	assert.Equal(t, "1497830399", pointInTimeInPast.Process("eow"))
// }

// func TestFormatPastOffsetBOM(t *testing.T) {
// 	assert.Equal(t, "2017-06-01 00:00:00 +0000 UTC", pointInTimeInPast.Process("bom"))
// 	assert.Equal(t, "1496275200", pointInTimeInPast.Process("bom"))
// }

// func TestFormatPastOffsetEOM(t *testing.T) {
// 	assert.Equal(t, "2017-06-30 23:59:59.999999999 +0000 UTC", pointInTimeInPast.Process("eom"))
// 	assert.Equal(t, "1498867199", pointInTimeInPast.Process("eom"))
// }
