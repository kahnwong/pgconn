package internal

import (
	"testing"
)

func TestGetAccountsWithMockData(t *testing.T) {
	// Save original ConnMap and restore after test
	originalConnMap := ConnMap
	defer func() { ConnMap = originalConnMap }()

	// Set up test data
	ConnMap = map[string]map[string]map[string]Connection{
		"account1": {
			"db1": {
				"user1": Connection{Username: "user1"},
			},
		},
		"account2": {
			"db2": {
				"user2": Connection{Username: "user2"},
			},
		},
	}

	accounts := GetAccounts()

	// Check that we got 2 accounts
	if len(accounts) != 2 {
		t.Errorf("Expected 2 accounts, got %d", len(accounts))
	}

	// Check that both accounts are present (order doesn't matter)
	accountMap := make(map[string]bool)
	for _, acc := range accounts {
		accountMap[acc] = true
	}

	if !accountMap["account1"] {
		t.Error("Expected 'account1' in accounts list")
	}
	if !accountMap["account2"] {
		t.Error("Expected 'account2' in accounts list")
	}
}

func TestGetDatabasesWithMockData(t *testing.T) {
	// Save original ConnMap and restore after test
	originalConnMap := ConnMap
	defer func() { ConnMap = originalConnMap }()

	// Set up test data
	ConnMap = map[string]map[string]map[string]Connection{
		"test-account": {
			"db1": {
				"user1": Connection{Username: "user1"},
			},
			"db2": {
				"user2": Connection{Username: "user2"},
			},
		},
	}

	databases := GetDatabases("test-account")

	// Check that we got 2 databases
	if len(databases) != 2 {
		t.Errorf("Expected 2 databases, got %d", len(databases))
	}

	// Check that both databases are present
	dbMap := make(map[string]bool)
	for _, db := range databases {
		dbMap[db] = true
	}

	if !dbMap["db1"] {
		t.Error("Expected 'db1' in databases list")
	}
	if !dbMap["db2"] {
		t.Error("Expected 'db2' in databases list")
	}
}

func TestGetRolesWithMockData(t *testing.T) {
	// Save original ConnMap and restore after test
	originalConnMap := ConnMap
	defer func() { ConnMap = originalConnMap }()

	// Set up test data
	ConnMap = map[string]map[string]map[string]Connection{
		"test-account": {
			"test-db": {
				"admin":    Connection{Username: "admin"},
				"readonly": Connection{Username: "readonly"},
				"writer":   Connection{Username: "writer"},
			},
		},
	}

	roles := GetRoles("test-account", "test-db")

	// Check that we got 3 roles
	if len(roles) != 3 {
		t.Errorf("Expected 3 roles, got %d", len(roles))
	}

	// Check that all roles are present
	roleMap := make(map[string]bool)
	for _, role := range roles {
		roleMap[role] = true
	}

	if !roleMap["admin"] {
		t.Error("Expected 'admin' in roles list")
	}
	if !roleMap["readonly"] {
		t.Error("Expected 'readonly' in roles list")
	}
	if !roleMap["writer"] {
		t.Error("Expected 'writer' in roles list")
	}
}
