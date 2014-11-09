package contract

import (
	"github.com/didiercrunch/filou/keys"
)

type PublicContract struct {
	AcceptedHashes           []string `yaml:"accepted_hashes" json:"accepted_hashes"`
	AcceptedBlockCyphers     []string `yaml:"accepted_block_cypher" json:"accepted_block_cypher"`
	AcceptedBlockCypherModes []string `yaml:"accepted_block_cypher_mode" json:"accepted_block_cypher_mode"`

	AcceptedAsynchronousEncryptionScheme []string `yaml:"accepted_asynchronous_encryption_scheme" json:"accepted_asynchronous_encryption_scheme"`
	AcceptedSignatureScheme              []string `yaml:"accepted_signature_scheme" json:"accepted_signature_scheme"`

	RSAPublicKey []*keys.RSAPublicKey `yaml:"rsa_public_key, omitempty" json:"rsa_public_key,omitempty"`
}
