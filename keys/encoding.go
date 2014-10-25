package keys

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
)

// All the method to encode/decode keys to json, yaml...

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
