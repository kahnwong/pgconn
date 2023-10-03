package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rolesCmd = &cobra.Command{
	Use:   "roles",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
