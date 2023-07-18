//go:build !windows

package files

import (
	"syscall"
)

func SetULimit(ulimit uint64) error {
	var rLimit syscall.Rlimit
	rLimit.Max = ulimit
	rLimit.Cur = ulimit
	return syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit)
}

func GetULimit() (syscall.Rlimit, error) {
	var rLimit syscall.Rlimit
	err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	return rLimit, err
}
