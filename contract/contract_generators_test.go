package contract

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"reflect"
	"testing"
)

func TestCreateRSAKey(t *testing.T) {
	kc := new(ContractsGenerator)
	if err := kc.createRSAKeys(100); err != nil {
		t.Error(err)
	}
	if kc.privateKey == nil {
		t.Fail()
	}
	if kc.publicKey == nil {
		t.Fail()
	}
}

func TestSaveContractAsJSON(t *testing.T) {
	cg := new(ContractsGenerator)
	if err := cg.createContracts(100); err != nil {
		t.Error(err)
		return
	}
	file, err := ioutil.TempFile(os.TempDir(), "prefix")

	if err != nil {
		t.Error(err)
		return
	}
	defer os.Remove(file.Name())
	if err := cg.saveContractAsJSON(cg.privateContract, file.Name()); err != nil {
		t.Error(err)
	}
	data, err := ioutil.ReadFile(file.Name())
	if err != nil {
		t.Error(err)
		return
	}
	privC := new(PrivateContract)
	if err := json.Unmarshal(data, privC); err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(privC, cg.privateContract) {
		t.Fail()
	}
}

func TestEnsureDirectoryExists(t *testing.T) {
	kc := new(ContractsGenerator)
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
