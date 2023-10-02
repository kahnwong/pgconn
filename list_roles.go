package main

import (
	"fmt"
	"os"
)

type ListRolesCommand struct{}

func (c *ListRolesCommand) Help() string {
	return "List available roles in a database"
}

func (c *ListRolesCommand) Run(args []string) int {
	if len(args) == 0 {
		fmt.Println("Please specify a database")
		os.Exit(1)
	}

	fmt.Printf("Database: %s\n", args[0])
	fmt.Println("Roles:")
	for _, v := range config {
		if v.Name == args[0] {
			for _, v := range v.Roles {
				fmt.Printf("    %s\n", v.Username)
			}
		}
	}
	return 0
}

func (c *ListRolesCommand) Synopsis() string {
	return "List available roles in a database"
}
