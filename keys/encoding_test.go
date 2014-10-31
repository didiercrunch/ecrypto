package keys

import (
	"encoding/json"
	"reflect"
	"testing"

	. "github.com/didiercrunch/filou/helper"
)

func TestRSAPrivateKeyJsonification(t *testing.T) {
	key := &RSAPrivateKey{B(1023), B(2112), B(328723), B(2332), 89}
	data, err := json.Marshal(key)
	if err != nil {
		t.Error(err)
		return
	}
	key2 := new(RSAPrivateKey)
	if err := json.Unmarshal(data, key2); err != nil {
		t.Error(err)
		return
	}
	if !reflect.DeepEqual(key, key2) {
		t.Fail()
	}
}

func TestRSAPublicKeyJsonification(t *testing.T) {
	key := &RSAPublicKey{B(1023), 89}
	data, err := json.Marshal(key)
	if err != nil {
		t.Error(err)
		return
	}
	key2 := new(RSAPublicKey)
	if err := json.Unmarshal(data, key2); err != nil {
		t.Error(err)
		return
	}
	if !reflect.DeepEqual(key, key2) {
		t.Fail()
	}
}

func TestRSAPublicKeyJsonificationFromString(t *testing.T) {
	data := []byte(`{"N":"3ff","E":89}`)
	key := new(RSAPublicKey)
	if err := json.Unmarshal(data, key); err != nil {
		t.Error(err)
		return
	}
	if !reflect.DeepEqual(key, &RSAPublicKey{B(1023), 89}) {
		t.Fail()
	}
}
