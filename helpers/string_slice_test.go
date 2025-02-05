package helpers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//
//type stringSliceDiffTest struct {
//	Name                 string
//	Slice1               []string
//	Slice2               []string
//	ExpectedOnlyInSlice1 []string
//}
//
//var testCasesStringSliceDiff = []stringSliceDiffTest{
//	{
//		"Same",
//		[]string{"a", "b", "c"},
//		[]string{"a", "b", "c"},
//		[]string{},
//	},
//	{
//		"Same with dupes",
//		[]string{"a", "b", "b", "c"},
//		[]string{"a", "b", "c", "c"},
//		[]string{},
//	},
//	{
//		"Same with empty entries",
//		[]string{"a", "b", "", "c"},
//		[]string{"a", "b", "", "c"},
//		[]string{},
//	},
//	{
//		"Slice1 all entry enptries",
//		[]string{"", "", "", ""},
//		[]string{"a", "b", "c"},
//		[]string{},
//	},
//	{
//		"Slice1 empty",
//		[]string{},
//		[]string{"a", "b", "", "c"},
//		[]string{},
//	},
//	{
//		"Slice2 all entry enptries",
//		[]string{"a", "b", "c"},
//		[]string{"", "", "", ""},
//		[]string{"a", "b", "c"},
//	},
//	{
//		"Slice2 empty",
//		[]string{"a", "b", "c"},
//		[]string{},
//		[]string{"a", "b", "c"},
//	},
//	{
//		"Unique element in each",
//		[]string{"a", "b", "c"},
//		[]string{"b", "c", "d"},
//		[]string{"a"},
//	},
//	{
//		"Unique element in each (2)",
//		[]string{"b", "c", "d"},
//		[]string{"a", "b", "c"},
//		[]string{"d"},
//	},
//	{
//		"Unique element in each, slice 1 longer",
//		[]string{"b", "c", "1", "2", "a"},
//		[]string{"b", "c", "d"},
//		[]string{"1", "2", "a"},
//	},
//	{
//		"Unique element in each, slice 2 longer",
//		[]string{"b", "c", "d"},
//		[]string{"b", "c", "1", "2", "a"},
//		[]string{"d"},
//	},
//	{
//		"Unique element in each, duplicate in slice 1",
//		[]string{"a", "b", "b", "c"},
//		[]string{"b", "c", "d"},
//		[]string{"a"},
//	}, {
//		"Unique element in each, duplicate (2)",
//		[]string{"b", "c", "d"},
//		[]string{"a", "b", "b", "c"},
//		[]string{"d"},
//	},
//}
//
//func TestStringSliceDiff(t *testing.T) {
//	for _, test := range testCasesStringSliceDiff {
//		onlyInSlice1 := StringSliceDiff(test.Slice1, test.Slice2)
//		if !reflect.DeepEqual(test.ExpectedOnlyInSlice1, onlyInSlice1) {
//			t.Errorf(`Test: '%s'' FAILED : onlyInSlice1 expected %v, got %v`, test.Name, test.ExpectedOnlyInSlice1, onlyInSlice1)
//		}
//	}
//}
//
//type removeFromStringSliceTest struct {
//	Name     string
//	Slice1   []string
//	Values   []string
//	Expected []string
//}
//
//var testCasesRemoveFromStringSlice = []removeFromStringSliceTest{
//	{
//		"Single",
//		[]string{"a", "b", "c"},
//		[]string{"a"},
//		[]string{"b", "c"},
//	},
//	{
//		"Multiple",
//		[]string{"a", "b", "c"},
//		[]string{"a", "b"},
//		[]string{"c"},
//	},
//	{
//		"Single not there",
//		[]string{"a", "b", "c"},
//		[]string{"z"},
//		[]string{"a", "b", "c"},
//	},
//	{
//		"Multiple",
//		[]string{"a", "b", "c"},
//		[]string{"z", "g"},
//		[]string{"a", "b", "c"},
//	},
//}
//
//func TestRemoveFromStringSlice(t *testing.T) {
//	for _, test := range testCasesRemoveFromStringSlice {
//		res := RemoveFromStringSlice(test.Slice1, test.Values...)
//		if !reflect.DeepEqual(test.Expected, res) {
//			t.Errorf(`Test: '%s'' FAILED : onlyInSlice1 expected %v, got %v`, test.Name, test.Expected, res)
//		}
//	}
//}
//
//type stringSliceDistinctTest struct {
//	Name     string
//	Slice    []string
//	Expected []string
//}
//
//var testCasesStringSliceDistinct = []stringSliceDistinctTest{
//	{
//		"no dupes",
//		[]string{"A", "B"},
//		[]string{"A", "B"},
//	},
//	{
//		"single dupe",
//		[]string{"A", "B", "B"},
//		[]string{"A", "B"},
//	},
//	{
//		"multiple dupes",
//		[]string{"A", "A", "A", "A", "A", "A", "A", "B", "B", "B", "B", "B", "B"},
//		[]string{"A", "B"},
//	},
//}
//
//func TestStringSliceDistinct(t *testing.T) {
//	for _, test := range testCasesStringSliceDistinct {
//		res := StringSliceDistinct(test.Slice)
//		if !reflect.DeepEqual(res, test.Expected) {
//			t.Errorf(`Test: '%s'' FAILED : expected %v, got %v`, test.Name, test.Expected, res)
//		}
//	}
//}
//
//type stringSliceHasDuplicatesTest struct {
//	Name     string
//	Slice    []string
//	Expected bool
//}
//
//var testCasesStringSliceHasDuplicates = []stringSliceHasDuplicatesTest{
//	{
//		"no dupes",
//		[]string{"A", "B"},
//		false,
//	},
//	{
//		"single dupe",
//		[]string{"A", "B", "B"},
//		true,
//	},
//	{
//		"multiple dupes",
//		[]string{"A", "A", "A", "A", "A", "A", "A", "B", "B", "B", "B", "B", "B"},
//		true,
//	},
//}
//
//func TestStringSliceHasDuplicates(t *testing.T) {
//	for _, test := range testCasesStringSliceHasDuplicates {
//		res := StringSliceHasDuplicates(test.Slice)
//		if !reflect.DeepEqual(res, test.Expected) {
//			t.Errorf(`Test: '%s'' FAILED : expected %v, got %v`, test.Name, test.Expected, res)
//		}
//	}
//}
//
//type stringSliceEqualIgnoreOrderTest struct {
//	Name     string
//	A        []string
//	B        []string
//	Expected bool
//}
//
//var testCaseStringSliceEqualIgnoreOrder = []stringSliceEqualIgnoreOrderTest{
//	{
//		"equal ignore order",
//		[]string{"A", "B", "c"},
//		[]string{"c", "A", "B"},
//		true,
//	},
//	{
//		"equal ignore order 2",
//		[]string{"feeder", "just", "a", "day"},
//		[]string{"just", "feeder", "a", "day"},
//		true,
//	},
//	{
//		"not equal",
//		[]string{"A", "B", "B"},
//		[]string{"A", "B", "C"},
//		false,
//	},
//	{
//		"not equal",
//		[]string{"feeders", "just", "a", "day"},
//		[]string{"just", "feeder", "a", "day"},
//		false,
//	},
//}
//
//func TestStringSliceEqualIgnoreOrder(t *testing.T) {
//	for _, test := range testCaseStringSliceEqualIgnoreOrder {
//		res := StringSliceEqualIgnoreOrder(test.A, test.B)
//		if !reflect.DeepEqual(res, test.Expected) {
//			t.Errorf(`Test: '%s'' FAILED : expected %v, got %v`, test.Name, test.Expected, res)
//		}
//	}
//}

