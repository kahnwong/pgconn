package connect

import (
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"os/exec"
	"syscall"
	"time"

	"github.com/kahnwong/pgconn/color"
	"github.com/kahnwong/pgconn/config"
	"github.com/kahnwong/pgconn/utils"
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
	minPort := 5432
	maxPort := 8000

	port := rand.IntN(maxPort-minPort) + minPort

	return port
}

func (c connection) InitProxy() *exec.Cmd {
	var proxyCmd string

	// create cmd
	switch c.ProxyKind {
	case "ssh":
		proxyCmd = fmt.Sprintf("ssh -N -L %d:%s:5432 %s", c.Port, c.Hostname, c.ProxyHost)
	case "cloud-sql-proxy":
		utils.CheckIfBinaryExists("cloud-sql-proxy")
		proxyCmd = fmt.Sprintf("cloud-sql-proxy %s --port %d --quiet", c.ProxyHost, c.ProxyPort)
	}

	// main
	cmd := exec.Command("/bin/sh", "-c", proxyCmd)
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	if err := cmd.Start(); err != nil {
		log.Printf("Failed to start the first process: %v\n", err)
		os.Exit(1)
	}

	time.Sleep(1 * time.Second) // important, so proxy has some time to start up

	return cmd
}

func (c connection) Connect() {
	// set hostname
	var connectHostname string
	if c.ProxyKind != "" {
		switch c.ProxyKind {
		case "ssh":
			connectHostname = "localhost"
		case "cloud-sql-proxy":
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
		log.Printf("Failed to start the second process: %v\n", err)
		os.Exit(1)
	}
}

func (c connection) KillProxyPid() {
	pgid, err := syscall.Getpgid(c.ProxyCmd.Process.Pid)
	if err == nil {
		err = syscall.Kill(-pgid, syscall.SIGKILL)
		if err != nil {
			log.Fatal(err)
		}
	}
}
