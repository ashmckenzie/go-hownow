package main_test

import (
	"testing"
	"time"

	hownow "github.com/ashmckenzie/go-hownow/hownow"
	"github.com/stretchr/testify/assert"
)

var pastDateTime = time.Date(2017, 6, 18, 17, 51, 49, 123456789, time.UTC)
var pointInTimeInPast = hownow.New(pastDateTime, -3)

func TestProcessPastOffsetBOD(t *testing.T) {
	assert.Equal(t, "2017-06-15 00:00:00 +0000 UTC", pointInTimeInPast.Process("bod", false), "bod")
	assert.Equal(t, "1497484800", pointInTimeInPast.Process("bod", true), "bod epoch")
}

func TestProcessPastOffsetBOW(t *testing.T) {
	assert.Equal(t, "2017-06-11 00:00:00 +0000 UTC", pointInTimeInPast.Process("bow", false), "bow")
	assert.Equal(t, "1497139200", pointInTimeInPast.Process("bow", true), "bow epoch")
}

func TestProcessPastOffsetEOD(t *testing.T) {
	assert.Equal(t, "2017-06-15 23:59:59.999999999 +0000 UTC", pointInTimeInPast.Process("eod", false), "eod")
	assert.Equal(t, "1497571199", pointInTimeInPast.Process("eod", true), "eod epoch")
}

func TestProcessPastOffsetEOW(t *testing.T) {
	assert.Equal(t, "2017-06-17 23:59:59.999999999 +0000 UTC", pointInTimeInPast.Process("eow", false), "eow")
	assert.Equal(t, "1497743999", pointInTimeInPast.Process("eow", true), "eow epoch")
}
