package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// init
var config = readConfig()

// readConfig
type Config struct {
	Account string `yaml:"account"`
	Dbs     []struct {
		Name     string `yaml:"name"`
		Hostname string `yaml:"hostname"`
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
}

func readConfig() []Config {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	filename := filepath.Join(homeDir, ".config", "pgconn", "db.yaml")

	// Check if the file exists
	_, err = os.Stat(filename)

	if os.IsNotExist(err) {
		fmt.Printf("File %s does not exist.\n", filename)
		os.Exit(1)
	}

	var configs []Config

	source, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("failed reading config file: %v\n", err)
	}

	err = yaml.Unmarshal(source, &configs)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	return configs
}

func getAccounts() []string {
	accounts := make([]string, 0)
	for _, v := range config {
		accounts = append(accounts, v.Account)
	}

	return accounts
}

func getDatabases(account string) []string {
	databases := make([]string, 0)
	for _, v := range config {
		if v.Account == account {
			for _, v := range v.Dbs {
				databases = append(databases, v.Name)
			}
		}
	}

	return databases
}

func getRoles(account string, database string) []string {
	roles := make([]string, 0)
	for _, v := range config {
		if v.Account == account {
			for _, v := range v.Dbs {
				if v.Name == database {
					for _, v := range v.Roles {
						roles = append(roles, v.Username)
					}
				}
			}
		}
	}

	return roles
}
