package internal

import (
	"log"

	cliBase "github.com/kahnwong/cli-base-sops"
)

// config
var ConnMap = initConnMap()

func initConnMap() map[string]map[string]map[string]Connection {
	config, err := cliBase.ReadYamlSops[Config]("~/.config/pgconn/pgconn.sops.yaml")
	if err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}
	return createConnMap(config)
}

// for cli
type Connection struct {
	Hostname  string
	ProxyKind string
	ProxyHost string
	Port      int
	Username  string
	Password  string
	Dbname    string
}

// raw
type Config struct {
	Pgconn []struct {
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
	} `yaml:"pgconn"`
}

// convert raw config to map
func createConnMap(config *Config) map[string]map[string]map[string]Connection {
	configMap := make(map[string]map[string]map[string]Connection)

	for _, a := range config.Pgconn {
		configMap[a.Account] = map[string]map[string]Connection{}

		for _, db := range a.Dbs {
			configMap[a.Account][db.Name] = map[string]Connection{}
			hostname := db.Hostname
			port := db.Port
			proxyKind := db.Proxy.Kind
			proxyHost := db.Proxy.Host

			for _, role := range db.Roles {
				configMap[a.Account][db.Name][role.Username] = Connection{
					Hostname:  hostname,
					ProxyKind: proxyKind,
					ProxyHost: proxyHost,
					Port:      port,
					Username:  role.Username,
					Password:  role.Password,
					Dbname:    role.Dbname,
				}
			}
		}
	}

	return configMap
}
