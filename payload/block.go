package payload

import (
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha512"
	"errors"
	"io"
)

var _ = sha512.New

type Block string

const AES256 Block = "aes256"

func hash512BytesOfData(h crypto.Hash, data io.Reader) ([]byte, error) {
	if !h.Available() {
		return nil, errors.New("unavailable hash function ")
	}
	hf := h.New()
	if _, err := io.CopyN(hf, data, 512); err != nil {
		return nil, err
	}
	return hf.Sum(nil), nil
}

func createRandomKey(lengthInBits int, rd io.Reader) ([]byte, error) {
	if ret, err := hash512BytesOfData(crypto.SHA512, rd); err != nil {
		return nil, err
	} else if lengthInBits > len(ret)*8 {
		return nil, errors.New("asked too long key.")
	} else {
		return ret[0 : lengthInBits/8], nil
	}
}

func (this Block) CreateBlock(rd io.Reader) (cipher.Block, error) {
	var err error
	key := make([]byte, 256/8)
	if _, err = rd.Read(key); err != nil {
		return nil, err
	}
	return aes.NewCipher(key)
}
