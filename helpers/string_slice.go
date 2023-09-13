package helpers

// StringSliceContains returns whether the string slice contains the given string
func StringSliceContains(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
