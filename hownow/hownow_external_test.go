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

func TestProcessFutureOffsetBOD(t *testing.T) {
	assert.Equal(t, "2017-06-21 00:00:00 +0000 UTC", pointInTimeInFuture.Process("bod", false), "bod")
	assert.Equal(t, "1498003200", pointInTimeInFuture.Process("bod", true), "bod epoch")
}

func TestProcessFutureOffsetEOD(t *testing.T) {
	assert.Equal(t, "2017-06-21 23:59:59.999999999 +0000 UTC", pointInTimeInFuture.Process("eod", false), "eod")
	assert.Equal(t, "1498089599", pointInTimeInFuture.Process("eod", true), "eod epoch")
}

func TestProcessFutureOffsetBOW(t *testing.T) {
	assert.Equal(t, "2017-06-19 00:00:00 +0000 UTC", pointInTimeInFuture.Process("bow", false), "bow")
	assert.Equal(t, "1497830400", pointInTimeInFuture.Process("bow", true), "bow epoch")
}

func TestProcessFutureOffsetEOW(t *testing.T) {
	assert.Equal(t, "2017-06-25 23:59:59.999999999 +0000 UTC", pointInTimeInFuture.Process("eow", false), "eow")
	assert.Equal(t, "1498435199", pointInTimeInFuture.Process("eow", true), "eow epoch")
}

func TestProcessFutureOffsetBOM(t *testing.T) {
	assert.Equal(t, "2017-06-01 00:00:00 +0000 UTC", pointInTimeInFuture.Process("bom", false), "bow")
	assert.Equal(t, "1496275200", pointInTimeInFuture.Process("bom", true), "bom epoch")
}

func TestProcessFutureOffsetEOM(t *testing.T) {
	assert.Equal(t, "2017-06-30 23:59:59.999999999 +0000 UTC", pointInTimeInFuture.Process("eom", false), "eom")
	assert.Equal(t, "1498867199", pointInTimeInFuture.Process("eom", true), "eom epoch")
}

func TestProcessBOD(t *testing.T) {
	assert.Equal(t, "2017-06-17 00:00:00 +0000 UTC", pointInTime.Process("bod", false), "bod")
	assert.Equal(t, "1497657600", pointInTime.Process("bod", true), "bod epoch")
}

func TestProcessEOD(t *testing.T) {
	assert.Equal(t, "2017-06-17 23:59:59.999999999 +0000 UTC", pointInTime.Process("eod", false), "eod")
	assert.Equal(t, "1497743999", pointInTime.Process("eod", true), "eod epoch")
}

func TestProcessBOW(t *testing.T) {
	assert.Equal(t, "2017-06-12 00:00:00 +0000 UTC", pointInTime.Process("bow", false), "bow")
	assert.Equal(t, "1497225600", pointInTime.Process("bow", true), "bow epoch")
}

func TestProcessEOW(t *testing.T) {
	assert.Equal(t, "2017-06-18 23:59:59.999999999 +0000 UTC", pointInTime.Process("eow", false), "eow")
	assert.Equal(t, "1497830399", pointInTime.Process("eow", true), "eow epoch")
}

func TestProcessBOM(t *testing.T) {
	assert.Equal(t, "2017-06-01 00:00:00 +0000 UTC", pointInTime.Process("bom", false), "bom")
	assert.Equal(t, "1496275200", pointInTime.Process("bom", true), "bom epoch")
}

func TestProcessEOM(t *testing.T) {
	assert.Equal(t, "2017-06-30 23:59:59.999999999 +0000 UTC", pointInTime.Process("eom", false), "eom")
	assert.Equal(t, "1498867199", pointInTime.Process("eom", true), "eom epoch")
}

func TestProcessPastOffsetBOD(t *testing.T) {
	assert.Equal(t, "2017-06-15 00:00:00 +0000 UTC", pointInTimeInPast.Process("bod", false), "bod")
	assert.Equal(t, "1497484800", pointInTimeInPast.Process("bod", true), "bod epoch")
}

func TestProcessPastOffsetEOD(t *testing.T) {
	assert.Equal(t, "2017-06-15 23:59:59.999999999 +0000 UTC", pointInTimeInPast.Process("eod", false), "eod")
	assert.Equal(t, "1497571199", pointInTimeInPast.Process("eod", true), "eod epoch")
}

func TestProcessPastOffsetBOW(t *testing.T) {
	assert.Equal(t, "2017-06-12 00:00:00 +0000 UTC", pointInTimeInPast.Process("bow", false), "bow")
	assert.Equal(t, "1497225600", pointInTimeInPast.Process("bow", true), "bow epoch")
}

func TestProcessPastOffsetEOW(t *testing.T) {
	assert.Equal(t, "2017-06-18 23:59:59.999999999 +0000 UTC", pointInTimeInPast.Process("eow", false), "eow")
	assert.Equal(t, "1497830399", pointInTimeInPast.Process("eow", true), "eow epoch")
}

func TestProcessPastOffsetBOM(t *testing.T) {
	assert.Equal(t, "2017-06-01 00:00:00 +0000 UTC", pointInTimeInPast.Process("bom", false), "bom")
	assert.Equal(t, "1496275200", pointInTimeInPast.Process("bom", true), "bom epoch")
}

func TestProcessPastOffsetEOM(t *testing.T) {
	assert.Equal(t, "2017-06-30 23:59:59.999999999 +0000 UTC", pointInTimeInPast.Process("eom", false), "eom")
	assert.Equal(t, "1498867199", pointInTimeInPast.Process("eom", true), "eom epoch")
}
