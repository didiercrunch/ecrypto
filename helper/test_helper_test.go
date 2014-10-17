package helper

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"math/rand"
	"reflect"
	"testing"
)

func TestRange(t *testing.T) {
	if !reflect.DeepEqual(Range(3), []byte{0, 1, 2}) {
		t.Fail()
	}
}

func TestFromRange(t *testing.T) {
	if !reflect.DeepEqual(FromRange(3, 5), []byte{3, 4}) {
		t.Fail()
	}
}

func TestB(t *testing.T) {
	if !reflect.DeepEqual(B(3), big.NewInt(3)) {
		t.Fail()
	}
}

func TestMockIoWriter(t *testing.T) {
	w := NewMockIoWriter()
	fmt.Fprint(w, "something to write")
	if !reflect.DeepEqual(w.Bytes(), []byte("something to write")) {
		t.Fail()
	} else if "something to write" != w.String() {
		t.Fail()
	} else if w.Length() != len("something to write") {
		t.Fail()
	} else if readerOutput, err := ioutil.ReadAll(w.Reader()); err != nil {
		t.Error(err)
	} else if "something to write" != string(readerOutput) {
		t.Error("bad reader")
	}
}

func TestIsValidJson(t *testing.T) {
	m := make(map[string]bool)
	m[`{}`] = true
	m[`[]`] = true
	m[`[1, "abc", {"aa": true, "2": [1,2,3]}]`] = true
	m[`{"aa": true, "2": [1,2,3]}`] = true

	m[`regular string`] = false
	m[`{"aa": true "2": [1,2,3]}`] = false // malformed json
	for json_, isJSON := range m {
		w := &MockIoWriter{[]byte(json_)}
		ok := w.IsValidJson()
		if ok && !isJSON {
			t.Error(json_, "is not json but has been detected as json")
		} else if !ok && isJSON {
			t.Error(json_, "is json but has not been detected as json")
		}
	}
}

func TestGetRSAPublicKeyAndGetRSAPrivateKey(t *testing.T) {
	privateKey := GetRSAPrivateKey(50)
	if err := privateKey.Validate(); err != nil {
		t.Error(err)
		return
	}
	publicKey := GetRSAPublicKey(50)
	if publicKey.E != privateKey.E || publicKey.N.Cmp(privateKey.N) != 0 {
		t.Error("private and public keys do not match", publicKey.N, privateKey.N)
	}

	if publicKey.N.BitLen() != 50 {
		t.Error("modulo has bad bit length")
	}
}

func TestGetNextFourRandomBytes(t *testing.T) {
	m := &MockRandomReader{0, rand.New(rand.NewSource(22))}
	rb := m.getNextFourRandomBytes()
	var expt []byte = []byte{233, 237, 107, 32}
	if len(rb) != 4 {
		t.Error("bad random length")
		return
	} else if !reflect.DeepEqual(expt, rb) {
		t.Error("unexpect randome data")
	}

	rb = m.getNextFourRandomBytes()
	expt = []byte{27, 131, 191, 20}
	if len(rb) != 4 {
		t.Error("bad random length")
		return
	} else if !reflect.DeepEqual(expt, rb) {
		t.Error("unexpect randome data")
	}
}

func TestMockRandomReader(t *testing.T) {
	m := NewMockRandomReader()
	dump := make([]byte, 4*10+3)
	if l, err := m.Read(dump); l != len(dump) || err != nil {
		t.Error("that is very bad :(")
		return
	}
	var expt []byte = []byte{73, 233, 188, 33, 31, 118, 98, 81, 66, 120, 254, 85, 138,
		36, 118, 80, 38, 83, 169, 57, 27, 61, 103, 122, 37, 164, 185, 40, 112, 66, 1,
		50, 185, 128, 156, 20, 8, 164, 37, 13, 182, 50, 175}
	if !reflect.DeepEqual(dump, expt) {
		t.Error("unexpected returned data\n", dump)
	}
}

func TestGetRandomBytes(t *testing.T) {
	rdBytes := NewMockRandomReader().GetRandomBytes(10)
	expt := make([]byte, 10)
	NewMockRandomReader().Read(expt)
	if !reflect.DeepEqual(expt, rdBytes) {
		t.Fail()
	}
}

func TestGetReadLength(t *testing.T) {
	m := NewMockRandomReader()
	dump := make([]byte, 4*10+3)
	if _, err := m.Read(dump); err != nil {
		t.Error(err)
		return
	}
	if m.GetReadLength() != 4*10+3 {
		t.Error("bad read length")
	}
}
