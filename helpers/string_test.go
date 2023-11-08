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
	"width = 35": {
		input:    "Elastic Compute Cloud (EC2) supports encryption at rest when using the Elastic Block Store (EBS) service.\n While disabled by default, forcing encryption at EBS volume creation is supported.",
		width:    35,
		expected: "Elastic Compute Cloud (EC2) suppor…\n While disabled by default, forcin…",
	},
	"width = 30": {
		input:    "Elastic Compute Cloud (EC2) supports encryption at rest when using the Elastic Block Store (EBS) service.\n While disabled by default.",
		width:    30,
		expected: "Elastic Compute Cloud (EC2) s…\n While disabled by default.",
	},
	"width = 29": {
		input:    "Elastic\n Compute Cloud (EC2)\n supports encryption\n at rest when using\n the Elastic Block Store\n (EBS) service.\n While disabled by default.",
		width:    29,
		expected: "Elastic\n Compute Cloud (EC2)\n supports encryption\n at rest when using\n the Elastic Block Store\n (EBS) service.\n While disabled by default.",
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
		input:    "[38;5;32myo,[0m [38;5;82m世界![0m",
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

type trimBlankLinesTest struct {
	input    string
	expected string
}

var testCasesTrimBlankLines = map[string]trimBlankLinesTest{
	"no blanklines": {
		input:    "foo bar",
		expected: "foo bar",
	},
	"single blanklines": {
		input:    "foo\n\nbar\n\nfoobar",
		expected: "foo\nbar\nfoobar",
	},
	"multiple blanklines": {
		input:    "foo\n\n\n\n\n\n\n\n\n\nbar\n\nfoobar",
		expected: "foo\nbar\nfoobar",
	},
}

func TestTrimBlankLines(t *testing.T) {
	for name, test := range testCasesTrimBlankLines {
		output := TrimBlankLines(test.input)
		if output != test.expected {
			t.Errorf("Test: '%s'' FAILED : \nexpected:\n %s \ngot:\n %s\n", name, test.expected, output)
		}
	}
}