func TestRemoveFromStringSlice(t *testing.T) {
	type args struct {
		slice  []string
		values []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Remove single element",
			args: args{slice: []string{"apple", "banana", "cherry"}, values: []string{"banana"}},
			want: []string{"apple", "cherry"},
		},
		{
			name: "Remove multiple elements",
			args: args{slice: []string{"apple", "banana", "cherry", "date"}, values: []string{"banana", "date"}},
			want: []string{"apple", "cherry"},
		},
		{
			name: "Remove all elements",
			args: args{slice: []string{"apple", "banana"}, values: []string{"apple", "banana"}},
			want: []string{},
		},
		{
			name: "Remove non-existent element",
			args: args{slice: []string{"apple", "banana"}, values: []string{"cherry"}},
			want: []string{"apple", "banana"}, // No change
		},
		{
			name: "Remove from empty slice",
			args: args{slice: []string{}, values: []string{"banana"}},
			want: []string{},
		},
		{
			name: "Remove empty values (no-op)",
			args: args{slice: []string{"apple", "banana"}, values: []string{}},
			want: []string{"apple", "banana"},
		},
		{
			name: "Remove duplicate occurrences",
			args: args{slice: []string{"apple", "banana", "banana", "cherry"}, values: []string{"banana"}},
			want: []string{"apple", "cherry"},
		},
		{
			name: "Remove empty string",
			args: args{slice: []string{"apple", "banana", "", "cherry"}, values: []string{""}},
			want: []string{"apple", "banana", "cherry"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(
				t,
				tt.want,
				RemoveFromStringSlice(tt.args.slice, tt.args.values...),
				"RemoveFromStringSlice(%v, %v)",
				tt.args.slice,
				tt.args.values,
			)
		})
	}
}

