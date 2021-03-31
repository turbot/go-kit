package files

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/danwakefield/fnmatch"
)

// An ImportMode controls the behavior of the Import method.
type ListOption uint

const (
	Files ListOption = 1 << iota
	Directories
	Recursive
	AllFlat              = Files | Directories
	AllRecursive         = Files | Directories | Recursive
	FilesRecursive       = Files | Recursive
	DirectoriesRecursive = Directories | Recursive
	FilesFlat            = Files
	DirectoriesFlat      = Directories
)

type ListFilesOptions struct {
	Options ListOption
	Include []string
	Exclude []string
}

// ListFiles :: recursively list files and or folders
func ListFiles(listPath string, opts *ListFilesOptions) ([]string, error) {
	// check folder exists
	if _, err := os.Stat(listPath); os.IsNotExist(err) {
		return nil, nil
	}
	if opts == nil {
		opts = &ListFilesOptions{Options: Files & Directories & Recursive}
	}
	// if no include list provided, default to including everything
	if len(opts.Include) == 0 {
		opts.Include = []string{"*"}
	}

	if opts.Options&Recursive != 0 {
		return listFilesRecursive(listPath, opts)
	}
	return listFilesFlat(listPath, opts)
}

// InclusionsFromExtensions :: take a list of file extensions and convert into a .gitgnore format inclusions list
func InclusionsFromExtensions(extensions []string) []string {
	// build include string from extensions
	var includeStrings []string
	for _, extension := range extensions {
		includeStrings = append(includeStrings, fmt.Sprintf("**/*%s", extension))
	}
	return includeStrings
}

func listFilesRecursive(listPath string, opts *ListFilesOptions) ([]string, error) {
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

func listFilesFlat(path string, opts *ListFilesOptions) ([]string, error) {
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

func shouldIncludeEntry(path string, entry os.FileInfo, opts *ListFilesOptions) bool {
	if entry.IsDir() {
		// if this is a directory and we are not including directories, exclude
		if opts.Options&Directories == 0 {
			return false
		}
	} else {
		// if this is a file and we are not including files, exclude
		if opts.Options&Files == 0 {
			return false
		}
	}

	// if the entry matches any of the exclude patterns, exclude
	for _, excludePattern := range opts.Exclude {
		if fnmatch.Match(excludePattern, path, 0) {
			return false
		}
	}
	// if the entry matches ANY of the include patterns, include
	shouldInclude := false
	for _, includePattern := range opts.Include {
		if fnmatch.Match(includePattern, path, 0) {
			shouldInclude = true
		}
	}

	return shouldInclude
}
