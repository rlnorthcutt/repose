package main

import (
	"flag"
	"fmt"
	"os"
)

// Defining a new public type 'Command'
type Command int

// Defining a global varaiable for Command
var command Command

// **********  Public Command Methods  **********

// Initializes a new ZenForge project.
// It creates the proper folder structure and starter files.
func (c *Command) Init() string {
	fmt.Println("Initializing ZenForge project")
	return ""
}

// Creates new content based on the provided content type and filename.
// It requires two arguments: content type and filename.
// If the required arguments are not provided, it prints the usage information.
func (c *Command) New() {
	if len(os.Args) != 4 {
		logger.Info("Usage: zenforge new [CONTENTTYPE] [FILENAME]")
		return
	}
	contentType := os.Args[2]
	fileName := os.Args[3]

	// Determine the path and content based on contentType
	var path string
	var content string

	// Example: Customize path and content based on contentType
	switch contentType {
	case "file":
		path = fileName // Assuming fileName includes the path
		content = ""    // Default content for a new file
	case "directory":
		path = fileName // Directory path
		content = ""    // No content needed for directory
	default:
		logger.Error("Unknown content type: %s", contentType)
		return
	}

	// Create the file or directory
	if err := filesystem.Create(path, content); err != nil {
		logger.Error("Failed to create %s: %v", path, err)
		return
	}

	fmt.Printf("Successfully created new %s: %s\n", contentType, path)
}

// Generates a new project with demo content and templates to create a new site.
func (c *Command) Demo() string {
	logger.Info("Generating demo content")
	return ""
}

// Builds the ZenForge site based on the current project default values.
// It uses command-line flags to modify the root directory and config file.
// If there is an error parsing the command flags, it prints an error message.
func (c *Command) Build() {
	buildCmd := flag.NewFlagSet("build", flag.ExitOnError)
	root := buildCmd.String("r", ".", "Directory to use as root of the project")
	config := buildCmd.String("c", "config.yml", "Path to configuration file")

	if err := buildCmd.Parse(os.Args[2:]); err != nil {
		logger.Error("Error parsing build command flags:", err)
		return
	}

	logger.Info("Building site from %s with config %s\n", *root, *config)
}

// Starts serving the ZenForge site for local preview.
func (c *Command) Preview() string {
	fmt.Printf("ZenForge site")
	return ""
}

func (c *Command) Help() string {
	response := `ZenForge Commands:
    init  - Initialize a new ZenForge project
    new   - Create new content. Usage: zenforge new [CONTENTTYPE] [FILENAME]
    demo  - Generate demo content
    build - Build the site. Options: -r [ROOT], -c [CONFIG]
    preview - Setup a local server to preview the site
    help  - Show this help message\n`
	logger.Info(response)
	return ""
}

// **********  Private Command Methods  **********
