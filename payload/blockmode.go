package payload

import (
	"crypto/cipher"
	"errors"
	"fmt"
	"io"
)

type BlockStream func(block cipher.Block, iv []byte) cipher.Stream

type BlockMode string

const OFB BlockMode = "ofb"

func (this BlockMode) GetEncryptorStream() (BlockStream, error) {
	switch this {
	case OFB:
		return cipher.NewOFB, nil
	default:
		err := fmt.Sprintf("BlockMode '%s' is not available", this)
		return nil, errors.New(err)
	}

}

func createIV(b cipher.Block, rd io.Reader) ([]byte, error) {
	ret := make([]byte, b.BlockSize())
	if _, err := rd.Read(ret); err != nil {
		return nil, err
	}
	return ret, nil
}

func (this BlockMode) GetCipherStream(block cipher.Block, rd io.Reader) (stream cipher.Stream, iv []byte, err error) {
	iv, err = createIV(block, rd)
	if err != nil {
		return nil, nil, err
	}
	streamCreator, err := this.GetEncryptorStream()
	if err != nil {
		return nil, nil, err
	}
	return streamCreator(block, iv), iv, nil
}
