package shared

import (
	"errors"
	"fmt"
	"os"
	"path"
)

const VERSION = "0.0.1"

var filouDir string

func ensureEcryptoDirectoryIsOkay(dir string) error {
	stat, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return os.Mkdir(dir, DIRECTORY_PERMISSION)
	} else if err != nil {
		return err
	} else if !stat.IsDir() {
		return errors.New(dir + " is not a directory")
	} else if perm := stat.Mode().Perm(); perm != DIRECTORY_PERMISSION {
		msg := fmt.Sprintf("root folder has bad file permission.  should be set to %v  but vas %v", DIRECTORY_PERMISSION, perm)
		return errors.New(msg)
	}
	return nil

}

func init() {
	filouDir = path.Join(os.Getenv("HOME"), ".filou")
	if err := ensureEcryptoDirectoryIsOkay(filouDir); err != nil {
		fmt.Println("cannot work with eCrypto directory", err)
		os.Exit(1)

	}
}

func GetEcryptoDir() string {
	return filouDir
}
