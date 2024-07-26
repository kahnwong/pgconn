package connect

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"syscall"
	"time"

	"github.com/kahnwong/pgconn/color"

	"github.com/kahnwong/pgconn/config"
)

var connMap = config.ConnMap

type connection struct {
	config.Connection
	ProxyPort int
	ProxyCmd  *exec.Cmd
}

func (c connection) SetProxyPort() int {
	// prevent port conflict in case
	// simultaneously connecting to proxied db
	if c.ProxyKind != "" {
		minPort := 5432
		maxPort := 8000

		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		port := r.Intn(maxPort-minPort+1) + minPort

		return port
	} else {
		return c.Port
	}
}

func (c connection) InitProxy() *exec.Cmd {
	var proxyCmd string

	if c.ProxyKind != "" {
		// create cmd
		if c.ProxyKind == "ssh" {
			proxyCmd = fmt.Sprintf("ssh -N -L %d:%s:5432 %s", c.Port, c.Hostname, c.ProxyHost)
		} else if c.ProxyKind == "cloud-sql-proxy" {
			checkIfBinaryExists("cloud-sql-proxy")
			proxyCmd = fmt.Sprintf("cloud-sql-proxy %s --port %d --quiet", c.ProxyHost, c.ProxyPort)
		}

		// main
		cmd := exec.Command("/bin/sh", "-c", proxyCmd)
		cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
		if err := cmd.Start(); err != nil {
			fmt.Printf("Failed to start the first process: %v\n", err)
			os.Exit(1)
		}

		time.Sleep(1 * time.Second) // important, so proxy has some time to start up

		return cmd
	} else {
		return nil
	}
}

func (c connection) Connect() {
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
	fmt.Printf("Port: %s\n", color.Green(c.ProxyPort))

	// connect
	connectionString := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable", c.Username, c.Password, connectHostname, c.ProxyPort, c.Dbname)
	cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("pgcli \"%s\"", connectionString))
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("Failed to start the second process: %v\n", err)
		os.Exit(1)
	}
}
