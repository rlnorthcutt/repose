package main

import (
	"flag"
	"log"
	"os"
)

// @TODO: Move this to a separate config.go file
// Config struct to hold the configuration values
type Config struct {
	// Sitename is name of the site - funny right?
	Sitename string `yaml:"sitename"`
	// Author is the creator of the site & used as the default author for content
	Author string `yaml:"author"`
	// Editor is the text editor that should be used to edit content
	// Leave blank to not auto-open the editor - defaults to nano
	Editor string `yaml:"editor"`
	// ContentDirectory is the directory where the content is stored
	// Defaults to "content"
	ContentDirectory string `yaml:"contentDirectory"`
	// OutputDirectory is the directory where the generated site is stored
	// Defaults to "web"
	OutputDirectory string `yaml:"outputDirectory"`
	// URL is the URL for the site
	URL string `yaml:"url"`
	// PreviewURL is the URL for the local preview server
	// Defaults to "http://localhost:8080"
	PreviewURL string `yaml:"previewUrl"`
	// Theme is the theme to use for the site
	// Defaults to picocss, but can be bootstrap or tailwind
	Theme string `yaml:"theme"`
}

// Create a global config variable so it can be accessed from anywhere
var config Config

// Func main should be as small as possible and do as little as possible by convention
func main() {
	// Check if a command is provided, immediately exit if not.
	if len(os.Args) < 2 {
		logger.Warn("Expected a command - type `repose help` to get options.")
		os.Exit(0)
	}

	// Generate our config based on the config supplied
	// by the user in the flags
	configPath, rootPath, err := parseFlags()
	if err != nil {
		log.Fatal(err)
	}

	// Set rootPath and configPath for command
	// @TODO look at moving this to main instead of command
	command.SetRootPath(rootPath)
	command.SetConfigPath(configPath)

	// Load config for specific commands
	commandName := os.Args[1]
	switch commandName {
	case "new", "build", "preview":
		var err error
		config, err = config.ReadConfig()
		if err != nil {
			logger.Warn("No config file found. You need to run `repose init` first.")
			os.Exit(0)
		}
	}

	// Dispatch the command
	dispatchCommand(commandName, config)
}

// readSiteConfig reads the site configuration from the config file
// We use this instead of loading the YAML modules to keep the size down
func (c *Config) ReadConfig() (Config, error) {
	// Read the entire config file content
	data, err := os.ReadFile(command.configPath)
	if err != nil {
		return Config{}, err
	}

	// Use readYAML to parse the content
	yamlMap, err := filesystem.ParseYml(string(data))
	if err != nil {
		return Config{}, err
	}

	// Populate the Config struct with data from the map
	config.Sitename = yamlMap["sitename"]
	config.Author = yamlMap["author"]
	config.Editor = yamlMap["editor"]
	config.ContentDirectory = yamlMap["contentDirectory"]
	config.OutputDirectory = yamlMap["outputDirectory"]
	config.URL = yamlMap["url"]
	config.Theme = yamlMap["theme"]

	return config, nil
}

// **********  Private Main methods  **********

// ParseFlags will create and parse the CLI flags
// and return the path to be used elsewhere
func parseFlags() (string, string, error) {
	// Set the default values for the flags
	defaultConfigPath := "config.yml"
	configPath := ""
	rootPath := ""

	// Parse flags
	flag.StringVar(&configPath, "config", defaultConfigPath, "Path to the configuration file")
	flag.StringVar(&configPath, "c", defaultConfigPath, "Path to the configuration file (shorthand)")
	flag.StringVar(&rootPath, "root", rootPath, "Root path of the project")
	flag.StringVar(&rootPath, "r", rootPath, "Root path of the project (shorthand)")
	flag.Parse()

	// Ensure default config path is used if flag is not set
	if configPath == "" {
		configPath = defaultConfigPath
	}

	// Return the configuration and root paths
	return configPath, rootPath, nil
}

// dispatchCommand will take the command name and dispatch it to the correct function
func dispatchCommand(commandName string, config Config) {
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
