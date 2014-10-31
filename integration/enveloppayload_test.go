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

type mockEncryptor struct{}

func (this *mockEncryptor) Encrypt(data []byte) ([]byte, error) {
	// cesar encryption!
	ret := make([]byte, len(data))
	for i, d := range data {
		ret[i] = d + 3 // modulo is implicit
	}
	return ret, nil
}

type mockSigner struct{}

var random = helper.NewMockRandomReader()

func generateRSAKeyPair() (*rsa.PrivateKey, *rsa.PublicKey) {
	if key, err := rsa.GenerateKey(random, 2048); err != nil {
		panic(err)
	} else {
		return key, &key.PublicKey
	}
}

func getEncryterAndSigner() *asymetric.RsaOaepPss {
	alicePrivateKey, _ := generateRSAKeyPair()
	_, bobPubKey := generateRSAKeyPair()
	return &asymetric.RsaOaepPss{alicePrivateKey, bobPubKey, crypto.SHA512, random}
}

func (this *mockSigner) Sign(data []byte) ([]byte, error) {
	// a slice with the sum of the data plus a "secret" value
	var sum byte = 36
	for _, d := range data {
		sum += d
	}
	return []byte{sum}, nil
}

func TestPayloadCanBeUseInEnvelop(t *testing.T) {
	encryterAndSigner := getEncryterAndSigner()
	payload_ := payload.GetDefaultPayload(bytes.NewBufferString("some data to encrypt"))
	envelop := envelop.NewEnveloper(encryterAndSigner, payload_, encryterAndSigner)
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
