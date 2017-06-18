package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var s string

func TestParseBOD(t *testing.T) {
	assert.Regexp(t, `^\d{4}-\d\d-\d\d 00:00:00 \+\d{4} \w+$`, parse([]string{"bod"}), "parse bod")
}

func TestParseBOW(t *testing.T) {
	assert.Regexp(t, `^\d{4}-\d\d-\d\d 00:00:00 \+\d{4} \w+$`, parse([]string{"bow"}), "parse bow")
}

func TestParseEOD(t *testing.T) {
	assert.Regexp(t, `^\d{4}-\d\d-\d\d 23:59:59.999999999 \+\d{4} \w+$`, parse([]string{"eod"}), "parse eod")
}

func TestParseEOW(t *testing.T) {
	assert.Regexp(t, `^\d{4}-\d\d-\d\d 23:59:59.999999999 \+\d{4} \w+$`, parse([]string{"eow"}), "parse eow")
}

func TestFormatNoNewLine(t *testing.T) {
	assert.Equal(t, format("test", false), "test", "format")
}

func TestFormatNewLine(t *testing.T) {
	assert.Equal(t, format("test", true), "test\n", "format + new line")
}
