package helpers

import (
	"reflect"
	"testing"
)

type stringSliceContainsTest struct {
	Name     string
	Slice    []string
	Value    string
	Expected bool
}

var testCasesStringSliceContains = []stringSliceContainsTest{
	{
		"contains",
		[]string{"A", "B"},
		"A",
		true,
	},
	{
		"does not contains",
		[]string{"A", "B"},
		"Z",
		false,
	},
}

func TestStringSliceContains(t *testing.T) {
	for _, test := range testCasesStringSliceContains {
		res := StringSliceContains(test.Slice, test.Value)
		if !reflect.DeepEqual(res, test.Expected) {
			t.Errorf(`Test: '%s'' FAILED : expected %v, got %v`, test.Name, test.Expected, res)
		}
	}
}
