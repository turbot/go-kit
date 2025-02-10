package helpers

import (
	"bytes"
	"encoding/csv"
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

// TabifyStringSlice adds the provided tab string to beginning of each line of string
func TabifyStringSlice(strs []string, tab string) []string {
	for i, line := range strs {
		strs[i] = fmt.Sprintf("%s%s", tab, line)
	}
	return strs
}

// TruncateString limits the string to the given length, adding an ellipsis if the string is being truncated,
// also handles newlines
func TruncateString(str string, length int) string {
	// NOTE: this function does not handle (i.e. ignore) ansi color codes in strings
	newStrings := strings.Split(str, "\n")
	for i, strnew := range newStrings {
		if PrintableLength(strnew) <= length {
			newStrings[i] = strnew
		}
		if length <= 1 {
			return ""
		}
		// now limit each printable runes and add an ellipsis
		if PrintableLength(strnew) > length {
			newStrings[i] = fmt.Sprintf("%s…", LimitPrintableRunes(strnew, length-1))
		}
	}
	str = strings.Join(newStrings, "\n")
	return str
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

// SplitByRune uses the CSV decoder to parse out the tokens - even if they are quoted and/or escaped
func SplitByRune(str string, r rune) []string {
	csvDecoder := csv.NewReader(strings.NewReader(str))
	csvDecoder.Comma = r
	csvDecoder.LazyQuotes = true
	csvDecoder.TrimLeadingSpace = true
	split, _ := csvDecoder.Read()
	return split
}

// SplitByWhitespace splits by the ' ' rune
func SplitByWhitespace(str string) []string {
	return SplitByRune(str, ' ')
}

// Resize resizes the string with the given length. It ellipses with '…' when the string's length exceeds
// the desired length or pads spaces to the right of the string when length is smaller than desired
func Resize(s string, length uint) string {
	n := int(length)
	if len(s) == n {
		return s
	}
	// Pads only when length of the string smaller than len needed
	s = PadRight(s, n, ' ')
	if len(s) > n {
		b := []byte(s)
		var buf bytes.Buffer
		for i := 0; i < n-1; i++ {
			buf.WriteByte(b[i])
		}
		buf.WriteString("…")
		s = buf.String()
	}
	return s
}

// PadRight returns a new string of a specified length in which the end of the current string is padded with spaces or with a specified Unicode character.
func PadRight(str string, length int, pad byte) string {
	if len(str) >= length {
		return str
	}
	buf := bytes.NewBufferString(str)
	for i := 0; i < length-len(str); i++ {
		buf.WriteByte(pad)
	}
	return buf.String()
}

// TrimBlankLines removes any empty lines from the string
func TrimBlankLines(str string) string {
	lines := strings.Split(str, "\n")
	strippedLines := RemoveFromStringSlice(lines, "")
	return strings.Join(strippedLines, "\n")
}

// IsOnlyNumeric returns true if the string only contains numeric characters
func IsOnlyNumeric(s string) bool {
	return strings.Trim(s, "0123456789") == ""
}
