package helpers

import (
	"os/user"
	"path/filepath"
	"strings"

	"github.com/turbot/go-kit/files"
)

// Tildefy converts ~ to home directory
// Deprecated: use files.Tildefy
func Tildefy(filePath string) (string, error) {
	if filePath == "~" || strings.HasPrefix(filePath, "~/") {
		usr, err := user.Current()
		if err != nil {
			return "", err
		}
		var homeDir = usr.HomeDir
		if filePath == "~" {
			// In case of "~", which won't be caught by the "else if"
			filePath = homeDir
		} else if strings.HasPrefix(filePath, "~/") {
			// Use strings.HasPrefix so we don't match paths like
			// "/something/~/something/"
			filePath = filepath.Join(homeDir, filePath[2:])
		}
	}
	return filepath.Abs(filePath)

}

// FileExists checks if a file exists and is not a directory
// Deprecated: use files.FileExists
var FileExists = files.FileExists

// func FileExists(filename string) bool {
// 	info, err := os.Stat(filename)
// 	if os.IsNotExist(err) {
// 		return false
// 	}
// 	return !info.IsDir()
// }
