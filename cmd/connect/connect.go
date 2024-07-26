package connect

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/kahnwong/pgconn/utils"

	"github.com/spf13/cobra"
)

func connectionInfoGet(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	var autocompleteOptions []string

	if len(args) == 0 { // account
		autocompleteOptions = utils.GetAccounts()
	} else if len(args) == 1 { // database
		autocompleteOptions = utils.GetDatabases(args[0])
	} else if len(args) == 2 { // role
		autocompleteOptions = utils.GetRoles(args[0], args[1])
	}

	return autocompleteOptions, cobra.ShellCompDirectiveNoFileComp
}

var Cmd = &cobra.Command{
	Use:               "connect [database] [role]",
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
		connInfo := connMap[args[0]][args[1]][args[2]]
		c := connection{connInfo}

		var proxyCmd *exec.Cmd
		c.Port = c.SetProxyPort()

		if c.ProxyKind != "" {
			proxyCmd = createProxy(c)
		}

		// connect via pgcli
		connectDB(c)

		// clean up proxy PID
		if c.ProxyKind != "" {
			killProxyPid(proxyCmd)
		}
	},
}

func init() {
	checkIfBinaryExists("pgcli")
}