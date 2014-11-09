package contract

import (
	"errors"

	"crypto/rsa"
	"github.com/didiercrunch/filou/helper"
)

type GetRSAPrivateKey func() (*rsa.PrivateKey, error)

type AcceptedContract struct {
	Hash            string
	BlockCipher     string
	BlockCipherMode string

	AsynchronousEncryptionScheme string
	SignatureScheme              string

	RsaPublicKey  *rsa.PublicKey
	RsaPrivateKey GetRSAPrivateKey
}

func listContains(lst []string, elm string) bool {
	for _, elm_ := range lst {
		if elm == elm_ {
			return true
		}
	}
	return false
}

func getFirstItemInList1ThatIsAlsoInList2(lst1, lst2 []string) (string, error) {
	for _, elm := range lst1 {
		if listContains(lst2, elm) {
			return elm, nil
		}
	}
	return "", errors.New("no items in list1 is in list2")
}

func getAsynchronousEncryptionScheme(prvCtr *PrivateContract, pubCtr *PublicContract) (string, error) {
	if b, err := getFirstItemInList1ThatIsAlsoInList2(pubCtr.AcceptedAsynchronousEncryptionScheme, prvCtr.AcceptedAsynchronousEncryptionScheme); err != nil {
		return "", errors.New("cannot find a asynchronous encryption scheme that satisfy both contracts")
	} else {
		return b, nil
	}
}

func getSignatureScheme(prvCtr *PrivateContract, pubCtr *PublicContract) (string, error) {
	if b, err := getFirstItemInList1ThatIsAlsoInList2(pubCtr.AcceptedSignatureScheme, prvCtr.AcceptedSignatureScheme); err != nil {
		return "", errors.New("cannot find a signature scheme that satisfy both contracts")
	} else {
		return b, nil
	}
}

func getBestBlockCipher(prvCtr *PrivateContract, pubCtr *PublicContract) (string, error) {
	if b, err := getFirstItemInList1ThatIsAlsoInList2(pubCtr.AcceptedBlockCyphers, prvCtr.AcceptedBlockCyphers); err != nil {
		return "", errors.New("cannot find a block cypher that satisfy both contracts")
	} else {
		return b, nil
	}
}

func getBestBlockCipherMode(prvCtr *PrivateContract, pubCtr *PublicContract) (string, error) {
	if m, err := getFirstItemInList1ThatIsAlsoInList2(pubCtr.AcceptedBlockCypherModes, prvCtr.AcceptedBlockCypherModes); err != nil {
		return "", errors.New("cannot find a block cypher mode that satisfy both contracts")
	} else {
		return m, nil
	}
}

func getBestHash(prvCtr *PrivateContract, pubCtr *PublicContract) (string, error) {
	if h, err := getFirstItemInList1ThatIsAlsoInList2(pubCtr.AcceptedHashes, prvCtr.AcceptedHashes); err != nil {
		return "", errors.New("cannot find a hash that satisfy both contracts")
	} else {
		return h, nil
	}
}

func GetAcceptedContract(prvCtr *PrivateContract, pubCtr *PublicContract) (*AcceptedContract, error) {
	ret := new(AcceptedContract)
	errs := make([]string, 0, 4)
	var err error
	if ret.Hash, err = getBestHash(prvCtr, pubCtr); err != nil {
		errs = append(errs, err.Error())
	}
	if ret.BlockCipher, err = getBestBlockCipher(prvCtr, pubCtr); err != nil {
		errs = append(errs, err.Error())
	}
	if ret.BlockCipherMode, err = getBestBlockCipherMode(prvCtr, pubCtr); err != nil {
		errs = append(errs, err.Error())
	}
	if ret.AsynchronousEncryptionScheme, err = getAsynchronousEncryptionScheme(prvCtr, pubCtr); err != nil {
		errs = append(errs, err.Error())
	}
	if ret.SignatureScheme, err = getSignatureScheme(prvCtr, pubCtr); err != nil {
		errs = append(errs, err.Error())
	}

	return ret, helper.JoinAsError(errs)
}
