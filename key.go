package main

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"errors"
	"github.com/didiercrunch/ecrypto/shared"
	"os"
	"path"
)

type KeyGenerator struct {
	publicKey  *PublicKey
	privateKey *PrivateKey
}

func (this *KeyGenerator) ensureDirectoryExists(dir string) error {
	stat, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return os.Mkdir(dir, 0700)
	} else if err != nil {
		return err
	} else if !stat.IsDir() {
		return errors.New(dir + " is not a directory")
	}
	return nil

}

func (this *KeyGenerator) saveKeyAsJSON(key interface{}, filepath string) error {
	w, err := os.Create(filepath)
	if err != nil {
		return err
	}
	enc := json.NewEncoder(w)
	return enc.Encode(key)
}

func (this *KeyGenerator) saveKeys(dirPath string) error {
	if err := this.ensureDirectoryExists(dirPath); err != nil {
		return err
	}
	if err := this.saveKeyAsJSON(this.publicKey, path.Join(dirPath, "publickey.json")); err != nil {
		return err
	}
	return this.saveKeyAsJSON(this.privateKey, path.Join(dirPath, "privatekey.json"))
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

func (this *KeyGenerator) CreateNewKey(size int) error {
	if err := this.createRSAKey(size); err != nil {
		return err
	}
	return this.saveKeys(shared.GetEcryptoDir())
}
