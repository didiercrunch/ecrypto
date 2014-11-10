package asymetric

import (
	_ "crypto/sha512"

	"crypto"
	"crypto/rsa"
	"testing"
)

func TestRsaOaepEncrypt(t *testing.T) {
	bobPrivKey, bobPubKey := generateKeyPair()

	r := &RsaOaepEncrypter{nil, bobPubKey, crypto.SHA384, random}

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

func TestRsaOaepEncryptDecrypt(t *testing.T) {
	bobPrivKey, bobPubKey := generateKeyPair()

	message := "this is a secret message"

	rAlice := &RsaOaepEncrypter{nil, bobPubKey, crypto.SHA384, random}
	rBob := &RsaOaepEncrypter{bobPrivKey, nil, crypto.SHA384, random}

	if ciphertext, err := rAlice.Encrypt([]byte(message)); err != nil {
		t.Error(err)
		return
	} else if decryptedText, err := rBob.Decrypt(ciphertext); err != nil {
		t.Error(err)
	} else if string(decryptedText) != message {
		t.Error("the decryption of the message does not fit the original message")
	}
}

func TestRsaOaepAvailableHashes(t *testing.T) {
	hashes := map[crypto.Hash]bool{crypto.SHA224: true, crypto.SHA256: true, crypto.SHA384: true, crypto.SHA512: true}
	var i crypto.Hash
	for i = 0; i < 22; i++ {
		r := &RsaOaepEncrypter{Hash: i}
		if _, err := r.getHash(); err != nil && hashes[i] {
			t.Error("should support", i, "hash type but does not")
		} else if err == nil && !hashes[i] {
			t.Error("should not support", i, "hash type but does")
		}
	}
}
