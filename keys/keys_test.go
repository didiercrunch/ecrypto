package keys

import (
	"crypto/rsa"
	"math/big"
	"testing"

	. "github.com/didiercrunch/filou/helper"
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
