package main

import (
	"testing"

	"github.com/didiercrunch/filou/contract"
)

func getMockContract() *contract.PublicContract {
	c := new(contract.PublicContract)
	c.AcceptedHashes = []string{"hash_function_1", "hash_function_2"}
	return c
}

func getMockContractAndFile() *contract.PublicContract {

	return nil

}

func TestGetFileContract(t *testing.T) {
}
