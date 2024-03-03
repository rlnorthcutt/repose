package main

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
)

// Controls the build command
type Builder struct {
	rootPath    string
	contentDir  string
	outputDir   string
	templateDir string
	templates   *template.Template
	dirsMap     map[string]DirectoryInfo
}

// Defining a global varaiable for build command
var buildCommand Builder

// Holds information about a directory during processing
// Keyed by the full path to the directory
type DirectoryInfo struct {
	Path     string     // The relative path to the content directory
	NumFiles int        // The number of files in the directory
	HasIndex bool       //	Whether the directory has an index file
	Files    []FileInfo // A slice of FileInfo structs for each file in the directory
}

// Holds information about a file during processing
// Keyed by the full path to the file
type FileInfo struct {
	Name        string                 // The name of the file (no extension)
	Path        string                 // The relative path to the file relative to the content directory
	OutputPath  string                 // The path and file name for the output file
	FileType    string                 // The type of file (e.g. "md", "html")
	ContentType string                 // The type of content (e.g. "page", "post", "project")
	MetaData    map[string]interface{} // Metadata extracted from the file
	Content     template.HTML          // The content of the file
}

// PageData holds data to pass into templates
// This is used to build the full page content
type PageData struct {
	SiteName string                 // The name of the site
	Logo     template.HTML          // The site logo
	Title    string                 // The title of the page
	Content  template.HTML          // The content of the page
	Metadata map[string]interface{} // Metadata for the page
}

// **********  Public Command Methods  **********

// Generates the site from the content and template files
func (b *Builder) BuildSite() error {
	// Initialize the templates
	err := b.initTemplates()
	if err != nil {
		return err
	}

	// Generate the files and dirs for the content directory
	dirsMap, err := b.walkContentDir()
	if err != nil {
		logger.Error("Error walking content directory: ", err)
		return err
	}

	// Reset the output directory before writing new files
	// @TODO: refactor to only delete files and directories that need to be deleted
	b.resetOutputDirectory()

	// Render the files
	// @TODO: refactor to only render files that need to be rendered
	err = b.renderFiles(dirsMap)
	if err != nil {
		return err
	}

	// Build index files
	err = b.buildIndexFiles(dirsMap)
	if err != nil {
		return err
	}

	return nil
}

// Set the root path and comomon directories for commands
func (b *Builder) SetRootPath(path string) {
	if path == "" {
		path = "."
	}
	b.rootPath = path
	b.contentDir = filepath.Join(path, config.ContentDirectory)
	b.outputDir = filepath.Join(path, config.OutputDirectory)
	b.templateDir = filepath.Join(path, "template")
}

// **********  Private Command Methods  **********

// Walk the content directory and build the files and dirs maps
// Returns a map of files and a map of directories
func (b *Builder) walkContentDir() (map[string]DirectoryInfo, error) {
	// Create maps to hold the files and directories
	b.dirsMap = make(map[string]DirectoryInfo)

	contentPath := filepath.Join(b.rootPath, b.contentDir)
	logger.Detail("Walking content directory: " + contentPath)

	// Walk the content directory and build the files and dirs maps
	// We update both maps as we walk the directory one time
	err := filepath.Walk(contentPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("error accessing path %q: %v", path, err)
		}

		logger.Detail("Processing path: " + path)

		// Check if this is a directory or a file
		isDir, err := filesystem.IsDir(path)
		if err != nil {
			return err
		}

		if isDir {
			// Process the directory
			if err := b.processDir(path); err != nil {
				return err
			}
		} else {
			// Process the file
			if err := b.processFile(path); err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err // More descriptive error handling
	}

	return b.dirsMap, nil
}

// Process a single directory, updating the directory information map.
func (b *Builder) processDir(path string) error {
	// Check if the directory is already in the map
	if _, exists := b.dirsMap[path]; !exists {
		// Get the relative path from the root directory using filepath.Rel
		relPath, err := filepath.Rel(b.contentDir, path)
		if err != nil {
			return fmt.Errorf("error getting relative path: %s, error: %v", path, err)
		}

		// Add the directory to the map
		b.dirsMap[path] = DirectoryInfo{
			Path:     relPath,
			NumFiles: 0,
			HasIndex: false,
			Files:    []FileInfo{},
		}
	}

	return nil
}

