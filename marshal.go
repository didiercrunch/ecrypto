package main

import (
	"crypto/rsa"
	"math/big"
)

const (
	DEFAULT_BLOCK_CYPHER      = "AES"
	DEFAULT_BLOCK_CYPHER_MODE = "CBC"
	DEFAULT_HASH_FUNCTION     = "SHA256"
)

type PrivateKey struct {
	Name        string
	Description string
	Type        string // rsa, el gamal,...
	Version     string
	Key         interface{}
	Accept      *AcceptMethods
}

type PublicKey struct {
	Name        string
	Description string
	Type        string // rsa, el gamal,...
	Version     string
	Key         interface{}
	Accept      *AcceptMethods
}

type RSAPublicKey struct {
	N *big.Int
	E int
}

type RSAPrivateKey struct {
	N *big.Int
	P *big.Int
	Q *big.Int
	D *big.Int
	E int
}

type AcceptMethods struct {
	BlocCypher      string
	BlockCypherMode string
	HashFunction    string
}

func GetDefaultRSAPublicKey(publicKey *rsa.PublicKey) *PublicKey {
	key := &RSAPublicKey{publicKey.N, publicKey.E}
	accept := &AcceptMethods{
		DEFAULT_BLOCK_CYPHER,
		DEFAULT_BLOCK_CYPHER_MODE,
		DEFAULT_HASH_FUNCTION,
	}
	return &PublicKey{Type: "RSA", Version: VERSION, Key: key, Accept: accept}
}

func GetDefaultRSAPrivateKey(publicKey *rsa.PrivateKey) (*PrivateKey, error) {
	if err := publicKey.Validate(); err != nil {
		return nil, err
	}
	key := &RSAPrivateKey{publicKey.N, publicKey.Primes[0], publicKey.Primes[1], publicKey.D, publicKey.E}
	accept := &AcceptMethods{
		DEFAULT_BLOCK_CYPHER,
		DEFAULT_BLOCK_CYPHER_MODE,
		DEFAULT_HASH_FUNCTION,
	}
	return &PrivateKey{Type: "RSA", Version: VERSION, Key: key, Accept: accept}, nil
}
