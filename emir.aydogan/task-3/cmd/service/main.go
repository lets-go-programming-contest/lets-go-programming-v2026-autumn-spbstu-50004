package main

import (
	"flag"
	"fmt"

	"github.com/emiraydgn/task-3/internal/cbr"
	"github.com/emiraydgn/task-3/internal/config"
	"github.com/emiraydgn/task-3/internal/jsonout"
)

const defaultConfigPath = "config.yaml"

func main() {
	configPath := flag.String("config", defaultConfigPath, "path to YAML configuration file")
	flag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		panic(fmt.Errorf("load config: %w", err))
	}

	currencies, err := cbr.Load(cfg.InputFile)
	if err != nil {
		panic(fmt.Errorf("load currencies: %w", err))
	}

	if err := jsonout.Save(cfg.OutputFile, currencies); err != nil {
		panic(fmt.Errorf("save JSON: %w", err))
	}
}
