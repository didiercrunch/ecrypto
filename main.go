package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/didiercrunch/filou/keygenerator"
	"os"
)

var createKeyCommand = cli.Command{
	Name:  "createkey",
	Usage: "create a pair of public/private key",
	Flags: []cli.Flag{
		cli.IntFlag{Name: "size, s", Value: 2048, Usage: "key size in bits"},
	},
	Action: func(c *cli.Context) {
		keyGenerator := new(keygenerator.KeyGenerator)
		fmt.Println("creating key")
		if err := keyGenerator.CreateNewKey(c.Int("size")); err != nil {
			fmt.Println("error while creating the key\n", err)
		}
	},
}

var encryptFileCommand = cli.Command{
	Name:  "encrypt",
	Usage: "encrypt file or directory",
	Flags: []cli.Flag{
		cli.StringFlag{Name: "file, s", Value: "", Usage: "the file or directory to encrypt."},
		cli.StringFlag{Name: "publicKey, k", Value: "", Usage: "the public key to use to encrypt"},
	},
	Action: func(c *cli.Context) {
		keyGenerator := new(keygenerator.KeyGenerator)
		fmt.Println("creating key")
		if err := keyGenerator.CreateNewKey(c.Int("size")); err != nil {
			fmt.Println("error while creating the key\n", err)
		}
	},
}

func main() {
	app := cli.NewApp()
	app.Name = "filou"
	app.Usage = "encrypt files"
	app.Commands = []cli.Command{
		createKeyCommand,
		encryptFileCommand,
	}

	app.Run(os.Args)
}
