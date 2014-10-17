package payload

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/didiercrunch/ecrypto/helper"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestPayloadGetHashMethod(t *testing.T) {
	p := new(Payload)
	testCases := make(map[string]Hash)
	testCases["sha512"] = SHA512
	for name, hashfunc := range testCases {
		p.HashMethod = hashfunc
		if returnedName := p.GetHashMethod(); returnedName != name {
			t.Error("bad hash name.  supposed to be ", name, " but received", returnedName)
		}
	}
}

func TestGetStreamAndCreateMetadata(t *testing.T) {
	p := new(Payload)
	p.Block = AES256
	p.BlockMode = OFB
	p.Random = helper.NewMockRandomReader()
	if _, err := p.getStreamAndCreateMetadata(); err != nil {
		t.Error(err)
	} else if len(p.Key) != 256/8 {
		t.Fail()
	} else if len(p.IV) != 128/8 {
		t.Fail()
	}
}

func TestEncrypt(t *testing.T) {
	p := new(Payload)
	p.Block = AES256
	p.HashMethod = SHA512
	p.BlockMode = OFB
	p.Random = helper.NewMockRandomReader()
	data := bytes.NewBufferString("nice string")
	stream, err := p.getStreamAndCreateMetadata()
	if err != nil {
		t.Error(err)
		return
	}
	reader, err := p.encrypt(data, stream)
	if err != nil {
		t.Error(err)
		return
	}
	go func(t *testing.T) {
		if outputdata, err := ioutil.ReadAll(reader); err != nil {
			t.Error(err)
		} else if !reflect.DeepEqual(outputdata, []byte{0, 124, 45, 17, 57, 48, 19, 154, 180, 15, 43}) {
			t.Error("bad output data")
		}
	}(t)
}

func TestGetMode(t *testing.T) {
	p := new(Payload)
	p.BlockMode = OFB
	if p.GetMode() != "ofb" {
		t.Error("bad block mode")
	}
}

func TestGetAlgorithm(t *testing.T) {
	p := new(Payload)
	p.Block = AES256
	if p.GetAlgorithm() != "aes256" {
		t.Error("bad block")
	}
}

func TestPayloadWriter(t *testing.T) {
	pw := getPayloadWriter()
	go func(t *testing.T) { // Need this section!
		if _, err := ioutil.ReadAll(pw.Reader); err != nil {
			t.Error(err)
		}
	}(t)
	if w, err := pw.getDataWriter(); err != nil {
		t.Error(err)
	} else {
		fmt.Fprint(w, "some data")
	}

	if w, err := pw.getMetadataWriter(); err != nil {
		t.Error(err)
	} else {
		fmt.Fprint(w, "some meta data")
	}

	if err := pw.Close(); err != nil {
		t.Error(err)
	}
}

func TestWriteMetadata(t *testing.T) {
	p := new(Payload)
	p.Key = []byte{0, 1, 0, 1}
	p.IV = []byte{10, 11, 10, 11}
	w := helper.NewMockIoWriter()
	if err := p.writeMetadata(w); err != nil {
		t.Error(err)
		return
	} else if !w.IsValidJson() {
		t.Error(w, "is not valid json")
	}
}

func TestGetPayloadData(t *testing.T) {
	p := new(Payload)
	p.DataSource = bytes.NewBufferString("some data to encrypt")
	p.Block = AES256
	p.BlockMode = OFB
	p.HashMethod = SHA512
	p.Random = helper.NewMockRandomReader()
	r, err := p.GetPayloadData()
	if err != nil {
		t.Error(err)
		return
	}
	errorC := make(chan error)
	go func(errorC chan error) {
		if _, err := ioutil.ReadAll(r); err != nil {
			errorC <- err
		}
		errorC <- nil
	}(errorC)

	if err := <-errorC; err != nil {
		t.Error(err)
	} else if len(p.GetHash()) != 512/8 {
		fmt.Println(p.GetHash())
		errorC <- errors.New("hash has not been computed")
	}

}
