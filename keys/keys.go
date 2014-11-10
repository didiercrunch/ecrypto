package keys

import (
	"crypto/rsa"
	"math/big"

	"github.com/didiercrunch/filou/shared"
)

const (
	DEFAULT_BLOCK_CYPHER      = "AES"
	DEFAULT_BLOCK_CYPHER_MODE = "CBC"
	DEFAULT_HASH_FUNCTION     = "SHA256"
)

type PrivateKey struct {
	Version string
	Type    string // rsa, el gamal,...
	Key     *RSAPrivateKey
	Accept  *AcceptMethods
}

type PublicKey struct {
	Type    string // rsa, el gamal,...
	Version string
	Key     *RSAPublicKey
	Accept  *AcceptMethods
}

type RSAPublicKey struct {
	N *big.Int
	E int
}

type RSAPrivateKey struct {
	N *big.Int `json:"n,omitempty"`
	P *big.Int `json:"p,omitempty"`
	Q *big.Int `json:"q,omitempty"`
	D *big.Int `json:"d,omitempty"`
	E int      `json:"e,omitempty"`
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
	return &PublicKey{Type: "RSA", Version: shared.VERSION, Key: key, Accept: accept}
}

func NewRSAPublicKey(key *rsa.PublicKey) *RSAPublicKey {
	return &RSAPublicKey{key.N, key.E}
}

func NewRSAPrivateKey(key *rsa.PrivateKey) *RSAPrivateKey {
	return &RSAPrivateKey{key.N, key.Primes[0], key.Primes[1], key.D, key.E}
}

func GetDefaultRSAPrivateKey(privateKey *rsa.PrivateKey) (*PrivateKey, error) {
	if err := privateKey.Validate(); err != nil {
		return nil, err
	}
	key := NewRSAPrivateKey(privateKey)
	accept := &AcceptMethods{
		DEFAULT_BLOCK_CYPHER,
		DEFAULT_BLOCK_CYPHER_MODE,
		DEFAULT_HASH_FUNCTION,
	}
	return &PrivateKey{Type: "RSA", Version: shared.VERSION, Key: key, Accept: accept}, nil
}
