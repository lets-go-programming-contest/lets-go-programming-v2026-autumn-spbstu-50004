package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"ahrameev.vladislav/task-3/pkg/config"
	"ahrameev.vladislav/task-3/pkg/currency"
)

const dirPerm = 0o755

func main() {
	configPath := flag.String("config", "", "path to config file")

	flag.Parse()

	if *configPath == "" {
		panic("config flag is required")
	}

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		panic(fmt.Sprintf("failed to load config: %v", err))
	}

	xmlData, err := os.ReadFile(cfg.InputFile)
	if err != nil {
		panic(fmt.Sprintf("failed to read input file: %v", err))
	}

	valCurs, err := currency.ParseXML(xmlData)
	if err != nil {
		panic(fmt.Sprintf("failed to parse XML: %v", err))
	}

	results, err := currency.ToResults(valCurs.Valutes)
	if err != nil {
		panic(fmt.Sprintf("failed to convert results: %v", err))
	}

	currency.SortByValueDescending(results)

	outputDir := filepath.Dir(cfg.OutputFile)
	if err := os.MkdirAll(outputDir, dirPerm); err != nil {
		panic(fmt.Sprintf("failed to create output directory: %v", err))
	}

	if err := currency.SaveToJSON(results, cfg.OutputFile); err != nil {
		panic(fmt.Sprintf("failed to save JSON: %v", err))
	}

	fmt.Println("Done!")
}
