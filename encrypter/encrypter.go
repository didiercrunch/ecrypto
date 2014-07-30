package encrypter

import (
	"github.com/didiercrunch/ecrypto/keys"
)

type Encrypter struct {
	PublicKey  keys.PublicKey
	PrivateKey keys.PrivateKey
	InputPath  string
	OutputPath string
}

func (this *Encrypter) Encrypt() error {
	// encrypt file
}
