package cmd

import "golang.org/x/exp/maps"

func getAccounts() []string {
	accounts := maps.Keys(config)

	return accounts
}

func getDatabases(account string) []string {
	databases := maps.Keys(config[account])

	return databases
}

func getRoles(account string, database string) []string {
	roles := maps.Keys(config[account][database])

	return roles
}
