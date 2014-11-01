package files

import (
	"io"

	"github.com/didiercrunch/filou/contract"
	"github.com/didiercrunch/filou/envelop"
	"github.com/didiercrunch/filou/payload"
)

func EncryptFile(reader io.Reader, writer io.Writer, acceptedContract *contract.AcceptedContract) error {
	payload_ := payload.GetDefaultPayload(reader)
	envelop := envelop.NewEnveloper(nil, payload_, nil)
	return envelop.WriteToWriter(writer)
}
