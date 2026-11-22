package main

import (
	"flag"
	"fmt"

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

	if err := cbr.WriteJSON(cfg.OutputFile, rates); err != nil {
		panic(fmt.Errorf("json write error: %w", err))
	}
}
