package cmd

import (
	"fmt"
	"os"

	"github.com/kahnwong/pgconn/internal"
	"github.com/spf13/cobra"
)

func connectionInfoGet(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	var autocompleteOptions []string

	if len(args) == 0 { // account
		autocompleteOptions = internal.GetAccounts()
	} else if len(args) == 1 { // database
		autocompleteOptions = internal.GetDatabases(args[0])
	} else if len(args) == 2 { // role
		autocompleteOptions = internal.GetRoles(args[0], args[1])
	}

	return autocompleteOptions, cobra.ShellCompDirectiveNoFileComp
}

var connectCmd = &cobra.Command{
	Use:               "connect [account] [database] [role]",
	Short:             "Connect to a database with specified role",
	ValidArgsFunction: connectionInfoGet,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please specify an account, database and role")
			os.Exit(1)
		} else if len(args) == 1 {
			fmt.Println("Please specify a database and role")
			os.Exit(1)
		} else if len(args) == 2 {
			fmt.Println("Please specify a role")
			os.Exit(1)
		} else if len(args) > 3 {
			fmt.Println("`connect` only requires three argument")
			os.Exit(1)
		}

		// main
		connInfo := internal.ConnMap[args[0]][args[1]][args[2]]
		c := internal.Pgconn{Connection: connInfo}

		if c.ProxyKind != "" {
			c.ProxyPort = c.SetProxyPort()
			c.ProxyCmd = c.InitProxy()
			c.Connect()
			c.KillProxyPid()
		} else {
			c.ProxyPort = c.Port
			c.Connect()
		}
	},
}

func init() {
	rootCmd.AddCommand(connectCmd)
}
