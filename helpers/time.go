package helpers

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// ParseTime parses a string into a time.Time object attempting to parse it as a variety of formats
func ParseTime(input string) (time.Time, error) {
	var timeFormats = []string{
		time.RFC3339Nano,             // "2006-01-02T15:04:05.999999999Z07:00"
		time.RFC3339,                 // "2006-01-02T15:04:05Z07:00"
		time.RFC1123Z,                // "Mon, 02 Jan 2006 15:04:05 -0700"
		time.RFC1123,                 // "Mon, 02 Jan 2006 15:04:05 MST"
		time.RFC850,                  // "Monday, 02-Jan-06 15:04:05 MST"
		time.RFC822Z,                 // "02 Jan 06 15:04 -0700"
		time.RFC822,                  // "02 Jan 06 15:04 MST"
		time.UnixDate,                // "Mon Jan _2 15:04:05 MST 2006"
		time.RubyDate,                // "Mon Jan 02 15:04:05 -0700 2006"
		time.ANSIC,                   // "Mon Jan _2 15:04:05 2006"
		"02/Jan/2006:15:04:05 -0700", // Nginx/Apache standard log format
		"02/Jan/2006:15:04:05 MST",   // Nginx/Apache alternative (MST timezone)
		time.DateTime,                // "2006-01-02 15:04:05" (ISO-like format without timezone)
		"2006-01-02 15:04:05 -0700",  // ISO-like format with offset
		"2006-01-02 15:04:05 MST",    // ISO-like format with named timezone
		"2006/01/02 15:04:05",        // ISO-like format with different separator
		"2006/01/02 15:04:05 -0700",  // ISO-like format with different separator and offset
		"2006/01/02 15:04:05 MST",    // ISO-like format with different separator and named timezone
		time.DateOnly,                // "2006-01-02" (Date-only format)
		"2006-01-02T15:04:05",        // ISO 8601 datetime
		"2006-01-02T15:04:05.000",    // ISO 8601 datetime with milliseconds
	}

	var t time.Time
	var err error

	// short-circuit for unix seconds/milliseconds/nanoseconds
	if IsOnlyNumeric(input) {
		var unixTime int64
		unixTime, err = strconv.ParseInt(input, 10, 64)
		if err == nil {
			switch len(input) {
			case 9, 10:
				t = time.Unix(unixTime, 0) // Unix seconds
				return t, nil
			case 12, 13:
				t = time.UnixMilli(unixTime) // Unix milliseconds
				return t, nil
			case 18, 19:
				t = time.Unix(0, unixTime) // Unix nanoseconds
				return t, nil
			}
		}
	}

	// parse time with each format until we match
	for _, timeFormat := range timeFormats {
		t, err = time.Parse(timeFormat, input)
		if err == nil {

			if strings.Contains(timeFormat, "MST") {
				t = adjustToNamedTimezone(t, input)
			}
			return t, nil
		}
	}

	return time.Time{}, fmt.Errorf("unrecognized timestamp format: %s", input)
}

func adjustToNamedTimezone(t time.Time, originalTs string) time.Time {
	// Extract timezone string from input
	parts := strings.Fields(originalTs)
	for _, part := range parts {

		// Try Go's location database as a fallback
		loc, err := time.LoadLocation(part)
		if err == nil {
			return time.Date(
				t.Year(), t.Month(), t.Day(),
				t.Hour(), t.Minute(), t.Second(), t.Nanosecond(),
				loc,
			)
		}
	}
	return t
}
