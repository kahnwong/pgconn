package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get a list of databases or roles for a given database",
	Long:  `Get a list of databases or roles for a given database`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Please specify a command: [databases | roles]")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
