package asymetric

import (
	_ "crypto/sha256"
	_ "crypto/sha512"

	"crypto"
	"crypto/rsa"
	"errors"
	"fmt"
	"io"
)

type RsaOaepPss struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
	Hash       crypto.Hash

	Random io.Reader
}

func (this *RsaOaepPss) getHash() (crypto.Hash, error) {
	if !this.Hash.Available() {
		return 0, errors.New(fmt.Sprintf("hash function %v is unabailable", this.Hash))
	} else {
		return this.Hash, nil
	}
}

func (this *RsaOaepPss) Encrypt(data []byte) ([]byte, error) {
	label := make([]byte, 0)
	if hash, err := this.getHash(); err != nil {
		return nil, err
	} else {
		return rsa.EncryptOAEP(hash.New(), this.Random, this.PublicKey, data, label)
	}

}

func (this *RsaOaepPss) Decrypt(ciphertext []byte) ([]byte, error) {
	label := make([]byte, 0)
	if hash, err := this.getHash(); err != nil {
		return nil, err
	} else {
		return rsa.DecryptOAEP(hash.New(), this.Random, this.PrivateKey, ciphertext, label)
	}
}

func (this *RsaOaepPss) Sign(hashedData []byte) ([]byte, error) {
	ret, err := rsa.SignPSS(this.Random, this.PrivateKey, this.Hash, hashedData, nil)
	if err != nil {
		fmt.Println(">>> ", len(hashedData))
	}
	return ret, err
}

func (this *RsaOaepPss) VerifySignature(hashedData []byte, signature []byte) error {
	return rsa.VerifyPSS(this.PublicKey, this.Hash, hashedData, signature, nil)
}
