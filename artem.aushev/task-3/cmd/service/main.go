package main

import (
	"errors"
	"flag"
	"fmt"
	"log"

	"task-3/internal/config"
	"task-3/internal/currency"
)

var errConfigPathRequired = errors.New("config file path is required")

func main() {
	configFilePath := flag.String("config", "", "path to configuration file")
	flag.Parse()

	if configFilePath == nil || *configFilePath == "" {
		log.Println(errConfigPathRequired)

		return
	}

	appConfig, loadError := config.Load(*configFilePath)
	if loadError != nil {
		log.Println(fmt.Errorf("failed to load configuration from %q: %w", *configFilePath, loadError))

		return
	}

	parsedCurrencies, parseError := currency.ParseXML(appConfig.InputFile)
	if parseError != nil {
		log.Println(fmt.Errorf("failed to parse currencies from XML file %q: %w", appConfig.InputFile, parseError))

		return
	}

	currency.SortByValueDesc(parsedCurrencies)

	if saveError := currency.SaveJSON(appConfig.OutputFile, parsedCurrencies); saveError != nil {
		log.Println(fmt.Errorf("failed to save sorted currencies to JSON file %q: %w", appConfig.OutputFile, saveError))

		return
	}
}
