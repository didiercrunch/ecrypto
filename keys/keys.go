package keys

import (
	"crypto/rsa"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/didiercrunch/ecrypto/shared"
	"math/big"
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

type encodingRSAKey struct {
	N string
	P string
	Q string
	D string
	E int
}

func (this *encodingRSAKey) CreateFromRSAPrivateKey(key *RSAPrivateKey) *encodingRSAKey {
	this.N = fmt.Sprintf("%x", key.N)
	this.P = fmt.Sprintf("%x", key.P)
	this.Q = fmt.Sprintf("%x", key.Q)
	this.D = fmt.Sprintf("%x", key.D)
	this.E = key.E
	return this
}

func (this *encodingRSAKey) CreateFromRSAPublicKey(key *RSAPublicKey) *encodingRSAKey {
	this.N = fmt.Sprintf("%x", key.N)
	this.E = key.E
	return this
}

func (this *encodingRSAKey) SetRSAPrivateKey(key *RSAPrivateKey) (*RSAPrivateKey, error) {
	if i, ok := new(big.Int).SetString(this.N, 16); !ok {
		return nil, errors.New("cannot parse string has hexadecimal string")
	} else {
		key.N = i
	}

	if i, ok := new(big.Int).SetString(this.P, 16); !ok {
		return nil, errors.New("cannot parse string has hexadecimal string")
	} else {
		key.P = i
	}

	if i, ok := new(big.Int).SetString(this.Q, 16); !ok {
		return nil, errors.New("cannot parse string has hexadecimal string")
	} else {
		key.Q = i
	}

	if i, ok := new(big.Int).SetString(this.D, 16); !ok {
		return nil, errors.New("cannot parse string has hexadecimal string")
	} else {
		key.D = i
	}
	key.E = this.E
	return key, nil
}

func (this *encodingRSAKey) SetRSAPublicKey(key *RSAPublicKey) (*RSAPublicKey, error) {
	if i, ok := new(big.Int).SetString(this.N, 16); !ok {
		return nil, errors.New("cannot parse string has hexadecimal string")
	} else {
		key.N = i
	}
	key.E = this.E
	return key, nil
}

func (this *RSAPrivateKey) MarshalJSON() ([]byte, error) {
	e := new(encodingRSAKey)
	return json.Marshal(e.CreateFromRSAPrivateKey(this))
}

func (this *RSAPrivateKey) UnmarshalJSON(data []byte) error {
	obj := new(encodingRSAKey)
	if err := json.Unmarshal(data, obj); err != nil {
		return err
	}
	_, err := obj.SetRSAPrivateKey(this)
	return err
}

func (this *RSAPublicKey) MarshalJSON() ([]byte, error) {
	e := new(encodingRSAKey)
	return json.Marshal(e.CreateFromRSAPublicKey(this))
}

func (this *RSAPublicKey) UnmarshalJSON(data []byte) error {
	obj := new(encodingRSAKey)
	if err := json.Unmarshal(data, obj); err != nil {
		return err
	}
	_, err := obj.SetRSAPublicKey(this)
	return err
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
	return &PrivateKey{Type: "RSA", Version: shared.VERSION, Key: key, Accept: accept}, nil
}
