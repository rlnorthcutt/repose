package main

import (
	"flag"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	siteName         string `yaml:"siteName"`
	author           string `yaml:"author"`
	editor           string `yaml:"editor"`
	contentDirectory string `yaml:"contentDirectory"`
	outputDirectory  string `yaml:"outputDirectory"`
	previewURL       string `yaml:"previewURL"`
}

var config Config

func main() {
	// Check if a command is provided, immediately exit if not.
	if len(os.Args) < 2 {
		logger.Warn("Expected a command - type `repose help` to get options.")
		os.Exit(0)
	}

	configPath := "config.yml"
	rootPath := ""

	// Parse flags if they exist
	flag.StringVar(&configPath, "config", configPath, "")
	flag.StringVar(&configPath, "c", configPath, "")
	flag.StringVar(&rootPath, "root", rootPath, "")
	flag.StringVar(&rootPath, "r", rootPath, "")
	flag.Parse()

	// Set rootPath and configPath for command
	command.SetRootPath(rootPath)
	command.SetConfigPath(configPath)

	// Load config for specific commands
	commandName := os.Args[1]
	switch commandName {
	case "new", "build", "preview":
		var err error
		config, err = readConfig()
		if err != nil {
			//logger.Fatal("Error reading config: %v", err)
			logger.Info("No config file found. You need to run `repose init` first.")
			os.Exit(0)
		}
	}

	// Dispatch the command
	switch commandName {
	case "init":
		command.Init()
	case "new":
		command.New(config)
	case "demo":
		command.Demo()
	case "build":
		command.Build(config)
	case "preview":
		command.Preview(config)
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
func readConfig() (Config, error) {

	data, err := os.ReadFile(command.configPath)
	if err != nil {
		return config, err
	}

	if err := yaml.Unmarshal(data, &config); err != nil {
		return config, err
	}

	return config, nil
}
