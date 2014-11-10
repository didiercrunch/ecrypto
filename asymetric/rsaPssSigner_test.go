package asymetric

import (
	_ "crypto/sha512"

	"crypto"
	"crypto/rsa"
	"testing"
)

//var random = helper.NewMockRandomReader()

//func generateKeyPair() (*rsa.PrivateKey, *rsa.PublicKey) {
//	if key, err := rsa.GenerateKey(random, 1024); err != nil {
//		panic(err)
//	} else {
//		return key, &key.PublicKey
//	}
//}

func TestRsaPssSign(t *testing.T) {
	alicePrivKey, alicePubKey := generateKeyPair()
	h := crypto.SHA384.New()
	h.Write([]byte("this is a secret message"))
	hashedMessage := h.Sum(nil)

	rAlice := &RsaPssSigner{alicePrivKey, nil, crypto.SHA384, random}
	signature, err := rAlice.Sign(hashedMessage)
	if err != nil {
		t.Error(err)
		return
	}
	if err := rsa.VerifyPSS(alicePubKey, crypto.SHA384, hashedMessage, signature, nil); err != nil {
		t.Error(err)
	}
}

func TestRsaPssSignAndVerify(t *testing.T) {
	alicePrivKey, alicePubKey := generateKeyPair()
	h := crypto.SHA384.New()
	h.Write([]byte("this is a secret message"))
	hashedMessage := h.Sum(nil)

	rAlice := &RsaPssSigner{alicePrivKey, nil, crypto.SHA384, random}
	rBob := &RsaPssSigner{nil, alicePubKey, crypto.SHA384, random}
	if signature, err := rAlice.Sign(hashedMessage); err != nil {
		t.Error(err)
	} else if err := rBob.VerifySignature(hashedMessage, signature); err != nil {
		t.Error(err)
	}
}

func TestRsaPssAvailableHashes(t *testing.T) {
	hashes := map[crypto.Hash]bool{crypto.SHA224: true, crypto.SHA256: true, crypto.SHA384: true, crypto.SHA512: true}
	var i crypto.Hash
	for i = 0; i < 22; i++ {
		r := &RsaPssSigner{Hash: i}
		if _, err := r.getHash(); err != nil && hashes[i] {
			t.Error("should support", i, "hash type but does not")
		} else if err == nil && !hashes[i] {
			t.Error("should not support", i, "hash type but does")
		}
	}
}
