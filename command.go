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

// **********  Command functions  **********************************************

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
	fmt.Printf("Creating new %s with filename %s\n", contentType, fileName)
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
