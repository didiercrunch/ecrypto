package helper

import (
	"errors"
	"testing"
)

func TestJoinAsError(t *testing.T) {
	var errs []string = nil
	if JoinAsError(errs) != nil {
		t.Error("expected no errors")
	}
	errs = make([]string, 0)
	if JoinAsError(errs) != nil {
		t.Error("expected no errors")
	}

	errs = []string{"a", "b"}
	if JoinAsError(errs).Error() != "a\nb" {
		t.Error("bad error")
	}
}

func TestJoinErrors(t *testing.T) {
	var errs []error = nil
	if JoinErrors(errs) != nil {
		t.Error("expected no errors")
	}
	errs = make([]error, 0)
	if JoinErrors(errs) != nil {
		t.Error("expected no errors")
	}

	errs = []error{errors.New("a"), errors.New("b")}
	if JoinErrors(errs).Error() != "a\nb" {
		t.Error("bad error")
	}
}
