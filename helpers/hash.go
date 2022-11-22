package helpers

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"hash/fnv"
	"io"
	"os"
)

// FileMD5Hash streams a file into a MD5 hasher and returns the resultant MD5 hash bytes
// This DOES NOT read the whole file into memory
func FileMD5Hash(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	hasher := md5.New()
	if _, err := io.Copy(hasher, file); err != nil {
		return nil, err
	}
	return hasher.Sum(nil), nil
}

// FileHash streams a file into a MD5 hasher and returns the resultant MD5 hash
// as a HEX encoded string. Uses FileMD5Hash under-the-hood
func FileHash(filePath string) (string, error) {
	hash, err := FileMD5Hash(filePath)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hash), nil
}

// StringFnvHash returns a FNV1 hash of the given string.
// returns 0 if there's an error
func StringFnvHash(s string) uint32 {
	b := bytes.NewBufferString(s)
	h := fnv.New32a()
	if _, err := io.Copy(h, b); err != nil {
		return 0
	}
	return h.Sum32()
}

// GetMD5Hash returns the MD5 hash of the given string
func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	h := hex.EncodeToString(hash[:])
	return h
}
