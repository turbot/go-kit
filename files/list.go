package files

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type ListFlag uint

const (
	// Files lists only files
	Files ListFlag = 1 << iota
	// Directories lists only directories
	Directories
	// Recursively traverses directories
	Recursive
	// Ignores empty directories
	NotEmpty
	AllFlat              = Files | Directories
	AllRecursive         = Files | Directories | Recursive
	FilesRecursive       = Files | Recursive
	DirectoriesRecursive = Directories | Recursive
	FilesFlat            = Files
	DirectoriesFlat      = Directories
)

type ListOptions struct {
	// .gitignore (fnmatch) format patterns for file inclusions and exclusions
	Include []string
	Exclude []string
	Flags   ListFlag
	// max results
	MaxResults int
}

// ListFiles returns path of files and or folders under listPath
// inclusions/exclusions/recursion is defined by opts
func ListFiles(listPath string, opts *ListOptions) ([]string, error) {
	return ListFilesWithContext(context.Background(), listPath, opts)
}

func ListFilesWithContext(ctx context.Context, listPath string, opts *ListOptions) ([]string, error) {
	// check folder exists
	if _, err := os.Stat(listPath); os.IsNotExist(err) {
		return nil, nil
	}
	if opts == nil {
		opts = &ListOptions{Flags: Files & Directories & Recursive}
	}

	//check if the listPath is the path to a file in the system
	if FileExists(listPath) {
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

	if opts.Flags&Recursive != 0 {
		return listFilesRecursive(ctx, listPath, opts)
	}
	return listFilesFlat(ctx, listPath, opts)
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
// Note: it is expected the pattern will be absolute, i.e including the base path: /tmp/foo/**/*.json
func ShouldIncludePath(path string, include, exclude []string) bool {
	// if the entry matches any of the exclude patterns, exclude
	for _, excludePattern := range exclude {
		if Match(excludePattern, path) {
			return false
		}
	}
	// if no include list provided, default to including everything
	if len(include) == 0 {
		return true
	}

	// if the entry matches ANY of the include patterns, include
	for _, includePattern := range include {
		if Match(includePattern, path) {
			return true
		}
	}
	return false
}

func listFilesRecursive(ctx context.Context, listPath string, opts *ListOptions) ([]string, error) {
	var res []string
	count := 0 // initialize a counter to keep track of the number of files returned
	err := filepath.WalkDir(listPath,
		func(filePath string, entry fs.DirEntry, err error) error {
			// handle context cancellations
			if ctx.Err() != nil {
				log.Println("[INFO] context canceled")
				return fs.SkipAll
			}
			if err != nil {
				if _, ok := err.(*os.PathError); ok {
					// ignore path errors - this may be for a file which has been removed during the walk
					return nil
				}
				if entry.IsDir() {
					// TO DO LOG??
					return fs.SkipDir
				}
				return err
			}
			// ignore list path itself
			if filePath == listPath {
				return nil
			}
			// should we include this file?
			if shouldIncludeEntry(listPath, filePath, entry, opts) {
				res = append(res, filePath)
				count++ // increment the counter
				// check the number of files reached, if MaxResults is reached,
				// stop walking the directory
				if opts.MaxResults > 0 {
					if count == opts.MaxResults {
						// TODO: go 1.20 has a fs.SkipAll error which we can use
						return io.EOF
					}
				}
			} else if entry.IsDir() && !shouldSearchInDir(listPath, filePath, opts) {
				return fs.SkipDir
			}
			return nil
		})
	// set err to nil in case of EOF error so that the code doesn't exit from here
	if err == io.EOF {
		err = nil
	}
	// skipAll sends a nil error in case of context cancellation - we want to capture the context cancellation error
	if err == nil {
		err = ctx.Err()
	}
	return res, err
}

func listFilesFlat(ctx context.Context, listPath string, opts *ListOptions) ([]string, error) {

	entries, err := os.ReadDir(listPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read folder %s: %v", listPath, err)
	}

	var matches []string
	for _, entry := range entries {
		// handle context cancellations
		if ctx.Err() != nil {
			return nil, ctx.Err()
		}
		filePath := filepath.Join(listPath, entry.Name())
		if err != nil {
			continue
		}
		if shouldIncludeEntry(listPath, filePath, entry, opts) {
			matches = append(matches, filePath)
		}
	}
	return matches, nil
}

// should the list results include this entry, based on the list options
func shouldIncludeEntry(listPath, filePath string, entry fs.DirEntry, opts *ListOptions) bool {
	if entry.IsDir() {
		// if this is a directory and we are not including directories, exclude
		if opts.Flags&Directories == 0 {
			return false
		}

		// this is a directory and we are including directories
		// check if we can include empty directories
		if opts.Flags&NotEmpty != 0 {
			ls, err := os.ReadDir(filePath)
			if err != nil {
				// if dir read fails, skip this directory since the listing code will probably also fail to read it
				return false
			}

			// do not include this if it's empty
			if len(ls) == 0 {
				return false
			}
		}
	} else {
		// if this is a file and we are not including files, exclude
		if opts.Flags&Files == 0 {
			return false
		}
	}

	return ShouldIncludePath(filePath, ResolveGlobRoots(opts.Include, listPath), ResolveGlobRoots(opts.Exclude, listPath))
}

// shouldSearchInDir returns whether the specified directory satisfies the inclusion and exclusion options
// (in .gitignore format)
func shouldSearchInDir(listPath, dirPath string, opts *ListOptions) bool {
	// if no include list provided, default to including everything
	if len(opts.Include) == 0 {
		return true
	}
	include := ResolveGlobRoots(opts.Include, listPath)

	pathLen := len(strings.Split(dirPath, string(os.PathSeparator)))

	// if the entry matches ANY of the include patterns, include
	for _, includePattern := range include {
		// trim the include pattern to be no longer than the path length
		includeSegments := strings.Split(includePattern, string(os.PathSeparator))
		if len(includeSegments) > pathLen {
			includePattern = strings.Join(includeSegments[:pathLen], string(os.PathSeparator))
		}

		if Match(includePattern, dirPath, WithAsDir) {
			return true
		}
	}
	return false
}

// ResolveGlobRoots resolve the glob patter for each of the given root paths
func ResolveGlobRoots(pattern []string, rootPaths ...string) []string {
	// create an array of the maximum capacity, so that
	// appends do not have to increase the size of the
	// underlying array
	res := make([]string, 0, len(pattern)*len(rootPaths))
	for _, p := range pattern {
		// if the glob is already absolute, add it and continue
		if strings.HasPrefix(p, string(os.PathSeparator)) {
			res = append(res, p)
			continue
		}
		for _, rootPath := range rootPaths {
			res = append(res, filepath.Join(rootPath, p))
		}
	}
	return res
}

// SplitPath splits the given path using the os.PathSeparator
func SplitPath(path string) []string {
	return strings.Split(path, string(os.PathSeparator))
}

// GlobRoot takes in a glob and tries to resolve the prefix of the glob
// such that the prefix exists in the filesystem.
//
// If the given glob is relative, then GlobRoot converts it into an absolute path
// before attempting resolution.
//
// If the given glob is can be resolved to an existing file in the system, then
// the parent directory of the file along with the full path of the file is returned
func GlobRoot(glob string) (string, string, error) {
	// we cannot work with an empty input
	if len(glob) == 0 {
		return "", "", errors.New("cannot accept empty path")
	}

	// if the glob exists as a file
	if FileExists(glob) {
		// return the absolute path to the
		absolutePath, _ := Tildefy(glob)
		// find out the parent directory
		parentDirectory := filepath.Dir(absolutePath)
		// return them
		return parentDirectory, absolutePath, nil
	}

	workingDirectory, err := os.Getwd()
	if err != nil {
		return "", "", err
	}

	// if the first contains * or ** - we prefix the working directory
	// then replace that with the current working directory
	// we are (more-or-less) sure that go-getter - the resource fetching library under the hood
	// does not have any getter which accepts an input with a * in the first path segment
	// return immediately, since we are sure that the first segment is THE glob
	firstSegment := SplitPath(glob)[0]
	if strings.Contains(firstSegment, "*") {
		glob, err = Tildefy(glob)
		if err != nil {
			return "", "", err
		}
		return workingDirectory, glob, nil
	}

	// if the first segment is a ".",
	// then replace that with the current
	// working directory as well
	if firstSegment == "." {
		glob, err = Tildefy(glob)
		if err != nil {
			return "", "", err
		}
	}

	// assume that the root is the absolute glob
	root := glob

	for {
		// clean the path given
		root = filepath.Clean(root)

		// resolve the ~ ($HOME) and get the root as an absolute path
		// if an absolute path was given, Tildefy does not change the output
		absoluteRoot, err := Tildefy(root)
		if err != nil {
			return "", "", err
		}

		// stat the current root
		if DirectoryExists(absoluteRoot) {
			// path exists in the file system
			t, _ := Tildefy(glob)
			return absoluteRoot, t, nil
		}

		// split the path into dir/glob
		// the root becomes the directory
		root = filepath.Dir(root)

		// filepath.Dir returns a "." instead of an empty path
		if root == "." {
			// return the original glob
			return "", glob, nil
		}
	}
}
