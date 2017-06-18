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
	assert.Equal(t, pointInTime.Process("bod", false), "2013-11-18 00:00:00 +0000 UTC", "bod")
	assert.Equal(t, pointInTime.Process("bod", true), "1384732800", "bod epoch")
}

func TestProcessBOW(t *testing.T) {
	assert.Equal(t, pointInTime.Process("bow", false), "2013-11-17 00:00:00 +0000 UTC", "bow")
	assert.Equal(t, pointInTime.Process("bow", true), "1384646400", "bow epoch")
}

func TestProcessEOD(t *testing.T) {
	assert.Equal(t, pointInTime.Process("eod", false), "2013-11-18 23:59:59.999999999 +0000 UTC", "eod")
	assert.Equal(t, pointInTime.Process("eod", true), "1384819199", "eod epoch")
}

func TestProcessEOW(t *testing.T) {
	assert.Equal(t, pointInTime.Process("eow", false), "2013-11-23 23:59:59.999999999 +0000 UTC", "eow")
	assert.Equal(t, pointInTime.Process("eow", true), "1385251199", "eow epoch")
}
