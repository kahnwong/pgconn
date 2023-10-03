package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
	"time"

	"github.com/spf13/cobra"
)

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please specify a database and role")
			os.Exit(1)
		} else if len(args) == 1 {
			fmt.Println("Please specify a role")
			os.Exit(1)
		}

		// get db config
		dbConfig := getConnectionInfo(args[0], args[1])

		// start proxy process if necessary
		var proxyCmd *exec.Cmd
		if dbConfig.ProxyKind != "" {
			proxyCmd = createProxy(dbConfig.Hostname, dbConfig.ProxyKind, dbConfig.ProxyHost)
		}

		// connect via pgcli
		var connectHostname string
		if dbConfig.ProxyKind != "" {
			if dbConfig.ProxyKind == "ssh" {
				connectHostname = "localhost"
			} else if dbConfig.ProxyKind == "cloud-sql-proxy" {
				connectHostname = "127.0.0.1"
			}
		} else {
			connectHostname = dbConfig.Hostname
		}
		dbCmd := connectDb(connectHostname, dbConfig.Username, dbConfig.Password, dbConfig.Dbname)

		time.Sleep(1 * time.Second) // important, so proxy has some time to start up

		if err := dbCmd.Run(); err != nil {
			fmt.Printf("Failed to start the second process: %v\n", err)
			os.Exit(1)
		}

		// clean up proxy PID
		if dbConfig.ProxyKind != "" {
			pgid, err := syscall.Getpgid(proxyCmd.Process.Pid)
			if err == nil {
				err = syscall.Kill(-pgid, syscall.SIGKILL)
				if err != nil {
					log.Fatal(err)
				}
			}
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

type Connection struct {
	Hostname  string
	Username  string
	Password  string
	Dbname    string
	ProxyKind string
	ProxyHost string
}

func getConnectionInfo(name string, role string) Connection {
	var dbConfig Connection

	for _, db := range config {
		if db.Name == name {
			dbConfig.Hostname = db.Hostname
			dbConfig.ProxyKind = db.Proxy.Kind
			dbConfig.ProxyHost = db.Proxy.Host

			for _, dbRole := range db.Roles {
				if dbRole.Username == role {
					dbConfig.Username = dbRole.Username
					dbConfig.Password = dbRole.Password
					dbConfig.Dbname = dbRole.Dbname
				}
			}
		}
	}

	return dbConfig
}

func createProxy(hostname string, proxyKind string, proxyHost string) *exec.Cmd {
	var proxyCmd string
	if proxyKind == "ssh" {
		proxyCmd = fmt.Sprintf("ssh -N -L 5432:%s:5432 %s", hostname, proxyHost)
	} else if proxyKind == "cloud-sql-proxy" {
		proxyCmd = fmt.Sprintf("cloud-sql-proxy %s --quiet", proxyHost)
	}

	cmd := exec.Command("/bin/sh", "-c", proxyCmd)
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	if err := cmd.Start(); err != nil {
		fmt.Printf("Failed to start the first process: %v\n", err)
		os.Exit(1)
	}

	return cmd
}

func connectDb(hostname string, username string, password string, dbname string) *exec.Cmd {
	connectionString := fmt.Sprintf("postgresql://%s:%s@%s:5432/%s?sslmode=disable", username, password, hostname, dbname)
	cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("pgcli \"%s\"", connectionString))
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd
}
