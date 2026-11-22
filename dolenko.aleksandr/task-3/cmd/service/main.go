package main

import (
	"flag"
	"fmt"
)

const defaultConfigPath = "config.yaml"

func main() {
	configPath := flag.String("config", defaultConfigPath, "path to config file")
	flag.Parse()

	fmt.Println("Config file:", *configPath)
}
