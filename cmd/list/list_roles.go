package list

import (
	"fmt"
	"os"

	"github.com/kahnwong/pgconn/color"
	"github.com/kahnwong/pgconn/utils"
	"golang.org/x/exp/slices"

	"github.com/spf13/cobra"
)

func RoleGet(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	var autocomplete []string

	if len(args) == 0 {
		autocomplete = utils.GetAccounts()
	} else if len(args) == 1 {
		autocomplete = utils.GetDatabases(args[0])
	}

	return autocomplete, cobra.ShellCompDirectiveNoFileComp
}

var rolesCmd = &cobra.Command{
	Use:               "roles",
	Short:             "Get a list of roles for a given database",
	ValidArgsFunction: RoleGet,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please specify an account")
			os.Exit(1)
		} else if len(args) == 1 {
			fmt.Println("Please specify a database")
			os.Exit(1)
		} else if len(args) > 2 {
			fmt.Println("`list databases` only requires two argument")
			os.Exit(1)
		}

		isValidAccount := slices.Contains(utils.GetAccounts(), args[0])
		isValidDatabase := slices.Contains(utils.GetDatabases(args[0]), args[1])

		if isValidAccount && isValidDatabase {
			fmt.Printf("%s %s\n", color.Green("Account:"), args[0])
			fmt.Printf("%s %s\n", color.Green("Database:"), args[1])

			fmt.Printf("%s\n", color.Blue("Roles:"))

			for _, v := range utils.GetRoles(args[0], args[1]) {
				fmt.Printf("  - %s\n", v)
			}
		} else {
			fmt.Println("Please specify an available account and database")
			os.Exit(1)
		}
	},
}
