package payload

import (
	"encoding/json"
	"github.com/didiercrunch/ecrypto/helper"
	"reflect"
	"strings"
	"testing"
)

func TestJsonifyMetadata(t *testing.T) {
	m := &Metadata{helper.Range(10), helper.Range(12)}
	marshaled, err := json.Marshal(m)
	if err != nil {
		t.Error(err)
		return
	}
	if !strings.Contains(string(marshaled), `"000102030405060708090a0b"`) {
		t.Error(`marshaled data should containe "000102030405060708090a0b"`, "\n", string(marshaled))
	}
	um := new(Metadata)
	if err := json.Unmarshal(marshaled, um); err != nil {
		t.Error(err)
		return
	}
	if !reflect.DeepEqual(m, um) {
		t.Error(m, " is not ", um)
	}

}
