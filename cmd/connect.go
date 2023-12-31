package cmd

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"syscall"
	"time"

	"github.com/spf13/cobra"
)

func connectionInfoGet(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	var autocompleteOptions []string

	if len(args) == 0 { // account
		autocompleteOptions = getAccounts()
	} else if len(args) == 1 { // database
		autocompleteOptions = getDatabases(args[0])
	} else if len(args) == 2 { // role
		autocompleteOptions = getRoles(args[0], args[1])
	}

	return autocompleteOptions, cobra.ShellCompDirectiveNoFileComp
}

var connectCmd = &cobra.Command{
	Use:               "connect [database] [role]",
	Short:             "Connect to a database with specified role",
	Long:              `Connect to a database with specified role`,
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

		// get db config
		connInfo := config[args[0]][args[1]][args[2]]

		// start proxy process if necessary
		var proxyCmd *exec.Cmd
		if connInfo.ProxyKind != "" {
			var port int
			proxyCmd, port = CreateProxy(connInfo)
			connInfo.Port = port
		}

		// connect via pgcli
		ConnectDB(connInfo)

		// clean up proxy PID
		if connInfo.ProxyKind != "" {
			cleanup(proxyCmd)
		}
	},
}

func init() {
	rootCmd.AddCommand(connectCmd)
}

type ConnectCommand struct{}

func (c *ConnectCommand) Help() string {
	return "Connect to a database"
}

func (c *ConnectCommand) Synopsis() string {
	return "Connect to a database"
}

// functions
func CreateProxy(c Connection) (*exec.Cmd, int) {
	// set port
	// random port for ssh port forwarding
	min := 5432
	max := 8000

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	port := r.Intn(max-min+1) + min

	// create cmd
	var proxyCmd string
	if c.ProxyKind == "ssh" {
		proxyCmd = fmt.Sprintf("ssh -N -L %d:%s:5432 %s", port, c.Hostname, c.ProxyHost)
	} else if c.ProxyKind == "cloud-sql-proxy" {
		// check if cloud-sql-proxy exists
		binaryName := "cloud-sql-proxy"
		_, err := exec.LookPath(binaryName)
		if err != nil {
			fmt.Printf("Binary '%s' not found in the PATH\n", binaryName)
			os.Exit(1)
		}

		proxyCmd = fmt.Sprintf("cloud-sql-proxy %s --port %d --quiet", c.ProxyHost, port)
	}

	// main
	cmd := exec.Command("/bin/sh", "-c", proxyCmd)
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	if err := cmd.Start(); err != nil {
		fmt.Printf("Failed to start the first process: %v\n", err)
		os.Exit(1)
	}

	time.Sleep(1 * time.Second) // important, so proxy has some time to start up

	return cmd, port
}

func ConnectDB(c Connection) *exec.Cmd {
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

func cleanup(cmd *exec.Cmd) {
	pgid, err := syscall.Getpgid(cmd.Process.Pid)
	if err == nil {
		err = syscall.Kill(-pgid, syscall.SIGKILL)
		if err != nil {
			log.Fatal(err)
		}
	}
}
