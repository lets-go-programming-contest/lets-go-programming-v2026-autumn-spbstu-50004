package main

import (
	"task-3/internal/config"
	"task-3/internal/currency"
)

func main() {
	cfg := config.LoadFromArgs()

	if err := currency.Process(cfg.InputFile, cfg.OutputFile); err != nil {
		panic(err)
	}
}
