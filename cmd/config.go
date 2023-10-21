package cmd

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/exp/maps"
	"gopkg.in/yaml.v3"
)

// read raw config
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

// convert to map
var config = createConfigMap(readConfig())

type Connection struct {
	Hostname        string
	ProxyKind       string
	ProxyHost       string
	ProxyTargetPort int
	Username        string
	Password        string
	Dbname          string
}

func createConfigMap(config []Config) map[string]map[string]map[string]Connection {
	configMap := make(map[string]map[string]map[string]Connection)

	for _, a := range config {
		configMap[a.Account] = map[string]map[string]Connection{}

		for _, db := range a.Dbs {
			configMap[a.Account][db.Name] = map[string]Connection{}
			hostname := db.Hostname
			proxyKind := db.Proxy.Kind
			proxyHost := db.Proxy.Host

			for _, role := range db.Roles {
				// create random port for ssh port forwarding
				min := 5432
				max := 8000

				r := rand.New(rand.NewSource(time.Now().UnixNano()))
				port := r.Intn(max-min+1) + min

				// main
				configMap[a.Account][db.Name][role.Username] = Connection{
					Hostname:        hostname,
					ProxyKind:       proxyKind,
					ProxyHost:       proxyHost,
					ProxyTargetPort: port,
					Username:        role.Username,
					Password:        role.Password,
					Dbname:          role.Dbname,
				}

			}
		}
	}

	return configMap
}

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
