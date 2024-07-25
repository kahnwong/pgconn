package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/fatih/color"
	"github.com/kahnwong/pgconn/config"
)

func ConnectDB(c config.Connection) *exec.Cmd {
	// check if pgcli exists
	binaryName := "pgcli"
	_, err := exec.LookPath(binaryName)
	if err != nil {
		fmt.Printf("Binary '%s' not found in the PATH\n", binaryName)
		os.Exit(1)
	}

	// set hostname
	var connectHostname string
	if c.ProxyKind != "" {
		if c.ProxyKind == "ssh" {
			connectHostname = "localhost"
		} else if c.ProxyKind == "cloud-sql-proxy" {
			connectHostname = "127.0.0.1"
		}
	} else {
		connectHostname = c.Hostname
	}

	// print port
	green := color.New(color.FgGreen).SprintFunc()
	fmt.Printf("Port: %s\n", green(c.Port))

	// connect
	connectionString := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable", c.Username, c.Password, connectHostname, c.Port, c.Dbname)
	cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("pgcli \"%s\"", connectionString))
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("Failed to start the second process: %v\n", err)
		os.Exit(1)
	}

	return cmd
}
