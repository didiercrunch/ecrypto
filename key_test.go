package main

import (
	"testing"
)

func TestCreateRSAKey(t *testing.T) {
	kc := new(KeyGenerator)
	if err := kc.createRSAKey(100); err != nil {
		t.Error(err)
	}
	if kc.privateKey == nil {
		t.Fail()
	}
	if kc.publicKey == nil {
		t.Fail()
	}
}
