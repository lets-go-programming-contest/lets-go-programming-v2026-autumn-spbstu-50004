package main

import (
	"flag"
	"fmt"

	"github.com/SoundwaveAI/task-3/internal/cbr"
	"github.com/SoundwaveAI/task-3/internal/config"
	"github.com/SoundwaveAI/task-3/internal/currency"
	"github.com/SoundwaveAI/task-3/internal/storage"
)

const defaultConfigPath = "config.yaml"

func main() {
	cfgPath := flag.String("config", defaultConfigPath, "path to YAML configuration file")
	flag.Parse()

	cfg, err := config.ParseFile(*cfgPath)
	if err != nil {
		panic(fmt.Errorf("load configuration: %w", err))
	}

	rates, err := cbr.ExtractRates(cfg.InputFile)
	if err != nil {
		panic(fmt.Errorf("extract currency rates: %w", err))
	}

	currency.SortByValueDescending(rates)

	err = storage.SaveRates(cfg.OutputFile, rates)
	if err != nil {
		panic(fmt.Errorf("save json output: %w", err))
	}
}
