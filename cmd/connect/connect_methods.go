package connect

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"syscall"
	"time"

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
