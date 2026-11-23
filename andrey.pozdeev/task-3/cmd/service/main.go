package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ryouta25/task-3/internal/config"
	"github.com/ryouta25/task-3/internal/rates"
)

const (
	defaultConfigFile = "config.yaml"
	outputDirPerm     = 0o755
	outputFilePerm    = 0o600
)

func main() {
	cfgPath := flag.String("config", defaultConfigFile, "path to configuration YAML file")
	flag.Parse()

	settings, err := config.Parse(*cfgPath)
	if err != nil {
		panic(fmt.Errorf("config error: %w", err))
	}

	sortedRates, err := rates.FetchSortedRates(settings.InputPath)
	if err != nil {
		panic(fmt.Errorf("rates data error: %w", err))
	}

	if err := exportToJSON(settings.OutputPath, sortedRates); err != nil {
		panic(fmt.Errorf("export error: %w", err))
	}
}

func exportToJSON(destPath string, data []rates.Rate) error {
	cleanDest := filepath.Clean(destPath)

	if err := os.MkdirAll(filepath.Dir(cleanDest), outputDirPerm); err != nil {
		return fmt.Errorf("create target directory: %w", err)
	}

	payload, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("encode json payload: %w", err)
	}

	if err := os.WriteFile(cleanDest, payload, outputFilePerm); err != nil {
		return fmt.Errorf("write json file: %w", err)
	}

	return nil
}
