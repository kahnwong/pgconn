package main

import (
	"fmt"
	"os"

	"github.com/mitchellh/cli"
)

// init
var config = readConfig()

// main
func main() {
	c := cli.NewCLI("pgconn", "0.2.0")
	c.Args = os.Args[1:]

	c.Commands = map[string]cli.CommandFactory{
		"connect": func() (cli.Command, error) {
			return &ConnectCommand{}, nil
		},
		"list databases": func() (cli.Command, error) {
			return &ListDatabasesCommand{}, nil
		},
		"list roles": func() (cli.Command, error) {
			return &ListRolesCommand{}, nil
		},
	}

	c.Autocomplete = true

	exitStatus, err := c.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	}
	os.Exit(exitStatus)
}
