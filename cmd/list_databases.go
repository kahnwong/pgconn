package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var databasesCmd = &cobra.Command{
	Use:   "databases",
	Short: "Get a list of databases",
	Long:  `Get a list of databases`,
	Run: func(cmd *cobra.Command, args []string) {
		color.Green("Available databases:")
		for _, v := range getDatabases() {
			fmt.Printf("  - %s\n", v)
		}
	},
}

func init() {
	listCmd.AddCommand(databasesCmd)
}
