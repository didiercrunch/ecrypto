package main

import (
	"io"
)

type Params struct {
}

type Envelop struct {
	err    error
	params *Params
}

func ZipIoReader(reader io.Reader) {

}

func createEnvelop(reader io.Reader, writer io.WriteCloser, params *Params) error {
	env := &Envelop{nil, params}
	env.ZipDataForPayload()
	env.CreatePayloadMetadata()
	env.EncryptPayload()
	env.CreateEnvelopMetadata()
	env.EncryptPayloadKeyAndAddItToEncelop()
	return env.getError()
}
