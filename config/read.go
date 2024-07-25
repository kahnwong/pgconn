package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/getsops/sops/v3/decrypt"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Pgconn []struct {
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
	} `yaml:"pgconn"`
}

func readConfig() Config {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	filename := filepath.Join(homeDir, ".config", "pgconn", "pgconn.sops.yaml")

	// Check if the file exists
	_, err = os.Stat(filename)

	if os.IsNotExist(err) {
		fmt.Printf("File %s does not exist.\n", filename)
		os.Exit(1)
	}

	var config Config

	data, err := decrypt.File(filename, "yaml")
	if err != nil {
		fmt.Println(fmt.Printf("Failed to decrypt: %v", err))
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	return config
}