// processFile processes a single file, updating the directory information map.
func (b *Builder) processFile(path string) error {
	relPath, dir, contentType, fileName, fileType, err := filesystem.GetFileInfo(b.contentDir, path)
	if err != nil {
		return fmt.Errorf("error getting file info for %q: %v", path, err)
	}

	if contentType == "" {
		contentType = "page"
	}

	// Process the markdown file to extract HTML content and metadata
	renderedContent, metaData, err := b.processMarkdown(path)
	if err != nil {
		return fmt.Errorf("error processing markdown for %q: %v", relPath, err)
	}

	// Create the FileInfo struct
	fileInfo := FileInfo{
		Name:        fileName,
		Path:        relPath,
		OutputPath:  "/" + filepath.Join(dir, fileName+".html"),
		FileType:    fileType,
		ContentType: contentType,
		MetaData:    metaData,
		Content:     template.HTML(renderedContent),
	}

	// Update the directory info with the new file
	dirKey := filepath.Join(b.contentDir, dir)
	// Process the directory
	// @TODO: we are processing the directory twice - once here and once in processDir
	// This is ok for now since we do a check against the map, but can we set this up better?
	if err := b.processDir(dirKey); err != nil {
		return err
	}
	// Load the directory object from the map
	dirInfo, exists := (b.dirsMap)[dirKey]
	if !exists {
		return fmt.Errorf("directory %q not found in directory map", dir)
	}
	// Update the directory object with the new file
	dirInfo.NumFiles++
	dirInfo.Files = append(dirInfo.Files, fileInfo)
	if fileName == "index" {
		dirInfo.HasIndex = true
	}

	(b.dirsMap)[dirKey] = dirInfo

	return nil
}

// Render the files and write them to the output directory
func (b *Builder) renderFiles(dirsMap map[string]DirectoryInfo) error {
	// Loop through each directory in dirsMap
	for _, dirInfo := range dirsMap {
		// Loop through each file in the directory
		for _, file := range dirInfo.Files {
			// Remove the "content/" prefix from the file path so we can replace
			// it with the output directory
			trimmedPath := strings.TrimPrefix(file.Path, "content/")
			outputPath := filepath.Join(b.outputDir, trimmedPath)
			outputPath = strings.TrimSuffix(outputPath, filepath.Ext(outputPath)) + ".html"

			// Write the HTML content to the output directory
			if err := b.renderAndWriteFile(outputPath, file); err != nil {
				return err
			}
		}
	}

	return nil
}

// Process the markdown file and extract metadata
// @TODO: we have to use this twice - once for markdown and once for HTML, so lets memoize it
func (b *Builder) processMarkdown(filePath string) (htmlContent string, metaData map[string]interface{}, err error) {
	// @TODO: see if we need to adjust this for HTML files
	// @TODO: for html files - what about the metadata?
	// Create a new markdown parser with the meta extension
	markdown := goldmark.New(
		goldmark.WithExtensions(
			meta.Meta,
		),
	)

	// Read the MD file and process it
	content, err := filesystem.Read(filePath)
	if err != nil {
		return "", nil, fmt.Errorf("error reading markdown file %s: %w", filePath, err)
	}

	// Get the metadata from the markdown file
	var buf bytes.Buffer
	context := parser.NewContext()
	if err := markdown.Convert([]byte(content), &buf, parser.WithContext(context)); err != nil {
		return "", nil, fmt.Errorf("error converting markdown to HTML: %w", err)
	}

	// Extract metadata with type assertion
	metaDataMap := meta.Get(context)
	if metaDataMap == nil {
		// Handle the case where metadata is not present or not in the expected format
		metaDataMap = make(map[string]interface{}) // Initialize as empty if not present
	}

	htmlContent = buf.String()

	return htmlContent, metaDataMap, nil
}

// Render the HTML content with the template and write to the output directory
func (b *Builder) renderAndWriteFile(outputPath string, file FileInfo) error {
	// Extract the template name from outputPath or set a default
	templateFile := file.MetaData["template"].(string)
	if templateFile == "" {
		templateFile = "default.tmpl"
	}

	// Process the MD content with the template
	// This will be used to process the full page from the template
	templateContent, err := b.getTemplateContent(file, templateFile)
	if err != nil {
		return err
	}

	// Build PageData
	pageData := PageData{
		SiteName: config.Sitename,
		Logo:     logo50,
		Title:    file.MetaData["title"].(string),
		Content:  template.HTML(templateContent),
		Metadata: file.MetaData,
	}

	// Execute the full page template with the built PageData
	var output bytes.Buffer
	if err := b.templates.ExecuteTemplate(&output, "fullpage.tmpl", pageData); err != nil {
		return err
	}

	// Use filesystem.Create to write the output to the specified path
	// Assuming filesystem.Create takes a string path and byte slice as content
	return filesystem.Create(outputPath, output.String())
}

