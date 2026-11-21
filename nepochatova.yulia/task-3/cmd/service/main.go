package main

import (
	"flag"
	"fmt"

	"github.com/GenJune222/task-3/internal/cbr"
	"github.com/GenJune222/task-3/internal/config"
	"github.com/GenJune222/task-3/internal/currency"
	"github.com/GenJune222/task-3/internal/storage"
)

const defaultConfigPath = "config.yaml"

func main() {
	configPath := flag.String("config", defaultConfigPath, "path to YAML configuration file")
	flag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		panic(fmt.Errorf("load config: %w", err))
	}

	rates, err := cbr.ParseXML(cfg.InputFile)
	if err != nil {
		panic(fmt.Errorf("parse XML: %w", err))
	}

	currency.SortByValueDescending(rates)

	if err := storage.SaveJSON(cfg.OutputFile, rates); err != nil {
		panic(fmt.Errorf("save JSON: %w", err))
	}

	fmt.Printf("OK: %d currencies written to %s\n", len(rates), cfg.OutputFile)
}
