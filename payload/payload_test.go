package payload

import (
	"bytes"
	"crypto"
	"io/ioutil"
	"testing"
)

func TestPayloadGetHashMethod(t *testing.T) {
	p := new(Payload)
	testCases := make(map[string]crypto.Hash)
	testCases["md4"] = crypto.MD4
	testCases["md5"] = crypto.MD5
	testCases["sha1"] = crypto.SHA1
	testCases["sha224"] = crypto.SHA224
	testCases["sha256"] = crypto.SHA256
	testCases["sha384"] = crypto.SHA384
	testCases["sha512"] = crypto.SHA512
	testCases["md5sha1"] = crypto.MD5SHA1
	testCases["ripemd160"] = crypto.RIPEMD160
	for name, hashfunc := range testCases {
		p.Hash = hashfunc
		if returnedName := p.GetHashMethod(); returnedName != name {
			t.Error("bad hash name.  supposed to be ", name, " but received", returnedName)
		}
	}
}

func TestGetStream(t *testing.T) {
	p := new(Payload)
	if s := p.getStream(); s != nil {
		t.Fail()
	}
}

func TestEncrypt(t *testing.T) {
	if true {
		return
	}
	p := new(Payload)
	data := bytes.NewBufferString("nice string")
	reader := p.encrypt(data)
	if outputdata, err := ioutil.ReadAll(reader); err != nil {
		t.Error(err)
	} else if string(outputdata) != "nice string" {
		t.Fail()
	}

}