func TestStringSliceDiff(t *testing.T) {
	type args struct {
		slice1 []string
		slice2 []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Elements only in slice1",
			args: args{slice1: []string{"apple", "banana", "cherry"}, slice2: []string{"banana"}},
			want: []string{"apple", "cherry"},
		},
		{
			name: "All elements in slice1 are in slice2",
			args: args{slice1: []string{"apple", "banana"}, slice2: []string{"apple", "banana"}},
			want: []string{}, // No elements unique to slice1
		},
		{
			name: "No common elements",
			args: args{slice1: []string{"apple", "banana"}, slice2: []string{"cherry", "date"}},
			want: []string{"apple", "banana"}, // slice1 is unchanged
		},
		{
			name: "Empty slice1",
			args: args{slice1: []string{}, slice2: []string{"apple", "banana"}},
			want: []string{}, // Nothing in slice1 to compare
		},
		{
			name: "Empty slice2",
			args: args{slice1: []string{"apple", "banana"}, slice2: []string{}},
			want: []string{"apple", "banana"}, // Everything in slice1 remains
		},
		{
			name: "Both slices empty",
			args: args{slice1: []string{}, slice2: []string{}},
			want: []string{}, // Nothing to return
		},
		{
			name: "Duplicates in slice1",
			args: args{slice1: []string{"apple", "banana", "banana", "cherry"}, slice2: []string{"banana"}},
			want: []string{"apple", "cherry"},
		},
		{
			name: "Removing empty string",
			args: args{slice1: []string{"apple", "banana", "", "cherry"}, slice2: []string{""}},
			want: []string{"apple", "banana", "cherry"},
		},
		{
			name: "Case-sensitive comparison",
			args: args{slice1: []string{"Apple", "banana"}, slice2: []string{"apple"}},
			want: []string{"Apple", "banana"}, // "Apple" ≠ "apple"
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StringSliceDiff(tt.args.slice1, tt.args.slice2)
			assert.Equalf(t, tt.want, got, "StringSliceDiff(%v, %v)", tt.args.slice1, tt.args.slice2)
		})
	}
}

