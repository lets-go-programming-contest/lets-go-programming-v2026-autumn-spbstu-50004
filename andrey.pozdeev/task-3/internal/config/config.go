package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

var errEmptyConfigPaths = errors.New("input-file and output-file must be specified")

type Settings struct {
	InputPath  string `yaml:"input-file"`
	OutputPath string `yaml:"output-file"`
}

func Parse(configPath string) (Settings, error) {
	cleanPath := filepath.Clean(configPath)

	rawBytes, err := os.ReadFile(cleanPath)
	if err != nil {
		return Settings{}, fmt.Errorf("read yaml config: %w", err)
	}

	var cfg Settings
	if err := yaml.Unmarshal(rawBytes, &cfg); err != nil {
		return Settings{}, fmt.Errorf("decode yaml config: %w", err)
	}

	if cfg.InputPath == "" || cfg.OutputPath == "" {
		return Settings{}, errEmptyConfigPaths
	}

	cfg.InputPath = filepath.Clean(cfg.InputPath)
	cfg.OutputPath = filepath.Clean(cfg.OutputPath)

	return cfg, nil
}
