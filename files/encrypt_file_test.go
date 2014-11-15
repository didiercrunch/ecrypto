package files

import (
	"bytes"
	"crypto/rsa"
	"fmt"
	"testing"

	"github.com/didiercrunch/filou/contract"
	"github.com/didiercrunch/filou/helper"
)

func getRSAPrivateKey() contract.GetRSAPrivateKey {
	return func() (*rsa.PrivateKey, error) {
		return helper.GetRSAPrivateKey(2048), nil
	}
}

func getSensibleAcceptedContract() *contract.AcceptedContract {
	ret := new(contract.AcceptedContract)
	ret.Hash = "sha512"
	ret.BlockCipher = "aes256"
	ret.BlockCipherMode = "ofb"
	ret.AsynchronousEncryptionScheme = "rsa_oaep"
	ret.SignatureScheme = "rsa_pss"
	ret.RsaPrivateKey = getRSAPrivateKey()
	ret.RsaPublicKey = helper.GetRSAPublicKey(2048)
	return ret
}

func TestEncryptFile(t *testing.T) {
	buffer := bytes.NewBufferString("some secret text")
	w := helper.NewMockIoWriter()
	c := getSensibleAcceptedContract()
	if err := EncryptFile(buffer, w, c); err != nil {
		t.Error(err)
	}
	fmt.Println(w.String())
}
