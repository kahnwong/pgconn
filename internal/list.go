package internal

import "golang.org/x/exp/maps"

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
