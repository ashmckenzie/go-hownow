package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var s string

func TestParseBOD(t *testing.T) {
	assert.Regexp(t, `^\d{4}-\d\d-\d\d 00:00:00 \+\d{4} \w+\n$`, parse([]string{"bod"}), "parse bod")
}

func TestParseEOD(t *testing.T) {
	assert.Regexp(t, `^\d{4}-\d\d-\d\d 23:59:59.999999999 \+\d{4} \w+\n$`, parse([]string{"eod"}), "parse eod")
}

func TestParseBOW(t *testing.T) {
	assert.Regexp(t, `^\d{4}-\d\d-\d\d 00:00:00 \+\d{4} \w+\n$`, parse([]string{"bow"}), "parse bow")
}

func TestParseEOW(t *testing.T) {
	assert.Regexp(t, `^\d{4}-\d\d-\d\d 23:59:59.999999999 \+\d{4} \w+\n$`, parse([]string{"eow"}), "parse eow")
}

func TestParseBOM(t *testing.T) {
	assert.Regexp(t, `^\d{4}-\d\d-\d\d 00:00:00 \+\d{4} \w+\n$`, parse([]string{"bom"}), "parse bom")
}

func TestParseEOM(t *testing.T) {
	assert.Regexp(t, `^\d{4}-\d\d-\d\d 23:59:59.999999999 \+\d{4} \w+\n$`, parse([]string{"eom"}), "parse eom")
}

func TestFormatNoNewLine(t *testing.T) {
	assert.Equal(t, "test", format("test", false), "format")
}

func TestFormatNewLine(t *testing.T) {
	assert.Equal(t, "test\n", format("test", true), "format + new line")
}
