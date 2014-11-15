package helper

import (
	"bytes"
	"crypto/rsa"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/big"
	"math/rand"
	"os"
)

func B(x int64) *big.Int {
	return big.NewInt(x)
}

type MockRandomReader struct {
	readLength int
	random     *rand.Rand
}

func NewMockRandomReader() *MockRandomReader {
	return &MockRandomReader{0, rand.New(rand.NewSource(99))}
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
	this.readLength += len(p)
	return len(p), nil
}

func (this *MockRandomReader) GetReadLength() int {
	return this.readLength
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

func (this *MockIoWriter) IsValidJson() bool {
	m := make(map[string]interface{})
	l := make([]interface{}, 0, 100)
	if json.Unmarshal(this.Bytes(), &m) == nil {
		return true
	} else if json.Unmarshal(this.Bytes(), &l) == nil {
		return true
	}
	return false
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

// return a tmp file in the default tmp directory.  If fails, the function
// will panic
func GetTmpEmptyFile() *os.File {
	if f, err := ioutil.TempFile("", ""); err != nil {
		panic(err)
	} else {
		return f
	}
}

//  returned a closed tmp file with text in it.
func GetTmpFileWithText(text string) *os.File {
	f := GetTmpEmptyFile()
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	fmt.Fprint(f, text)
	return f
}

//  returns a path that can be safely use as tmp file
func GetTmpFileName() string {
	f := GetTmpEmptyFile()
	if err := f.Close(); err != nil {
		panic(err)
	}
	if err := os.Remove(f.Name()); err != nil {
		panic(err)
	}
	return f.Name()
}

//  create a tmp dir and return its path.  If failing, the method will panic
func GetTmpEmptyDir() string {
	if dirName, err := ioutil.TempDir("", ""); err != nil {
		panic(err)
	} else {
		return dirName
	}
}
