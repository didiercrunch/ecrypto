package encrypter

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	. "github.com/didiercrunch/filou/helper"
	"reflect"
	"strings"
	"testing"
)

func NewMockAESOFBEncrypter() *AESOFBEncrypter {
	ret := new(AESOFBEncrypter)
	ret.randomReader = NewMockRandomReader()
	return ret
}

func TestCreateSymetricKey(t *testing.T) {
	c := NewMockAESOFBEncrypter()
	key, e := c.createSymetricKey(512)
	if e != nil {
		t.Error(e)
	} else if len(key) != 512 {
		t.Error("bad key lenght")
	} else if !reflect.DeepEqual(key, NewMockRandomReader().GetRandomBytes(512)) {
		t.Fail()
	}
}

func TestGetCypherBlock(t *testing.T) {
	c := NewMockAESOFBEncrypter()
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
	src := Range(16)
	dst := make([]byte, 16)
	expected := []byte{30, 248, 157, 9, 161, 60, 109, 82, 221, 152, 171, 167, 137,
		239, 177, 191}
	expectedKey := []byte{73, 233, 188, 33, 31, 118, 98, 81, 66, 120, 254, 85, 138, 36,
		118, 80, 38, 83, 169, 57, 27, 61, 103, 122, 37, 164, 185, 40, 112, 66, 1, 50}

	block.Encrypt(dst, src)
	if !reflect.DeepEqual(dst, expected) {
		t.Error("bad aes encryption", dst)
	}
	if !reflect.DeepEqual(c.key, expectedKey) {
		t.Error("bad key", c.key)
	}
}

func TestNewEncrypter(t *testing.T) {
	c := NewAESOFBEncrypter()
	if c.randomReader != rand.Reader {
		t.Fail()
	}
}

func TestCreateIV(t *testing.T) {
	m := NewMockAESOFBEncrypter()
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
	expt := NewMockRandomReader().GetRandomBytes(16)
	if !reflect.DeepEqual(iv, expt) {
		t.Error("bad iv", iv)
	}
}

func TestCreateStreamMode(t *testing.T) {
	m := NewMockAESOFBEncrypter()
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
	exptIv := NewMockRandomReader().GetRandomBytes(16)
	if !reflect.DeepEqual(m.iv, exptIv) {
		t.Error("bad iv", m.iv, " is not ", exptIv)
	}
}

func TestCreateFileHandlerAndCypherBlock(t *testing.T) {
	m := NewMockAESOFBEncrypter()
	if m.iv != nil {
		t.Error("expected nil iv at first")
	}

	_, err := m.createFileHandlerAndCypherBlock()
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(m.key, NewMockRandomReader().GetRandomBytes(32)) {
		t.Error("bad key", m.key)
	}
	if !reflect.DeepEqual(m.iv, NewMockRandomReader().GetRandomBytes(32 + 16)[32:32+16]) {
		t.Error("bad iv", m.iv)
	}
}

func TestSimpleEncryptReaderToWriter(t *testing.T) {
	reader := strings.NewReader("some plain text")
	writer := NewMockIoWriter()
	m := NewMockAESOFBEncrypter()
	if err := m.EncryptReaderToWriter(reader, writer); err != nil {
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
	writer := NewMockIoWriter()
	m := NewMockAESOFBEncrypter()
	if err := m.EncryptReaderToWriter(reader, writer); err != nil {
		t.Error(err)
	}
	if plaintext2, err := decryptEncryptedTextWithAESAlgorithmAndOFBMode(m.key, m.iv, writer.Bytes()); err != nil {
		t.Error(err)
	} else if string(plaintext2) != "some plain text" {
		t.Error("bad encryption/decryption process")
	}
}
