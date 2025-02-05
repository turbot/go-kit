package helpers

import (
	"sort"

	"golang.org/x/exp/maps"
)

// SortedMapKeys returns the sorted keys of the map `m`
func SortedMapKeys[V any](m map[string]V) []string {
	var keys = maps.Keys(m)
	sort.Strings(keys)
	return keys
}

// SliceToLookup converts a slice into a lookup
func SliceToLookup[K comparable](src []K) map[K]struct{} {
	var result = make(map[K]struct{}, len(src))
	for _, k := range src {
		result[k] = struct{}{}
	}
	return result
}

// FilterMap builds a map based on `src`, but using only keys specified in `keys`
func FilterMap[K comparable, V any](src map[K]V, keys []K) map[K]V {
	var result = make(map[K]V, len(keys))
	for _, k := range keys {
		if v, ok := src[k]; ok {
			result[k] = v
		}
	}
	return result
}
