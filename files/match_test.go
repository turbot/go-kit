package files

import "testing"

type matchTestCase struct {
	pattern  string
	file     string
	expected bool
}

var matchTests = map[string]matchTestCase{
	"dirname": {
		pattern:  "*/foo",
		file:     "a/foo",
		expected: true,
	},
	"wildcard dir, filename": {
		pattern:  "*/foo.spc",
		file:     "a/foo.spc",
		expected: true,
	},
	"wildcard dir, filename mismatch": {
		pattern:  "*/foo.spc",
		file:     "a/bar.spc",
		expected: false,
	},
	"wildcard dir, wildcard filename 1": {
		pattern:  "*/*.spc",
		file:     "a/foo.spc",
		expected: true,
	},
	"wildcard dir, wildcard filename 2": {
		pattern:  "*/*.spc",
		file:     "a/bar.spc",
		expected: true,
	},
	"wildcard dir, wildcard filename, nested file": {
		pattern:  "*/*.spc",
		file:     "a/b/bar.spc",
		expected: false,
	},
	"wildcard recursive dir, wildcard filename, nested file": {
		pattern:  "a/**/*.spc",
		file:     "a/b/bar.spc",
		expected: true,
	},
	"wildcard recursive dir, wildcard filename, deeply nested file": {
		pattern:  "a/**/*.spc",
		file:     "a/b/c/bar.spc",
		expected: true,
	},
	"wildcard recursive dir, wildcard filename with defined parents, deeply nested file - match": {
		pattern:  "a/**/c/*.spc",
		file:     "a/b/c/bar.spc",
		expected: true,
	},
}

func TestMatch(t *testing.T) {
	for name, test := range matchTests {
		actual := Match(test.pattern, test.file)

		if actual != test.expected {
			t.Errorf("Test: '%s'' FAILED : expected:\n\n%v\n\ngot:\n\n%v", name, test.expected, actual)
		}
	}
}
