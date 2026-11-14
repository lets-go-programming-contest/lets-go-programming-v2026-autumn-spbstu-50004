package main

import (
	"flag"

	"github.com/ekaterinaivanc/task-3/internal/cbr"
	"github.com/ekaterinaivanc/task-3/internal/config"
	"github.com/ekaterinaivanc/task-3/internal/jsonout"
	"github.com/ekaterinaivanc/task-3/internal/valute"
)

func main() {
	configPath := flag.String("config", "", "path to YAML configuration file")

	flag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		panic(err)
	}

	items, err := cbr.ReadFile(cfg.InputFile)
	if err != nil {
		panic(err)
	}

	valute.SortByValueDesc(items)

	err = jsonout.Write(cfg.OutputFile, items)
	if err != nil {
		panic(err)
	}
}
