package helpers

func ByteMapToStringMap(m map[string][]byte) map[string]string {
	res := make(map[string]string, len(m))
	for k, v := range m {
		res[k] = string(v)
	}
	return res
}
