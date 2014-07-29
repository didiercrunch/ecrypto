package main

import (
	"crypto/rsa"
	"math/big"
	"testing"
)

func b(x int64) *big.Int {
	return big.NewInt(x)
}

func TestGetDefaultRSAPublicKey(t *testing.T) {
	key := &rsa.PublicKey{b(31 * 11), 3}
	pk := GetDefaultRSAPublicKey(key)
	niceKey, ok := pk.Key.(*RSAPublicKey)
	if !ok {
		t.Error("bad key type")
		return
	}
	if big.NewInt(31*11).Cmp(niceKey.N) != 0 {
		t.Error("bad N")
	}
	if 3 != niceKey.E {
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
	niceKey, ok := pk.Key.(*RSAPrivateKey)
	if !ok {
		t.Error("bad key type")
		return
	}
	if niceKey.D.Int64() != 3 {
		t.Error("bad D")
	}
	if niceKey.E != 7 {
		t.Error("bad D")
	}
	if niceKey.P.Int64() != 3 {
		t.Error("bad D")
	}
	if niceKey.Q.Int64() != 11 {
		t.Error("bad D")
	}
}
