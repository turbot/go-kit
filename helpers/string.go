package helpers

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/acarl005/stripansi"
)

// Tabify adds the provided tab string to beginning of each line of string
func Tabify(str string, tab string) string {
	split := strings.Split(str, "\n")
	for i, line := range split {
		split[i] = fmt.Sprintf("%s%s", tab, line)
	}
	return strings.Join(split, "\n")
}

// TruncateString limits the string to the given length, adding an ellipsis if the string is being truncated
func TruncateString(str string, length int) string {
	// NOTE: this function does not handle (i.e. ignore) ansi color codes in strings
	if PrintableLength(str) <= length {
		return str
	}
	if length <= 1 {
		return ""
	}
	// now limit the printable runes and add an ellipsis
	return fmt.Sprintf("%s…", LimitPrintableRunes(str, length-1))
}

// LimitPrintableRunes limits the string to the given number of runes
func LimitPrintableRunes(s string, n int) string {
	if n <= 0 || s == "" {
		return ""
	}

	for i, r := range s {
		// if this is a printable character, count it
		if unicode.IsGraphic(r) {
			n--
		}
		if n == -1 {
			return s[:i]
		}
	}
	// s has n or less runes
	return s
}

// PrintableLength returns the number of printable runes in a string (after stripping ansi codes)
func PrintableLength(str string) int {
	cleaned := Clean(str)
	return len([]rune(cleaned))
}

// Clean removes non-printing characters from the string
func Clean(str string) string {
	// first strip ansi codes from string
	str = stripansi.Strip(str)
	// no strip non printing characters
	return strings.Map(func(r rune) rune {
		if unicode.IsGraphic(r) {
			return r
		}
		return -1
	}, str)
}
