package payload

import (
	"crypto"
	"crypto/cipher"
	"fmt"
	"io"
)

var _ = fmt.Print

/*
	type Payload interface {
		GetHash() []byte
		GetKey() []byte
		GetAlgorithm() string
		GetMode() string
		GetHashMethod() string
		GetPayloadData() io.Reader
	}
*/

type BlockMode string

const OFB BlockMode = "ofb"

type BlockModeStream func(block cipher.Block, iv []byte) cipher.Stream

type Payload struct {
	DataSource io.Reader
	Hash       crypto.Hash
	Block      Block
	BlockMode  BlockMode
}

func (this *Payload) GetHashMethod() string {
	switch this.Hash {
	case crypto.MD5:
		return "md5"
	case crypto.MD4:
		return "md4"
	case crypto.SHA1:
		return "sha1"
	case crypto.MD5SHA1:
		return "md5sha1"
	case crypto.RIPEMD160:
		return "ripemd160"
	case crypto.SHA224:
		return "sha224"
	case crypto.SHA256:
		return "sha256"
	case crypto.SHA384:
		return "sha384"
	case crypto.SHA512:
		return "sha512"
	default:
		panic("unknown hash method")
	}
}

func (this *Payload) GetPayloadData() io.Reader {
	//  zip data source

	// create metadata
	// zip the zipped data and metadata

	// encryptp everything togeter
	return nil
}

func (this *Payload) getStream() cipher.Stream {
	return nil //this.BlockMode
}

func (this *Payload) encrypt(reader io.Reader) io.Reader {
	ret, writer := io.Pipe()
	go func() {
		cipherWriter := &cipher.StreamWriter{S: this.getStream(), W: writer}
		io.Copy(cipherWriter, reader)
		cipherWriter.Close()
	}()
	return ret

}
