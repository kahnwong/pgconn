package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rolesCmd = &cobra.Command{
	Use:   "roles",
	Short: "Get a list of roles for a given database",
	Long:  `Get a list of roles for a given database`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please specify a database")
			os.Exit(1)
		}

		fmt.Printf("Database: %s\n", args[0])
		fmt.Println("Roles:")
		for _, v := range config {
			if v.Name == args[0] {
				for _, v := range v.Roles {
					fmt.Printf("    %s\n", v.Username)
				}
			}
		}
	},
}

func init() {
	listCmd.AddCommand(rolesCmd)
}
