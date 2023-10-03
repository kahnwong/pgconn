package cmd

import (
	"fmt"
	"os"

	"golang.org/x/exp/slices"

	"github.com/spf13/cobra"
)

func DatabaseGet(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return getDatabases(), cobra.ShellCompDirectiveNoFileComp
}

var rolesCmd = &cobra.Command{
	Use:               "roles",
	Short:             "Get a list of roles for a given database",
	Long:              `Get a list of roles for a given database`,
	ValidArgsFunction: DatabaseGet,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please specify a database")
			os.Exit(1)
		}

		isValidDatabase := slices.Contains(getDatabases(), args[0]) // true

		if isValidDatabase {
			fmt.Printf("Database: %s\n", args[0])
			fmt.Println("Roles:")

			for _, v := range getRoles(args[0]) {
				fmt.Printf("    %s\n", v)
			}

		} else {
			fmt.Println("Please specify an available database")
			os.Exit(1)
		}
	},
}

func init() {
	listCmd.AddCommand(rolesCmd)
}
