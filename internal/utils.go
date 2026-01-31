package internal

import (
	"log"
	"os"
	"os/exec"

	"golang.org/x/exp/maps"
)

func GetAccounts() []string {
	accounts := maps.Keys(ConnMap)

	return accounts
}

func GetDatabases(account string) []string {
	databases := maps.Keys(ConnMap[account])

	return databases
}

func GetRoles(account string, database string) []string {
	roles := maps.Keys(ConnMap[account][database])

	return roles
}

func CheckIfBinaryExists(binaryName string) {
	_, err := exec.LookPath(binaryName)
	if err != nil {
		log.Printf("Binary '%s' not found in the PATH\n", binaryName)
		os.Exit(1)
	}
}
