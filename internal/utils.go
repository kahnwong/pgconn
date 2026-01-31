package internal

import (
	"log"
	"os"
	"os/exec"
)

func CheckIfBinaryExists(binaryName string) {
	_, err := exec.LookPath(binaryName)
	if err != nil {
		log.Printf("Binary '%s' not found in the PATH\n", binaryName)
		os.Exit(1)
	}
}
