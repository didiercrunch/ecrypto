package main

import (
	"crypto/rand"
	"crypto/rsa"
	"errors"
)

type KeyGenerator struct {
	publicKey  *PublicKey
	privateKey *PrivateKey
}

func (this *KeyGenerator) createRSAKey(size int) error {
	rsaPrivateKey, err := rsa.GenerateKey(rand.Reader, size)
	if err != nil {
		return errors.New("cannot generate the rsa public/private key pair\n" + err.Error())
	}

	if this.privateKey, err = GetDefaultRSAPrivateKey(rsaPrivateKey); err != nil {
		return errors.New("cannot generate the rsa public/private key pair\n" + err.Error())
	}
	this.publicKey = GetDefaultRSAPublicKey(&rsaPrivateKey.PublicKey)
	return nil
}

func (this *KeyGenerator) CreateNewKey(size int, password string) error {
	if err := this.createRSAKey(size); err != nil {
		return err
	}
	return nil
}
