package contract

import (
	"github.com/didiercrunch/filou/keys"
)

type PrivateContract struct {
	PublicContract
	RSAPrivateKey []*keys.RSAPrivateKey `yaml:"rsa_private_key,omitempty" json:"rsa_private_key,omitempty"`
}
