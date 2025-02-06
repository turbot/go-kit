package files

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"testing"
)

type resolveGlobRootTest struct {
	patterns  []string
	rootPaths []string
	expected  interface{}
}

var resolveGlobRootTests = map[string]resolveGlobRootTest{
	"all absolute": {
		patterns:  []string{"/an/absolute/pattern"},
		rootPaths: []string{"/an/absolute/root"},
		expected:  []string{"/an/absolute/pattern"},
	},
	"multiple absolute patterns": {
		patterns:  []string{"/an/absolute/pattern", "/another/absolute/pattern"},
		rootPaths: []string{"/an/absolute/root"},
		expected:  []string{"/an/absolute/pattern", "/another/absolute/pattern"},
	},
	"absolute and relative patterns": {
		patterns:  []string{"/an/absolute/pattern", "a/relative/pattern"},
		rootPaths: []string{"/an/absolute/root"},
		expected:  []string{"/an/absolute/pattern", "/an/absolute/root/a/relative/pattern"},
	},
	"relative patterns": {
		patterns:  []string{"a/relative/pattern", "another/relative/pattern"},
		rootPaths: []string{"/an/absolute/root"},
		expected:  []string{"/an/absolute/root/a/relative/pattern", "/an/absolute/root/another/relative/pattern"},
	},
	"relative patterns, multiple roots": {
		patterns:  []string{"a/relative/pattern", "another/relative/pattern"},
		rootPaths: []string{"/an/absolute/root", "/another/absolute/root"},
		expected:  []string{"/an/absolute/root/a/relative/pattern", "/another/absolute/root/a/relative/pattern", "/an/absolute/root/another/relative/pattern", "/another/absolute/root/another/relative/pattern"},
	},
	"absolute and relative patterns, multiple roots": {
		patterns:  []string{"/an/absolute/pattern", "a/relative/pattern"},
		rootPaths: []string{"/an/absolute/root", "/another/absolute/root"},
		expected:  []string{"/an/absolute/pattern", "/an/absolute/root/a/relative/pattern", "/another/absolute/root/a/relative/pattern"},
	},
}

func TestResolveGlobRoots(t *testing.T) {
	for name, test := range resolveGlobRootTests {
		actual := ResolveGlobRoots(test.patterns, test.rootPaths...)
		expected := test.expected
		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("Test: '%s'' FAILED : expected:\n\n%s\n\ngot:\n\n%s\n\n", name, expected, actual)
		}
	}
}

type globRootTestExpected struct {
	root string
	glob string
}

func (t globRootTestExpected) String() string {
	return fmt.Sprintf("{root: %s ; glob: %s}", t.root, t.glob)
}

type globRootTest struct {
	path        string
	expected    globRootTestExpected
	expectError bool
}

var wd, _ = os.Getwd()

var splitTests = map[string]globRootTest{
	"$PWD/test_data/list_test1/config/aws.spc": {
		path:     filepath.Join(wd, "test_data/list_test1/config/aws.spc"),
		expected: globRootTestExpected{root: filepath.Join(wd, "test_data/list_test1/config"), glob: filepath.Join(wd, "test_data/list_test1/config/aws.spc")},
	},
	"**/*.tf": {
		path:     "**/*.tf",
		expected: globRootTestExpected{root: wd, glob: filepath.Join(wd, "**/*.tf")},
	},
	"test_data/list_test1/config/*.spc": {
		path:     "test_data/list_test1/config/*.spc",
		expected: globRootTestExpected{root: filepath.Join(wd, "test_data/list_test1/config"), glob: filepath.Join(wd, "test_data/list_test1/config/*.spc")},
	},
	"./*.tf": {
		path:     "./*.tf",
		expected: globRootTestExpected{root: wd, glob: filepath.Join(wd, "*.tf")},
	},
	"./terr/**/*.tf": {
		path:     "./terr/**/*.tf",
		expected: globRootTestExpected{root: wd, glob: filepath.Join(wd, "terr/**/*.tf")},
	},
}

func TestGlobRoot(t *testing.T) {
	for name, test := range splitTests {
		root, glob, err := GlobRoot(test.path)

		if err != nil && !test.expectError {
			t.Errorf("Test: '%s'' FAILED with unexpected error: %v", name, err)
		}

		actual := globRootTestExpected{
			root: root,
			glob: glob,
		}
		expected := test.expected
		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("Test: '%s'' FAILED : expected:\n\n%s\n\ngot:\n\n%s\n\n", name, expected, actual)
		}
	}
}

type listFilesTest struct {
	expected interface{}
	options  *ListOptions
	preRun   func()
	postRun  func()
	source   string
}

