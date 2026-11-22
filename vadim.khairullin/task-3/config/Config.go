package config

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

const DefaultConfig = "config.yaml"

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func ParseConfig() (Config, error) {
	configPath := flag.String("config", DefaultConfig, "Path to config file")
	flag.Parse()

	file, err := os.Open(*configPath)
	if err != nil {
		return Config{}, fmt.Errorf("failed to open config file: %w", err)
	}

	defer func() {
		err := file.Close()
		if err != nil {
			panic(fmt.Errorf("failed to close file: %w", err))
		}
	}()

	var conf Config

	decoder := yaml.NewDecoder(file)

	err = decoder.Decode(&conf)
	if err != nil {
		return conf, fmt.Errorf("failed to decode YAML file: %w", err)
	}

	return conf, nil
}
