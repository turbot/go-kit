package helpers

// MergeStringMaps :: merge 2 string maps, returning a new map
func MergeStringMaps(a, b map[string]string) map[string]string {
	var dst = make(map[string]string)
	for k, v := range a {
		dst[k] = v
	}
	for k, v := range b {
		dst[k] = v
	}
	return dst
}

// CloneStringMap :: clone a string map - return a copy of the map
func CloneStringMap(src map[string]string) map[string]string {
	var dst = make(map[string]string)

	for k, v := range src {
		dst[k] = v
	}
	return dst
}
