package helpers

import (
	"sort"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func StringSliceEqualIgnoreOrder(a, b []string) bool {
	less := func(a, b string) bool { return a < b }
	return cmp.Diff(a, b, cmpopts.SortSlices(less)) == ""

}

// StringSliceContains returns whether the string slice contains the given string
func StringSliceContains(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

// StringSliceDiff returns values which only exists in the fist string slice
func StringSliceDiff(slice1, slice2 []string) (onlyInSlice1 []string) {
	onlyInSlice1 = []string{}
	if len(slice1) == 0 {
		return
	}
	if len(slice2) == 0 {
		onlyInSlice1 = slice1
		return
	}

	sort.Strings(slice1)
	sort.Strings(slice2)

	idx := 0
	for _, item := range slice1 {
		// ignore empty
		if item == "" {
			continue
		}

		for slice2[idx] < item && idx+1 < len(slice2) {
			idx++
		}
		if slice2[idx] != item {
			onlyInSlice1 = append(onlyInSlice1, item)
		}
	}
	return
}
