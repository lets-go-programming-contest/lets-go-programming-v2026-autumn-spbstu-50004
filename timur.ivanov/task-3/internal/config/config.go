package config

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

const defaultConfigPath = "config.yaml"

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func LoadFromArgs() *Config {
	configPath := flag.String("config", defaultConfigPath, "path to config file")
	flag.Parse()

	cfg, err := Load(*configPath)
	if err != nil {
		panic(err)
	}

	return cfg
}

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading config: %w", err)
	}

	var cfg Config

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("error checking YAML: %w", err)
	}

	return &cfg, nil
}
