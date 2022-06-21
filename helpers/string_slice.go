package helpers

import "sort"

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

// RemoveFromStringSlice removes the given string from the string slice
func RemoveFromStringSlice(slice []string, values ...string) []string {
	var res []string
	for _, item := range slice {
		var remove bool
		for _, value := range values {
			if item == value {
				remove = true
				break
			}
		}
		if !remove {
			res = append(res, item)
		}
	}
	return res
}

// StringSliceDistinct returns a slice with the unique elements the input string slice
func StringSliceDistinct(slice []string) []string {
	var res []string
	countMap := make(map[string]int)
	for _, item := range slice {
		countMap[item]++
		// if this is the first time we have come across this item, add to res
		if countMap[item] == 1 {
			res = append(res, item)
		}
	}
	return res
}

// StringSliceHasDuplicates returns whether a string slice has duplicate elements
func StringSliceHasDuplicates(slice []string) bool {
	return len(slice) > len(StringSliceDistinct(slice))
}
