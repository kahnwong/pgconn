package utils

import (
	"github.com/kahnwong/pgconn/config"
	"golang.org/x/exp/maps"
)

var connMap = config.ConnMap

func GetAccounts() []string {
	accounts := maps.Keys(connMap)

	return accounts
}

func GetDatabases(account string) []string {
	databases := maps.Keys(connMap[account])

	return databases
}

func GetRoles(account string, database string) []string {
	roles := maps.Keys(connMap[account][database])

	return roles
}
