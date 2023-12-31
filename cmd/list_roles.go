package cmd

import (
	"fmt"
	"os"

	"golang.org/x/exp/slices"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func RoleGet(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	var autocomplete []string

	if len(args) == 0 {
		autocomplete = getAccounts()
	} else if len(args) == 1 {
		autocomplete = getDatabases(args[0])
	}

	return autocomplete, cobra.ShellCompDirectiveNoFileComp
}

var rolesCmd = &cobra.Command{
	Use:               "roles",
	Short:             "Get a list of roles for a given database",
	Long:              `Get a list of roles for a given database`,
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

		isValidAccount := slices.Contains(getAccounts(), args[0])
		isValidDatabase := slices.Contains(getDatabases(args[0]), args[1])

		if isValidAccount && isValidDatabase {
			green := color.New(color.FgGreen).SprintFunc()

			fmt.Printf("%s %s\n", green("Account:"), args[0])
			fmt.Printf("%s %s\n", green("Database:"), args[1])

			color.Blue("Roles:")

			for _, v := range getRoles(args[0], args[1]) {
				fmt.Printf("  - %s\n", v)
			}
		} else {
			fmt.Println("Please specify an available account and database")
			os.Exit(1)
		}
	},
}

func init() {
	listCmd.AddCommand(rolesCmd)
}
