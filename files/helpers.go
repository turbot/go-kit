package files

import (
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

// Tildefy converts ~ to home directory
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
// FileExists uses `os.Stat` under the hood and returns `false` if `os.Stat` returns an error
func FileExists(filename string) bool {
	if info, err := os.Stat(filename); err == nil {
		return !info.IsDir()
	}
	return false
}

// DirectoryExists checks if a path exists and is a directory
// DirectoryExists uses `os.Stat` under the hood and returns `false` if `os.Stat` returns an error
func DirectoryExists(dirname string) bool {
	if info, err := os.Stat(dirname); err == nil {
		return info.IsDir()
	}
	return false
}
