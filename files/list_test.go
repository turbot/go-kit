package files

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

type ListFilesTest struct {
	source   string
	options  *ListFilesOptions
	expected interface{}
}

var testCasesListFiles = map[string]ListFilesTest{
	"AllRecursive, exclude **/a*, **/*.swp, **/.steampipe*": {
		source: "test_data/list_test1",
		options: &ListFilesOptions{
			Options: AllRecursive,
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
		options: &ListFilesOptions{
			Options: AllRecursive,
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
			"test_data/list_test1/a.swp",
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
		options: &ListFilesOptions{
			Options: AllFlat,
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
		options: &ListFilesOptions{
			Options: FilesFlat,
		},
		expected: []string{
			"test_data/list_test1/a.swp",
		},
	},
	"DirectoriesFlat": {
		source: "test_data/list_test1",
		options: &ListFilesOptions{
			Options: DirectoriesFlat,
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
		options: &ListFilesOptions{
			Options: DirectoriesRecursive,
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
		options: &ListFilesOptions{
			Options: DirectoriesRecursive,
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
		options: &ListFilesOptions{
			Options: FilesRecursive,
		},
		expected: []string{
			"test_data/list_test1/.steampipe/mods/github.com/turbot/m1/mod.sp",
			"test_data/list_test1/.steampipe/mods/github.com/turbot/m1/q1.sp",
			"test_data/list_test1/.steampipe/mods/github.com/turbot/m2/mod.sp",
			"test_data/list_test1/.steampipe/mods/github.com/turbot/m2/q1.sp",
			"test_data/list_test1/a/mod.sp",
			"test_data/list_test1/a/q1.sp",
			"test_data/list_test1/a/q2.sp",
			"test_data/list_test1/a.swp",
			"test_data/list_test1/b/mod.sp",
			"test_data/list_test1/b/q1.sp",
			"test_data/list_test1/b/q2.sp",
			"test_data/list_test1/config/aws.spc",
			"test_data/list_test1/config/default.spc",
		},
	},
	"FilesRecursive, exclude  **/.steampipe*": {
		source: "test_data/list_test1",
		options: &ListFilesOptions{
			Options: FilesRecursive,
			Exclude: []string{"**/.steampipe*"},
		},
		expected: []string{
			"test_data/list_test1/a/mod.sp",
			"test_data/list_test1/a/q1.sp",
			"test_data/list_test1/a/q2.sp",
			"test_data/list_test1/a.swp",
			"test_data/list_test1/b/mod.sp",
			"test_data/list_test1/b/q1.sp",
			"test_data/list_test1/b/q2.sp",
			"test_data/list_test1/config/aws.spc",
			"test_data/list_test1/config/default.spc",
		},
	},
	"FilesRecursive, include exclude  **/.steampipe* **/*.sp": {
		source: "test_data/list_test1",
		options: &ListFilesOptions{
			Options: FilesRecursive,
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
			return
		}

		if test.expected == "ERROR" {
			t.Errorf("Test: '%s'' FAILED - expected error", name)
		}

		// now remove loacl path from files for expectation testing (as expectations are relative)
		localDirectory, err := os.Getwd()
		if err != nil {
			t.Errorf("failed to get working directory %v", err)
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
