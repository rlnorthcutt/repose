package main

import (
	"fmt"
	"os"
	"path/filepath"
)

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

// Define the name of the config file
const ConfigFile = "config.yml"

// **********  Public Config Methods  **********

// Loads the site configuration from the config file
// We use this instead of loading the YAML modules to keep the size down
func (c *Config) Load() (Config, error) {
	// Read the entire config file content
	configPath := filepath.Join(command.rootPath, ConfigFile)
	data, err := os.ReadFile(configPath)
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
	config.PreviewURL = yamlMap["previewUrl"]

	return config, nil
}

// Create initializes the site configuration file
func (c *Config) Create(installDir string) error {
	logger.Info("Initializing config file")
	sitename := command.promptForInput("Enter the site name", "Repose site")
	author := command.promptForInput("Enter the author's name", "Creator")
	editor := command.promptForInput("Enter the editor ('none' for no editing)", "nano")
	url := command.promptForInput("Enter the site URL", "mysite.com")
	theme := command.promptForInput("Enter the CSS theme to use (pico, bootstrap, tailwind, none)", "pico")

	configContent := fmt.Sprintf(configTemplate, sitename, author, editor, "content", "web", url, "http://localhost:8080", theme)

	// Create the filepath
	configPath := filepath.Join(installDir, ConfigFile)

	// Write the configContent to config.yml
	if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
		logger.Error("Error writing config file:", err)
		return err
	}

	logger.Success("Config initialized successfully.")

	return nil
}

// The template for the config file
const configTemplate = `sitename: %s
author: %s
editor: %s
contentDirectory: %s
outputDirectory: %s
url: %s
previewUrl: %s
theme: %s
`
