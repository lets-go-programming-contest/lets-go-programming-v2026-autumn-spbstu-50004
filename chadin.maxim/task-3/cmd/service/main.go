package main

import (
	"flag"

	"github.com/maxchadin/task-3/internal/config"
	"github.com/maxchadin/task-3/internal/currency"
)

const defaultConfigPath = "config.yaml"

func main() {
	configPath := flag.String("config", defaultConfigPath, "path to yaml config file")
	flag.Parse()

	cfg := config.Load(*configPath)
	currencies := currency.ParseFromXML(cfg.InputFile)

	currency.SortDesc(currencies)
	currency.SaveToJSON(cfg.OutputFile, currencies)
}
