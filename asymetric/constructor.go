package asymetric

import (
	"github.com/didiercrunch/filou/contract"
)

type PublicKeyEncryptor interface {
	Encrypt(data []byte) ([]byte, error)
}

type Signer interface {
	Sign(data []byte) ([]byte, error)
}

func GetSigner(contract_ contract.AcceptedContract) (Signer, error) {
	return nil, nil
}

func GetPublicKeyEncryptor(contract_ contract.AcceptedContract) (Signer, error) {
	return nil, nil
}
