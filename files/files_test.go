package files

import (
	"fmt"
	"io"
	"os"
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

func removeFile(name string, t *testing.T) {
	if err := os.Remove(name); err != nil {
		t.Error(err)
	}
}
func TestEncryptPathWrongPath(t *testing.T) {
	pe := new(PathEncryptor)

	filePath := "/something_that_does_not_exists"
	if err := pe.EncryptPathToWriter(filePath, nil, nil); err == nil {
		t.Error("should say the file does not exists")
	}
}

func TestEncryptPathForDirectory(t *testing.T) {
	dirName := helper.GetTmpEmptyDir()
	fileEncryptor, dirEncryptor := new(mockEncryptor), new(mockEncryptor)
	pe := &PathEncryptor{fileEncryptor.Encrypt, dirEncryptor.Encrypt}

	if err := pe.EncryptPathToWriter(dirName, nil, nil); err != nil {
		t.Error(err)
	} else if fileEncryptor.N != 0 {
		t.Error("should have 0", fileEncryptor.N)
	} else if dirEncryptor.N != 1 {
		t.Error("should have 1", dirEncryptor.N)
	}
}

func TestEncryptPathToWriter(t *testing.T) {
	file := helper.GetTmpFileWithText("foo bar")
	defer removeFile(file.Name(), t)
	fileEncryptor, dirEncryptor := new(mockEncryptor), new(mockEncryptor)
	pe := &PathEncryptor{fileEncryptor.Encrypt, dirEncryptor.Encrypt}

	if err := pe.EncryptPathToWriter(file.Name(), nil, nil); err != nil {
		t.Error(err)
	} else if fileEncryptor.N != 1 {
		t.Error("should have 1", fileEncryptor.N)
	} else if dirEncryptor.N != 0 {
		t.Error("should have 0", dirEncryptor.N)
	}
}

func TestGetDefaultPathToEncryptFile(t *testing.T) {
	cases := map[string]string{
		"/foo/bar.txt":    "/foo/bar.filou",  // regular file
		"/foo/bar.tar.gz": "/foo/bar.filou",  // file with two extensions
		"/foo/bar":        "/foo/bar.filou",  // directory
		"/foo/.bar":       "/foo/.bar.filou", // hidden directory
		"/foo/.bar.txt":   "/foo/.bar.filou", // hidden file
	}
	pe := new(PathEncryptor)
	for input, expectedOutput := range cases {
		if output := pe.getDefaultPathToEncryptFile(input); output != expectedOutput {
			t.Error("expected", expectedOutput, "but received", output)
		}
	}
}

func TestEncryptPathToOtherPath(t *testing.T) {
	fromFile := helper.GetTmpFileWithText("foo bar")
	defer removeFile(fromFile.Name(), t)

	toFile := helper.GetTmpFileName()
	defer removeFile(toFile, t)

	fileEncryptor, dirEncryptor := new(mockEncryptor), new(mockEncryptor)
	pe := &PathEncryptor{fileEncryptor.Encrypt, dirEncryptor.Encrypt}

	if err := pe.EncryptPathToOtherPath(fromFile.Name(), toFile, nil); err != nil {
		t.Error(err)
	} else if fileEncryptor.N != 1 {
		t.Error("should have 1", fileEncryptor.N)
	} else if dirEncryptor.N != 0 {
		t.Error("should have 0", dirEncryptor.N)
	}
}
