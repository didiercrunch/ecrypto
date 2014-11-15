package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/didiercrunch/filou/contract"
)

var createKeyCommand = cli.Command{
	Name:  "create-contracts",
	Usage: "create a pair of public/private key encapsulated in contracts",
	Flags: []cli.Flag{
		cli.IntFlag{Name: "size, s", Value: 2048, Usage: "key size in bits"},
	},
	Action: func(c *cli.Context) {
		keyGenerator := new(contract.ContractsGenerator)
		fmt.Println("creating key ...")
		if err := keyGenerator.CreateContracts(c.Int("size")); err != nil {
			fmt.Println("error while creating the key\n", err)
		}
	},
}

var encryptFileCommand = cli.Command{
	Name:  "encrypt",
	Usage: "encrypt file or directory",
	Flags: []cli.Flag{
		cli.StringFlag{Name: "file, s", Value: "", Usage: "the file or directory to encrypt."},
		cli.StringFlag{Name: "publicContract, c", Value: "", Usage: "the public contract to use to encrypt"},
	},
	Action: func(c *cli.Context) {
		if _, err := GetContract(c.String("publicContract")); err != nil {
			fmt.Println("error getting the public contract", err)
			return
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
