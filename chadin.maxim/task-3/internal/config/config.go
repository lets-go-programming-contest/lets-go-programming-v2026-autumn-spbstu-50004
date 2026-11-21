package config

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func Load(configPath string) Config {
	if configPath == "" {
		panic("flag --config is required")
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		panic(fmt.Errorf("failed to read config file: %w", err))
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		panic(fmt.Errorf("failed to parse yaml config: %w", err))
	}

	cfg.InputFile = strings.TrimSpace(cfg.InputFile)
	cfg.OutputFile = strings.TrimSpace(cfg.OutputFile)

	if cfg.InputFile == "" || cfg.OutputFile == "" {
		panic("input-file or output-file is empty in config")
	}

	return cfg
}
