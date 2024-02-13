package helpers

import (
	"reflect"
	"time"
)

// AnySliceToTypedSlice determines whether input is []any and if so converts to an array of the underlying type
func AnySliceToTypedSlice(input any) any {
	result := input
	switch result.(type) { //nolint:gocritic // TODO fix this gocritic lint suggestion
	case []any:
		val := reflect.ValueOf(result)
		if val.Kind() == reflect.Slice {
			if val.Len() == 0 {
				// if array is empty we cannot know the underlying type
				// just return empty string array
				result = []string{}
			} else {
				// convert into an array of the appropriate type
				elem := val.Index(0).Interface()
				switch elem.(type) {
				case int16:
					result = ToTypedSlice[int16](result.([]any))
				case int32:
					result = ToTypedSlice[int32](result.([]any))
				case int64:
					result = ToTypedSlice[int64](result.([]any))
				case float32:
					result = ToTypedSlice[float32](result.([]any))
				case float64:
					result = ToTypedSlice[float64](result.([]any))
				case string:
					result = ToTypedSlice[string](result.([]any))
				case time.Time:
					result = ToTypedSlice[time.Time](result.([]any))
				}
			}
		}
	}
	return result
}

// ToTypedSlice converts []any to []T
func ToTypedSlice[T any](input []any) []T {
	res := make([]T, len(input))
	for i, val := range input {
		res[i] = val.(T)
	}
	return res
}

// AppendSliceUnique appends elements from slice2 to slice1, omitting duplicates.
func AppendSliceUnique[T comparable](slice1, slice2 []T) []T {
	// Map existing elements of slice1 for quick lookup
	exists := SliceToLookup(slice1)
	result := make([]T, len(slice1))
	copy(result, slice1)

	// Check each element in slice2; if not a duplicate, append it to the result
	for _, item := range slice2 {
		if _, dupe := exists[item]; !dupe {
			result = append(result, item)
			exists[item] = struct{}{}
		}
	}

	return result
}
