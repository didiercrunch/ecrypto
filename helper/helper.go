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
