package files

import (
	"fmt"
	"io"
	"testing"

	"github.com/didiercrunch/filou/contract"
	"github.com/didiercrunch/filou/helper"
)

var _ = fmt.Print

type mockEncryptor struct {
	N int
}

func (this *mockEncryptor) Encrypt(reader io.Reader, writer io.Writer, acceptedContract *contract.AcceptedContract) error {
	this.N += 1
	return nil
}

func TestEncryptPathWrongPath(t *testing.T) {
	pe := new(PathEncryptor)

	filePath := "/something_that_does_not_exists"
	if err := pe.EncryptPath(filePath, nil, nil); err == nil {
		t.Error("should say the file does not exists")
	}
}

func TestEncryptPathForDirectory(t *testing.T) {
	dirName := helper.GetTmpEmptyDir()
	fileEncryptor, dirEncryptor := new(mockEncryptor), new(mockEncryptor)
	pe := &PathEncryptor{fileEncryptor.Encrypt, dirEncryptor.Encrypt}

	if err := pe.EncryptPath(dirName, nil, nil); err != nil {
		t.Error(err)
	} else if fileEncryptor.N != 0 {
		t.Error("should have 0", fileEncryptor.N)
	} else if dirEncryptor.N != 1 {
		t.Error("should have 1", dirEncryptor.N)
	}
}

func TestEncryptPathForFile(t *testing.T) {
	file := helper.GetTmpEmptyFile()
	defer func() {
		if err := file.Close(); err != nil {
			t.Error(err)
		}
	}()
	fileEncryptor, dirEncryptor := new(mockEncryptor), new(mockEncryptor)
	pe := &PathEncryptor{fileEncryptor.Encrypt, dirEncryptor.Encrypt}

	if err := pe.EncryptPath(file.Name(), nil, nil); err != nil {
		t.Error(err)
	} else if fileEncryptor.N != 1 {
		t.Error("should have 1", fileEncryptor.N)
	} else if dirEncryptor.N != 0 {
		t.Error("should have 0", dirEncryptor.N)
	}
}
