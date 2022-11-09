package files

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

type SplitPathTestExpected struct {
	root string
	glob string
}

func (t SplitPathTestExpected) String() string {
	return fmt.Sprintf("{root: %s ; glob: %s}", t.root, t.glob)
}

type SplitPathTest struct {
	path        string
	expected    SplitPathTestExpected
	expectError bool
}

var wd, _ = os.Getwd()

var splitTests = map[string]SplitPathTest{
	"$PWD/test_data/list_test1/config/aws.spc": {
		path:     filepath.Join(wd, "test_data/list_test1/config/aws.spc"),
		expected: SplitPathTestExpected{root: filepath.Join(wd, "test_data/list_test1/config"), glob: filepath.Join(wd, "test_data/list_test1/config/aws.spc")},
	},
	"**/*.tf": {
		path:     "**/*.tf",
		expected: SplitPathTestExpected{root: wd, glob: filepath.Join(wd, "**/*.tf")},
	},
	"test_data/list_test1/config/*.spc": {
		path:     "test_data/list_test1/config/*.spc",
		expected: SplitPathTestExpected{root: filepath.Join(wd, "test_data/list_test1/config"), glob: filepath.Join(wd, "test_data/list_test1/config/*.spc")},
	},
	"./*.tf": {
		path:     "./*.tf",
		expected: SplitPathTestExpected{root: wd, glob: filepath.Join(wd, "*.tf")},
	},
	"./terr/**/*.tf": {
		path:     "./terr/**/*.tf",
		expected: SplitPathTestExpected{root: wd, glob: filepath.Join(wd, "terr/**/*.tf")},
	},
}

func TestSplitPath(t *testing.T) {
	for name, test := range splitTests {
		root, glob, err := GlobRoot(test.path)

		if err != nil && !test.expectError {
			t.Errorf("Test: '%s'' FAILED with unexpected error: %v", name, err)
		}

		actual := SplitPathTestExpected{
			root: root,
			glob: glob,
		}
		expected := test.expected
		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("Test: '%s'' FAILED : expected:\n\n%s\n\ngot:\n\n%s\n\n", name, expected, actual)
		}
	}
}

type ListFilesTest struct {
	source   string
	options  *ListOptions
	expected interface{}
}

