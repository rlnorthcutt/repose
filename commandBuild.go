package main

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/russross/blackfriday/v2"
)

// PageData holds data to pass into templates
type PageData struct {
	SiteName string
	Title    string
	Content  template.HTML
	Template string
	Metadata map[string]interface{}
	Logo     template.HTML
}

// Defining a new public type 'Template'
type Builder int

// Defining a global varaiable for build command
var builder Builder

func (b *Builder) BuildSite() error {
	// Define your base directories
	contentDir := "./" + config.ContentDirectory
	webDir := "./" + config.OutputDirectory
	templateDir := "./template"

	// Parse templates
	tmpl, err := template.ParseGlob(filepath.Join(templateDir, "*.tmpl"))
	if err != nil {
		return err
	}

	// Walk the content directory
	return filepath.Walk(contentDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Process only markdown files
		if strings.HasSuffix(path, ".md") {
			fmt.Println("Processing", path)
			return b.processMarkdownFile(path, contentDir, webDir, tmpl)
		}

		return nil
	})
}

func (b *Builder) processMarkdownFile(filePath, contentDir, webDir string, tmpl *template.Template) error {
	// Read markdown file
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	// Split the content to separate front matter from Markdown content
	parts := bytes.SplitN(data, []byte("---"), 3)
	if len(parts) < 3 {
		return err
	}

	// Parse the front matter into a map
	// Assuming `parts[1]` contains the front matter as a string
	metadataString := string(parts[1]) // Ensure it's a string, adjust as necessary

	// Parse the front matter into a map using readYAML
	metadataMap, err := filesystem.ParseYml(metadataString)
	if err != nil {
		return err
	}

	// Convert map[string]string to map[string]interface{}
	metadata := make(map[string]interface{})
	for k, v := range metadataMap {
		metadata[k] = v
	}

	// Extract the title from metadata if it exists for consistency with PageData
	title, _ := metadata["title"].(string)
	templateFile, ok := metadata["template"].(string)
	if !ok {
		templateFile = "default.tmpl"
	}

	// Convert markdown content to HTML
	markdownContent := parts[2]
	htmlContent := blackfriday.Run(markdownContent)

	// Create a relative path for the output file
	relPath, err := filepath.Rel(contentDir, filePath)
	if err != nil {
		return err
	}
	outputPath := filepath.Join(webDir, strings.TrimSuffix(relPath, filepath.Ext(relPath))+".html")

	// Ensure the output directory exists
	if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
		return err
	}

	// Open output file
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// Apply the template
	pageData := PageData{
		SiteName: config.Sitename,
		Title:    title,
		Content:  template.HTML(htmlContent),
		Metadata: metadata,
		Template: templateFile,
		Logo:     template.HTML(logo50), // Convert the SVG string to template.HTML
	}

	return tmpl.ExecuteTemplate(outputFile, "page.tmpl", pageData)
}
