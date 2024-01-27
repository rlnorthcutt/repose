package main

import (
	"flag"
	"fmt"
	"os"
)

// Defining a new type 'logger' as an int
type logger int

// Declaring a global variable 'log' of type 'logger'
var log logger

func main() {
	// Check if a command is provided
	if len(os.Args) < 2 {
		log.Warn("Expected a command - type `zenforge help` to get options.")
		os.Exit(0)
	}

	// Determine the command
	switch os.Args[1] {
	case "init":
		initCommand()
	case "new":
		newCommand()
	case "demo":
		demoCommand()
	case "build":
		buildCommand()
	case "serve":
		previewCommand()
	case "help":
		helpCommand()
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
	}
}

// **********  Command functions  **********************************************

// Initializes a new ZenForge project.
// It creates the proper folder structure and starter files.
func initCommand() string {
	fmt.Println("Initializing ZenForge project")
	return ""
}

// Creates new content based on the provided content type and filename.
// It requires two arguments: content type and filename.
// If the required arguments are not provided, it prints the usage information.
func newCommand() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: zenforge new [CONTENTTYPE] [FILENAME]")
		return
	}
	contentType := os.Args[2]
	fileName := os.Args[3]
	fmt.Printf("Creating new %s with filename %s\n", contentType, fileName)
}

// Generates a new project with demo content and templates to create a new site.
func demoCommand() string {
	fmt.Printf("Generating demo content")
	return ""
}

// Builds the ZenForge site based on the current project default values.
// It uses command-line flags to modify the root directory and config file.
// If there is an error parsing the command flags, it prints an error message.
func buildCommand() {
	buildCmd := flag.NewFlagSet("build", flag.ExitOnError)
	root := buildCmd.String("r", ".", "Directory to use as root of the project")
	config := buildCmd.String("c", "config.yml", "Path to configuration file")

	if err := buildCmd.Parse(os.Args[2:]); err != nil {
		fmt.Println("Error parsing build command flags:", err)
		return
	}

	fmt.Printf("Building site from %s with config %s\n", *root, *config)
}

// Starts serving the ZenForge site for local preview.
func previewCommand() string {
	fmt.Printf("ZenForge site")
	return ""
}

func helpCommand() string {
	return `ZenForge Commands:
    init  - Initialize a new ZenForge project
    new   - Create new content. Usage: zenforge new [CONTENTTYPE] [FILENAME]
    demo  - Generate demo content
    build - Build the site. Options: -r [ROOT], -c [CONFIG]
    preview - Setup a local server to preview the site
    help  - Show this help message`
}

// *****************************************************************************

// **********  Logger output functions  ****************************************
// Warn method for the logger type
// This method formats and prints a warning message with yellow color
func (l *logger) Warn(format string, value ...any) {
	fmt.Printf("\u001B[0;33m[WARN]\u001B[0;39m "+format, value...)
}

// Err method for the logger type
// This method formats and prints an error message with red color
func (l *logger) Err(format string, value ...any) {
	fmt.Printf("\u001B[0;31m[ERROR]\u001B[0;39m "+format, value...)
}

// Info method for the logger type
// This method formats and prints an info message with green color
func (l *logger) Info(format string, value ...any) {
	fmt.Printf("\u001B[0;32m[INFO]\u001B[0;39m "+format, value...)
}

// Fatal method for the logger type
// This method formats and prints a fatal message and exits the program
func (l *logger) Fatal(format string, value ...any) {
	fmt.Printf("\u001B[0;31m[FATAL]\u001B[0;39m "+format, value...)
}

// *****************************************************************************
