package payload

import (
	"archive/zip"
	"crypto/cipher"
	"crypto/rand"
	"encoding/json"
	"io"

	"github.com/didiercrunch/filou/contract"
	"github.com/didiercrunch/filou/helper"
)

type BlockModeStream func(block cipher.Block, iv []byte) cipher.Stream

type Payload struct {
	DataSource io.Reader
	Random     io.Reader
	HashMethod Hash
	Block      Block
	BlockMode  BlockMode

	IV   []byte
	Key  []byte
	Hash []byte
}

func GetDefaultPayload(dataSource io.Reader) *Payload {
	return &Payload{
		DataSource: dataSource,
		Random:     rand.Reader,
		HashMethod: SHA512,
		Block:      AES256,
		BlockMode:  OFB,
	}
}

func GetPayload(dataSource io.Reader, acceptedContract *contract.AcceptedContract) (*Payload, error) {
	ret := &Payload{DataSource: dataSource, Random: rand.Reader}
	errs := make([]error, 0)
	var err error
	if ret.HashMethod, err = GetHashByHashName(acceptedContract.Hash); err != nil {
		errs = append(errs, err)
	}
	if ret.Block, err = GetBlockByBlockName(acceptedContract.BlockCipher); err != nil {
		errs = append(errs, err)
	}
	if ret.BlockMode, err = GetBlockModeByName(acceptedContract.BlockCipherMode); err != nil {
		errs = append(errs, err)
	}

	return ret, helper.JoinErrors(errs)
}

type payloadWriter struct {
	Reader     io.Reader
	Pipewriter io.WriteCloser
	ZipWriter  *zip.Writer
}

func getPayloadWriter() *payloadWriter {
	p := new(payloadWriter)
	p.Reader, p.Pipewriter = io.Pipe()
	p.ZipWriter = zip.NewWriter(p.Pipewriter)
	return p
}

func (this *payloadWriter) getDataWriter() (io.Writer, error) {
	return this.ZipWriter.Create("data")
}

func (this *payloadWriter) getMetadataWriter() (io.Writer, error) {
	return this.ZipWriter.Create("metadata")
}

func (this *payloadWriter) Close() error {
	if err := this.ZipWriter.Close(); err != nil {
		return err
	}
	return this.Pipewriter.Close()

}

func (this *Payload) GetMode() string {
	return string(this.BlockMode)
}

func (this *Payload) GetHash() []byte {
	return this.Hash
}

func (this *Payload) GetKey() []byte {
	return this.Key
}

func (this *Payload) GetAlgorithm() string {
	return string(this.Block)
}

func (this *Payload) GetHashMethod() string {
	return string(this.HashMethod)
}

func (this *Payload) GetPayloadData() (io.Reader, error) {
	pw := getPayloadWriter()
	stream, err := this.getStreamAndCreateMetadata()
	if err != nil {
		return nil, err
	}
	go func(payload *Payload, pw *payloadWriter) {
		if mdw, err := pw.getMetadataWriter(); err != nil {
			panic(err)
		} else if err := payload.writeMetadata(mdw); err != nil {
			panic(err)
		}

		if mw, err := pw.getDataWriter(); err != nil {
			panic(err)
		} else if _, err := io.Copy(mw, payload.DataSource); err != nil {
			panic(err)
		}
		if err := pw.Close(); err != nil {
			panic(err)
		}
	}(this, pw)

	return this.encrypt(pw.Reader, stream)
}

func (this *Payload) writeMetadata(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(this.getMetadata())
}

func (this *Payload) getMetadata() *Metadata {
	return &Metadata{IV: this.IV, Key: this.Key}
}

func (this *Payload) getStreamAndCreateMetadata() (stream cipher.Stream, err error) {
	var block cipher.Block
	block, this.Key, err = this.Block.CreateBlock(this.Random)
	if err != nil {
		return
	}
	stream, this.IV, err = this.BlockMode.GetCipherStream(block, this.Random)
	return
}

func (this *Payload) computeHash(reader io.Reader) (io.Reader, error) {
	readerHash := make(chan []byte)
	read, err := this.HashMethod.HashOnTheWay(reader, readerHash)
	if err != nil {
		return nil, err
	}
	go func(readerHash chan []byte) {
		this.Hash = <-readerHash
	}(readerHash)

	return read, nil
}

func (this *Payload) encrypt(reader io.Reader, stream cipher.Stream) (io.Reader, error) {
	ret, writer := io.Pipe()
	go func(stream cipher.Stream) {
		cipherWriter := &cipher.StreamWriter{S: stream, W: writer}
		io.Copy(cipherWriter, reader)
		cipherWriter.Close()
	}(stream)
	return this.computeHash(ret)
}
