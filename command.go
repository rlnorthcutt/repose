package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

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
// @TODO: modffy this to ask the user for config values
func (c *Command) Init() string {
	if err := c.createNewProjectFiles(c.rootPath); err != nil {
		log.Fatal("Error creating site structure: ", err)
	}
	return ""
}

// Creates new content based on the provided content type and filename.
// It requires two arguments: content type and filename.
// The content type defines the path, so it can also include a subfolder
func (c *Command) New(config Config) {
	if len(os.Args) < 4 {
		logger.Warn("Missing arguments. Usage: repose new [CONTENTTYPE] [FILENAME]")
		return
	} else if len(os.Args) > 4 {
		logger.Warn("File name cannot contain spaces. Usage: repose new [CONTENTTYPE] [FILENAME]")
		return
	}

	typeDirectory := os.Args[2]
	fileNameParam := os.Args[3]

	if err := c.createNewContent(config, typeDirectory, fileNameParam); err != nil {
		logger.Error(err.Error())
	}
}

// Generates a new project with demo content and templates to create a new site.
// @TODO: create demo content so this works
func (c *Command) Demo() string {
	logger.Info("Generating demo content")
	return ""
}

// Builds the Repose site based on the current project default values.
// It uses command-line flags to modify the root directory and config file.
// If there is an error parsing the command flags, it prints an error message.
func (c *Command) Build(config Config) {
	logger.Info("Building site from %s with config %s\n", c.rootPath, c.configPath)
	if err := builder.BuildSite(); err != nil {
		fmt.Println("Error building site:", err)
	}
	logger.Success("Site built successfully")
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

// @TODO: see if we can now remove this
func (c *Command) SetRootPath(path string) {
	c.rootPath = path
}

func (c *Command) SetConfigPath(path string) {
	c.configPath = path
}

// **********  Private Command Methods  **********

// processFileName takes a name and returns the filename (with extension) and title.
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
	content = strings.Replace(content, "{author}", config.Author, -1)

	return content
}

// createNewProjectFiles creates the default files and directories for a new project.
func (c *Command) createNewProjectFiles(rootPath string) error {
	// Create the config file
	if err := c.initConfig(rootPath); err != nil {
		return err
	}

	// Set the output for the root path
	installDir := rootPath
	if rootPath == "" {
		installDir = "this directory"
	}

	// Create the project directory structure
	logger.Info("Creating new project in %s", installDir)
	dirs := []string{"content", "template", "web", "web/asset", "web/asset/css", "web/asset/js", "web/asset/img"}
	for _, dir := range dirs {
		dirPath := filepath.Join(rootPath, dir)
		if _, err := os.Stat(dirPath); os.IsNotExist(err) {
			if err := os.Mkdir(dirPath, 0755); err != nil {
				return err
			}
		}
	}

	// Create the default files
	indexMD := c.defaultContent("default", "Your homepage")
	files := []struct {
		Name    string
		Content string
	}{
		{"template/default.tmpl", DefaultTemplate},
		{"template/page.tmpl", PageTemplate},
		{"template/header.tmpl", HeaderTemplate},
		{"template/navigation.tmpl", NavigationTemplate},
		{"template/footer.tmpl", FooterTemplate},
		{"content/index.md", indexMD},
		{"content/test.md", MarkdownTest},
	}
	for _, f := range files {
		filePath := filepath.Join(rootPath, f.Name)
		cleanContent := strings.TrimSpace(f.Content)
		if err := filesystem.Create(filePath, cleanContent); err != nil {
			return err
		}
	}

	logger.Success("Repose project created in %s", installDir)

	return nil
}

// createNewContent creates a new content file in the specified directory.
func (c *Command) createNewContent(config Config, typeDirectory, fileNameParam string) error {
	fileName, title := c.processFileName(fileNameParam)

	// Determine content type
	contentType := typeDirectory
	if strings.Contains(typeDirectory, "/") {
		contentType = strings.Split(typeDirectory, "/")[0]
	}

	// Construct the path
	logger.Info("Creating new %s in %s", contentType, config.ContentDirectory)
	path := filepath.Join(config.ContentDirectory, typeDirectory, fileName)

	// Get default content
	content := c.defaultContent(contentType, title)

	// Create the file or directory
	if err := filesystem.Create(path, content); err != nil {
		return fmt.Errorf("failed to create %s: %v", path, err)
	}

	logger.Success("Successfully created new %s: %s\n", contentType, path)

	// Check if the template exists
	templateName := contentType + ".tmpl"
	found, err := filesystem.ExistsRecursive(templateName, "template")
	if err != nil {
		fmt.Println("Error searching for template:", err)
		return nil
	}

	// Ask the user to create the template file if it doesn't exist
	if !found {
		logger.Warn("Template file not found: %s\n", templateName)
		fmt.Println("Do you want to create this template? (yes/no)")

		// Read the user's response
		reader := bufio.NewReader(os.Stdin)
		response, err := reader.ReadString('\n')
		if err != nil {
			logger.Error("Error reading response:", err)
			return nil
		}

		// Trim whitespace and newline character
		response = strings.TrimSpace(response)
		// If yes, then create the template file
		if strings.ToLower(response) == "yes" {
			logger.Info("Creating template file: %s", templateName)
			path := "template/" + templateName
			if err := filesystem.Create(path, DefaultTemplate); err != nil {
				logger.Error("Error creating template:", err)
				return nil
			}
		}
		logger.Success("Template created successfully.")
	}

	// Check if the editor is set and not empty, then open the file with it
	if config.Editor != "" || config.Editor == "none" {
		if err := c.openFileInEditor(config.Editor, path); err != nil {
			// Log the error but do not fail the entire operation
			logger.Error("Failed to open file in editor: %v", err)
		}
	}

	return nil
}

// openFileInEditor opens the specified file in the given editor.
func (c *Command) openFileInEditor(editor, filePath string) error {
	logger.Info("Opening file in editor: %s", editor)
	// Pause for 1 second before opening the editor
	time.Sleep(1 * time.Second)

	cmd := exec.Command(editor, filePath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (c *Command) initConfig(installDir string) error {
	sitename := c.promptForInput("Enter the site name", "Repose site")
	author := c.promptForInput("Enter the author's name", "Creator")
	editor := c.promptForInput("Enter the editor", "nano")
	contentDirectory := c.promptForInput("Enter the content directory", "content")
	outputDirectory := c.promptForInput("Enter the output directory", "web")
	url := c.promptForInput("Enter the site URL", "mysite.com")
	previewUrl := c.promptForInput("Enter the preview URL", "http://localhost:8080")

	configContent := fmt.Sprintf(`sitename: %s
author: %s
editor: %s
contentDirectory: %s
outputDirectory: %s
url: %s
previewUrl: %s
`, sitename, author, editor, contentDirectory, outputDirectory, url, previewUrl)

	// Create the filepath
	if installDir != "" {
		installDir = installDir + "/"
	}
	configPath := installDir + "config.yml"

	// Write the configContent to config.yml
	if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
		fmt.Println("Error writing config file:", err)
		return err
	}

	fmt.Println("Config initialized successfully.")
	return nil
}

func (c *Command) promptForInput(prompt, defaultValue string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s [%s]: ", prompt, defaultValue)

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return defaultValue
	}

	input = strings.TrimSpace(input)
	if input == "" {
		return defaultValue
	}

	return input
}
