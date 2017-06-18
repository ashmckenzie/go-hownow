package main_test

import (
	"testing"
	"time"

	hownow "github.com/ashmckenzie/go-hownow/hownow"
	"github.com/stretchr/testify/assert"
)

var futureDateTime = time.Date(2017, 6, 18, 17, 51, 49, 123456789, time.UTC)
var pointInTimeInFuture = hownow.New(futureDateTime, 3)

func TestProcessFutureOffsetBOD(t *testing.T) {
	assert.Equal(t, "2017-06-21 00:00:00 +0000 UTC", pointInTimeInFuture.Process("bod", false), "bod")
	assert.Equal(t, "1498003200", pointInTimeInFuture.Process("bod", true), "bod epoch")
}

func TestProcessFutureOffsetBOW(t *testing.T) {
	assert.Equal(t, "2017-06-19 00:00:00 +0000 UTC", pointInTimeInFuture.Process("bow", false), "bow")
	assert.Equal(t, "1497830400", pointInTimeInFuture.Process("bow", true), "bow epoch")
}

func TestProcessFutureOffsetEOD(t *testing.T) {
	assert.Equal(t, "2017-06-21 23:59:59.999999999 +0000 UTC", pointInTimeInFuture.Process("eod", false), "eod")
	assert.Equal(t, "1498089599", pointInTimeInFuture.Process("eod", true), "eod epoch")
}

func TestProcessFutureOffsetEOW(t *testing.T) {
	assert.Equal(t, "2017-06-25 23:59:59.999999999 +0000 UTC", pointInTimeInFuture.Process("eow", false), "eow")
	assert.Equal(t, "1498435199", pointInTimeInFuture.Process("eow", true), "eow epoch")
}
