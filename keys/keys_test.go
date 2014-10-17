package keys

import (
	"crypto/rsa"
	"encoding/json"
	. "github.com/didiercrunch/filou/helper"
	"math/big"
	"reflect"
	"testing"
)

func TestGetDefaultRSAPublicKey(t *testing.T) {
	key := &rsa.PublicKey{B(31 * 11), 3}
	pk := GetDefaultRSAPublicKey(key)

	if big.NewInt(31*11).Cmp(pk.Key.N) != 0 {
		t.Error("bad N")
	}
	if 3 != pk.Key.E {
		t.Error("bad e")
	}
}

func TestGetDefaultRSAPrivateKey(t *testing.T) {
	privKey := &rsa.PrivateKey{
		rsa.PublicKey{B(3 * 11), 7},
		B(3),
		[]*big.Int{B(3), B(11)},
		rsa.PrecomputedValues{},
	}
	pk, err := GetDefaultRSAPrivateKey(privKey)
	if err != nil {
		t.Error(err)
	}
	if pk.Key.D.Int64() != 3 {
		t.Error("bad D")
	}
	if pk.Key.E != 7 {
		t.Error("bad D")
	}
	if pk.Key.P.Int64() != 3 {
		t.Error("bad D")
	}
	if pk.Key.Q.Int64() != 11 {
		t.Error("bad D")
	}
}

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
