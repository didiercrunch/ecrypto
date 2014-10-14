package helper

import (
	"bytes"
	"crypto/rsa"
	"encoding/binary"
	"io"
	"math/big"
	"math/rand"
)

func B(x int64) *big.Int {
	return big.NewInt(x)
}

type MockRandomReader struct {
	random *rand.Rand
}

func NewMockRandomReader() *MockRandomReader {
	return &MockRandomReader{rand.New(rand.NewSource(99))}
}

func (this *MockRandomReader) getNextFourRandomBytes() []byte {
	ret := make([]byte, 4)
	binary.LittleEndian.PutUint32(ret, uint32(this.random.Int31()))
	return ret
}

func (this *MockRandomReader) Read(p []byte) (n int, err error) {
	for i := 0; i < len(p); i += 4 {
		copy(p[i:], this.getNextFourRandomBytes())
	}
	return len(p), nil
}

func (this *MockRandomReader) GetRandomBytes(size int) []byte {
	ret := make([]byte, size)
	this.Read(ret)
	return ret

}

type MockIoWriter struct {
	state []byte
}

func NewMockIoWriter() *MockIoWriter {
	return &MockIoWriter{make([]byte, 0, 1024)}
}

func (this *MockIoWriter) Write(p []byte) (n int, err error) {
	this.state = append(this.state, p...)
	return len(p), nil
}

func (this *MockIoWriter) String() string {
	return string(this.state)
}

func (this *MockIoWriter) Length() int {
	return len(this.state)
}

func (this *MockIoWriter) Bytes() []byte {
	return this.state
}

func (this *MockIoWriter) Reader() io.Reader {
	return bytes.NewReader(this.state)
}

func (this *MockIoWriter) ReaderAt() io.ReaderAt {
	return bytes.NewReader(this.state)
}

func Range(length int) []byte {
	ret := make([]byte, length)
	for i, _ := range ret {
		ret[i] = byte(i)
	}
	return ret
}

func FromRange(from, to int) []byte {
	ret := make([]byte, to-from)
	for i, _ := range ret {
		ret[i] = byte(i + from)
	}
	return ret
}

func GetRSAPublicKey(size int) *rsa.PublicKey {
	return &GetRSAPrivateKey(size).PublicKey
}

func GetRSAPrivateKey(size int) *rsa.PrivateKey {
	key, err := rsa.GenerateKey(NewMockRandomReader(), size)
	if err != nil {
		panic(err)
	}
	return key

}