func (b *Builder) buildIndexFiles(dirsMap map[string]DirectoryInfo) error {
	logger.Info("Building index files")
	// Loop through each directory in dirsMap
	for contentPath, dirInfo := range dirsMap {
		logger.Detail("Processing directory: " + contentPath)
		// If the directory does not have an index file, create one
		if !dirInfo.HasIndex && dirInfo.NumFiles > 0 {

			// Get the content type from the first file in the directory
			contentType := dirInfo.Files[0].ContentType
			if contentType == "content" {
				continue
			}
			logger.Detail("Building index file for " + contentType + "s")

			// Generate the list content for the index file
			var listContent bytes.Buffer
			if err := b.templates.ExecuteTemplate(&listContent, "list.tmpl", dirInfo); err != nil {
				return err
			}

			// Build PageData for the full page
			pageData := PageData{
				SiteName: config.Sitename,
				Logo:     logo50,
				Title:    "All " + contentType + "s",
				Content:  template.HTML(listContent.String()),
				Metadata: dirInfo.Files[0].MetaData,
			}

			// Execute the full page template with the built PageData
			var output bytes.Buffer
			if err := b.templates.ExecuteTemplate(&output, "fullpage.tmpl", pageData); err != nil {
				logger.Error("Error executing template: ", err)
				return err
			}

			// Remove the "content/" prefix from the file path so we can replace
			// it with the output directory
			trimmedPath := strings.TrimPrefix(dirInfo.Path, "content/")
			outputPath := filepath.Join(b.outputDir, trimmedPath, "index.html")
			logger.Detail("Writing index file to " + outputPath)

			if err := filesystem.Create(outputPath, output.String()); err != nil {
				return err
			}
		}
	}
	return nil
}

// Process the content in the pageData struct to generate templated contend
func (b *Builder) getTemplateContent(file FileInfo, templateFile string) (template.HTML, error) {
	// Process the template in the metadata with the content in the metadata
	var tmplContent bytes.Buffer
	if err := b.templates.ExecuteTemplate(&tmplContent, templateFile, file); err != nil {
		return "", err
	}

	return template.HTML(tmplContent.String()), nil
}

// Parse the templates and store them in a global variable
func (b *Builder) initTemplates() error {
	var err error
	b.templates, err = template.ParseGlob(filepath.Join(b.templateDir, "*.tmpl"))
	if err != nil {
		return fmt.Errorf("failed to load templates: %w", err)
	}
	return nil
}

// deletes all files and directories in the output folder except the assets directory.
// @TODO: this seems way too big and complex - find a better way to do this
func (b *Builder) resetOutputDirectory() error {
	logger.Info("Resetting output directory")
	webDir := b.outputDir
	assetsDir := filepath.Join(webDir, "assets")

	// Initialize a slice to keep track of directories
	// NOTE: we have to do this because if we delete the directory before the file
	// Then the walk will end and we will get an error trying to create the files
	// Later. So we collect the directories to delete and then delete them after
	var dirsToDelete []string

	// First pass: Delete files and collect directories
	err := filepath.Walk(webDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Normalize paths for comparison
		normalizedPath := filepath.ToSlash(path)
		normalizedAssetsDir := filepath.ToSlash(assetsDir)

		// Skip the assets directory and its contents
		if strings.HasPrefix(normalizedPath, normalizedAssetsDir+"/") {
			return nil
		}

		if info.IsDir() {
			// Collect directories for later deletion, skipping the webDir itself
			if normalizedPath != filepath.ToSlash(webDir) {
				dirsToDelete = append(dirsToDelete, path)
			}
		} else {
			// Delete the file
			return os.Remove(path)
		}

		return nil
	})

	if err != nil {
		return err
	}

	// Sort directories in reverse order to ensure we delete child directories before their parents
	sort.Sort(sort.Reverse(sort.StringSlice(dirsToDelete)))

	// Second pass: Attempt to delete collected directories
	for _, dir := range dirsToDelete {
		// Attempt to remove the directory (will fail if not empty)
		err := os.Remove(dir)
		if err != nil && !os.IsNotExist(err) {
			logger.Warn("Failed to delete directory (may not be empty): %s", dir)
			// Optionally, you can log the error or handle it as needed
		}
	}

	return nil
}
