package config

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

var (
	errInputFileEmpty  = errors.New("input-file is empty")
	errOutputFileEmpty = errors.New("output-file is empty")
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func Load(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("read config: %w", err)
	}

	var cfg Config

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return Config{}, fmt.Errorf("parse config: %w", err)
	}

	if strings.TrimSpace(cfg.InputFile) == "" {
		return Config{}, errInputFileEmpty
	}

	if strings.TrimSpace(cfg.OutputFile) == "" {
		return Config{}, errOutputFileEmpty
	}

	return cfg, nil
}
