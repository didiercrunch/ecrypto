package contract

import (
	"github.com/didiercrunch/filou/keys"
)

type PublicContract struct {
	AcceptedHashes           []string             `yaml:"accepted_hashes" json:"accepted_hashes"`
	AcceptedBlockCyphers     []string             `yaml:"accepted_block_cypher" json:"accepted_block_cypher"`
	AcceptedBlockCypherModes []string             `yaml:"accepted_block_cypher_mode" json:"accepted_block_cypher_mode"`
	RSAPublicKey             []*keys.RSAPublicKey `yaml:"rsa_public_key, omitempty" json:"rsa_public_key,omitempty"`
}
