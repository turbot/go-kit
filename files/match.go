package files

import (
	"fmt"
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
	// For example, "**/foo" matches file or directory "foo" anywhere
	// "**/foo/bar" matches file or directory "bar" anywhere that is directly under directory "foo".
	// "**/*.json" matches any file with the extension .json
	if strings.HasPrefix(pattern, dblAsterisks) {
		return evalLeadingDblAsterisk(pattern, value)
	}

	// A trailing "/**" matches everything inside. For example, "abc/**"
	// matches all files inside directory "abc", relative to the location
	// of the .gitignore file, with infinite depth.
	if strings.HasSuffix(pattern, dblAsterisks) {
		return evalTrailingDblAsterisk(pattern, value)
	}

	// A slash followed by two consecutive asterisks then a slash matches
	// zero or more directories. For example, "a/**/b" matches "a/b",
	// /"a/x/b", "a/x/y/b" and so on.
	patternParts := strings.Split(pattern, dblAsterisks)
	for i, patternPart := range patternParts {
		switch i {
		case 0:
			patternPart = strings.TrimSuffix(patternPart, string(os.PathSeparator))
			if !prefixMatches(patternPart, value) {
				return false
			}
		case len(patternParts) - 1: // last part
			patternPart = strings.TrimPrefix(patternPart, string(os.PathSeparator))

			// if the pattern part has no directories, trim value of directory
			partSegments := strings.Split(patternPart, string(os.PathSeparator))
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

			return fnmatch.Match(patternPart, value, fnmatch.FNM_PATHNAME)
		default:
			if !strings.Contains(value, patternPart) {
				return false
			}
		}

		// trim evaluated text
		index := strings.Index(value, patternPart) + len(patternPart)
		value = value[index:]
	}

	// Other consecutive asterisks are considered invalid.
	return false
}

func evalLeadingDblAsterisk(pattern string, value string) bool {
	prefix := fmt.Sprintf("%s%c", dblAsterisks, os.PathSeparator)
	// just check each part of the path matches
	// work back through the pattern and the value - each segment must match
	trimmedPattern := strings.TrimPrefix(pattern, prefix)
	trimmedValue := strings.TrimPrefix(value, string(os.PathSeparator))
	return suffixMatches(trimmedPattern, trimmedValue)
}

func suffixMatches(trimmedPattern string, trimmedValue string) bool {
	patternParts := strings.Split(trimmedPattern, string(os.PathSeparator))
	valueParts := strings.Split(trimmedValue, string(os.PathSeparator))
	patternLen := len(patternParts)
	valueLen := len(valueParts)
	for i := 1; i <= patternLen; i++ {
		// if we have run out of value parts, fail
		if i > valueLen {
			return false
		}
		patternPart := patternParts[patternLen-i]
		valuePart := valueParts[valueLen-i]

		// not yet handled
		if patternPart == dblAsterisks {
			panic("Match does not currently handle globs with more than 1 '**'")
		}
		if !fnmatch.Match(patternPart, valuePart, fnmatch.FNM_PATHNAME) {
			return false
		}
	}
	return true
}

func evalTrailingDblAsterisk(pattern string, value string) bool {
	// check each part of the path matches
	// work formward through the pattern and the value - each segment must match
	suffix := fmt.Sprintf("%c%s", os.PathSeparator, dblAsterisks)
	// just check each part of the path matches
	// work back through the pattern and the value - each segment must match
	trimmedPattern := strings.TrimSuffix(pattern, suffix)
	trimmedValue := strings.TrimSuffix(value, string(os.PathSeparator))

	return prefixMatches(trimmedPattern, trimmedValue)
}

func prefixMatches(trimmedPattern string, trimmedValue string) bool {
	patternParts := strings.Split(trimmedPattern, string(os.PathSeparator))
	valueParts := strings.Split(trimmedValue, string(os.PathSeparator))
	patternLen := len(patternParts)
	valueLen := len(valueParts)
	for i := 0; i < patternLen; i++ {
		// if we have run out of value parts, fail
		if i >= valueLen {
			return false
		}
		patternPart := patternParts[i]
		valuePart := valueParts[i]

		// not yet handled
		if patternPart == dblAsterisks {
			panic("Match does not currently handle globs with more than 1 '**'")
		}
		if !fnmatch.Match(patternPart, valuePart, fnmatch.FNM_PATHNAME) {
			return false
		}
	}
	return true
}