var testCasesListFiles = map[string]ListFilesTest{
	"AllRecursive, exclude **/a*, **/*.swp, **/.steampipe*": {
		source: "test_data/list_test1",
		options: &ListOptions{
			Flags:   AllRecursive,
			Exclude: []string{"**/a*", "**/*.swp", "**/.steampipe*"},
		},
		expected: []string{
			"test_data/list_test1/b",
			"test_data/list_test1/b/mod.sp",
			"test_data/list_test1/b/q1.sp",
			"test_data/list_test1/b/q2.sp",
			"test_data/list_test1/config",
			"test_data/list_test1/config/default.spc",
		},
	},
	"AllRecursive": {
		source: "test_data/list_test1",
		options: &ListOptions{
			Flags: AllRecursive,
		},
		expected: []string{
			"test_data/list_test1/.steampipe",
			"test_data/list_test1/.steampipe/mods",
			"test_data/list_test1/.steampipe/mods/github.com",
			"test_data/list_test1/.steampipe/mods/github.com/turbot",
			"test_data/list_test1/.steampipe/mods/github.com/turbot/m1",
			"test_data/list_test1/.steampipe/mods/github.com/turbot/m1/mod.sp",
			"test_data/list_test1/.steampipe/mods/github.com/turbot/m1/q1.sp",
			"test_data/list_test1/.steampipe/mods/github.com/turbot/m2",
			"test_data/list_test1/.steampipe/mods/github.com/turbot/m2/mod.sp",
			"test_data/list_test1/.steampipe/mods/github.com/turbot/m2/q1.sp",
			"test_data/list_test1/a",
			"test_data/list_test1/a/mod.sp",
			"test_data/list_test1/a/q1.sp",
			"test_data/list_test1/a/q2.sp",
			"test_data/list_test1/b",
			"test_data/list_test1/b/mod.sp",
			"test_data/list_test1/b/q1.sp",
			"test_data/list_test1/b/q2.sp",
			"test_data/list_test1/config",
			"test_data/list_test1/config/aws.spc",
			"test_data/list_test1/config/default.spc",
		},
	},
	"AllFlat": {
		source: "test_data/list_test1",
		options: &ListOptions{
			Flags: AllFlat,
		},
		expected: []string{
			"test_data/list_test1/.steampipe",
			"test_data/list_test1/a",
			"test_data/list_test1/a.swp",
			"test_data/list_test1/b",
			"test_data/list_test1/config",
		},
	},
	"FilesFlat": {
		source: "test_data/list_test1",
		options: &ListOptions{
			Flags: FilesFlat,
		},
		expected: []string{
			"test_data/list_test1/a.swp",
		},
	},
	"DirectoriesFlat": {
		source: "test_data/list_test1",
		options: &ListOptions{
			Flags: DirectoriesFlat,
		},
		expected: []string{
			"test_data/list_test1/.steampipe",
			"test_data/list_test1/a",
			"test_data/list_test1/b",
			"test_data/list_test1/config",
		},
	},
	"DirectoriesRecursive": {
		source: "test_data/list_test1",
		options: &ListOptions{
			Flags: DirectoriesRecursive,
		},
		expected: []string{
			"test_data/list_test1/.steampipe",
			"test_data/list_test1/.steampipe/mods",
			"test_data/list_test1/.steampipe/mods/github.com",
			"test_data/list_test1/.steampipe/mods/github.com/turbot",
			"test_data/list_test1/.steampipe/mods/github.com/turbot/m1",
			"test_data/list_test1/.steampipe/mods/github.com/turbot/m2",
			"test_data/list_test1/a",
			"test_data/list_test1/b",
			"test_data/list_test1/config",
		},
	},
	"DirectoriesRecursive, exclude  **/.steampipe*": {
		source: "test_data/list_test1",
		options: &ListOptions{
			Flags:   DirectoriesRecursive,
			Exclude: []string{"**/.steampipe*"},
		},
		expected: []string{
			"test_data/list_test1/a",
			"test_data/list_test1/b",
			"test_data/list_test1/config",
		},
	},
	"FilesRecursive": {
		source: "test_data/list_test1",
		options: &ListOptions{
			Flags: FilesRecursive,
		},
		expected: []string{
			"test_data/list_test1/.steampipe/mods/github.com/turbot/m1/mod.sp",
			"test_data/list_test1/.steampipe/mods/github.com/turbot/m1/q1.sp",
			"test_data/list_test1/.steampipe/mods/github.com/turbot/m2/mod.sp",
			"test_data/list_test1/.steampipe/mods/github.com/turbot/m2/q1.sp",
			"test_data/list_test1/a/mod.sp",
			"test_data/list_test1/a/q1.sp",
			"test_data/list_test1/a/q2.sp",
			"test_data/list_test1/b/mod.sp",
			"test_data/list_test1/b/q1.sp",
			"test_data/list_test1/b/q2.sp",
			"test_data/list_test1/config/aws.spc",
			"test_data/list_test1/config/default.spc",
		},
	},
	"FilesRecursive, exclude  **/.steampipe*": {
		source: "test_data/list_test1",
		options: &ListOptions{
			Flags:   FilesRecursive,
			Exclude: []string{"**/.steampipe*"},
		},
		expected: []string{
			"test_data/list_test1/a/mod.sp",
			"test_data/list_test1/a/q1.sp",
			"test_data/list_test1/a/q2.sp",
			"test_data/list_test1/b/mod.sp",
			"test_data/list_test1/b/q1.sp",
			"test_data/list_test1/b/q2.sp",
			"test_data/list_test1/config/aws.spc",
			"test_data/list_test1/config/default.spc",
		},
	},
	"FilesRecursive, include exclude  **/.steampipe* **/*.sp": {
		source: "test_data/list_test1",
		options: &ListOptions{
			Flags:   FilesRecursive,
			Exclude: []string{"**/.steampipe*"},
			Include: []string{"**/*.sp"},
		},
		expected: []string{
			"test_data/list_test1/a/mod.sp",
			"test_data/list_test1/a/q1.sp",
			"test_data/list_test1/a/q2.sp",
			"test_data/list_test1/b/mod.sp",
			"test_data/list_test1/b/q1.sp",
			"test_data/list_test1/b/q2.sp",
		},
	},
	"FilesRecursive, include exclude **/github.com/**/mod.sp none": {
		source: "test_data/list_test1",
		options: &ListOptions{
			Flags:   FilesRecursive,
			Exclude: []string{},
			Include: []string{"**/github.com/**/mod.sp"},
		},
		expected: []string{
			"test_data/list_test1/.steampipe/mods/github.com/turbot/m1/mod.sp",
			"test_data/list_test1/.steampipe/mods/github.com/turbot/m2/mod.sp",
		},
	},
	"Selective": {
		source: "test_data/list_test1",
		options: &ListOptions{
			Flags:   FilesRecursive,
			Exclude: []string{},
			Include: []string{"**/github.com/**/mod.sp", "**/test_data/list_test1/a/mod.sp"},
		},
		expected: []string{
			"test_data/list_test1/.steampipe/mods/github.com/turbot/m1/mod.sp",
			"test_data/list_test1/.steampipe/mods/github.com/turbot/m2/mod.sp",
			"test_data/list_test1/a/mod.sp",
		},
	},
	"Single file with include - expects error": {
		source: "test_data/list_test1/config/aws.spc",
		options: &ListOptions{
			Flags:   FilesRecursive,
			Exclude: []string{},
			Include: []string{"*"},
		},
		expected: "ERROR",
	},
}

func TestListFiles(t *testing.T) {
	for name, test := range testCasesListFiles {
		listPath, err := filepath.Abs(test.source)
		if err != nil {
			t.Errorf("failed to build absolute list filepath from %s", test.source)
		}

		files, err := ListFiles(listPath, test.options)

		if err != nil {
			if test.expected != "ERROR" {
				t.Errorf("Test: '%s'' FAILED with unexpected error: %v", name, err)
			}
			continue
		}

		if test.expected == "ERROR" {
			t.Errorf("Test: '%s'' FAILED - expected error. got %v", name, files)
			continue
		}

		// now remove local path from files for expectation testing (as expectations are relative)
		localDirectory, err := os.Getwd()
		if err != nil {
			t.Errorf("failed to get working directory %v", err)
			continue
		}

		for i, f := range files {
			rel, err := filepath.Rel(localDirectory, f)
			if err != nil {
				t.Errorf("failed to convert %s to a relatyive path for verification: %v", f, err)
			}
			files[i] = rel
		}

		if !reflect.DeepEqual(test.expected, files) {
			fmt.Printf("")
			t.Errorf("Test: '%s'' FAILED : expected:\n\n%s\n\ngot:\n\n%s", name, test.expected, files)
		}
	}
}
