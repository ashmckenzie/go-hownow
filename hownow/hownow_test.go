package hownow_test

import (
	"testing"
	"time"

	hownow "github.com/ashmckenzie/go-hownow/hownow"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	futureDateTime := time.Date(2017, 6, 18, 17, 51, 49, 123456789, time.UTC)
	pointInTimeInFuture := hownow.New(futureDateTime, 3)

	assert.IsType(t, hownow.PointInTime{}, pointInTimeInFuture)
	assert.Equal(t, 3, pointInTimeInFuture.Offset)
}
