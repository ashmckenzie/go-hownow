package main_test

import (
	"testing"
	"time"

	hownow "github.com/ashmckenzie/go-hownow/hownow"
	"github.com/stretchr/testify/assert"
)

var currentDateTime = time.Date(2017, 6, 17, 17, 51, 49, 12345678, time.UTC)
var pointInTime = hownow.New(currentDateTime, 0)

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
