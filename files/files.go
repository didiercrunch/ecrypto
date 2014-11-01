package files

import (
	"fmt"
	"io"
	"os"

	"github.com/didiercrunch/filou/contract"
)

var _ = fmt.Print

type encryptor func(reader io.Reader, writer io.Writer, acceptedContract *contract.AcceptedContract) error

type PathEncryptor struct {
	EncryptFile      encryptor
	EncryptDirectory encryptor
}

func (this *PathEncryptor) EncryptPath(path string, writer io.Writer, acceptedContract *contract.AcceptedContract) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}

	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	if fd, err := f.Stat(); err != nil {
		return err
	} else if fd.IsDir() {
		return this.EncryptDirectory(f, writer, acceptedContract)
	} else {
		return this.EncryptFile(f, writer, acceptedContract)
	}
}
