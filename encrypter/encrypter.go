package encrypter

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"github.com/didiercrunch/ecrypto/keys"
	"io"
)

type FileEncrypter struct {
	PublicKey    keys.PublicKey
	PrivateKey   keys.PrivateKey
	InputPath    string
	OutputPath   string
	randomReader io.Reader

	iv  []byte
	key []byte
}

func NewEncrypter() *FileEncrypter {
	ret := new(FileEncrypter)
	ret.randomReader = rand.Reader
	return ret
}

func (this *FileEncrypter) createSymetricKey(size int) ([]byte, error) {
	key := make([]byte, size)
	_, err := io.ReadFull(this.randomReader, key)
	return key, err
}

func (this *FileEncrypter) getCypherBlock() (cipher.Block, error) {
	var err error
	this.key, err = this.createSymetricKey(32)
	if err != nil {
		return nil, err
	}
	return aes.NewCipher(this.key)
}

func (this *FileEncrypter) createIV(block cipher.Block) (iv []byte, err error) {
	return this.createSymetricKey(block.BlockSize())
}

func (this *FileEncrypter) createStreamMode(block cipher.Block) (stream cipher.Stream, err error) {
	this.iv, err = this.createIV(block)
	if err != nil {
		return nil, err
	}
	return cipher.NewOFB(block, this.iv), err
}

func (this *FileEncrypter) createFileHandlerAndCypherBlock() (stream cipher.Stream, err error) {
	cypherBlock, err := this.getCypherBlock()
	if err != nil {
		return nil, err
	}
	return this.createStreamMode(cypherBlock)
}

func (this *FileEncrypter) encrypeReaderToWriter(clearInput io.Reader, encryptedOutput io.Writer) error {
	stream, err := this.createFileHandlerAndCypherBlock()
	if err != nil {
		return err
	}
	writer := &cipher.StreamWriter{S: stream, W: encryptedOutput}
	if _, err := io.Copy(writer, clearInput); err != nil {
		return err
	}
	return nil
}

func (this *FileEncrypter) Encrypt() error {
	return nil
}
