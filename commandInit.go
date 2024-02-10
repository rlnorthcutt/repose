package main

import (
	"os"
	"path/filepath"
	"strings"
)

type Init int

var initCommand Init

// Defining a new public type 'FileContent'
type FileContent struct {
	Name    string
	Content string
}

// **********  Public Command Methods  **********

// Creates the default files and directories for a new project.
func (i *Init) CreateNewProjectFiles(rootPath string) error {
	// Create the config file
	if err := config.Create(rootPath); err != nil {
		return err
	}

	// Set the output for the root path
	installDir := rootPath
	if rootPath == "" || rootPath == "." {
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

	// Load the new config file
	config, _ = config.Load()

	// Get the template constants and files names
	files := i.getTemplateContents(config)

	// Loop over the files and create them
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

// **********  Private Command Methods  **********

// Returns the default template contents based on the config
func (i *Init) getTemplateContents(config Config) []FileContent {
	// Make sure the theme is set to "none" if it's not "pico", "bootstrap", or "tailwind"
	if config.Theme != "pico" && config.Theme != "bootstrap" && config.Theme != "tailwind" {
		config.Theme = "none"
	}

	// Define the template themes
	templateThemes := map[string]map[string]string{
		"pico": {
			"default":    DefaultTemplate_pico,
			"page":       PageTemplate_pico,
			"header":     HeaderTemplate_pico,
			"navigation": NavigationTemplate_pico,
			"footer":     FooterTemplate_pico,
			"css":        css_pico,
		},
		"bootstrap": {
			"default":    DefaultTemplate_bootstrap,
			"page":       PageTemplate_bootstrap,
			"header":     HeaderTemplate_bootstrap,
			"navigation": NavigationTemplate_bootstrap,
			"footer":     FooterTemplate_bootstrap,
			"css":        css_bootstrap,
		},
		"tailwind": {
			"default":    DefaultTemplate_tailwind,
			"page":       PageTemplate_tailwind,
			"header":     HeaderTemplate_tailwind,
			"navigation": NavigationTemplate_tailwind,
			"footer":     FooterTemplate_tailwind,
			"css":        css_tailwind,
		},
		"none": {
			"default":    DefaultTemplate_none,
			"page":       PageTemplate_none,
			"header":     HeaderTemplate_none,
			"navigation": NavigationTemplate_none,
			"footer":     FooterTemplate_none,
			"css":        css_none,
		},
	}

	// Access the correct theme's templates
	themeTemplates := templateThemes[config.Theme]

	indexMD := command.defaultContent("default", "Your homepage")
	files := []FileContent{
		{"template/default.tmpl", themeTemplates["default"]},
		{"template/fullpage.tmpl", themeTemplates["page"]},
		{"template/header.tmpl", themeTemplates["header"]},
		{"template/navigation.tmpl", themeTemplates["navigation"]},
		{"template/footer.tmpl", themeTemplates["footer"]},
		{"content/index.md", indexMD},
		{"content/test.md", MarkdownTest},
		{"web/asset/css/styles.css", themeTemplates["css"]},
	}

	return files
}
