package files

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/danwakefield/fnmatch"
	"github.com/turbot/go-kit/helpers"
)

type ListFlag uint

const (
	Files ListFlag = 1 << iota
	Directories
	Recursive
	AllFlat              = Files | Directories
	AllRecursive         = Files | Directories | Recursive
	FilesRecursive       = Files | Recursive
	DirectoriesRecursive = Directories | Recursive
	FilesFlat            = Files
	DirectoriesFlat      = Directories
)

type ListOptions struct {
	Flags ListFlag
	// .gitignore (fnmatch) format patterns for file inclusions and exclusions
	Include []string
	Exclude []string
}

// ListFiles returns path of files and or folders under listPath
// inclusions/exclusions/recursion is defined by optds
func ListFiles(listPath string, opts *ListOptions) ([]string, error) {
	// check folder exists
	if _, err := os.Stat(listPath); os.IsNotExist(err) {
		return nil, nil
	}
	if opts == nil {
		opts = &ListOptions{Flags: Files & Directories & Recursive}
	}

	//check if the listPath is the path to a file in the system
	if helpers.FileExists(listPath) {
		// if we are not listing files with a path to a file
		if opts.Flags&Files == 0 {
			// it's an error
			return nil, fmt.Errorf("if the path is a file, then you must set the Files ListFlag")
		}
		// there should not be an include
		if len(opts.Include)+len(opts.Exclude) > 0 {
			return nil, fmt.Errorf("if the path is a file, then you must not specify include/exclude")
		}
		// split up into the parent directory and the file name
		dir := filepath.Dir(listPath)
		// the parent directory becomes the directory we want to walk
		listPath = dir
		// and the file name becomes the filter
		opts.Include = []string{listPath}
	}

	// if no include list provided, default to including everything
	if len(opts.Include) == 0 {
		opts.Include = []string{"*"}
	}

	if opts.Flags&Recursive != 0 {
		return listFilesRecursive(listPath, opts)
	}
	return listFilesFlat(listPath, opts)
}

// InclusionsFromExtensions takes a list of file extensions and convert into a .gitgnore format inclusions list
func InclusionsFromExtensions(extensions []string) []string {
	// build include string from extensions
	var includeStrings []string
	for _, extension := range extensions {
		includeStrings = append(includeStrings, fmt.Sprintf("**/*%s", extension))
	}
	return includeStrings
}

// InclusionsFromFiles takes a list of file names convert into a .gitgnore format inclusions list
func InclusionsFromFiles(filenames []string) []string {
	// build include string from extensions
	var includeStrings []string
	for _, extension := range filenames {
		includeStrings = append(includeStrings, fmt.Sprintf("**/%s", extension))
	}
	return includeStrings
}

// ShouldIncludePath returns whether the specified file path satisfies the inclusion and exclusion options
// (in .gitignore format)
func ShouldIncludePath(path string, include, exclude []string) bool {
	// if no include list provided, default to including everything
	if len(include) == 0 {
		include = []string{"*"}
	}
	// if the entry matches any of the exclude patterns, exclude
	for _, excludePattern := range exclude {
		if fnmatch.Match(excludePattern, path, 0) {
			return false
		}
	}
	// if the entry matches ANY of the include patterns, include
	shouldInclude := false
	for _, includePattern := range include {
		if fnmatch.Match(includePattern, path, 0) {
			// why does this not return?
			shouldInclude = true
		}
	}
	return shouldInclude
}

func listFilesRecursive(listPath string, opts *ListOptions) ([]string, error) {
	var res []string
	err := filepath.Walk(listPath,
		func(path string, entry os.FileInfo, err error) error {
			if err != nil {
				if _, ok := err.(*os.PathError); ok {
					// ignore path errors - this may be for a file which has been removed during the walk
					return nil
				}
				return err
			}
			// ignore list path itself
			if path == listPath {
				return nil
			}

			// should we include this file?
			if shouldIncludeEntry(path, entry, opts) {
				res = append(res, path)
			}

			return nil
		})
	return res, err
}

func listFilesFlat(path string, opts *ListOptions) ([]string, error) {
	entries, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read folder %s: %v", path, err)
	}

	matches := []string{}
	for _, entry := range entries {
		path := filepath.Join(path, entry.Name())
		if shouldIncludeEntry(path, entry, opts) {
			matches = append(matches, path)
		}
	}
	return matches, nil
}

// should the list results include this entry, based on the list options
func shouldIncludeEntry(path string, entry os.FileInfo, opts *ListOptions) bool {
	if entry.IsDir() {
		// if this is a directory and we are not including directories, exclude
		if opts.Flags&Directories == 0 {
			return false
		}
	} else {
		// if this is a file and we are not including files, exclude
		if opts.Flags&Files == 0 {
			return false
		}
	}

	return ShouldIncludePath(path, opts.Include, opts.Exclude)
}

func PathToRootAndGlob(path string) (root string, glob string, e error) {
	// if the first is * or ** - we prefix WD
	// an alternate case is if the first segment CONTAINS a * (e.g: terra*/*.tf)
	// we would like to prefix with WD, but first, we have to make sure that is not a
	// valid go-getter input
	splitList := filepath.SplitList(path)
	if len(splitList) > 0 && (splitList[0] == "*" || splitList[0] == "**") {
		wd, err := os.Getwd()
		if err != nil {
			return "", "", err
		}
		path = filepath.Join(wd, path)
	}

	// assume that the root is the path
	root = path
	// and the glob is blank (empty by default anyway - just mentioning explicitly)
	glob = ""

	if len(path) == 0 {
		e = errors.New("cannot accept empty path")
		return
	}

	for {
		//
		if len(root) == 0 {
			return root, glob, nil
		}
		// stat the path given
		_, err := os.Stat(root)
		if err == nil {
			// path exists in the file system
			return root, glob, nil
		}
		// does not exist
		// split the path into dir/glob
		dir, base := filepath.Split(root)
		// the base gets prepended into globComponents - which we will return later
		glob = filepath.Join(base, glob)

		// if the `dir` as a trailing slash, remove it
		for {
			if !strings.HasSuffix(dir, "/") {
				break
			}
			// loose the last character
			dir = dir[:len(dir)-1]
		}

		// the root becomes the directory
		root = dir
	}
}
