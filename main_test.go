package main

import (
	"github.com/codegangsta/cli"
	"testing"
)

func TestCreateKeyCommand(t *testing.T) {
	c := createKeyCommand
	hasBeenCalled := false
	c.Action = func(c *cli.Context) {
		if c.Int("size") != 22 {
			t.Error("wrong size")
		}
		if c.String("password") != "money" {
			t.Error("wrong password")
		}
		hasBeenCalled = true
	}
	app := cli.NewApp()
	app.Commands = []cli.Command{c}
	args := []string{"file", "createkey", "-size", "22", "-password", "money"}
	app.Run(args)
	if !hasBeenCalled {
		t.Error("sub command *createkey* has not been called")
	}
}
