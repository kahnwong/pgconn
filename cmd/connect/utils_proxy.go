package connect

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"syscall"
	"time"

	"github.com/kahnwong/pgconn/config"
)

func createProxy(c config.Connection) (*exec.Cmd, int) {
	// set port
	// random port for ssh port forwarding
	minPort := 5432
	maxPort := 8000

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	port := r.Intn(maxPort-minPort+1) + minPort

	// create cmd
	var proxyCmd string
	if c.ProxyKind == "ssh" {
		proxyCmd = fmt.Sprintf("ssh -N -L %d:%s:5432 %s", port, c.Hostname, c.ProxyHost)
	} else if c.ProxyKind == "cloud-sql-proxy" {
		// check if cloud-sql-proxy exists
		checkIfBinaryExists("cloud-sql-proxy")

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

func killProxyPid(cmd *exec.Cmd) {
	pgid, err := syscall.Getpgid(cmd.Process.Pid)
	if err == nil {
		err = syscall.Kill(-pgid, syscall.SIGKILL)
		if err != nil {
			log.Fatal(err)
		}
	}
}
