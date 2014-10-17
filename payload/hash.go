package payload

import (
	cryptoSHA512 "crypto/sha512"
	"errors"
	"fmt"
	"hash"
	"io"
)

type Hash string

const SHA512 Hash = "sha512"

func (this Hash) New() (hash.Hash, error) {
	switch this {
	case SHA512:
		return cryptoSHA512.New(), nil
	default:
		msg := fmt.Sprintf("hash '%s' is not supported")
		return nil, errors.New(msg)
	}
}

func (this Hash) HashOnTheWay(data io.Reader, resultingHash chan []byte) (io.Reader, error) {
	h, err := this.New()
	pipeReader, pipeWriter := io.Pipe()
	if err != nil {
		return nil, err
	}
	go func(resultingHash chan []byte) {
		ou1, ou2 := CopyToTwoWriters(h, pipeWriter, data)
		if ou1.Err != nil {
			panic(ou1.Err)
		}
		if ou2.Err != nil {
			panic(ou2.Err)
		}
		if err := pipeWriter.Close(); err != nil {
			panic(err)
		}
		resultingHash <- h.Sum(nil)
	}(resultingHash)
	return pipeReader, nil
}
