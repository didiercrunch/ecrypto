package files

import (
	"io"

	"github.com/didiercrunch/filou/asymetric"
	"github.com/didiercrunch/filou/contract"
	"github.com/didiercrunch/filou/envelop"
	"github.com/didiercrunch/filou/helper"
	"github.com/didiercrunch/filou/payload"
)

func getEncryptorAndSigner(acceptedContract *contract.AcceptedContract) (asymetric.PublicKeyEncryptor, asymetric.Signer, error) {
	errs := make([]error, 0, 2)
	signer, err := asymetric.GetSigner(acceptedContract)
	if err != nil {
		errs = append(errs, err)
	}

	encryptor, err := asymetric.GetPublicKeyEncryptor(acceptedContract)
	if err != nil {
		errs = append(errs, err)
	}
	return encryptor, signer, helper.JoinErrors(errs)
}

func EncryptFile(reader io.Reader, writer io.Writer, acceptedContract *contract.AcceptedContract) error {
	encryptor, signer, err := getEncryptorAndSigner(acceptedContract)
	if err != nil {
		return err
	}

	payload_, err := payload.GetPayload(reader, acceptedContract)
	if err != nil {
		return err
	}

	envelop := envelop.NewEnveloper(encryptor, payload_, signer)
	return envelop.WriteToWriter(writer)
}
