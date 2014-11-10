package asymetric

import (
	"crypto/rsa"
	"testing"

	"github.com/didiercrunch/filou/contract"
)

func TestGetSignerEmptyContract(t *testing.T) {
	c := new(contract.AcceptedContract)
	if _, err := GetSigner(c); err == nil {
		t.Error("exceptec an error here")
	}
}

func TestGetSignerBadName(t *testing.T) {
	c := new(contract.AcceptedContract)
	c.SignatureScheme = "bigtits.com"
	if _, err := GetSigner(c); err == nil {
		t.Error("exceptec an error here")
	}
}

func TestGetSigner(t *testing.T) {
	c := new(contract.AcceptedContract)
	priv, pub := generateKeyPair()

	c.Hash = "sha256"
	c.SignatureScheme = "rsa_pss"
	c.RsaPublicKey = pub

	c.RsaPrivateKey = func() (*rsa.PrivateKey, error) {
		return priv, nil
	}
	if _, err := GetSigner(c); err != nil {
		t.Error(err)
	}
}

func TestGetPublicKeyEncryptor(t *testing.T) {
	c := new(contract.AcceptedContract)
	priv, pub := generateKeyPair()

	c.Hash = "sha256"
	c.AsynchronousEncryptionScheme = "rsa_oaep"
	c.RsaPublicKey = pub

	c.RsaPrivateKey = func() (*rsa.PrivateKey, error) {
		return priv, nil
	}
	if _, err := GetPublicKeyEncryptor(c); err != nil {
		t.Error(err)
	}
}

func TestGetPublicKeyEncryptorBadName(t *testing.T) {
	c := new(contract.AcceptedContract)
	c.AsynchronousEncryptionScheme = "bigtits.com"
	if _, err := GetPublicKeyEncryptor(c); err == nil {
		t.Error("exceptec an error here")
	}
}
