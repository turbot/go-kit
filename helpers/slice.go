package helpers

import (
	"reflect"
)

// AnySliceToTypedSlice converts a slice of `any` (`[]interface{}`) into a strongly typed slice
// If mixed types are detected, it returns the original `[]any` to prevent panics.
func AnySliceToTypedSlice(input any) any {
	val := reflect.ValueOf(input)

	// If input is NOT a slice, return it unchanged
	if val.Kind() != reflect.Slice {
		return input
	}

	// Handle empty slices safely
	if val.Len() == 0 {
		return []string{} // Default empty slice type
	}

	// Detect if the slice has mixed types
	elemType := reflect.TypeOf(val.Index(0).Interface())
	for i := 1; i < val.Len(); i++ {
		if reflect.TypeOf(val.Index(i).Interface()) != elemType {
			return input // 🚨 Mixed types detected, return as []any
		}
	}

	// Create a new slice of the inferred type
	typedSlice := reflect.MakeSlice(reflect.SliceOf(elemType), val.Len(), val.Len())

	// Convert elements and set them properly
	for i := 0; i < val.Len(); i++ {
		item := val.Index(i).Interface()

		// Convert item explicitly to the correct type
		convertedItem := reflect.ValueOf(item).Convert(elemType)

		// Now set it safely
		typedSlice.Index(i).Set(convertedItem)
	}

	return typedSlice.Interface()
}

// AppendSliceUnique appends elements from slice2 to slice1, ensuring uniqueness.
func AppendSliceUnique[T comparable](slice1, slice2 []T) []T {
	// Create a map for fast lookup (O(1) average time complexity)
	exists := make(map[T]struct{}, len(slice1)+len(slice2))

	// Copy slice1 into result and mark existing elements
	result := make([]T, 0, len(slice1)+len(slice2))
	for _, item := range slice1 {
		result = append(result, item)
		exists[item] = struct{}{}
	}

	// Append only unique items from slice2
	for _, item := range slice2 {
		if _, found := exists[item]; !found {
			result = append(result, item)
			exists[item] = struct{}{}
		}
	}

	return result
}
