package payload

import (
	"bytes"
	cryptoSHA512 "crypto/sha512"
	"encoding/hex"
	"errors"
	"github.com/didiercrunch/ecrypto/helper"
	"hash"
	"io"
	"io/ioutil"
	"reflect"
	"testing"
)

var __ = ioutil.Discard

func TestUnsupportedHash(t *testing.T) {
	unsupportedHashes := []Hash{"md5"}
	for _, h := range unsupportedHashes {
		if _, err := h.New(); err == nil {
			t.Error("should have an error here")
		}
	}
}

func assertHashIsOkay(h hash.Hash, expectedHash hash.Hash) error {
	h.Reset()
	expectedHash.Reset()
	h.Write(helper.Range(1025))
	expectedHash.Write(helper.Range(1025))
	if !reflect.DeepEqual(h.Sum(nil), expectedHash.Sum(nil)) {
		return errors.New("bad hash")
	}
	return nil
}

func TestSHA512(t *testing.T) {
	if h, err := SHA512.New(); err != nil {
		t.Error(err)
		return
	} else if err = assertHashIsOkay(h, cryptoSHA512.New()); err != nil {
		t.Error(err)
	}
}

func TestHashOnTheWay(t *testing.T) {
	data := bytes.NewBufferString("some data to hash")
	expectedHash, err := hex.DecodeString("c874e4001dbd67d770c47bc1bac090be069" +
		"2ffcd5bb995de7df96a7100f05e061b85307657e6595d25ffd6931793892dd0215e64" +
		"da97bb73c969b5fcfd2aed49")
	_ = expectedHash
	if err != nil {
		t.Error(err)
		return
	}
	outputHash := make(chan []byte)
	var outputReader io.Reader
	if outputReader, err = SHA512.HashOnTheWay(data, outputHash); err != nil {
		t.Error(err)
	}
	errorC := make(chan error)
	go func(errorC chan error) {
		if output, err := ioutil.ReadAll(outputReader); err != nil {
			errorC <- err
		} else if string(output) != "some data to hash" {
			errorC <- errors.New("output is not 'some data to hash' but: " + string(output))
		}
		errorC <- nil

	}(errorC)
	if err := <-errorC; err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(<-outputHash, expectedHash) {
		t.Error("bad hash\nexpected: ", expectedHash, "\nreceived:", outputHash)
	}
}
