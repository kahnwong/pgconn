package connect

import (
	"log"
	"os/exec"
	"syscall"
)

func killProxyPid(cmd *exec.Cmd) {
	pgid, err := syscall.Getpgid(cmd.Process.Pid)
	if err == nil {
		err = syscall.Kill(-pgid, syscall.SIGKILL)
		if err != nil {
			log.Fatal(err)
		}
	}
}
