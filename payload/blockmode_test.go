package payload

import (
	"crypto/cipher"
	"crypto/des"
	"github.com/didiercrunch/filou/helper"
	"testing"
)

func getAnyBlock() cipher.Block {
	block, err := des.NewTripleDESCipher(helper.Range(24))
	if err != nil {
		panic(err)
	}
	return block
}

func TestGetEncryptorStreamError(t *testing.T) {
	var foo BlockMode = "some_shit"
	if _, err := foo.GetEncryptorStream(); err == nil {
		t.Error("should have an error here")
	} else if err.Error() != "BlockMode 'some_shit' is not available" {
		t.Error("bad error message")
	}
}

func TestOFB(t *testing.T) {
	_, err := OFB.GetEncryptorStream()
	if err != nil {
		t.Error(err)
	}
}

func TestCreateIV(t *testing.T) {
	block := getAnyBlock()
	if iv, err := createIV(block, helper.NewMockRandomReader()); err != nil {
		t.Error(err)
	} else if len(iv) != block.BlockSize() {
		t.Error("bad iv length")
	}
}

func TestGetCipherStream(t *testing.T) {
	block := getAnyBlock()
	if _, _, err := OFB.GetCipherStream(block, helper.NewMockRandomReader()); err != nil {
		t.Error(err)
	}
}
