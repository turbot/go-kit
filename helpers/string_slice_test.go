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

type stringSliceDiffTest struct {
	Name                 string
	Slice1               []string
	Slice2               []string
	ExpectedOnlyInSlice1 []string
}

var testCasesStringSliceDiff = []stringSliceDiffTest{
	{
		"Same",
		[]string{"a", "b", "c"},
		[]string{"a", "b", "c"},
		[]string{},
	},
	{
		"Same with dupes",
		[]string{"a", "b", "b", "c"},
		[]string{"a", "b", "c", "c"},
		[]string{},
	},
	{
		"Same with empty entries",
		[]string{"a", "b", "", "c"},
		[]string{"a", "b", "", "c"},
		[]string{},
	},
	{
		"Slice1 all entry enptries",
		[]string{"", "", "", ""},
		[]string{"a", "b", "c"},
		[]string{},
	},
	{
		"Slice1 empty",
		[]string{},
		[]string{"a", "b", "", "c"},
		[]string{},
	},
	{
		"Slice2 all entry enptries",
		[]string{"a", "b", "c"},
		[]string{"", "", "", ""},
		[]string{"a", "b", "c"},
	},
	{
		"Slice2 empty",
		[]string{"a", "b", "c"},
		[]string{},
		[]string{"a", "b", "c"},
	},
	{
		"Unique element in each",
		[]string{"a", "b", "c"},
		[]string{"b", "c", "d"},
		[]string{"a"},
	},
	{
		"Unique element in each (2)",
		[]string{"b", "c", "d"},
		[]string{"a", "b", "c"},
		[]string{"d"},
	},
	{
		"Unique element in each, slice 1 longer",
		[]string{"b", "c", "1", "2", "a"},
		[]string{"b", "c", "d"},
		[]string{"1", "2", "a"},
	},
	{
		"Unique element in each, slice 2 longer",
		[]string{"b", "c", "d"},
		[]string{"b", "c", "1", "2", "a"},
		[]string{"d"},
	},
	{
		"Unique element in each, duplicate in slice 1",
		[]string{"a", "b", "b", "c"},
		[]string{"b", "c", "d"},
		[]string{"a"},
	}, {
		"Unique element in each, duplicate (2)",
		[]string{"b", "c", "d"},
		[]string{"a", "b", "b", "c"},
		[]string{"d"},
	},
}

func TestStringSliceDiff(t *testing.T) {
	for _, test := range testCasesStringSliceDiff {
		onlyInSlice1 := StringSliceDiff(test.Slice1, test.Slice2)
		if !reflect.DeepEqual(test.ExpectedOnlyInSlice1, onlyInSlice1) {
			t.Errorf(`Test: '%s'' FAILED : onlyInSlice1 expected %v, got %v`, test.Name, test.ExpectedOnlyInSlice1, onlyInSlice1)
		}
	}
}

type removeFromStringSliceTest struct {
	Name     string
	Slice1   []string
	Values   []string
	Expected []string
}

var testCasesRemoveFromStringSlice = []removeFromStringSliceTest{
	{
		"Single",
		[]string{"a", "b", "c"},
		[]string{"a"},
		[]string{"b", "c"},
	},
	{
		"Multiple",
		[]string{"a", "b", "c"},
		[]string{"a", "b"},
		[]string{"c"},
	},
	{
		"Single not there",
		[]string{"a", "b", "c"},
		[]string{"z"},
		[]string{"a", "b", "c"},
	},
	{
		"Multiple",
		[]string{"a", "b", "c"},
		[]string{"z", "g"},
		[]string{"a", "b", "c"},
	},
}

func TestRemoveFromStringSlice(t *testing.T) {
	for _, test := range testCasesRemoveFromStringSlice {
		res := RemoveFromStringSlice(test.Slice1, test.Values...)
		if !reflect.DeepEqual(test.Expected, res) {
			t.Errorf(`Test: '%s'' FAILED : onlyInSlice1 expected %v, got %v`, test.Name, test.Expected, res)
		}
	}
}

type stringSliceDistinctTest struct {
	Name     string
	Slice    []string
	Expected []string
}

var testCasesStringSliceDistinct = []stringSliceDistinctTest{
	{
		"no dupes",
		[]string{"A", "B"},
		[]string{"A", "B"},
	},
	{
		"single dupe",
		[]string{"A", "B", "B"},
		[]string{"A", "B"},
	},
	{
		"multiple dupes",
		[]string{"A", "A", "A", "A", "A", "A", "A", "B", "B", "B", "B", "B", "B"},
		[]string{"A", "B"},
	},
}

func TestStringSliceDistinct(t *testing.T) {
	for _, test := range testCasesStringSliceDistinct {
		res := StringSliceDistinct(test.Slice)
		if !reflect.DeepEqual(res, test.Expected) {
			t.Errorf(`Test: '%s'' FAILED : expected %v, got %v`, test.Name, test.Expected, res)
		}
	}
}

type stringSliceHasDuplicatesTest struct {
	Name     string
	Slice    []string
	Expected bool
}

var testCasesStringSliceHasDuplicates = []stringSliceHasDuplicatesTest{
	{
		"no dupes",
		[]string{"A", "B"},
		false,
	},
	{
		"single dupe",
		[]string{"A", "B", "B"},
		true,
	},
	{
		"multiple dupes",
		[]string{"A", "A", "A", "A", "A", "A", "A", "B", "B", "B", "B", "B", "B"},
		true,
	},
}

func TestStringSliceHasDuplicates(t *testing.T) {
	for _, test := range testCasesStringSliceHasDuplicates {
		res := StringSliceHasDuplicates(test.Slice)
		if !reflect.DeepEqual(res, test.Expected) {
			t.Errorf(`Test: '%s'' FAILED : expected %v, got %v`, test.Name, test.Expected, res)
		}
	}
}

type stringSliceEqualIgnoreOrderTest struct {
	Name     string
	A        []string
	B        []string
	Expected bool
}

var testCaseStringSliceEqualIgnoreOrder = []stringSliceEqualIgnoreOrderTest{
	{
		"equal ignore order",
		[]string{"A", "B", "c"},
		[]string{"c", "A", "B"},
		true,
	},
	{
		"equal ignore order 2",
		[]string{"feeder", "just", "a", "day"},
		[]string{"just", "feeder", "a", "day"},
		true,
	},
	{
		"not equal",
		[]string{"A", "B", "B"},
		[]string{"A", "B", "C"},
		false,
	},
	{
		"not equal",
		[]string{"feeders", "just", "a", "day"},
		[]string{"just", "feeder", "a", "day"},
		false,
	},
}

func TestStringSliceEqualIgnoreOrder(t *testing.T) {
	for _, test := range testCaseStringSliceEqualIgnoreOrder {
		res := StringSliceEqualIgnoreOrder(test.A, test.B)
		if !reflect.DeepEqual(res, test.Expected) {
			t.Errorf(`Test: '%s'' FAILED : expected %v, got %v`, test.Name, test.Expected, res)
		}
	}
}
