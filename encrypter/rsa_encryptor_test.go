package encrypter

import (
	"bytes"
	"crypto"
	"crypto/rsa"
	"crypto/sha512"
	"fmt"
	"github.com/didiercrunch/ecrypto/helper"
	"testing"
)

func TestHashIOReader(t *testing.T) {
	rs := new(RsaSha512)
	expt := "fabbbe22180f1f137cfdc9556d2570e775d1ae02a597ded43a72a40f9b485d500043b7b" +
		"e128fb9fcd982b83159a0d99aa855a9e7cc4240c00dc01a9bdf8218d7"
	input := bytes.NewBufferString("C is as portable as Stonehedge!!")
	if hashValue, err := rs.hashIOReader(input); err != nil {
		t.Error(err)
	} else if fmt.Sprintf("%x", hashValue) != expt {
		t.Errorf("Hash is not the same \n expected : %s\n"+
			"results  : %x  ", expt, hashValue)
	}
}

func TestSignStream(t *testing.T) {
	toSlice := func(data [sha512.Size]byte) []byte {
		ret := make([]byte, sha512.Size)
		for i, datum := range data {
			ret[i] = datum
		}
		return ret
	}
	privateKey := helper.GetRSAPrivateKey(1024)
	rs := &RsaSha512{privateKey}
	input := bytes.NewBufferString("something to sign")
	signature, err := rs.SignStream(input, helper.NewMockRandomReader())
	if err != nil {
		t.Error(err)
		return
	}
	hashed := toSlice(sha512.Sum512([]byte("something to sign")))
	if err := rsa.VerifyPSS(&privateKey.PublicKey, crypto.SHA512, hashed, signature, nil); err != nil {
		t.Error(err)
	}
}
