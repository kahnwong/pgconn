package main

import (
	"fmt"
)

type ListDatabasesCommand struct{}

func (c *ListDatabasesCommand) Help() string {
	return "List available databases"
}

func (c *ListDatabasesCommand) Run(args []string) int {
	fmt.Println("Available databases:")
	for _, v := range config {
		fmt.Printf("    %s\n", v.Name)
	}
	return 0
}

func (c *ListDatabasesCommand) Synopsis() string {
	return "List available databases"
}
