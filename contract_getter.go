package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"github.com/didiercrunch/filou/contract"
)

func fileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func getFileContract(path string) (*contract.PublicContract, error) {
	ret := new(contract.PublicContract)
	if data, err := ioutil.ReadFile(path); err != nil {
		return nil, err
	} else if err = json.Unmarshal(data, ret); err != nil {
		return nil, err
	} else {
		return ret, nil
	}
}

func GetContract(location string) (*contract.PublicContract, error) {
	switch {
	case fileExists(location):
		return getFileContract(location)
	default:
		return nil, errors.New("could not find public contract at " + location)
	}
}
