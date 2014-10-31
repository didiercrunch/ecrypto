package contract

import (
	"github.com/didiercrunch/filou/keys"
)

type PublicContract struct {
	AcceptedHashes      []string             `yaml:"accepted_hashes" json:"accepted_hashes"`
	AcceptedBlockCypher []string             `yaml:"accepted_block_cypher" json:"accepted_block_cypher"`
	RSAPublicKey        []*keys.RSAPublicKey `yaml:"rsa_public_key, omitempty" json:"rsa_public_key,omitempty"`
}
