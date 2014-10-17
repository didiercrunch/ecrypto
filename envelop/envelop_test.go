package envelop

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
	"github.com/didiercrunch/filou/helper"
	"io"
	"reflect"
	"testing"
)

var _ = fmt.Print

type mockPayload struct {
	MethodsCall []string
}

func newMockPayload() *mockPayload {
	return &mockPayload{make([]string, 0)}
}

func (this *mockPayload) addMethodCall(methodName string) {
	this.MethodsCall = append(this.MethodsCall, methodName)
}

//  assert firstToBeCalled is called before secondToBeCalled.  Calls can be made
//  between them and defore firstToBeCalled and after secondToBeCalled.
func (this *mockPayload) assertMethodCallOrder(firstToBeCalled, secondToBeCalled string) error {
	var firstToBeCalledIdx, secondToBeCalledIdx = -1, -1
	for i, name := range this.MethodsCall {
		if firstToBeCalledIdx == -1 && name == firstToBeCalled {
			firstToBeCalledIdx = i
		}
		if secondToBeCalledIdx == -1 && name == secondToBeCalled {
			secondToBeCalledIdx = i
		}
	}
	if firstToBeCalledIdx == -1 {
		return errors.New(firstToBeCalled + " is never called")
	}
	if secondToBeCalledIdx == -1 {
		return errors.New(secondToBeCalled + " is never called")
	}
	if firstToBeCalledIdx >= secondToBeCalledIdx {
		return errors.New(firstToBeCalled + " is called after " + secondToBeCalled)
	}
	return nil
}

func (this *mockPayload) GetKey() []byte {
	this.addMethodCall("GetKey")
	return []byte{1, 2, 3}
}

func (this *mockPayload) GetHash() []byte {
	this.addMethodCall("GetHash")
	return []byte{10, 9, 8}
}

func (this *mockPayload) GetAlgorithm() string {
	this.addMethodCall("GetAlgorithm")
	symetricEncryptionAlgorithm := "DES"
	return symetricEncryptionAlgorithm
}

func (this *mockPayload) GetMode() string {
	this.addMethodCall("GetMode")
	symetricEncryptionMode := "OFB"
	return symetricEncryptionMode
}

func (this *mockPayload) GetHashMethod() string {
	this.addMethodCall("GetHashMethod")
	theHashMethodUsedInThePayLoad := "SHA1"
	return theHashMethodUsedInThePayLoad
}

func (this *mockPayload) GetPayloadData() io.Reader {
	this.addMethodCall("GetPayloadData")
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
		payload:   newMockPayload(),
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
		payload:   newMockPayload(),
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
		payload:   newMockPayload(),
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
	env := &Enveloper{payload: newMockPayload()}
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
		payload:   newMockPayload(),
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
		payload:   newMockPayload(),
		signer:    &mockSigner{},
		encryptor: &mockEncryptor{},
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

func TestMethodOrder(t *testing.T) {
	mockPayload := newMockPayload()
	env := &Enveloper{
		payload:   mockPayload,
		signer:    &mockSigner{},
		encryptor: &mockEncryptor{},
	}
	w := helper.NewMockIoWriter()
	if err := env.WriteToWriter(w); err != nil {
		t.Error(err)
	}
	if err := mockPayload.assertMethodCallOrder("GetPayloadData", "GetHashMethod"); err != nil {
		t.Error(err)
	}
}
