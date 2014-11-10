package helper

import (
	"crypto"
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

func TestGetHashFunctionByLowerCaseName(t *testing.T) {
	m := map[string]crypto.Hash{
		"md4":       crypto.MD4,
		"md5":       crypto.MD5,
		"sha1":      crypto.SHA1,
		"sha224":    crypto.SHA224,
		"sha256":    crypto.SHA256,
		"sha384":    crypto.SHA384,
		"sha512":    crypto.SHA512,
		"md5sha1":   crypto.MD5SHA1,
		"ripemd160": crypto.RIPEMD160,
	}
	for hashName, exptedHash := range m {
		if hash, err := GetHashFunctionByLowerCaseName(hashName); err != nil {
			t.Error(err)
		} else if hash != exptedHash {
			t.Error("bad hash found for ", hashName)
		}
	}
}
