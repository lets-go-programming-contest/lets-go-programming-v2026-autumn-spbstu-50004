package main

import (
	"flag"
	"os"

	"em.maxim/task-3/internal/config"
	"em.maxim/task-3/internal/currency"
)

func main() {
	cfgPath := flag.String("config", "", "config path")
	flag.Parse()

	if *cfgPath == "" {
		panic("missing config path")
	}

	cfg, err := config.Load(*cfgPath)
	if err != nil {
		panic(err)
	}

	if cfg.InputFile == "" || cfg.OutputFile == "" {
		panic("missing input-file or output-file")
	}

	xmlData, err := os.ReadFile(cfg.InputFile)
	if err != nil {
		panic(err)
	}

	rates, err := currency.ParseXML(xmlData)
	if err != nil {
		panic(err)
	}

	if err := currency.SaveJSON(cfg.OutputFile, rates); err != nil {
		panic(err)
	}
}
