package encrypter

import (
	"io"
)

type StreamEncrypter interface {
	EncrypeReaderToWriter(clearInput io.Reader, encryptedOutput io.Writer)
}

type StreamSigner interface {
	SignStream(input io.Reader) ([]byte, error)
}
