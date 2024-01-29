package main

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Sitename         string `yaml:"sitename"`
	Author           string `yaml:"author"`
	Editor           string `yaml:"editor"`
	ContentDirectory string `yaml:"contentDirectory"`
	OutputDirectory  string `yaml:"outputDirectory"`
}

var config Config

func main() {
	// Check if a command is provided, immediately exit if not.
	if len(os.Args) < 2 {
		logger.Warn("Expected a command - type `sitestat help` to get options.")
		os.Exit(0)
	}

	// Load the config
	//  @TODO redo this so we only need it for  build & new
	config, err := readConfig("config.yml")
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}
	_ = config // to bypass "declared but not used" error

	// Dispatch the command
	switch os.Args[1] {
	case "init":
		command.Init()
	case "new":
		command.New()
	case "demo":
		command.Demo()
	case "build":
		command.Build()
	case "preview":
		command.Preview()
	case "update":
		command.Update()
	case "help":
		command.Help()
	default:
		logger.Error("Unknown command: %s\n", os.Args[1])
	}

}

// **********  Private Main methods  **********

// readConfig loads the configuration from the yml file
func readConfig(configPath string) (*Config, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
