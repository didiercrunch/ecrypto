package helper

import (
	"crypto"
	"errors"
	"strings"
)

func JoinAsError(msg []string) error {
	if len(msg) > 0 {
		return errors.New(strings.Join(msg, "\n"))
	}
	return nil
}

func JoinErrors(errs []error) error {
	msgs := make([]string, 0, len(errs))
	for _, err := range errs {
		if err != nil {
			msgs = append(msgs, err.Error())
		}
	}
	return JoinAsError(msgs)
}

func GetHashFunctionByLowerCaseName(hashName string) (crypto.Hash, error) {
	switch hashName {
	case "md4":
		return crypto.MD4, nil
	case "md5":
		return crypto.MD5, nil
	case "sha1":
		return crypto.SHA1, nil
	case "sha224":
		return crypto.SHA224, nil
	case "sha256":
		return crypto.SHA256, nil
	case "sha384":
		return crypto.SHA384, nil
	case "sha512":
		return crypto.SHA512, nil
	case "md5sha1":
		return crypto.MD5SHA1, nil
	case "ripemd160":
		return crypto.RIPEMD160, nil
	default:
		return 0, errors.New("unknown lowercase hash " + hashName)

	}
}
