package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get a list of accounts, databases or roles",
	Long:  `Get a list of accounts, databases or roles`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Please specify a command: [accounts | databases | roles]")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
