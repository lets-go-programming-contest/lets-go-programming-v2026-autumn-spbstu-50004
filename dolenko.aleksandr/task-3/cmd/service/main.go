package main

import (
	"flag"
	"fmt"

	"github.com/xweads/task-3/internal/config"
)

const defaultConfigPath = "config.yaml"

func main() {
	configPath := flag.String("config", defaultConfigPath, "path to config file")
	flag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		panic(err)
	}

	fmt.Println("Input file:", cfg.InputFile)
	fmt.Println("Output file:", cfg.OutputFile)
}
