package cmd

import (
	"fmt"
	"os"

	"golang.org/x/exp/slices"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func AccountGet(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	var autocomplete []string

	if len(args) == 0 {
		autocomplete = getAccounts()
	}

	return autocomplete, cobra.ShellCompDirectiveNoFileComp
}

var databasesCmd = &cobra.Command{
	Use:               "databases",
	Short:             "Get a list of databases for a given account",
	Long:              `Get a list of databases for a given account`,
	ValidArgsFunction: AccountGet,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please specify an account")
			os.Exit(1)
		} else if len(args) > 1 {
			fmt.Println("`list databases` only requires one argument")
			os.Exit(1)
		}

		isValidAccount := slices.Contains(getAccounts(), args[0])
		if isValidAccount {
			green := color.New(color.FgGreen).SprintFunc()

			fmt.Printf("%s %s\n", green("Account:"), args[0])
			color.Blue("Databases:")

			for _, v := range getDatabases(args[0]) {
				fmt.Printf("  - %s\n", v)
			}
		} else {
			fmt.Println("Please specify an available account")
			os.Exit(1)
		}
	},
}

func init() {
	listCmd.AddCommand(databasesCmd)
}
