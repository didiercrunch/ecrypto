package keys

import (
	"crypto/rsa"
	"encoding/json"
	"math/big"
	"reflect"
	"testing"
)

func b(x int64) *big.Int {
	return big.NewInt(x)
}

func TestGetDefaultRSAPublicKey(t *testing.T) {
	key := &rsa.PublicKey{b(31 * 11), 3}
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
		rsa.PublicKey{b(3 * 11), 7},
		b(3),
		[]*big.Int{b(3), b(11)},
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
	key := &RSAPrivateKey{b(1023), b(2112), b(328723), b(2332), 89}
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
	key := &RSAPublicKey{b(1023), 89}
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
