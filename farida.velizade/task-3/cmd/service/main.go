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

	data, err := json.MarshalIndent(rates, "", "  ")
	if err != nil {
		panic(fmt.Errorf("json marshal error: %w", err))
	}

	if err = os.MkdirAll(filepath.Dir(cfg.OutputFile), 0755); err != nil {
		panic(fmt.Errorf("create dir error: %w", err))
	}

	if err = os.WriteFile(cfg.OutputFile, data, 0600); err != nil {
		panic(fmt.Errorf("write file error: %w", err))
	}
}