func TestStringSliceDistinct(t *testing.T) {
	type args struct {
		slice []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "No duplicates",
			args: args{slice: []string{"apple", "banana", "cherry"}},
			want: []string{"apple", "banana", "cherry"}, // Unchanged
		},
		{
			name: "With duplicates",
			args: args{slice: []string{"apple", "banana", "banana", "cherry", "apple"}},
			want: []string{"apple", "banana", "cherry"}, // Order preserved, duplicates removed
		},
		{
			name: "All duplicates",
			args: args{slice: []string{"apple", "apple", "apple"}},
			want: []string{"apple"}, // Only one instance remains
		},
		{
			name: "Empty slice",
			args: args{slice: []string{}},
			want: []string{}, // Should return an empty slice
		},
		{
			name: "Single element",
			args: args{slice: []string{"apple"}},
			want: []string{"apple"}, // Unchanged
		},
		{
			name: "Case-sensitive distinct",
			args: args{slice: []string{"apple", "Apple", "banana", "Banana"}},
			want: []string{"apple", "Apple", "banana", "Banana"}, // "apple" ≠ "Apple"
		},
		{
			name: "Empty string as a distinct value",
			args: args{slice: []string{"apple", "", "banana", "", "cherry"}},
			want: []string{"apple", "", "banana", "cherry"}, // Empty string retained
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, StringSliceDistinct(tt.args.slice), "StringSliceDistinct(%v)", tt.args.slice)
		})
	}
}

func TestStringSliceEqualIgnoreOrder(t *testing.T) {
	type args struct {
		a []string
		b []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Same elements, same order",
			args: args{a: []string{"apple", "banana", "cherry"}, b: []string{"apple", "banana", "cherry"}},
			want: true,
		},
		{
			name: "Same elements, different order",
			args: args{a: []string{"apple", "banana", "cherry"}, b: []string{"cherry", "banana", "apple"}},
			want: true,
		},
		{
			name: "Different elements",
			args: args{a: []string{"apple", "banana"}, b: []string{"apple", "cherry"}},
			want: false,
		},
		{
			name: "Different lengths",
			args: args{a: []string{"apple", "banana", "cherry"}, b: []string{"apple", "banana"}},
			want: false,
		},
		{
			name: "Both empty",
			args: args{a: []string{}, b: []string{}},
			want: true,
		},
		{
			name: "One empty, one non-empty",
			args: args{a: []string{"apple", "banana"}, b: []string{}},
			want: false,
		},
		{
			name: "Duplicates in one slice",
			args: args{a: []string{"apple", "banana", "banana"}, b: []string{"apple", "banana"}},
			want: false,
		},
		{
			name: "Case-sensitive comparison",
			args: args{a: []string{"apple", "Banana"}, b: []string{"apple", "banana"}},
			want: false, // "Banana" ≠ "banana"
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, StringSliceEqualIgnoreOrder(tt.args.a, tt.args.b), "StringSliceEqualIgnoreOrder(%v, %v)", tt.args.a, tt.args.b)
		})
	}
}

func TestStringSliceHasDuplicates(t *testing.T) {
	type args struct {
		slice []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "No duplicates",
			args: args{slice: []string{"apple", "banana", "cherry"}},
			want: false,
		},
		{
			name: "One duplicate",
			args: args{slice: []string{"apple", "banana", "apple"}},
			want: true,
		},
		{
			name: "Multiple duplicates",
			args: args{slice: []string{"apple", "banana", "cherry", "banana", "cherry"}},
			want: true,
		},
		{
			name: "All duplicates",
			args: args{slice: []string{"apple", "apple", "apple"}},
			want: true,
		},
		{
			name: "Single element",
			args: args{slice: []string{"apple"}},
			want: false,
		},
		{
			name: "Empty slice",
			args: args{slice: []string{}},
			want: false,
		},
		{
			name: "Case-sensitive check",
			args: args{slice: []string{"apple", "Apple"}},
			want: false, // "apple" ≠ "Apple"
		},
		{
			name: "Includes empty string duplicates",
			args: args{slice: []string{"", "apple", "", "banana"}},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, StringSliceHasDuplicates(tt.args.slice), "StringSliceHasDuplicates(%v)", tt.args.slice)
		})
	}
}
