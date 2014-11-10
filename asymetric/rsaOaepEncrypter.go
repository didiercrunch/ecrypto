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

type RsaOaepEncrypter struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
	Hash       crypto.Hash

	Random io.Reader
}

func (this *RsaOaepEncrypter) getHash() (crypto.Hash, error) {
	if !this.Hash.Available() {
		return 0, errors.New(fmt.Sprintf("hash function %v is unabailable", this.Hash))
	} else {
		return this.Hash, nil
	}
}

func (this *RsaOaepEncrypter) Encrypt(data []byte) ([]byte, error) {
	label := make([]byte, 0)
	if hash, err := this.getHash(); err != nil {
		return nil, err
	} else {
		return rsa.EncryptOAEP(hash.New(), this.Random, this.PublicKey, data, label)
	}

}

func (this *RsaOaepEncrypter) Decrypt(ciphertext []byte) ([]byte, error) {
	label := make([]byte, 0)
	if hash, err := this.getHash(); err != nil {
		return nil, err
	} else {
		return rsa.DecryptOAEP(hash.New(), this.Random, this.PrivateKey, ciphertext, label)
	}
}
