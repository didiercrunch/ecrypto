package payload

import (
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/didiercrunch/filou/helper"
)

type Blocker func([]byte) (cipher.Block, error)

func assertBlockOkay(block cipher.Block, key []byte, blocker Blocker) error {
	expectedBlock, err := blocker(key)
	if err != nil {
		return err
	}
	if expectedBlock.BlockSize() != block.BlockSize() {
		return errors.New(fmt.Sprintf("bad block size.  Expected %v but was %v", expectedBlock.BlockSize(), block.BlockSize()))
	}
	src := helper.Range(block.BlockSize())
	encrypt := make([]byte, block.BlockSize())
	decrypt := make([]byte, block.BlockSize())
	block.Encrypt(encrypt, src)
	expectedBlock.Decrypt(decrypt, encrypt)
	if !reflect.DeepEqual(decrypt, src) {
		return errors.New("expected encryption is wrong")
	}
	return nil
}

func TestCreateBlock(t *testing.T) {
	random := helper.NewMockRandomReader()
	if block, key, err := AES256.CreateBlock(random); err != nil {
		t.Error(err)
	} else if err := assertBlockOkay(block, key, aes.NewCipher); err != nil {
		t.Error(err)
	}
}

func Testhash512BytesOfData(t *testing.T) {
	if _, err := hash512BytesOfData(crypto.MD5, helper.NewMockRandomReader()); err == nil {
		t.Error("should never allow MD5 hash")
	}
	randomGenerator := helper.NewMockRandomReader()
	if rd, err := hash512BytesOfData(crypto.SHA512, randomGenerator); err != nil {
		t.Error(err)
	} else if len(rd) != 512/8 {
		t.Fail()
	} else if l := randomGenerator.GetReadLength(); l != 512 {
		t.Error(fmt.Sprint("Was suppose to read 512 bytes of data but read", l, "bytes of data"))
	}
}

func TestCreateRandomKey(t *testing.T) {
	if key, err := createRandomKey(128, helper.NewMockRandomReader()); err != nil {
		t.Error(err)
	} else if len(key) != 128/8 {
		t.Error("bad key length.  Expected 128 found ", len(key)*8)
	}

}
