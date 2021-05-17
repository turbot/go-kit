package helpers

import (
	"testing"
)

type truncateTest struct {
	input    string
	width    int
	expected string
}

var testCasesTruncate = map[string]truncateTest{
	"width > string length": {
		input:    "foo bar",
		width:    10,
		expected: "foo bar",
	},
	"width == string length": {
		input:    "foo bar",
		width:    7,
		expected: "foo bar",
	},
	"width = string length-1": {
		input:    "foo bar",
		width:    6,
		expected: "foo b…",
	},
	"width = 2": {
		input:    "foo bar",
		width:    2,
		expected: "f…",
	},
	"width = 1": {
		input: "foo bar",
		width: 1,
		// return an empty string for 1 char
		expected: "",
	},
	"unicode - width > string length": {
		input:    "foo bar",
		width:    10,
		expected: "foo bar",
	},
	"unicode - width == string length": {
		input:    "yo, 世界!",
		width:    7,
		expected: "yo, 世界!",
	},
	"unicode - width = string length-1": {
		input:    "yo, 世界!",
		width:    6,
		expected: "yo, 世…",
	},
	"unicode - width = 2": {
		input:    "yo, 世界!",
		width:    2,
		expected: "y…",
	},
	"unicode - width = 1": {
		input: "yo, 世界!",
		width: 1,
		// return an empty string for 1 char
		expected: "",
	},
}

func TestTruncate(t *testing.T) {
	for name, test := range testCasesTruncate {
		output := TruncateString(test.input, test.width)
		if output != test.expected {
			t.Errorf("Test: '%s'' FAILED : \nexpected:\n %s \ngot:\n %s\n", name, test.expected, output)
		}
	}
}

type cleanTest struct {
	input    string
	expected string
}

var testCasesClean = map[string]cleanTest{
	"unicode, colour codes - width == string length": {
		input: "[38;5;32myo,[0m [38;5;82m世界![0m",
		expected: "yo, 世界!",
	},
}

func TestClean(t *testing.T) {
	for name, test := range testCasesClean {
		output := Clean(test.input)
		if output != test.expected {
			t.Errorf("Test: '%s'' FAILED : \nexpected:\n %s \ngot:\n %s\n", name, test.expected, output)
		}
	}
}
