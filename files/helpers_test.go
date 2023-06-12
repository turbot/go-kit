package files

import (
	"path/filepath"
	"testing"
)

type fileExistsTest struct {
	path     string
	expected bool
}

var fileExistsTests = map[string]fileExistsTest{
	"self": {
		path:     filepath.Join(wd, "helpers_test.go"),
		expected: true,
	},
	"long SQL": {
		// use a long SQL query which is one of the inputs that FileExists receives when resolving arguments to 'steampipe query'
		path:     "create table all_columns (nullcolumn CHAR(2), booleancolumn boolean, textcolumn1 CHAR(20), textcolumn2 VARCHAR(20),  textcolumn3 text, integercolumn1 smallint, integercolumn2 int, integercolumn3 SERIAL, integercolumn4 bigint,  integercolumn5 bigserial, numericColumn numeric(6,4), realColumn real, floatcolumn float,  date1 DATE,  time1 TIME,  timestamp1 TIMESTAMP, timestamp2 TIMESTAMPTZ, interval1 INTERVAL, array1 text[], jsondata jsonb, jsondata2 json, uuidcolumn UUID, ipAddress inet, macAddress macaddr, cidrRange cidr, xmlData xml, currency money)",
		expected: false,
	},
}

func TestFileExists(t *testing.T) {
	for k, fet := range fileExistsTests {
		exists := FileExists(fet.path)
		if exists != fet.expected {
			t.Errorf("Test: '%s'' FAILED - expected %v. got %v", k, fet.expected, exists)
			continue
		}
	}
}
