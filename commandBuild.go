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

	// Initialize a map to track directories
	dirMap := make(map[string][]os.FileInfo)

	// Walk the content directory
	err = filepath.Walk(contentDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip the root content directory itself
		if path == contentDir {
			return nil
		}

		dir := filepath.Dir(path)
		if info.IsDir() {
			// Ensure the directory is tracked
			if _, exists := dirMap[dir]; !exists {
				dirMap[dir] = []os.FileInfo{}
			}
			return nil
		}

		// Track files in their respective directories
		dirMap[dir] = append(dirMap[dir], info)

		// Process only markdown files
		if strings.HasSuffix(path, ".md") {
			fmt.Println("Processing", path)
			return b.processMarkdownFile(path, contentDir, webDir, tmpl)
		}

		return nil
	})
	if err != nil {
		return err
	}

	// Generate index.html for directories without one
	for dir, files := range dirMap {
		if !b.hasIndexFile(files) {
			fmt.Println("Generating index.html for", dir)
			if err := b.generateIndexHTML(dir, files); err != nil {
				return err
			}
		}
	}

	return nil
}

// hasIndexFile checks if the given slice of FileInfo contains an index file.
func (b Builder) hasIndexFile(files []os.FileInfo) bool {
	for _, file := range files {
		if file.Name() == "index.md" || file.Name() == "index.html" {
			return true
		}
	}
	return false
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

func (b *Builder) generateIndexHTML(dirPath string, items []os.FileInfo) error {
	// Calculate the relative path of dirPath from the content directory
	// @TODO : change command rootpath to always default with trailing slash
	// @TODO : check where paths are being created, and change them to filepath.join
	if command.rootPath == "" {
		command.rootPath = "."
	}
	contentDir := filepath.Join(command.rootPath, "content")
	relPath, err := filepath.Rel(contentDir, dirPath)
	if err != nil {
		return err // Handle the error appropriately
	}

	// Construct the output path by joining the root path, output directory, and the relative path
	outputPath := filepath.Join(command.rootPath, config.OutputDirectory, relPath)

	// Ensure the output directory exists
	if err := os.MkdirAll(outputPath, 0755); err != nil {
		return err
	}

	indexPath := filepath.Join(outputPath, "index.html")
	var links []string

	for _, item := range items {
		if item.IsDir() {
			continue // Skip directories
		}
		itemName := item.Name()
		// Assuming you convert .md files to .html
		if strings.HasSuffix(itemName, ".md") {
			itemName = strings.TrimSuffix(itemName, ".md") + ".html"
		}
		links = append(links, "<li><a href=\""+itemName+"\">"+itemName+"</a></li>")
	}

	htmlContent := "<html><body><ul>" + strings.Join(links, "") + "</ul></body></html>"
	return os.WriteFile(indexPath, []byte(htmlContent), 0644)
}
