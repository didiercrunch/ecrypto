package files

import (
	"io"

	"github.com/didiercrunch/filou/contract"
	"github.com/didiercrunch/filou/envelop"
	"github.com/didiercrunch/filou/payload"
)

func EncryptFile(reader io.Reader, writer io.Writer, acceptedContract *contract.AcceptedContract) error {
	payload_, err := payload.GetPayload(reader, acceptedContract)
	if err != nil {
		return err
	}

	//signer :=
	envelop := envelop.NewEnveloper(nil, payload_, nil)
	return envelop.WriteToWriter(writer)
}
