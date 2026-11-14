package config

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

var (
	errEmptyConfigPath       = errors.New("Configuration file path is empty")
	errMultipleYAMLDocuments = errors.New("Configuration file must contain exactly one YAML document")
	errMissingInputFile      = errors.New("Required field input-file is missing")
	errMissingOutputFile     = errors.New("Required field output-file is missing")
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func Load(path string) (Config, error) {
	path = strings.TrimSpace(path)
	if path == "" {
		return Config{}, errEmptyConfigPath
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("Read configuration file %q: %w", path, err)
	}

	cfg, err := decode(data)
	if err != nil {
		return Config{}, fmt.Errorf("Decode configuration file %q: %w", path, err)
	}

	return validate(cfg)
}

func decode(data []byte) (Config, error) {
	decoder := yaml.NewDecoder(bytes.NewReader(data))
	decoder.KnownFields(true)

	var cfg Config
	if err := decoder.Decode(&cfg); err != nil {
		return Config{}, fmt.Errorf("Decode YAML: %w", err)
	}

	var extra any
	if err := decoder.Decode(&extra); !errors.Is(err, io.EOF) {
		if err == nil {
			return Config{}, errMultipleYAMLDocuments
		}
		return Config{}, fmt.Errorf("Unexpected trailing YAML content: %w", err)
	}

	return cfg, nil
}

func validate(cfg Config) (Config, error) {
	cfg.InputFile = strings.TrimSpace(cfg.InputFile)
	cfg.OutputFile = strings.TrimSpace(cfg.OutputFile)

	if cfg.InputFile == "" {
		return Config{}, errMissingInputFile
	}
	if cfg.OutputFile == "" {
		return Config{}, errMissingOutputFile
	}

	return cfg, nil
}
