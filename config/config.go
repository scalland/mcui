package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type MCConfig struct {
	MCHost string `yaml:"mc_host"`
	MCPort int    `yaml:"mc_port"`
}

type AppConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Config struct {
	Memcached MCConfig  `yaml:"memcached"`
	App       AppConfig `yaml:"app"`
}

func LoadConfig(path string) (Config, error) {
	var cfg Config

	data, err := os.ReadFile(path)
	if err != nil {
		return cfg, fmt.Errorf("failed to read config file: %s", err.Error())
	}
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return cfg, fmt.Errorf("failed to parse config: %s", err.Error())
	}
	return cfg, nil
}
