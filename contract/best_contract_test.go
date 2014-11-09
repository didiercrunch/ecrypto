package contract

import (
	"testing"
)

func TestGetBestHash(t *testing.T) {
	prvCtr := new(PrivateContract)
	pubCtr := new(PublicContract)

	pubCtr.AcceptedHashes = []string{"SHA384", "SHA256", "MD5"}
	prvCtr.AcceptedHashes = []string{"SHA384", "SHA256"}
	if h, err := getBestHash(prvCtr, pubCtr); err != nil {
		t.Error(err)
	} else if h != "SHA384" {
		t.Error("bad choosen hash: ", h)
	}
	prvCtr.AcceptedHashes = []string{"SHA1", "MD4"}
	if h, err := getBestHash(prvCtr, pubCtr); err == nil {
		t.Error("should have an error here")
	} else if h != "" {
		t.Error("should not have choosen any hash function here")
	}
}

func TestGetFirstItemInList1ThatIsAlsoInList2(t *testing.T) {
	lst1 := []string{"a", "b", "c"}
	lst2 := []string{"b", "c"}
	if c, err := getFirstItemInList1ThatIsAlsoInList2(lst1, lst2); c != "b" || err != nil {
		t.Fail()
	}
	lst2 = []string{"d", "e"}
	if c, err := getFirstItemInList1ThatIsAlsoInList2(lst1, lst2); c != "" || err == nil {
		t.Fail()
	}
}

func TestGetAcceptedContract(t *testing.T) {
	prvCtr := new(PrivateContract)
	pubCtr := new(PublicContract)

	pubCtr.AcceptedHashes = []string{"SHA384", "SHA256", "MD5"}
	prvCtr.AcceptedHashes = []string{"SHA384", "SHA256"}

	pubCtr.AcceptedBlockCyphers = []string{"aes", "des"}
	prvCtr.AcceptedBlockCyphers = []string{"aes"}

	pubCtr.AcceptedBlockCypherModes = []string{"ofb"}
	prvCtr.AcceptedBlockCypherModes = []string{"ofb"}

	pubCtr.AcceptedAsynchronousEncryptionScheme = []string{"rsa_oaep"}
	prvCtr.AcceptedAsynchronousEncryptionScheme = []string{"rsa_oaep"}

	pubCtr.AcceptedSignatureScheme = []string{"rsa_pss"}
	prvCtr.AcceptedSignatureScheme = []string{"rsa_pss"}

	ac, err := GetAcceptedContract(prvCtr, pubCtr)
	if err != nil {
		t.Error(err)
		return
	}
	if ac.Hash != "SHA384" {
		t.Error("bad hash")
	}
	if ac.BlockCipher != "aes" {
		t.Error("bad block cipher")
	}
	if ac.SignatureScheme != "rsa_pss" {
		t.Error("bad signature scheme")
	}
	if ac.AsynchronousEncryptionScheme != "rsa_oaep" {
		t.Error("bad asynchronous encryption scheme")
	}

}
