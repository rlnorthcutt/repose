package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Defining a new public type 'Command'
type Command int

// Defining a global varaiable for Command
var command Command

// **********  Public Command Methods  **********

// Initializes a new SiteStat project.
// It creates the proper folder structure and starter files.
func (c *Command) Init() string {
	fmt.Println("Initializing SiteStat project")
	return ""
}

// Creates new content based on the provided content type and filename.
// It requires two arguments: content type and filename.
// The content type defines the path, so it can also include a subfolder
func (c *Command) New() {
	// If the required arguments are not provided, print the usage information.
	if len(os.Args) != 4 {
		logger.Info("Usage: sitestat new [CONTENTTYPE] [FILENAME]")
		return
	}
	contentDirectory := config.ContentDirectory
	typeDirectory := os.Args[2]
	fileNameParam := os.Args[3]
	fileName, title := processFileName(fileNameParam)

	// We allow the user to include a subdirectory in the content type param
	// Extract the first part as the type if contentType has a '/'
	var contentType string
	if strings.Contains(contentType, "/") {
		contentType = strings.Split(contentType, "/")[0]
	} else {
		contentType = typeDirectory
	}

	// Construct the path
	path := filepath.Join(contentDirectory, typeDirectory, fileName)

	// Get default content
	content := defaultContent(contentType, title)

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

// Builds the SiteStat site based on the current project default values.
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

// Starts serving the SiteStat site for local preview.
func (c *Command) Preview() string {
	fmt.Printf("SiteStat site")
	return ""
}

// Updates the SiteStat binary in the current directory
func (c *Command) Update() string {
	fmt.Printf("SiteStat update placeholder")
	return ""
}

func (c *Command) Help() string {
	response := `SiteStat Commands:
    init  - Initialize a new SiteStat project
    new   - Create new content. Usage: sitestat new [CONTENTTYPE] [FILENAME]
    build - Build the site. Options: -r [ROOT], -c [CONFIG]
    preview - Setup a local server to preview the site
    demo  - Generate demo content
	update - Update the sitestat binary
    help  - Show this help message\n`
	logger.Info(response)
	return ""
}

// **********  Private Command Methods  **********

func processFileName(fileName string) (string, string) {
	// Check the file extension
	ext := filepath.Ext(fileName)
	if ext != ".md" && ext != ".html" {
		// If no extension or any other extension, make it .md
		fileName = fileName + ".md"
		ext = ".md" // Update the extension to .md
	}

	fileNameWithoutExt := strings.TrimSuffix(fileName, ext)

	// Convert fileName to a title
	replaceWithSpaces := strings.Replace(strings.Replace(fileNameWithoutExt, "-", " ", -1), "_", " ", -1)
	caser := cases.Title(language.English)
	title := caser.String(replaceWithSpaces)

	return fileName, title
}

// defaultContent returns default content based on the content type.
func defaultContent(contentType string, title string) string {
	content := `---
title: "{title}"
description: "{contentType} about {title}"
tags: []
image: 
index: true
author: "{author}"
publish_date: 
template: "{contentType}.tmpl"
---
	
# {title}

`

	// Replace placeholders with actual values
	content = strings.Replace(content, "{title}", title, -1)
	content = strings.Replace(content, "{contentType}", contentType, -1)
	content = strings.Replace(content, "{author}", config.Author, -1)

	return content
}
