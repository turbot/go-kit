package helpers

import (
	"testing"
	"time"
)

type parseTimeTest struct {
	name     string
	input    string
	expected time.Time
	wantErr  bool
}

var testCasesParseTime = []parseTimeTest{
	{
		name:     "RFC3339Nano",
		input:    "2024-02-10T15:04:05.123456789Z",
		expected: time.Date(2024, 2, 10, 15, 4, 5, 123456789, time.UTC),
		wantErr:  false,
	},
	{
		name:     "RFC3339",
		input:    "2024-02-10T15:04:05Z",
		expected: time.Date(2024, 2, 10, 15, 4, 5, 0, time.UTC),
		wantErr:  false,
	},
	{
		name:     "RFC1123Z",
		input:    "Sat, 10 Feb 2024 15:04:05 GMT",
		expected: time.Date(2024, 2, 10, 15, 4, 5, 0, time.UTC),
		wantErr:  false,
	},
	{
		name:     "RFC1123NamedTimezone",
		input:    "Sat, 10 Feb 2024 15:04:05 UTC",
		expected: time.Date(2024, 2, 10, 15, 4, 5, 0, time.UTC),
		wantErr:  false,
	},
	{
		name:     "RFC850",
		input:    "Saturday, 10-Feb-24 15:04:05 UTC",
		expected: time.Date(2024, 2, 10, 15, 4, 5, 0, time.UTC),
		wantErr:  false,
	},
	{
		name:     "RFC822",
		input:    "10 Feb 24 15:04 UTC",
		expected: time.Date(2024, 2, 10, 15, 4, 0, 0, time.UTC),
		wantErr:  false,
	},
	{
		name:     "UnixDate",
		input:    "Tue Feb 10 15:04:05 UTC 2024",
		expected: time.Date(2024, 2, 10, 15, 4, 5, 0, time.UTC),
		wantErr:  false,
	},
	{
		name:     "RubyDate",
		input:    "Sat Feb 10 15:04:05 UTC 2024",
		expected: time.Date(2024, 2, 10, 15, 4, 5, 0, time.UTC),
		wantErr:  false,
	},
	{
		name:     "NginxApacheStandard",
		input:    "10/Feb/2024:15:04:05 -0500",
		expected: time.Date(2024, 2, 10, 15, 4, 5, 0, time.FixedZone("-0500", -5*60*60)),
		wantErr:  false,
	},
	{
		name:     "UnixMilliseconds",
		input:    "1707583205123",
		expected: time.UnixMilli(1707583205123),
		wantErr:  false,
	},
	{
		name:     "NginxApacheNamedTimezone",
		input:    "10/Feb/2024:15:04:05 EST",
		expected: time.Date(2024, 2, 10, 15, 4, 5, 0, time.FixedZone("EST", -5*60*60)),
		wantErr:  false,
	},
	{
		name:     "ISOFormatWithOffset",
		input:    "2024-02-10 15:04:05 -0700",
		expected: time.Date(2024, 2, 10, 15, 4, 5, 0, time.FixedZone("-0700", -7*60*60)),
		wantErr:  false,
	},
	{
		name:     "ISOFormatWithNamedTimezone",
		input:    "2024-02-10 15:04:05 UTC",
		expected: time.Date(2024, 2, 10, 15, 4, 5, 0, time.UTC),
		wantErr:  false,
	},
	{
		name:     "DateTimeNoTimezone",
		input:    "2024-02-10 15:04:05",
		expected: time.Date(2024, 2, 10, 15, 4, 5, 0, time.UTC),
		wantErr:  false,
	},
	{
		name:     "DateOnly",
		input:    "2024-02-10",
		expected: time.Date(2024, 2, 10, 0, 0, 0, 0, time.UTC),
		wantErr:  false,
	},
	{
		name:     "IncorrectOrder",
		input:    "10-02-2024 15:04:05",
		expected: time.Time{},
		wantErr:  true,
	},
	{
		name:     "AlternateSeparator",
		input:    "2024/02/10 15:04:05",
		expected: time.Date(2024, 02, 10, 15, 4, 5, 0, time.UTC),
		wantErr:  false,
	},
	{
		name:     "ShortenedYear",
		input:    "24-02-10 15:04:05",
		expected: time.Time{},
		wantErr:  true,
	},
	{
		name:     "OddOrdering",
		input:    "Saturday 10 Feb 2024 15:04:05",
		expected: time.Time{},
		wantErr:  true,
	},
	{
		name:     "TimeFirst",
		input:    "23:59:59 2024-02-10",
		expected: time.Time{},
		wantErr:  true,
	},
	{
		name:     "RFC1123ZWithOffset",
		input:    "Sun, 11 Feb 2024 03:15:30 +0100",
		expected: time.Date(2024, 2, 11, 3, 15, 30, 0, time.FixedZone("+0100", 1*60*60)),
		wantErr:  false,
	},
	{
		name:     "RandomString",
		input:    "invalid timestamp",
		expected: time.Time{},
		wantErr:  true,
	},
	{
		name:     "WrittenOutDate",
		input:    "10th of February, 2024",
		expected: time.Time{},
		wantErr:  true,
	},
	{
		name:     "NonParsableWord",
		input:    "Yesterday",
		expected: time.Time{},
		wantErr:  true,
	},
	{
		name:     "UnixNanoseconds",
		input:    "1707583205123456789",
		expected: time.Unix(0, 1707583205123456789),
		wantErr:  false,
	},
	{
		name:     "ImpossibleTime",
		input:    "2024-02-10T99:99:99Z",
		expected: time.Time{},
		wantErr:  true,
	},
	{
		name:     "UnixSecondsPre2001",
		input:    "534611045",
		expected: time.Unix(534611045, 0),
		wantErr:  false,
	},
	{
		name:     "UnixMillisecondsPre2001",
		input:    "534611045123",
		expected: time.UnixMilli(534611045123),
		wantErr:  false,
	},
	{
		name:     "UnixNanosecondsPre2001",
		input:    "534611045123456789",
		expected: time.Unix(0, 534611045123456789),
		wantErr:  false,
	},
	{
		name:     "UnixSeconds",
		input:    "1234567890",
		expected: time.Unix(1234567890, 0),
		wantErr:  false,
	},
	{
		name:     "RFC3339WithOffset",
		input:    "2024-02-10T15:04:05+02:00",
		expected: time.Date(2024, 2, 10, 15, 4, 5, 0, time.FixedZone("+0200", 2*60*60)),
		wantErr:  false,
	},
	{
		name:     "RFC1123ZWithOffset",
		input:    "Sat, 10 Feb 2024 15:04:05 +0100",
		expected: time.Date(2024, 2, 10, 15, 4, 5, 0, time.FixedZone("+0100", 1*60*60)),
		wantErr:  false,
	},
	{
		name:     "NginxApacheStandardRepeat",
		input:    "10/Feb/2024:15:04:05 -0500",
		expected: time.Date(2024, 2, 10, 15, 4, 5, 0, time.FixedZone("-0500", -5*60*60)),
		wantErr:  false,
	},
}

func TestParseTime(t *testing.T) {
	for _, tt := range testCasesParseTime {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseTime(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !result.Equal(tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
