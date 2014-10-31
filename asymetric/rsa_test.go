package asymetric

import (
	_ "crypto/sha512"

	"crypto"
	"crypto/rsa"
	"testing"

	"github.com/didiercrunch/filou/helper"
)

var random = helper.NewMockRandomReader()

func generateKeyPair() (*rsa.PrivateKey, *rsa.PublicKey) {
	if key, err := rsa.GenerateKey(random, 1024); err != nil {
		panic(err)
	} else {
		return key, &key.PublicKey
	}
}

func TestEncrypt(t *testing.T) {
	bobPrivKey, bobPubKey := generateKeyPair()

	r := &RsaOaepPss{nil, bobPubKey, crypto.SHA384, random}

	ciphertext, err := r.Encrypt([]byte("this is a secret message"))
	if err != nil {
		t.Error(err)
		return
	}
	orgTest, err := rsa.DecryptOAEP(crypto.SHA384.New(), random, bobPrivKey, ciphertext, []byte{})
	if err != nil {
		t.Error(err)
		return
	}
	if string(orgTest) != "this is a secret message" {
		t.Error("the decryption of the message does not fit the original message")
	}
}

func TestEncryptDecrypt(t *testing.T) {
	bobPrivKey, bobPubKey := generateKeyPair()

	message := "this is a secret message"

	rAlice := &RsaOaepPss{nil, bobPubKey, crypto.SHA384, random}
	rBob := &RsaOaepPss{bobPrivKey, nil, crypto.SHA384, random}

	if ciphertext, err := rAlice.Encrypt([]byte(message)); err != nil {
		t.Error(err)
		return
	} else if decryptedText, err := rBob.Decrypt(ciphertext); err != nil {
		t.Error(err)
	} else if string(decryptedText) != message {
		t.Error("the decryption of the message does not fit the original message")
	}
}

func TestSign(t *testing.T) {
	alicePrivKey, alicePubKey := generateKeyPair()
	h := crypto.SHA384.New()
	h.Write([]byte("this is a secret message"))
	hashedMessage := h.Sum(nil)

	rAlice := &RsaOaepPss{alicePrivKey, nil, crypto.SHA384, random}
	signature, err := rAlice.Sign(hashedMessage)
	if err != nil {
		t.Error(err)
		return
	}
	if err := rsa.VerifyPSS(alicePubKey, crypto.SHA384, hashedMessage, signature, nil); err != nil {
		t.Error(err)
	}
}

func TestSignAndVerify(t *testing.T) {
	alicePrivKey, alicePubKey := generateKeyPair()
	h := crypto.SHA384.New()
	h.Write([]byte("this is a secret message"))
	hashedMessage := h.Sum(nil)

	rAlice := &RsaOaepPss{alicePrivKey, nil, crypto.SHA384, random}
	rBob := &RsaOaepPss{nil, alicePubKey, crypto.SHA384, random}
	if signature, err := rAlice.Sign(hashedMessage); err != nil {
		t.Error(err)
	} else if err := rBob.VerifySignature(hashedMessage, signature); err != nil {
		t.Error(err)
	}
}

func TestAvailableHashes(t *testing.T) {
	hashes := map[crypto.Hash]bool{crypto.SHA224: true, crypto.SHA256: true, crypto.SHA384: true, crypto.SHA512: true}
	var i crypto.Hash
	for i = 0; i < 22; i++ {
		r := &RsaOaepPss{Hash: i}
		if _, err := r.getHash(); err != nil && hashes[i] {
			t.Error("should support", i, "hash type but does not")
		} else if err == nil && !hashes[i] {
			t.Error("should not support", i, "hash type but does")
		}
	}
}
