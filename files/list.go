package files

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/danwakefield/fnmatch"
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
