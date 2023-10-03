package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var databasesCmd = &cobra.Command{
	Use:   "databases",
	Short: "Get a list of databases",
	Long:  `Get a list of databases`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Available databases:")
		for _, v := range config {
			fmt.Printf("    %s\n", v.Name)
		}
	},
}

func init() {
	listCmd.AddCommand(databasesCmd)
}