var testCasesListFiles = map[string]listFilesTest{
	"AllRecursive, exclude a, a/*, *.swp, **/*.swp, .steampipe .steampipe/*": {
		source: "test_data/list_test1",
		options: &ListOptions{
			Flags:   AllRecursive,
			Exclude: []string{"a", "a/*", "*.swp", "**/*.swp", ".steampipe", ".steampipe/**"},
		},
		expected: []string{
			"test_data/list_test1/b",
			"test_data/list_test1/b.sp",
			"test_data/list_test1/b/mod.sp",
			"test_data/list_test1/b/q1.sp",
			"test_data/list_test1/b/q2.sp",
			"test_data/list_test1/config",
			"test_data/list_test1/config/aws.spc",
			"test_data/list_test1/config/default.spc",
		},
	},
	"SingleHopInclusion": {
		source: "test_data/list_test2",
		options: &ListOptions{
			Flags:   AllRecursive,
			Include: []string{"a2/*/*.sp"},
		},
		expected: []string{
			"test_data/list_test2/a2/b/mod.sp",
			"test_data/list_test2/a2/b/q1.sp",
			"test_data/list_test2/a2/b/q2.sp",
			"test_data/list_test2/a2/c/mod.sp",
			"test_data/list_test2/a2/c/q1.sp",
			"test_data/list_test2/a2/c/q2.sp",
		},
	},
	"TopLevelOnly": {
		source: "test_data/list_test2",
		options: &ListOptions{
			Flags:   AllRecursive,
			Include: []string{"*.sp"},
		},
		expected: []string{"test_data/list_test2/mod.sp"},
	},
	"TopLevelRecursive": {
		source: "test_data/list_test1",
		options: &ListOptions{
			Flags:   AllRecursive,
			Include: []string{"**/*.spc"},
		},
		expected: []string{
			"test_data/list_test1/config/aws.spc",
			"test_data/list_test1/config/default.spc",
		},
	},
	"RecursiveFromARoot": {
		source: "test_data/list_test2",
		options: &ListOptions{
			Flags:   AllRecursive,
			Include: []string{"a2/**/*.sp"},
		},
		expected: []string{
			"test_data/list_test2/a2/mod.sp",
			"test_data/list_test2/a2/q1.sp",
			"test_data/list_test2/a2/q2.sp",
			"test_data/list_test2/a2/b/mod.sp",
			"test_data/list_test2/a2/b/q1.sp",
			"test_data/list_test2/a2/b/q2.sp",
			"test_data/list_test2/a2/c/mod.sp",
			"test_data/list_test2/a2/c/q1.sp",
			"test_data/list_test2/a2/c/q2.sp",
			"test_data/list_test2/a2/c/d/mod.sp",
			"test_data/list_test2/a2/c/d/q1.sp",
			"test_data/list_test2/a2/c/d/q2.sp",
			"test_data/list_test2/a2/c/d/e/mod.sp",
			"test_data/list_test2/a2/c/d/e/q1.sp",
			"test_data/list_test2/a2/c/d/e/q2.sp",
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
			"test_data/list_test1/b.sp",
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
			"test_data/list_test1/b.sp",
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
	"DirectoriesRecursiveWithEmpty": {
		source: "test_data/list_test3",
		options: &ListOptions{
			Flags: DirectoriesRecursive,
		},
		preRun: func() {
			// create an empty directory
			_ = os.Mkdir(filepath.Join(wd, "test_data/list_test3/a/empty"), 0644)
		},
		postRun: func() {
			// remove empty directory
			_ = os.RemoveAll(filepath.Join(wd, "test_data/list_test3/a/empty"))
		},
		expected: []string{
			"test_data/list_test3/.steampipe",
			"test_data/list_test3/.steampipe/mods",
			"test_data/list_test3/.steampipe/mods/github.com",
			"test_data/list_test3/.steampipe/mods/github.com/turbot",
			"test_data/list_test3/.steampipe/mods/github.com/turbot/m1",
			"test_data/list_test3/.steampipe/mods/github.com/turbot/m2",
			"test_data/list_test3/a",
			"test_data/list_test3/a/empty",
			"test_data/list_test3/b",
		},
	},
	"DirectoriesRecursiveNotEmpty": {
		source: "test_data/list_test3",
		options: &ListOptions{
			Flags: DirectoriesRecursive | NotEmpty,
		},
		preRun: func() {
			// create an empty directory
			_ = os.Mkdir(filepath.Join(wd, "test_data/list_test3/a/empty"), 0644)
		},
		postRun: func() {
			// remove empty directory
			os.RemoveAll(filepath.Join(wd, "test_data/list_test3/a/empty"))
		},
		expected: []string{
			"test_data/list_test3/.steampipe",
			"test_data/list_test3/.steampipe/mods",
			"test_data/list_test3/.steampipe/mods/github.com",
			"test_data/list_test3/.steampipe/mods/github.com/turbot",
			"test_data/list_test3/.steampipe/mods/github.com/turbot/m1",
			"test_data/list_test3/.steampipe/mods/github.com/turbot/m2",
			"test_data/list_test3/a",
			"test_data/list_test3/b",
		},
	},
	"DirectoriesRecursive, exclude  .steampipe/*": {
		source: "test_data/list_test1",
		options: &ListOptions{
			Flags:   DirectoriesRecursive,
			Exclude: []string{".steampipe", ".steampipe/**"},
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
			"test_data/list_test1/a.swp",
			"test_data/list_test1/a/mod.sp",
			"test_data/list_test1/a/q1.sp",
			"test_data/list_test1/a/q2.sp",
			"test_data/list_test1/b.sp",
			"test_data/list_test1/b/mod.sp",
			"test_data/list_test1/b/q1.sp",
			"test_data/list_test1/b/q2.sp",
			"test_data/list_test1/config/aws.spc",
			"test_data/list_test1/config/default.spc",
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
	"FilesRecursive non recursive glob": {
		source: "test_data/list_test1",
		options: &ListOptions{
			Flags:   FilesRecursive,
			Include: []string{"*.sp"},
		},
		expected: []string{
			"test_data/list_test1/b.sp",
		},
	},
	"FilesRecursive, exclude  **/.steampipe/*": {
		source: "test_data/list_test1",
		options: &ListOptions{
			Flags:   FilesRecursive,
			Exclude: []string{".steampipe/**", ".steampipe/**/*"},
		},
		expected: []string{
			"test_data/list_test1/a.swp",
			"test_data/list_test1/a/mod.sp",
			"test_data/list_test1/a/q1.sp",
			"test_data/list_test1/a/q2.sp",
			"test_data/list_test1/b.sp",
			"test_data/list_test1/b/mod.sp",
			"test_data/list_test1/b/q1.sp",
			"test_data/list_test1/b/q2.sp",
			"test_data/list_test1/config/aws.spc",
			"test_data/list_test1/config/default.spc",
		},
	},
	"FilesRecursive, include  **/*.sp exclude .steampipe/**": {
		source: "test_data/list_test1",
		options: &ListOptions{
			Flags:   FilesRecursive,
			Exclude: []string{".steampipe/**"},
			Include: []string{"**/*.sp"},
		},
		expected: []string{
			"test_data/list_test1/a/mod.sp",
			"test_data/list_test1/a/q1.sp",
			"test_data/list_test1/a/q2.sp",
			"test_data/list_test1/b.sp",
			"test_data/list_test1/b/mod.sp",
			"test_data/list_test1/b/q1.sp",
			"test_data/list_test1/b/q2.sp",
		},
	},

	// tests disabled since we don't support multiple ** (yet)
	// "FilesRecursive, include **/github.com/**/mod.sp": {
	// 	source: "test_data/list_test1",
	// 	options: &ListOptions{
	// 		Flags:   FilesRecursive,
	// 		Exclude: []string{},
	// 		Include: []string{"**/github.com/**/mod.sp"},
	// 	},
	// 	expected: []string{
	// 		"test_data/list_test1/.steampipe/mods/github.com/turbot/m1/mod.sp",
	// 		"test_data/list_test1/.steampipe/mods/github.com/turbot/m2/mod.sp",
	// 	},
	// },
	// "Selective": {
	// 	source: "test_data/list_test1",
	// 	options: &ListOptions{
	// 		Flags:   FilesRecursive,
	// 		Exclude: []string{},
	// 		Include: []string{"**/github.com/**/mod.sp", "a/mod.sp"},
	// 	},
	// 	expected: []string{
	// 		"test_data/list_test1/.steampipe/mods/github.com/turbot/m1/mod.sp",
	// 		"test_data/list_test1/.steampipe/mods/github.com/turbot/m2/mod.sp",
	// 		"test_data/list_test1/a/mod.sp",
	// 	},
	// },
}

func TestListFiles(t *testing.T) {
	for name, test := range testCasesListFiles {
		if test.preRun != nil {
			test.preRun()
		}
		if test.postRun != nil {
			defer test.postRun()
		}
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

		sort.Strings(test.expected.([]string))
		sort.Strings(files)
		if !reflect.DeepEqual(test.expected, files) {
			fmt.Printf("")
			t.Errorf("Test: '%s'' FAILED : expected:\n\n%s\n\ngot:\n\n%s", name, test.expected, files)
		}
	}
}
