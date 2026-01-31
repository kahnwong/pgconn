package list

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/kahnwong/pgconn/internal"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"
)

func AccountGet(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	var autocomplete []string

	if len(args) == 0 {
		autocomplete = internal.GetAccounts()
	}

	return autocomplete, cobra.ShellCompDirectiveNoFileComp
}

var databasesCmd = &cobra.Command{
	Use:               "databases [account]",
	Short:             "Get a list of databases for a given account",
	ValidArgsFunction: AccountGet,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please specify an account")
			os.Exit(1)
		} else if len(args) > 1 {
			fmt.Println("`list databases` only requires one argument")
			os.Exit(1)
		}

		isValidAccount := slices.Contains(internal.GetAccounts(), args[0])
		if isValidAccount {
			fmt.Printf("%s %s\n", color.HiGreenString("Account:"), args[0])
			fmt.Printf("%s\n", color.BlueString("Databases:"))

			for _, v := range internal.GetDatabases(args[0]) {
				fmt.Printf("  - %s\n", v)
			}
		} else {
			fmt.Println("Please specify an available account")
			os.Exit(1)
		}
	},
}
