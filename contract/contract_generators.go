package contract

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"errors"
	"os"
	"path"

	"github.com/didiercrunch/filou/keys"
	"github.com/didiercrunch/filou/shared"
)

type ContractsGenerator struct {
	publicKey  *keys.RSAPublicKey
	privateKey *keys.RSAPrivateKey

	privateContract *PrivateContract
	publicContract  *PublicContract
}

func (this *ContractsGenerator) ensureDirectoryExists(dir string) error {
	stat, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return os.Mkdir(dir, 0700)
	} else if err != nil {
		return err
	} else if !stat.IsDir() {
		return errors.New(dir + " is not a directory")
	}
	return nil
}

func (this *ContractsGenerator) saveContractAsJSON(contract interface{}, filepath string) error {
	w, err := os.Create(filepath)
	if err != nil {
		return err
	}
	enc := json.NewEncoder(w)
	return enc.Encode(contract)
}

func (this *ContractsGenerator) saveContracts(dirPath string) error {
	if err := this.ensureDirectoryExists(dirPath); err != nil {
		return err
	}
	if err := this.saveContractAsJSON(this.publicContract, path.Join(dirPath, "public_contract.json")); err != nil {
		return err
	}
	return this.saveContractAsJSON(this.privateContract, path.Join(dirPath, "private_contract.json"))
}

func (this *ContractsGenerator) createRSAKeys(size int) error {
	rsaPrivateKey, err := rsa.GenerateKey(rand.Reader, size)
	if err != nil {
		return errors.New("cannot generate the rsa public/private key pair\n" + err.Error())
	}

	this.privateKey = keys.NewRSAPrivateKey(rsaPrivateKey)
	this.publicKey = keys.NewRSAPublicKey(&rsaPrivateKey.PublicKey)
	return nil
}

func (this *ContractsGenerator) createPrivateContractWithGoodDefaultValues(key *keys.RSAPrivateKey) {
	this.privateContract = new(PrivateContract)
	this.privateContract.AcceptedAsynchronousEncryptionScheme = []string{"rsa_oaep"}
	this.privateContract.AcceptedSignatureScheme = []string{"rsa_pss"}
	this.privateContract.AcceptedBlockCypherModes = []string{"ofb"}
	this.privateContract.AcceptedBlockCyphers = []string{"aes"}
	this.privateContract.RSAPrivateKey = []*keys.RSAPrivateKey{key}
	this.privateContract.RSAPublicKey = []*keys.RSAPublicKey{}
}

func (this *ContractsGenerator) createPublicContractWithGoodDefaultValues(key *keys.RSAPublicKey) {
	this.publicContract = new(PublicContract)
	this.privateContract.AcceptedAsynchronousEncryptionScheme = []string{"rsa_oaep"}
	this.privateContract.AcceptedSignatureScheme = []string{"rsa_pss"}
	this.privateContract.AcceptedBlockCypherModes = []string{"ofb"}
	this.privateContract.AcceptedBlockCyphers = []string{"aes"}
	this.privateContract.AcceptedHashes = []string{"sha384", "sha512", "sha256"}
	this.privateContract.RSAPublicKey = []*keys.RSAPublicKey{key}
}

func (this *ContractsGenerator) createContracts(keySize int) error {
	if err := this.createRSAKeys(keySize); err != nil {
		return err
	}
	this.createPrivateContractWithGoodDefaultValues(this.privateKey)
	this.createPublicContractWithGoodDefaultValues(this.publicKey)
	return nil
}

func (this *ContractsGenerator) CreateContracts(size int) error {
	if err := this.createContracts(size); err != nil {
		return err
	}
	return this.saveContracts(shared.GetEcryptoDir())
}
