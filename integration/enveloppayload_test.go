package integration

import (
	"archive/zip"
	"bytes"
	"crypto"
	"crypto/rsa"
	"testing"

	"github.com/didiercrunch/filou/asymetric"
	"github.com/didiercrunch/filou/envelop"
	"github.com/didiercrunch/filou/helper"
	"github.com/didiercrunch/filou/payload"
)

var random = helper.NewMockRandomReader()

func generateRSAKeyPair() (*rsa.PrivateKey, *rsa.PublicKey) {
	if key, err := rsa.GenerateKey(random, 2048); err != nil {
		panic(err)
	} else {
		return key, &key.PublicKey
	}
}

func getEncryter() *asymetric.RsaOaepEncrypter {
	alicePrivateKey, _ := generateRSAKeyPair()
	_, bobPubKey := generateRSAKeyPair()
	return &asymetric.RsaOaepEncrypter{alicePrivateKey, bobPubKey, crypto.SHA512, random}
}

func getSigner() *asymetric.RsaPssSigner {
	alicePrivateKey, _ := generateRSAKeyPair()
	_, bobPubKey := generateRSAKeyPair()
	return &asymetric.RsaPssSigner{alicePrivateKey, bobPubKey, crypto.SHA512, random}
}

func TestPayloadCanBeUseInEnvelop(t *testing.T) {
	encrypter := getEncryter()
	signer := getSigner()
	payload_ := payload.GetDefaultPayload(bytes.NewBufferString("some data to encrypt"))
	envelop := envelop.NewEnveloper(encrypter, payload_, signer)
	w := helper.NewMockIoWriter()
	if err := envelop.WriteToWriter(w); err != nil {
		t.Error(err)
		return
	}
	r, err := zip.NewReader(w.ReaderAt(), int64(w.Length()))
	if err != nil {
		t.Error(err)
		return
	}
	if len(r.File) != 3 {
		t.Error("wrong file number")
	}
}
