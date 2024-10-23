package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Bot `yaml:"bot"`
}

type Bot struct {
	Token string `yaml:"token"`
	Debug bool   `yaml:"debug"`
}

func Parse(filepath string) (cfg *Config, err error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, &cfg)
	return
}
