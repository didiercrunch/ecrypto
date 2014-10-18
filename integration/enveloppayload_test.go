package integration

import (
	"bytes"
	"testing"

	"github.com/didiercrunch/filou/envelop"
	"github.com/didiercrunch/filou/payload"
)

type mockEncryptor struct{}

func (this *mockEncryptor) Encrypt(data []byte) ([]byte, error) {
	// cesar encryption!
	ret := make([]byte, len(data))
	for i, d := range data {
		ret[i] = d + 3 // modulo is implicit
	}
	return ret, nil
}

type mockSigner struct{}

func (this *mockSigner) Sign(data []byte) ([]byte, error) {
	// a slice with the sum of the data plus a "secret" value
	var sum byte = 36
	for _, d := range data {
		sum += d
	}
	return []byte{sum}, nil
}

func TestPayloadCanBeUseInEnvelop(t *testing.T) {
	payload_ := payload.GetDefaultPayload(bytes.NewBufferString("some data to encrypt"))
	envelop := envelop.NewEnveloper(&mockEncryptor{}, payload_, &mockSigner{})
	if envelop == nil {
		t.Fail()
	}
}
