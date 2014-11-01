package helper

import (
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
