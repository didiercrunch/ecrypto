package asymetric

import (
	_ "crypto/sha512"

	"crypto/rsa"

	"github.com/didiercrunch/filou/helper"
)

var random = helper.NewMockRandomReader()

func generateKeyPair() (*rsa.PrivateKey, *rsa.PublicKey) {
	if key, err := rsa.GenerateKey(random, 1024); err != nil {
		panic(err)
	} else {
		return key, &key.PublicKey
	}
}
