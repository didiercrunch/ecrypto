package asymetric

import (
	"crypto/rand"
	"errors"

	"github.com/didiercrunch/filou/contract"
	"github.com/didiercrunch/filou/helper"
)

type PublicKeyEncryptor interface {
	Encrypt(data []byte) ([]byte, error)
}

type Signer interface {
	Sign(data []byte) ([]byte, error)
}

func GetSigner(contract_ *contract.AcceptedContract) (Signer, error) {
	if contract_.SignatureScheme == "rsa_pss" {
		return getRsaPssSignerFromContract(contract_)
	}
	return nil, errors.New("signature scheme " + contract_.SignatureScheme + " not supported")
}

func GetPublicKeyEncryptor(contract_ *contract.AcceptedContract) (PublicKeyEncryptor, error) {
	if contract_.AsynchronousEncryptionScheme == "rsa_oaep" {
		return getRsaOaepEncryptorFromContract(contract_)
	} else {
		return nil, errors.New("asynchronous encryption scheme " + contract_.AsynchronousEncryptionScheme + " not supported")
	}
}

func getRsaPssSignerFromContract(con *contract.AcceptedContract) (*RsaPssSigner, error) {
	if privKey, err := con.RsaPrivateKey(); err != nil {
		return nil, err
	} else if hash, err := helper.GetHashFunctionByLowerCaseName(con.Hash); err != nil {
		return nil, err
	} else {
		return &RsaPssSigner{privKey, con.RsaPublicKey, hash, rand.Reader}, nil
	}
}

func getRsaOaepEncryptorFromContract(con *contract.AcceptedContract) (*RsaOaepEncrypter, error) {
	if privKey, err := con.RsaPrivateKey(); err != nil {
		return nil, err
	} else if hash, err := helper.GetHashFunctionByLowerCaseName(con.Hash); err != nil {
		return nil, err
	} else {
		return &RsaOaepEncrypter{privKey, con.RsaPublicKey, hash, rand.Reader}, nil
	}
}
