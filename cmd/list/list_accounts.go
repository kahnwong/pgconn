package list

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/kahnwong/pgconn/internal"
	"github.com/spf13/cobra"
)

var accountsCmd = &cobra.Command{
	Use:   "accounts",
	Short: "Get a list of accounts",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			fmt.Println("`list accounts` does not require an argument")
			os.Exit(1)
		}

		fmt.Printf("%s\n", color.HiGreenString("Available accounts:"))
		for _, v := range internal.GetAccounts() {
			fmt.Printf("  - %s\n", v)
		}
	},
}
