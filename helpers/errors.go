package helpers

import (
	"errors"
	"fmt"
	"strings"
)

func CombineErrorsWithPrefix(prefix string, errorList ...error) error {
	if len(errorList) == 0 {
		return nil
	}

	if len(errorList) == 1 {
		if len(prefix) == 0 {
			return errorList[0]
		} else {
			return fmt.Errorf("%s - %s", prefix, errorList[0].Error())
		}
	}

	combinedErrorString := []string{prefix}
	for _, e := range errorList {
		combinedErrorString = append(combinedErrorString, e.Error())
	}
	return errors.New(strings.Join(combinedErrorString, "\n\t"))
}

func CombineErrors(errors ...error) error {
	return CombineErrorsWithPrefix("", errors...)
}

// ToError formats the supplied value as an error (or just returns it if already an error)
func ToError(val interface{}) error {
	if e, ok := val.(error); ok {
		return e
	} else {
		return fmt.Errorf("%v", val)
	}
}
