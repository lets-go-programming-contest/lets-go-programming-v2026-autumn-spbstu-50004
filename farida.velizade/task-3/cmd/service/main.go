package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/velizadefarida/task-3/internal/cbr"
	"github.com/velizadefarida/task-3/internal/config"
)

const (
	writeAccess = 0o600
	mkdirAccess = 0o755
)

func main() {
	configPath := flag.String("config", "config.yaml", "path to configuration file")
	flag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		panic(fmt.Errorf("config error: %w", err))
	}

	rates, err := cbr.ParseXML(cfg.InputFile)
	if err != nil {
		panic(fmt.Errorf("cbr parse error: %w", err))
	}

	if err := exportToReceiver(cfg.OutputFile, rates); err != nil {
		panic(fmt.Errorf("json save error: %w", err))
	}
}

func exportToReceiver(target string, data []cbr.Currency) error {
	payload, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal indented json: %w", err)
	}

	if err := os.MkdirAll(filepath.Dir(target), mkdirAccess); err != nil {
		return fmt.Errorf("create output directory tree: %w", err)
	}

	if err := os.WriteFile(target, payload, writeAccess); err != nil {
		return fmt.Errorf("write final json payload: %w", err)
	}

	return nil
}
