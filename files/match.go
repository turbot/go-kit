package files

import (
	"github.com/danwakefield/fnmatch"
	"os"
	"path/filepath"
	"strings"
)

const dblAsterisks = "**"

func Match(pattern, value string) bool {
	if strings.Contains(pattern, dblAsterisks) {
		return evalDblAsterisk(pattern, value)
	} else {
		return fnmatch.Match(pattern, value, fnmatch.FNM_PATHNAME)
	}
}

func evalDblAsterisk(pattern, value string) bool {
	// A leading "**" followed by a slash means match in all directories.
	// For example, "**/foo" matches file or directory "foo" anywhere,
	// the same as pattern "foo". "**/foo/bar" matches file or directory
	// "bar" anywhere that is directly under directory "foo".
	if strings.HasPrefix(pattern, dblAsterisks) {
		pattern = strings.TrimPrefix(pattern, dblAsterisks)
		return strings.HasSuffix(value, pattern)
	}

	// A trailing "/**" matches everything inside. For example, "abc/**"
	// matches all files inside directory "abc", relative to the location
	// of the .gitignore file, with infinite depth.
	if strings.HasSuffix(pattern, dblAsterisks) {
		pattern = strings.TrimSuffix(pattern, dblAsterisks)
		return strings.HasPrefix(value, pattern)
	}

	// A slash followed by two consecutive asterisks then a slash matches
	// zero or more directories. For example, "a/**/b" matches "a/b",
	// /"a/x/b", "a/x/y/b" and so on.
	parts := strings.Split(pattern, dblAsterisks)
	for i, part := range parts {
		switch i {
		case 0:
			if !strings.HasPrefix(value, part) {
				return false
			}
		case len(parts) - 1: // last part
			part = strings.TrimPrefix(part, string(os.PathSeparator))

			// if the pattern part has no directories, tri value of directory
			partSegments := strings.Split(part, string(os.PathSeparator))
			if len(partSegments) == 1 {
				value = filepath.Base(value)
			} else {
				// trim value to the same number of path segments
				valueSegments := strings.Split(value, string(os.PathSeparator))
				lengthDiff := len(valueSegments) - len(partSegments)
				if lengthDiff >= 0 {
					value = strings.Join(valueSegments[lengthDiff:], string(os.PathSeparator))
				}
			}

			return fnmatch.Match(part, value, fnmatch.FNM_PATHNAME)
		default:
			if !strings.Contains(value, part) {
				return false
			}
		}

		// trim evaluated text
		index := strings.Index(value, part) + len(part)
		value = value[index:]
	}

	// Other consecutive asterisks are considered invalid.
	return false
}
