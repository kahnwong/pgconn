package cmd

import "golang.org/x/exp/maps"

func getAccounts() []string {
	accounts := maps.Keys(connMap)

	return accounts
}

func getDatabases(account string) []string {
	databases := maps.Keys(connMap[account])

	return databases
}

func getRoles(account string, database string) []string {
	roles := maps.Keys(connMap[account][database])

	return roles
}
