package helpers

import (
	"fmt"
	"strings"
)

func CombineErrorsWithPrefix(prefix string, errors ...error) error {
	if len(errors) == 0 {
		return nil
	}

	if len(errors) == 1 {
		if len(prefix) == 0 {
			return errors[0]
		} else {
			return fmt.Errorf("%s - %s", prefix, errors[0].Error())
		}
	}

	combinedErrorString := []string{prefix}
	for _, e := range errors {
		combinedErrorString = append(combinedErrorString, e.Error())
	}
	return fmt.Errorf(strings.Join(combinedErrorString, "\n\t"))
}

func CombineErrors(errors ...error) error {
	return CombineErrorsWithPrefix("", errors...)
}
