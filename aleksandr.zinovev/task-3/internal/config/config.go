package config

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

var (
	errEmptyConfigPath  = errors.New("config file path cannot be empty")
	errMultipleYAMLDocs = errors.New("config file must contain a single YAML document")
	errEmptyInputFile   = errors.New("input-file is required in config")
	errEmptyOutputFile  = errors.New("output-file is required in config")
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func Load(path string) (Config, error) {
	trimmedPath := strings.TrimSpace(path)
	if trimmedPath == "" {
		return Config{}, errEmptyConfigPath
	}

	data, err := os.ReadFile(filepath.Clean(trimmedPath))
	if err != nil {
		return Config{}, fmt.Errorf("read config file %q: %w", trimmedPath, err)
	}

	decoder := yaml.NewDecoder(bytes.NewReader(data))
	decoder.KnownFields(true)

	var cfg Config
	if err := decoder.Decode(&cfg); err != nil {
		return Config{}, fmt.Errorf("decode YAML: %w", err)
	}

	var extra any
	if err := decoder.Decode(&extra); !errors.Is(err, io.EOF) {
		if err == nil {
			return Config{}, errMultipleYAMLDocs
		}

		return Config{}, fmt.Errorf("unexpected trailing content in YAML: %w", err)
	}

	cfg.InputFile = strings.TrimSpace(cfg.InputFile)
	cfg.OutputFile = strings.TrimSpace(cfg.OutputFile)

	if cfg.InputFile == "" {
		return Config{}, errEmptyInputFile
	}

	if cfg.OutputFile == "" {
		return Config{}, errEmptyOutputFile
	}

	return cfg, nil
}
