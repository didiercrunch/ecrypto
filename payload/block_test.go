package payload

import (
	"crypto"
	"github.com/didiercrunch/ecrypto/helper"
	"testing"
)

func TestCreateBlock(t *testing.T) {
	random := helper.NewMockRandomReader()
	if block, err := AES256.CreateBlock(random); err != nil {
		t.Error(err)
	} else {
		_ = block
	}
}

func Testhash512BytesOfData(t *testing.T) {
	if _, err := hash512BytesOfData(crypto.MD5, helper.NewMockRandomReader()); err == nil {
		t.Error("should never allow MD5 hash")
	}
	rd := helper.NewMockRandomReader()
	if rd, err := hash512BytesOfData(crypto.SHA512, rd); err != nil {
		t.Error(err)
	} else if len(rd) != 512/8 {
		t.Fail()
	}
}

func TestCreateRandomKey(t *testing.T) {
	if key, err := createRandomKey(128, helper.NewMockRandomReader()); err != nil {
		t.Error(err)
	} else if len(key) != 128/8 {
		t.Error("bad key length.  Expected 128 found ", len(key)*8)
	}

}
