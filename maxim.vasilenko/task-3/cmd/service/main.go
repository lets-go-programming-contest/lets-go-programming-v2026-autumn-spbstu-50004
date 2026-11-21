package main

import (
	"flag"

	"github.com/Madmaxvas/task-3/internal/config"
	"github.com/Madmaxvas/task-3/internal/currency"
)

func main() {
	configPath := flag.String("config", "", "path to configuration file")
	flag.Parse()

	if *configPath == "" {
		panic("config file path is required")
	}

	cfg, err := config.Load(*configPath)
	if err != nil {
		panic(err)
	}

	currencies, err := currency.ParseXML(cfg.InputFile)
	if err != nil {
		panic(err)
	}

	currency.SortByValueDesc(currencies)

	if err := currency.SaveJSON(cfg.OutputFile, currencies); err != nil {
		panic(err)
	}
}
