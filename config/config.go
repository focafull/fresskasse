package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type FressKasseConfig struct {
	PORT        string
	ADDRESS     string
	DB_ADDRESS  string
	DB_USER     string
	DB_PASSWORD string
	DB          string
}

var config *FressKasseConfig

func Init(configPath string) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("ERROR: %s\n", err)
	}

	config = &FressKasseConfig{}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}
}

func GetConfig() *FressKasseConfig {
	return config
}
