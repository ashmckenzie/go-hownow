package main_test

import (
	"testing"
	"time"

	hownow "github.com/ashmckenzie/go-hownow/hownow"
	"github.com/stretchr/testify/assert"
)

var currentDateTime = time.Date(2013, 11, 18, 17, 51, 49, 123456789, time.UTC)
var pointInTime = hownow.New(currentDateTime, 0)

func TestProcessBOD(t *testing.T) {
	assert.Equal(t, "2013-11-18 00:00:00 +0000 UTC", pointInTime.Process("bod", false), "bod")
	assert.Equal(t, "1384732800", pointInTime.Process("bod", true), "bod epoch")
}

func TestProcessBOW(t *testing.T) {
	assert.Equal(t, "2013-11-17 00:00:00 +0000 UTC", pointInTime.Process("bow", false), "bow")
	assert.Equal(t, "1384646400", pointInTime.Process("bow", true), "bow epoch")
}

func TestProcessEOD(t *testing.T) {
	assert.Equal(t, "2013-11-18 23:59:59.999999999 +0000 UTC", pointInTime.Process("eod", false), "eod")
	assert.Equal(t, "1384819199", pointInTime.Process("eod", true), "eod epoch")
}

func TestProcessEOW(t *testing.T) {
	assert.Equal(t, "2013-11-23 23:59:59.999999999 +0000 UTC", pointInTime.Process("eow", false), "eow")
	assert.Equal(t, "1385251199", pointInTime.Process("eow", true), "eow epoch")
}
