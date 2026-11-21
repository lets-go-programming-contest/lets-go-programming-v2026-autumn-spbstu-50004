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
	errBlankConfigPath  = errors.New("configuration file path cannot be blank")
	errMultipleYAMLDocs = errors.New("configuration must contain a single YAML document")
	errEmptyInputPath   = errors.New("input-file field is required")
	errEmptyOutputPath  = errors.New("output-file field is required")
)

type Settings struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func Load(path string) (Settings, error) {
	trimmed := strings.TrimSpace(path)
	if trimmed == "" {
		return Settings{}, errBlankConfigPath
	}
	cleanPath := filepath.Clean(trimmed)

	raw, err := os.ReadFile(cleanPath)
	if err != nil {
		return Settings{}, fmt.Errorf("read config file %q: %w", cleanPath, err)
	}

	decoder := yaml.NewDecoder(bytes.NewReader(raw))
	decoder.KnownFields(true)

	var cfg Settings
	if err := decoder.Decode(&cfg); err != nil {
		return Settings{}, fmt.Errorf("decode YAML %q: %w", cleanPath, err)
	}

	var trailing any
	if err := decoder.Decode(&trailing); !errors.Is(err, io.EOF) {
		if err == nil {
			return Settings{}, errMultipleYAMLDocs
		}
		return Settings{}, fmt.Errorf("check trailing YAML in %q: %w", cleanPath, err)
	}

	return validate(cfg)
}

func validate(cfg Settings) (Settings, error) {
	cfg.InputFile = strings.TrimSpace(cfg.InputFile)
	cfg.OutputFile = strings.TrimSpace(cfg.OutputFile)

	if cfg.InputFile == "" {
		return Settings{}, errEmptyInputPath
	}
	if cfg.OutputFile == "" {
		return Settings{}, errEmptyOutputPath
	}

	cfg.InputFile = filepath.Clean(cfg.InputFile)
	cfg.OutputFile = filepath.Clean(cfg.OutputFile)
	return cfg, nil
}
