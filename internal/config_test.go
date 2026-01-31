package internal

import "testing"

func TestCreateConnMap(t *testing.T) {
	// Create a simple test config
	config := &Config{
		Pgconn: []struct {
			Account string `yaml:"account"`
			Dbs     []struct {
				Name     string `yaml:"name"`
				Hostname string `yaml:"hostname"`
				Port     int    `yaml:"port"`
				Proxy    struct {
					Kind string `yaml:"kind"`
					Host string `yaml:"host"`
				} `yaml:"proxy"`
				Roles []struct {
					Username string `yaml:"username"`
					Password string `yaml:"password"`
					Dbname   string `yaml:"dbname"`
				} `yaml:"roles"`
			} `yaml:"dbs"`
		}{
			{
				Account: "test-account",
				Dbs: []struct {
					Name     string `yaml:"name"`
					Hostname string `yaml:"hostname"`
					Port     int    `yaml:"port"`
					Proxy    struct {
						Kind string `yaml:"kind"`
						Host string `yaml:"host"`
					} `yaml:"proxy"`
					Roles []struct {
						Username string `yaml:"username"`
						Password string `yaml:"password"`
						Dbname   string `yaml:"dbname"`
					} `yaml:"roles"`
				}{
					{
						Name:     "test-db",
						Hostname: "localhost",
						Port:     5432,
						Proxy: struct {
							Kind string `yaml:"kind"`
							Host string `yaml:"host"`
						}{
							Kind: "ssh",
							Host: "proxy.example.com",
						},
						Roles: []struct {
							Username string `yaml:"username"`
							Password string `yaml:"password"`
							Dbname   string `yaml:"dbname"`
						}{
							{
								Username: "admin",
								Password: "secret",
								Dbname:   "mydb",
							},
						},
					},
				},
			},
		},
	}

	result := createConnMap(config)

	// Test that the map structure is correct
	if result == nil {
		t.Fatal("createConnMap() returned nil")
	}

	// Test account exists
	if _, ok := result["test-account"]; !ok {
		t.Error("Expected account 'test-account' not found in result")
	}

	// Test database exists
	if _, ok := result["test-account"]["test-db"]; !ok {
		t.Error("Expected database 'test-db' not found in result")
	}

	// Test role exists
	conn, ok := result["test-account"]["test-db"]["admin"]
	if !ok {
		t.Fatal("Expected role 'admin' not found in result")
	}

	// Test connection values
	if conn.Hostname != "localhost" {
		t.Errorf("Expected Hostname 'localhost', got '%s'", conn.Hostname)
	}
	if conn.Port != 5432 {
		t.Errorf("Expected Port 5432, got %d", conn.Port)
	}
	if conn.Username != "admin" {
		t.Errorf("Expected Username 'admin', got '%s'", conn.Username)
	}
	if conn.Password != "secret" {
		t.Errorf("Expected Password 'secret', got '%s'", conn.Password)
	}
	if conn.Dbname != "mydb" {
		t.Errorf("Expected Dbname 'mydb', got '%s'", conn.Dbname)
	}
	if conn.ProxyKind != "ssh" {
		t.Errorf("Expected ProxyKind 'ssh', got '%s'", conn.ProxyKind)
	}
	if conn.ProxyHost != "proxy.example.com" {
		t.Errorf("Expected ProxyHost 'proxy.example.com', got '%s'", conn.ProxyHost)
	}
}

func TestCreateConnMapEmpty(t *testing.T) {
	config := &Config{}
	result := createConnMap(config)

	if result == nil {
		t.Fatal("createConnMap() returned nil for empty config")
	}

	if len(result) != 0 {
		t.Errorf("Expected empty map, got map with %d entries", len(result))
	}
}
