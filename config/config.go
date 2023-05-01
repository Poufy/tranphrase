package config

import (
	"log"

	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	SERVER struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
}

var Cfg Config

func LoadConfig() {
	file, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	err = yaml.Unmarshal(file, &Cfg)
	if err != nil {
		log.Fatalf("Failed to unmarshal config file: %v", err)
	}
}
