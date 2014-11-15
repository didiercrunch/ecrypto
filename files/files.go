package files

import (
	"errors"
	"io"
	"os"
	"regexp"

	"github.com/didiercrunch/filou/contract"
)

type encryptor func(reader io.Reader, writer io.Writer, acceptedContract *contract.AcceptedContract) error

type PathEncryptor struct {
	EncryptFile      encryptor
	EncryptDirectory encryptor
}

func (this *PathEncryptor) fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func (this *PathEncryptor) getDefaultPathToEncryptFile(originalFileName string) string {
	isFileWithSingleExtension := regexp.MustCompile(`\.\w+$`)
	isFileWithDoubleExtension := regexp.MustCompile(`\.\w+\.\w+$`)
	isHiddenDirectory := regexp.MustCompile(`\/.\w+$`)
	isHiddenFile := regexp.MustCompile(`\/\.\w+\.\w+$`)

	if isHiddenFile.MatchString(originalFileName) {
		return isFileWithSingleExtension.ReplaceAllString(originalFileName, ".filou")
	}
	if isHiddenDirectory.MatchString(originalFileName) {
		return originalFileName + ".filou"
	}
	if isFileWithDoubleExtension.MatchString(originalFileName) {
		return isFileWithDoubleExtension.ReplaceAllString(originalFileName, ".filou")
	}
	if isFileWithSingleExtension.MatchString(originalFileName) {
		return isFileWithSingleExtension.ReplaceAllString(originalFileName, ".filou")
	}
	return originalFileName + ".filou"
}

//  create a file with permission set to 600
func (this *PathEncryptor) createPrivateFile(name string) (*os.File, error) {
	return os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
}

func (this *PathEncryptor) EncryptPathToDefaultFile(path string, acceptedContract *contract.AcceptedContract) error {
	outputPathName := this.getDefaultPathToEncryptFile(path)
	if _, err := os.Stat(outputPathName); !os.IsNotExist(err) {
		return errors.New("the default path to encrypt already exists.")
	}
	return nil
}

func (this *PathEncryptor) EncryptPathToOtherPath(fromPath, toPath string, acceptedContract *contract.AcceptedContract) error {
	if this.fileExists(toPath) {
		return errors.New("output file already exists")
	}
	if f, err := this.createPrivateFile(toPath); err != nil {
		return err
	} else {
		defer f.Close()
		return this.EncryptPathToWriter(fromPath, f, acceptedContract)
	}
}

func (this *PathEncryptor) EncryptPathToWriter(path string, writer io.Writer, acceptedContract *contract.AcceptedContract) error {
	if !this.fileExists(path) {
		return errors.New("input file does not exists")
	}
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
