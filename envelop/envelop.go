package envelop

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
)

type Payload interface {
	GetHash() []byte
	GetKey() []byte
	GetAlgorithm() string
	GetMode() string
	GetHashMethod() string
	GetPayloadData() (io.Reader, error)
}

type PublicKeyEncryptor interface {
	Encrypt(data []byte) ([]byte, error)
}

type Signer interface {
	Sign(data []byte) ([]byte, error)
}

func NewEnveloper(publicKeyEncryptor PublicKeyEncryptor, payload Payload, signer Signer) *Enveloper {
	return &Enveloper{
		publicKeyEncryptor: publicKeyEncryptor,
		payload:            payload,
		signer:             signer,
	}
}

type Enveloper struct {
	publicKeyEncryptor PublicKeyEncryptor
	payload            Payload
	signer             Signer
	EncryptedKey       []byte
	Signature          []byte
	Metadata           *Metadata
}

func (this *Enveloper) EncryptKey() ([]byte, error) {
	return this.publicKeyEncryptor.Encrypt(this.payload.GetKey())
}

func (this *Enveloper) SignPayload() ([]byte, error) {
	return this.signer.Sign(this.payload.GetHash())
}

func (this *Enveloper) CreateMetadata() *Metadata {
	md := new(Metadata)
	md.BlockAlgorithm = this.payload.GetAlgorithm()
	md.BlockMode = this.payload.GetMode()
	md.SignatureAlgorithm = "RSAPSS"
	md.HashMethod = this.payload.GetHashMethod()
	return md
}

func (this *Enveloper) CreateEnvelopCompletely() error {
	var err error
	if this.EncryptedKey, err = this.EncryptKey(); err != nil {
		return err
	} else if this.Signature, err = this.SignPayload(); err != nil {
		return err
	}
	this.Metadata = this.CreateMetadata()
	return nil
}

func (this *Enveloper) WritePayload(z *zip.Writer) error {
	dataWriter, err := z.Create("data")
	if err != nil {
		return err
	}
	if r, err := this.payload.GetPayloadData(); err != nil {
		return err
	} else if _, err = io.Copy(dataWriter, r); err != nil {
		return err
	}
	return nil
}

func (this *Enveloper) WriteMetadata(z *zip.Writer) error {
	dataWriter, err := z.Create("metadata.json")
	if err != nil {
		return err
	}
	encoder := json.NewEncoder(dataWriter)
	return encoder.Encode(this.Metadata)
}

func (this *Enveloper) WriteEncryptedKey(z *zip.Writer) error {
	dataWriter, err := z.Create("key.json")
	if err != nil {
		return err
	}
	encoder := json.NewEncoder(dataWriter)
	return encoder.Encode(&struct {
		Key string
	}{fmt.Sprintf("%s", this.EncryptedKey)})
}

func (this *Enveloper) WriteToWriter(w io.Writer) error {
	z := zip.NewWriter(w)
	defer z.Close()
	errc := make(chan error)
	go func() {
		errc <- this.WritePayload(z)
	}()

	if err := <-errc; err != nil {
		return err
	}
	if err := this.CreateEnvelopCompletely(); err != nil {
		return err
	}

	if err := this.WriteMetadata(z); err != nil {
		return err
	} else if err := this.WriteEncryptedKey(z); err != nil {
		return err
	}
	return nil
}
