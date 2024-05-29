package config

import (
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Host        string `yaml:"host"`
	Port        string `yaml:"port"`
	Username    string `yaml:"username"`
	Password    string `yaml:"password"`
	Database    string `yaml:"database"`
	Delay       int    `yaml:"delay"`
	MaxAttempts int    `yaml:"maxAttempts"`
}

func New() Config {
	yamlFile, err := os.ReadFile("../config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	var cfg Config
	err = yaml.Unmarshal(yamlFile, &cfg)
	if err != nil {
		log.Fatal(err)
	}
	return cfg
}
