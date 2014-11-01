package contract

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/didiercrunch/filou/helper"
	"github.com/didiercrunch/filou/keys"
)

func assertJSONMarshalling(c *PublicContract) error {
	ret := new(PublicContract)
	if data, err := json.Marshal(c); err != nil {
		return err
	} else if err = json.Unmarshal(data, ret); err != nil {
		return err
	} else if !reflect.DeepEqual(ret, c) {
		fmt.Println(string(data))
		return errors.New("bad serisalization")
	}
	return nil
}

func expectJson(c *PublicContract, jsonData string) error {
	ret := new(PublicContract)
	if err := json.Unmarshal([]byte(jsonData), ret); err != nil {
		return err
	} else if !reflect.DeepEqual(ret.RSAPublicKey, c.RSAPublicKey) {
		return errors.New("bad serisalization")
	}
	return nil
}

func TestSimpleJsonification(t *testing.T) {
	c := &PublicContract{
		AcceptedHashes:       []string{"sha246"},
		AcceptedBlockCyphers: []string{"aes"},
		RSAPublicKey:         []*keys.RSAPublicKey{&keys.RSAPublicKey{helper.B(12), 3}},
	}
	if err := assertJSONMarshalling(c); err != nil {
		t.Error(err)
	}
}

func TestEmptyRSAPublicKey(t *testing.T) {
	c := &PublicContract{
		AcceptedHashes:       []string{"sha246"},
		AcceptedBlockCyphers: []string{"aes"},
		RSAPublicKey:         []*keys.RSAPublicKey{&keys.RSAPublicKey{helper.B(12), 3}},
	}
	jsonData := `{
		"accepted_hashes": ["sha246"], 
		"accepted_block_cypher": ["aes"],
		"rsa_public_key": [{"n": "c", "e": 3}]
	}`
	if err := expectJson(c, jsonData); err != nil {
		t.Error(err)
	}

}
