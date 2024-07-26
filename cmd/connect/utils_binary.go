package connect

import (
	"fmt"
	"os"
	"os/exec"
)

func checkIfBinaryExists(binaryName string) {
	_, err := exec.LookPath(binaryName)
	if err != nil {
		fmt.Printf("Binary '%s' not found in the PATH\n", binaryName)
		os.Exit(1)
	}
}
