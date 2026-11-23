package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func Load(filePath string) (*Config, error) {
	fileBytes, readError := os.ReadFile(filePath)
	if readError != nil {
		return nil, fmt.Errorf("failed to read config file %q: %w", filePath, readError)
	}

	var parsedConfig Config

	if unmarshalError := yaml.Unmarshal(fileBytes, &parsedConfig); unmarshalError != nil {
		return nil, fmt.Errorf("failed to parse YAML config from %q: %w", filePath, unmarshalError)
	}

	return &parsedConfig, nil
}
