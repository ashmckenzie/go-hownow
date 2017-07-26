package hownow_test

import (
	"testing"
	"time"

	hownow "github.com/ashmckenzie/go-hownow/hownow"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	p := hownow.New(time.Date(2017, 6, 18, 17, 51, 49, 123456789, time.UTC), 3)

	assert.IsType(t, hownow.PointInTime{}, p)
	assert.Equal(t, 3, p.Offset)
}

func TestProcess_bod(t *testing.T) {
	p := hownow.New(time.Date(2017, 6, 18, 17, 51, 49, 123456789, time.UTC), 3)
	ti := p.Process("bod")

	assert.IsType(t, time.Time{}, ti)
	assert.Equal(t, int64(1498003200), ti.Unix())
}

func TestFormat_epoch(t *testing.T) {
	p := hownow.New(time.Date(2017, 6, 18, 17, 51, 49, 123456789, time.UTC), 3)
	ti := p.Process("bod")
	s := p.Format(ti, "", true)

	assert.Equal(t, "1498003200", s)
}

func TestFormat_epoch_and_format(t *testing.T) {
	p := hownow.New(time.Date(2017, 6, 18, 17, 51, 49, 123456789, time.UTC), 3)
	ti := p.Process("bod")
	s := p.Format(ti, "invalid", true)

	assert.Equal(t, "1498003200", s)
}

func TestFormat_format_invalid(t *testing.T) {
	p := hownow.New(time.Date(2017, 6, 18, 17, 51, 49, 123456789, time.UTC), 3)
	ti := p.Process("bod")
	s := p.Format(ti, "invalid", false)

	assert.Equal(t, "invalid", s)
}

func TestFormat_format_valid(t *testing.T) {
	p := hownow.New(time.Date(2017, 6, 18, 17, 51, 49, 123456789, time.UTC), 3)
	ti := p.Process("bod")
	s := p.Format(ti, "start 02/01/2006 end", false)

	assert.Equal(t, "start 21/06/2017 end", s)
}
