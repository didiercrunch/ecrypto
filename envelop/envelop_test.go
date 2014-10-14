package envelop

import (
	"archive/zip"
	"bytes"
	"github.com/didiercrunch/ecrypto/helper"
	"io"
	"reflect"
	"testing"
)

type mockPayload struct{}

func (this *mockPayload) GetKey() []byte {
	return []byte{1, 2, 3}
}

func (this *mockPayload) GetHash() []byte {
	return []byte{10, 9, 8}
}

func (this *mockPayload) GetAlgorithm() string {
	symetricEncryptionAlgorithm := "DES"
	return symetricEncryptionAlgorithm
}

func (this *mockPayload) GetMode() string {
	symetricEncryptionMode := "OFB"
	return symetricEncryptionMode
}

func (this *mockPayload) GetHashMethod() string {
	theHashMethodUsedInThePayLoad := "SHA1"
	return theHashMethodUsedInThePayLoad
}

func (this *mockPayload) GetPayloadData() io.Reader {
	someData := `some data must be here.  In reality, the data
	             would be the encrypted payload.  This will be
				the big chunk of the data.
				`
	return bytes.NewReader([]byte(someData))
}

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

func (this *mockSigner) Sign(data []byte) ([]byte, error) {
	// a slice with the sum of the data plus a "secret" value
	var sum byte = 36
	for _, d := range data {
		sum += d
	}
	return []byte{sum}, nil
}

func TestEncryptKey(t *testing.T) {
	env := &Enveloper{
		encryptor: &mockEncryptor{},
		payload:   &mockPayload{},
	}
	if enc, err := env.EncryptKey(); err != nil {
		t.Error(err)
	} else if !reflect.DeepEqual(enc, []byte{1 + 3, 2 + 3, 3 + 3}) {
		t.Error("bad signature")
	}
}

func TestCreateEnvelopCompletelyEncryptKey(t *testing.T) {
	env := &Enveloper{
		encryptor: &mockEncryptor{},
		payload:   &mockPayload{},
		signer:    &mockSigner{},
	}
	if err := env.CreateEnvelopCompletely(); err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(env.EncryptedKey, []byte{1 + 3, 2 + 3, 3 + 3}) {
		t.Error("CreateEnvelopCompletely has not encrypted the key")
	}
}

func TestSignPayload(t *testing.T) {
	env := &Enveloper{
		signer:  &mockSigner{},
		payload: &mockPayload{},
	}
	if sign, err := env.SignPayload(); err != nil {
		t.Error(err)
	} else if !reflect.DeepEqual([]byte{10 + 9 + 8 + 36}, sign) {
		t.Error("bad signature")
	}
}

func TestCreateEnvelopCompletelySignHash(t *testing.T) {
	env := &Enveloper{
		payload:   &mockPayload{},
		encryptor: &mockEncryptor{},
		signer:    &mockSigner{},
	}
	if err := env.CreateEnvelopCompletely(); err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(env.Signature, []byte{10 + 9 + 8 + 36}) {
		t.Error("CreateEnvelopCompletely has not signature")
	}
}

func TestCreateMetadata(t *testing.T) {
	env := &Enveloper{payload: &mockPayload{}}
	exptectedMetaData := &Metadata{
		BlockAlgorithm:     "DES",
		BlockMode:          "OFB",
		SignatureAlgorithm: "RSAPSS",
		HashMethod:         "SHA1",
	}

	if m := env.CreateMetadata(); !reflect.DeepEqual(m, exptectedMetaData) {
		t.Error("bad meta data")
	}
}

func TestCreateEnvelopCompletelyMetadata(t *testing.T) {
	env := &Enveloper{
		encryptor: &mockEncryptor{},
		payload:   &mockPayload{},
		signer:    &mockSigner{},
	}
	if err := env.CreateEnvelopCompletely(); err != nil {
		t.Error(err)
	}
	if env.Metadata == nil || env.Metadata.HashMethod != "SHA1" {
		t.Error("CreateEnvelopCompletely has not created metadata")
	}
}

func TestWriteToWriter(t *testing.T) {
	env := &Enveloper{
		payload: &mockPayload{},
	}
	w := helper.NewMockIoWriter()
	if err := env.WriteToWriter(w); err != nil {
		t.Error(err)
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
