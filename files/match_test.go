package files

import (
	"slices"
	"testing"
)

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
	"all hidden folders": {
		pattern:  "/Users/kai/Dev/github/turbot/steampipe/pkg/workspace/test_data/dependent_mod/.*",
		file:     "/Users/kai/Dev/github/turbot/steampipe/pkg/workspace/test_data/dependent_mod/.steampipe",
		expected: true,
	},
	"all nested hidden folders": {
		pattern:  "/Users/kai/Dev/github/turbot/steampipe/pkg/workspace/test_data/dependent_mod/**/.*",
		file:     "/Users/kai/Dev/github/turbot/steampipe/pkg/workspace/test_data/dependent_mod/foo/.steampipe",
		expected: true,
	},
	"all nested hidden folders 2": {
		pattern:  "/Users/kai/Dev/github/turbot/steampipe/pkg/workspace/test_data/dependent_mod/**/.*",
		file:     "/Users/kai/Dev/github/turbot/steampipe/pkg/workspace/test_data/dependent_mod/foo/bar/.steampipe",
		expected: true,
	},
	"all nested hidden folders under a specific path": {
		pattern:  "/Users/kai/Dev/github/turbot/steampipe/pkg/workspace/test_data/dependent_mod/**/a/.*",
		file:     "/Users/kai/Dev/github/turbot/steampipe/pkg/workspace/test_data/dependent_mod/foo/bar/a/.steampipe",
		expected: true,
	},
	"all nested hidden folders under a specific path (fails)": {
		pattern:  "/Users/kai/Dev/github/turbot/steampipe/pkg/workspace/test_data/dependent_mod/**/a/.*",
		file:     "/Users/kai/Dev/github/turbot/steampipe/pkg/workspace/test_data/dependent_mod/foo/bar/a/.steampipe",
		expected: true,
	},
	"everything in a hidden folder": {
		pattern:  "/Users/kai/Dev/github/turbot/steampipe/pkg/workspace/test_data/dependent_mod/.*/**",
		file:     "/Users/kai/Dev/github/turbot/steampipe/pkg/workspace/test_data/dependent_mod/.steampipe/mods/github.com/pskrbasu/steampipe-mod-m1@v4.0/dashboard.sp",
		expected: true,
	},
	"a recursive wildcard": {
		pattern:  "**/*.spc",
		file:     "a/b/c/bar.spc",
		expected: true,
	},
	"a recursive wildcard under a specific path": {
		pattern:  "**/a/b/c/*.spc",
		file:     "foo/a/b/c/bar.spc",
		expected: true,
	},
	"a recursive wildcard under a specific path (fails)": {
		pattern:  "**/a/b/c/*.spc",
		file:     "foo/a/b/bar.spc",
		expected: false,
	},
	"a recursive wildcard under a specific path (fails 2)": {
		pattern:  "**/a/b/c/*.spc",
		file:     "b/c/bar.spc",
		expected: false,
	},
	"everything in a hidden folder (fails)": {
		pattern:  "/Users/kai/Dev/github/turbot/steampipe/pkg/workspace/test_data/dependent_mod/.*/**",
		file:     "/Users/kai/Dev/github/turbot/steampipe/pkg/workspace/test_data/dependent_mod/dashboard.sp",
		expected: false,
	},
	// NOTE: when globbing on command line this would fail
	// however according to https://metacpan.org/dist/File-Globstar/view/lib/File/Globstar.pod#**
	// this _should_ include files at top level
	"all child folder sp files - top level file": {
		pattern:  "testdata/mods/single_mod_one_query/**/*.sp",
		file:     "testdata/mods/single_mod_one_query/query.sp",
		expected: true,
	},
	"all child folder sp files with wildcard in path": {
		pattern:  "testdata/m*/**/*.sp",
		file:     "testdata/mods/single_mod_one_query/query.sp",
		expected: true,
	},
}

func TestMatch(t *testing.T) {
	var testNames = []string{}
	for name, test := range matchTests {
		if len(testNames) > 0 && !slices.Contains(testNames, name) {
			continue
		}
		actual := Match(test.pattern, test.file)

		if actual != test.expected {
			t.Errorf("Test: '%s'' FAILED : expected:\n\n%v\n\ngot:\n\n%v", name, test.expected, actual)
		}
	}
}
