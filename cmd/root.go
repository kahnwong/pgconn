package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "pgconn",
	Version: "0.4.2",
	Short:   "pgcli wrapper to connect to PostgreSQL database specified in db.yaml",
	Long:    `pgcli wrapper to connect to PostgreSQL database specified in db.yaml. Proxy/tunnel connection is automatically created and killed when pgcli is exited. `,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
