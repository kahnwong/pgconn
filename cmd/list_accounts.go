package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var accountsCmd = &cobra.Command{
	Use:   "accounts",
	Short: "Get a list of accounts",
	Long:  `Get a list of accounts`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			fmt.Println("`list accounts` does not require an argument")
			os.Exit(1)
		}
		color.Green("Available accounts:")
		for _, v := range getAccounts() {
			fmt.Printf("  - %s\n", v)
		}
	},
}

func init() {
	listCmd.AddCommand(accountsCmd)
}
