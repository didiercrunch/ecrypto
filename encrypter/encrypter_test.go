package encrypter

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

var _ = fmt.Print

type mockIoWriter struct {
	state []byte
}

func (this *mockIoWriter) Write(p []byte) (n int, err error) {
	this.state = append(this.state, p...)
	return len(p), nil
}

func (this *mockIoWriter) String() string {
	return string(this.state)
}

func newMockIoWriter() *mockIoWriter {
	return &mockIoWriter{make([]byte, 0, 1024)}
}

func (this *mockIoWriter) Length() int {
	return len(this.state)
}

func (this *mockIoWriter) Bytes() []byte {
	return this.state
}

func range_(length int) []byte {
	ret := make([]byte, length)
	for i, _ := range ret {
		ret[i] = byte(i)
	}
	return ret
}
func fromRange(from, to int) []byte {
	ret := make([]byte, to-from)
	for i, _ := range ret {
		ret[i] = byte(i + from)
	}
	return ret
}

func createMockRandomData(length int) []byte {
	ret := make([]byte, length)
	for i, _ := range ret {
		ret[i] = byte(i)
	}
	return ret
}

func NewMockEncrypter() *FileEncrypter {
	ret := new(FileEncrypter)
	mockRandomData := createMockRandomData(1024)
	ret.randomReader = bytes.NewReader(mockRandomData)
	return ret
}

func TestCreateSymetricKey(t *testing.T) {
	c := NewMockEncrypter()
	key, e := c.createSymetricKey(512)
	if e != nil {
		t.Error(e)
	} else if len(key) != 512 {
		t.Error("bad key lenght")
	} else if !reflect.DeepEqual(key, createMockRandomData(512)) {
		t.Fail()
	}
}

func TestGetCypherBlock(t *testing.T) {
	c := NewMockEncrypter()
	if c.key != nil {
		t.Error("expected nil key at this point")
	}
	block, err := c.getCypherBlock()
	if err != nil {
		t.Error(err)
	}
	if block.BlockSize() != 16 {
		t.Fail()
	}
	src := createMockRandomData(16)
	dst := make([]byte, 16)
	expected := []byte{90, 110, 4, 87, 8, 251, 113, 150, 240, 46, 85, 61, 2, 195, 166, 146}

	block.Encrypt(dst, src)
	if !reflect.DeepEqual(dst, expected) {
		t.Error("bad aes encryption")
	}
	if !reflect.DeepEqual(c.key, range_(32)) {
		t.Error("bad key", c.key)
	}
}

func TestNewEncrypter(t *testing.T) {
	c := NewEncrypter()
	if c.randomReader != rand.Reader {
		t.Fail()
	}
}

func TestCreateIV(t *testing.T) {
	m := NewMockEncrypter()
	b, err := aes.NewCipher([]byte("example key 1234"))
	if err != nil {
		t.Error(err)
		return
	}
	iv, err := m.createIV(b)
	if err != nil {
		t.Error(err)
		return
	}
	expt := range_(16)
	if !reflect.DeepEqual(iv, expt) {
		t.Error("bad iv", iv)
	}
}

func TestCreateStreamMode(t *testing.T) {
	m := NewMockEncrypter()
	if m.iv != nil {
		t.Error("expected nil iv at first")
	}
	b, err := aes.NewCipher([]byte("example key 1234"))
	if err != nil {
		t.Error(err)
		return
	}
	_, err = m.createStreamMode(b)
	if err != nil {
		t.Error(err)
		return
	}
	exptIv := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	if !reflect.DeepEqual(m.iv, exptIv) {
		t.Error("bad iv", m.iv)
	}
}

func TestCreateFileHandlerAndCypherBlock(t *testing.T) {
	m := NewMockEncrypter()
	if m.iv != nil {
		t.Error("expected nil iv at first")
	}

	_, err := m.createFileHandlerAndCypherBlock()
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(m.key, range_(32)) {
		t.Error("bad key", m.key)
	}
	if !reflect.DeepEqual(m.iv, fromRange(32, 32+16)) {
		t.Error("bad iv", m.iv)
	}
}

func TestSimpleEncrypeReaderToWriter(t *testing.T) {
	reader := strings.NewReader("some plain text")
	writer := newMockIoWriter()
	m := NewMockEncrypter()
	if err := m.encrypeReaderToWriter(reader, writer); err != nil {
		t.Error(err)
	}
}

func decryptEncryptedTextWithAESAlgorithmAndOFBMode(key, iv, encryptedText []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	stream := cipher.NewOFB(block, iv)
	ret := make([]byte, len(encryptedText))
	stream.XORKeyStream(ret, encryptedText)
	return string(ret), nil
}

func TestEncryptDecrypt(t *testing.T) {
	reader := strings.NewReader("some plain text")
	writer := newMockIoWriter()
	m := NewMockEncrypter()
	if err := m.encrypeReaderToWriter(reader, writer); err != nil {
		t.Error(err)
	}
	if plaintext2, err := decryptEncryptedTextWithAESAlgorithmAndOFBMode(m.key, m.iv, writer.Bytes()); err != nil {
		t.Error(err)
	} else if string(plaintext2) != "some plain text" {
		t.Error("bad encryption/decryption process")
	}
}
