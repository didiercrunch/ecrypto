package main

import (
	"testing"

	"github.com/codegangsta/cli"
)

func TestCreateKeyCommand(t *testing.T) {
	c := createKeyCommand
	hasBeenCalled := false
	c.Action = func(c *cli.Context) {
		if c.Int("size") != 22 {
			t.Error("wrong size")
		}
		hasBeenCalled = true
	}
	app := cli.NewApp()
	app.Commands = []cli.Command{c}
	args := []string{"file", "createkey", "-size", "22"}
	app.Run(args)
	if !hasBeenCalled {
		t.Error("sub command *createkey* has not been called")
	}
}
