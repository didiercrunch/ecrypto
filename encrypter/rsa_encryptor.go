package encrypter

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha512"
	"io"
)

var _ = sha512.New

type RsaSha512 struct {
	PrivateKey *rsa.PrivateKey
}

func (this *RsaSha512) hashIOReader(input io.Reader) ([]byte, error) {
	hashFunction := crypto.SHA512.New()
	if _, err := io.Copy(hashFunction, input); err != nil {
		return nil, err
	}
	return hashFunction.Sum(nil), nil
}

func (this *RsaSha512) SignStream(input io.Reader, rand io.Reader) ([]byte, error) {
	hashed, err := this.hashIOReader(input)
	if err != nil {
		return nil, err
	}
	return rsa.SignPSS(rand, this.PrivateKey, crypto.SHA512, hashed, nil)
}
