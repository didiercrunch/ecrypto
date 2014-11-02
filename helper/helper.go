package helper

import (
	"errors"
	"strings"
)

func JoinAsError(msg []string) error {
	if len(msg) > 0 {
		return errors.New(strings.Join(msg, "\n"))
	}
	return nil
}

func JoinErrors(errs []error) error {
	msgs := make([]string, 0, len(errs))
	for _, err := range errs {
		if err != nil {
			msgs = append(msgs, err.Error())
		}
	}
	return JoinAsError(msgs)
}
