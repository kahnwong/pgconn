package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// readConfig
type Config struct {
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
}

func readConfig() []Config {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	filename := filepath.Join(homeDir, ".config", "pgconn", "db.yaml")

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
