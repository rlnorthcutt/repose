package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Defining a new public type 'Command'
type Command struct {
	rootPath   string
	configPath string
}

// Defining a global varaiable for Command
var command Command

// **********  Public Command Methods  **********

// Initializes a new Repose project.
// It creates the proper folder structure and starter files.
func (c *Command) Init() string {
	if err := createDirectoryStructure(c.rootPath); err != nil {
		log.Fatal("Error creating site structure: ", err)
	}
	logger.Info("Repose project created in %s", c.rootPath)
	return ""
}

// Creates new content based on the provided content type and filename.
// It requires two arguments: content type and filename.
// The content type defines the path, so it can also include a subfolder
func (c *Command) New(config Config) {
	// If the required arguments are not provided, print the usage information.
	if len(os.Args) != 4 {
		logger.Info("Usage: repose new [CONTENTTYPE] [FILENAME]")
		return
	}
	contentDirectory := config.contentDirectory
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

// Builds the Repose site based on the current project default values.
// It uses command-line flags to modify the root directory and config file.
// If there is an error parsing the command flags, it prints an error message.
func (c *Command) Build(config Config) {

	logger.Info("Building site from %s with config %s\n", *&c.rootPath, *&c.configPath)
}

// Starts serving the Repose site for local preview.
func (c *Command) Preview(config Config) string {
	fmt.Printf("Repose site")
	return ""
}

// Updates the Repose binary in the current directory
func (c *Command) Update() string {
	fmt.Printf("Repose update placeholder")
	return ""
}

func (c *Command) Help() string {
	response := `Repose Commands:
Usage: repose [OPTIONS] <COMMAND>

Commands:
	init    - Initialize a new Repose project
	new     - Create new content. Usage: repose new [CONTENTTYPE] [FILENAME]
	build   - Build the site.
	preview - Setup a local server to preview the site
	demo    - Generate demo content
	update  - Update the repose binary
	help    - Show this help message 
	
Options:
	-r, --root <ROOT> Directory to use as root of project (default: .)
	-c, --config <CONFIG> Path to configuration file (default: config.toml)
`
	logger.Info(response)
	return ""
}

func (c *Command) SetRootPath(path string) {
	c.rootPath = path
}

func (c *Command) SetConfigPath(path string) {
	c.configPath = path
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
	content = strings.Replace(content, "{author}", config.author, -1)

	return content
}

func createDirectoryStructure(rootPath string) error {
	for _, dir := range []string{"content", "template", "web"} {
		if err := os.Mkdir(filepath.Join(rootPath, dir), 0755); err != nil {
			return err
		}
	}

	files := []struct {
		Name    string
		Content []byte
	}{
		{"config.yml", []byte("url = \"http://localhost:8080\"\ntitle = \"My website\"\n")},
		{"template/default.html", []byte("<!DOCTYPE html>\n<head>\n\t<title>{{ .Title }}</title>\n</head>\n<body>\n{{ .Content }}\n</body>\n</html>")},
		{"content/index.md", []byte("+++\ntitle = \"Repose!\"\n+++\n\nWelcome to my website.\n")},
	}
	for _, f := range files {
		if err := os.WriteFile(filepath.Join(rootPath, f.Name), f.Content, 0655); err != nil {
			return err
		}
	}

	return nil
}
