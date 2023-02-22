package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	BaseURL     string   `yaml:"baseURL"`
	IgnoreFiles []string `yaml:"ignoreFiles"`
	Permalinks  map[string]string
}

func FromFile(filename string) (Config, error) {
	f, err := os.Open(filename)
	if err != nil {
		return Config{}, err
	}

	var c Config
	err = yaml.NewDecoder(f).Decode(&c)
	if err != nil {
		return Config{}, err
	}

	return c, nil
}
