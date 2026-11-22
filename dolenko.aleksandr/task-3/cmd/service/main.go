package main

import (
	"flag"
	"fmt"

	"github.com/xweads/task-3/internal/config"
	"github.com/xweads/task-3/internal/valcurs"
)

const defaultConfigPath = "config.yaml"

func main() {
	configPath := flag.String("config", defaultConfigPath, "path to config file")
	flag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		panic(err)
	}

	currencies, err := valcurs.Load(cfg.InputFile)
	if err != nil {
		panic(err)
	}

	valcurs.SortByValue(currencies.Valutes)

	fmt.Println("Sorted currencies:")

	for _, currency := range currencies.Valutes {
		fmt.Printf("%s: %.4f\n", currency.CharCode, currency.Value)
	}
}
