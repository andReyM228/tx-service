package config

import (
	"github.com/andReyM228/one/chain_client"
	"gopkg.in/yaml.v3"

	"log"
	"os"
)

type (
	Config struct {
		Chain  chain_client.ClientConfig `yaml:"chain"`
		HTTP   HTTP                      `yaml:"http"`
		Extra  Extra                     `yaml:"extra"`
		Rabbit Rabbit                    `yaml:"rabbit"`
	}

	HTTP struct {
		Port int `yaml:"port"`
	}

	Extra struct {
		Mnemonic string `yaml:"mnemonic"`
	}

	Rabbit struct {
		Url string `yaml:"url"`
	}
)

func ParseConfig() (Config, error) {
	file, err := os.ReadFile("./cmd/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	var cfg Config

	if err := yaml.Unmarshal(file, &cfg); err != nil {
		log.Fatal(err)
	}

	return cfg, nil
}
