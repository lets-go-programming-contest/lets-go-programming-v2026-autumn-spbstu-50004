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

func ParseFile(configPath string) (Settings, error) {
	trimmedPath := strings.TrimSpace(configPath)
	if trimmedPath == "" {
		return Settings{}, errBlankConfigPath
	}

	cleanPath := filepath.Clean(trimmedPath)

	raw, err := os.ReadFile(cleanPath)
	if err != nil {
		return Settings{}, fmt.Errorf("read config file %q: %w", cleanPath, err)
	}

	decoder := yaml.NewDecoder(bytes.NewReader(raw))
	decoder.KnownFields(true)

	var cfg Settings

	err = decoder.Decode(&cfg)
	if err != nil {
		return Settings{}, fmt.Errorf("decode YAML %q: %w", cleanPath, err)
	}

	var trailing any

	err = decoder.Decode(&trailing)
	if !errors.Is(err, io.EOF) {
		if err == nil {
			return Settings{}, errMultipleYAMLDocs
		}

		return Settings{}, fmt.Errorf("check trailing YAML in %q: %w", cleanPath, err)
	}

	return preparePaths(cfg)
}

func preparePaths(cfg Settings) (Settings, error) {
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
