package keygenerator

import (
	"encoding/json"
	"github.com/didiercrunch/ecrypto/keys"
	"io/ioutil"
	"os"
	"path"
	"reflect"
	"testing"
)

func TestCreateRSAKey(t *testing.T) {
	kc := new(KeyGenerator)
	if err := kc.createRSAKey(100); err != nil {
		t.Error(err)
	}
	if kc.privateKey == nil {
		t.Fail()
	}
	if kc.publicKey == nil {
		t.Fail()
	}
}

func TestSaveKeyAsJSON(t *testing.T) {
	kc := new(KeyGenerator)
	if err := kc.CreateNewKey(100); err != nil {
		t.Error(err)
		return
	}
	file, err := ioutil.TempFile(os.TempDir(), "prefix")

	if err != nil {
		t.Error(err)
		return
	}
	defer os.Remove(file.Name())
	if err := kc.saveKeyAsJSON(kc.privateKey, file.Name()); err != nil {
		t.Error(err)
	}
	data, err := ioutil.ReadFile(file.Name())
	if err != nil {
		t.Error(err)
		return
	}
	pk := new(keys.PrivateKey)
	if err := json.Unmarshal(data, pk); err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(pk.Key, kc.privateKey.Key) {
		t.Fail()
	}
}

func TestEnsureDirectoryExists(t *testing.T) {
	kc := new(KeyGenerator)
	fooDir := path.Join(os.TempDir(), "fds89d")
	defer os.Remove(fooDir)
	if err := kc.ensureDirectoryExists(fooDir); err != nil {
		t.Error(err)
		return
	}
	if _, err := os.Stat(fooDir); err != nil {
		t.Error(err)
		return
	}
	if err := kc.ensureDirectoryExists(fooDir); err != nil {
		t.Error(err)
	}

}
