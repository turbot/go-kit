package helpers

import (
	"path"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
)

// GetCallingFunction :: return the calling function
/* level is how far up the call stack to go. so

func f1(){
	f2()
}

func f2(){
	// returns "f2"
	helpers.GetCallingFunction(0)

	// returns "f1"
	helpers.GetCallingFunction(1)
}

*/
func GetCallingFunction(level int) string {
	fpcs := make([]uintptr, 1)
	// Skip 2 levels to get the caller
	n := runtime.Callers(level+2, fpcs)
	if n == 0 {
		return ""
	}

	caller := runtime.FuncForPC(fpcs[0] - 1)
	if caller == nil {
		return ""
	}
	return path.Ext(filepath.Base(caller.Name()))[1:]
}

// GetFunctionName :: return the name of the given function
func GetFunctionName(i interface{}) string {
	name := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	// split on '.' and return final element
	split := strings.Split(name, ".")
	return split[len(split)-1]
}
