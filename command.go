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
	if err := c.createNewProjectFiles(c.rootPath); err != nil {
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
	fileName, title := c.processFileName(fileNameParam)

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
	content := c.defaultContent(contentType, title)

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
	response := HelpText
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

func (c *Command) processFileName(fileName string) (string, string) {
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
func (c *Command) defaultContent(contentType string, title string) string {
	content := NewMD

	// Replace placeholders with actual values
	content = strings.Replace(content, "{title}", title, -1)
	content = strings.Replace(content, "{contentType}", contentType, -1)
	content = strings.Replace(content, "{author}", config.author, -1)

	return content
}

func (c *Command) createNewProjectFiles(rootPath string) error {
	// Create the project directory structure
	logger.Info("Creating new project in %s", rootPath)
	dirs := []string{"content", "template", "web"}
	for _, dir := range dirs {
		dirPath := filepath.Join(rootPath, dir)
		if _, err := os.Stat(dirPath); os.IsNotExist(err) {
			if err := os.Mkdir(dirPath, 0755); err != nil {
				return err
			}
		}
	}

	// Create the default files
	files := []struct {
		Name    string
		Content string
	}{
		{"config.yml", DefaultConfig},
		{"template/default.html", NewMD},
		{config.contentDirectory + "/index.md", DefaultHTML},
	}
	for _, f := range files {
		filePath := filepath.Join(rootPath, f.Name)
		if err := filesystem.Create(filePath, f.Content); err != nil {
			return err
		}
	}

	return nil
}
