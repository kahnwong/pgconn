package cmd

import (
	"os"

	"github.com/kahnwong/pgconn/cmd/connect"
	"github.com/kahnwong/pgconn/cmd/list"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pgconn",
	Short: "pgcli wrapper to connect to PostgreSQL database specified in pgconn.sops.yaml",
	Long:  `pgcli wrapper to connect to PostgreSQL database specified in pgconn.sops.yaml. Proxy/tunnel connection is automatically created and killed when pgcli is exited. `,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(list.Cmd)
	rootCmd.AddCommand(connect.Cmd)
}
