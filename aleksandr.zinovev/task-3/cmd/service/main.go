package main

import (
	"flag"
	"fmt"

	"aleksandr.zinovev/task-3/internal/config"
	"aleksandr.zinovev/task-3/internal/currency"
)

const defaultConfigPath = "config.yaml"

func main() {
	configPath := flag.String("config", defaultConfigPath, "path to config file")
	flag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		panic(fmt.Errorf("load config: %w", err))
	}

	valutes, err := currency.ParseXML(cfg.InputFile)
	if err != nil {
		panic(fmt.Errorf("parse xml: %w", err))
	}

	currency.SortByValueDesc(valutes)

	if err := currency.SaveJSON(cfg.OutputFile, valutes); err != nil {
		panic(fmt.Errorf("save json: %w", err))
	}
}
