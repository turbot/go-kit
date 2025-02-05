package helpers

import (
	"slices"
)

func StringSliceEqualIgnoreOrder(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	slices.Sort(a)
	slices.Sort(b)
	return slices.Equal(a, b)
}

// StringSliceDiff finds elements in slice1 that are not in slice2.
func StringSliceDiff(slice1, slice2 []string) []string {
	if len(slice1) == 0 {
		return []string{} // Return an empty slice instead of nil
	}
	if len(slice2) == 0 {
		return slices.Clone(slice1) // Return copy to preserve input order
	}

	// Convert slice2 into a map for O(1) lookups
	set := make(map[string]struct{}, len(slice2))
	for _, v := range slice2 {
		set[v] = struct{}{}
	}

	// Collect items from slice1 that aren't in slice2
	var onlyInSlice1 []string
	for _, v := range slice1 {
		if _, found := set[v]; !found && v != "" {
			onlyInSlice1 = append(onlyInSlice1, v)
		}
	}

	// Ensure we return an empty slice instead of nil
	if onlyInSlice1 == nil {
		return []string{}
	}
	return onlyInSlice1
}

// RemoveFromStringSlice removes the given values from the slice.
func RemoveFromStringSlice(slice []string, values ...string) []string {
	// Convert values to a set for O(1) lookups
	removeSet := make(map[string]struct{}, len(values))
	for _, v := range values {
		removeSet[v] = struct{}{}
	}

	// Use slices.DeleteFunc to filter
	return slices.DeleteFunc(slice, func(item string) bool {
		_, found := removeSet[item]
		return found
	})
}

func StringSliceDistinct(slice []string) []string {
	if len(slice) == 0 {
		return []string{} // Ensure empty slice instead of nil
	}

	seen := make(map[string]struct{}, len(slice))
	var res []string

	for _, v := range slice {
		if _, exists := seen[v]; !exists {
			seen[v] = struct{}{}
			res = append(res, v)
		}
	}

	return res
}

// StringSliceHasDuplicates returns whether a string slice has duplicate elements
func StringSliceHasDuplicates(slice []string) bool {
	seen := make(map[string]struct{}, len(slice))
	for _, item := range slice {
		if _, exists := seen[item]; exists {
			return true // Found a duplicate early
		}
		seen[item] = struct{}{}
	}
	return false
}
