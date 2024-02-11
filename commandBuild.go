package main

import (
	"bytes"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
)

// Defining a new public type 'Template'
type Builder int

// Defining a global varaiable for build command
var buildCommand Builder

// Holds information about a directory during processing
// Used to process the content directory structure
type DirectoryInfo struct {
	Path     string
	NumFiles int
	HasIndex bool
	Files    []FileInfo
}

// Holds information about a file during processing
// Used to process the content files
type FileInfo struct {
	Name        string
	Path        string
	FileType    string
	ContentType string
	Metadata    map[string]interface{}
}

// PageData holds data to pass into templates
// This is used to build the full page content
type PageData struct {
	SiteName        string
	Logo            template.HTML
	Title           string
	MdContent       template.HTML
	TemplateFile    string
	TemplateContent template.HTML
	Metadata        map[string]interface{}
}

var templates *template.Template

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
		panic(err)
	}

	// Process the files
	err = b.processFiles(dirsMap)
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

// **********  Private Command Methods  **********

// Walk the content directory and build the files and dirs maps
// Returns a map of files and a map of directories
func (b *Builder) walkContentDir() (map[string]DirectoryInfo, error) {
	// Create maps to hold the files and directories
	dirsMap := make(map[string]DirectoryInfo)

	// Walk the content directory and build the files and dirs maps
	// We update both maps as we walk the directory one time
	err := filepath.Walk(command.contentDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Get the relative path to the root directory
		relPath, err := filepath.Rel(command.rootPath, path)
		if err != nil {
			return err
		}

		// Split the path into directory and file name
		dir, fileName := filepath.Split(relPath)
		dir = strings.TrimSuffix(dir, "/") // Clean up trailing slash

		if info.IsDir() { // Process the directory
			// This should run the first time we encounter a directory
			if _, exists := dirsMap[dir]; !exists {
				dirsMap[dir] = DirectoryInfo{
					Path:     dir,
					NumFiles: 0,
					HasIndex: false,
					Files:    []FileInfo{},
				}
			}
		} else { // Process the file
			// Determine file type and content type
			fileType := filepath.Ext(fileName)
			fileType = strings.ToLower(fileType)
			fileType = strings.TrimPrefix(fileType, ".")
			// The first subdrectory under content is the content type
			contentType := strings.SplitN(dir, "/", 2)[0]

			// Only process md and html files
			if fileType == "md" || fileType == "html" {
				// Create a new FileInfo struct
				fileInfo := FileInfo{
					Name:        fileName,
					Path:        relPath,
					FileType:    fileType,
					ContentType: contentType,
				}

				// Update directory info
				dirInfo := dirsMap[dir]
				// Increment the number of files in the directory
				dirInfo.NumFiles++
				// Add the file to the directory info
				dirInfo.Files = append(dirInfo.Files, fileInfo)
				// Check if the file is an index file
				if fileName == "index.md" || fileName == "index.html" {
					dirInfo.HasIndex = true
				}

				dirsMap[dir] = dirInfo
			}
		}

		return nil
	})

	return dirsMap, err
}

func (b *Builder) processFiles(dirsMap map[string]DirectoryInfo) error {
	// Create a new markdown parser with the meta extension
	markdown := goldmark.New(
		goldmark.WithExtensions(
			meta.Meta,
		),
	)

	// Loop through each directory in dirsMap
	for _, dirInfo := range dirsMap {
		// Loop through each file in the directory
		for _, file := range dirInfo.Files {
			// Read the MD file and process it
			content, err := filesystem.Read(file.Path)
			if err != nil {
				return err
			}

			// Get the metadata from the markdown file
			var buf bytes.Buffer
			context := parser.NewContext()
			if err := markdown.Convert([]byte(content), &buf, parser.WithContext(context)); err != nil {
				return err
			}
			metaData := meta.Get(context)
			htmlContent := buf.String()

			// Remove the "content/" prefix from the file path so we can replace
			// it with the output directory
			trimmedPath := strings.TrimPrefix(file.Path, "content/")
			outputPath := filepath.Join(command.outputDir, trimmedPath)
			outputPath = strings.TrimSuffix(outputPath, filepath.Ext(outputPath)) + ".html"

			// Write the HTML content to the output directory
			if err := b.renderAndWriteFile(outputPath, htmlContent, metaData); err != nil {
				return err
			}
		}
	}

	return nil
}

func (b *Builder) renderAndWriteFile(outputPath string, contentHTML string, metaData map[string]interface{}) error {
	// Extract the template name from outputPath or set a default
	templateFile := metaData["template"].(string)
	if templateFile == "" {
		templateFile = "default.tmpl"
	}

	// Build PageData
	pageData := PageData{
		SiteName:     config.Sitename,
		Logo:         logo50,
		Title:        metaData["title"].(string),
		MdContent:    template.HTML(contentHTML),
		TemplateFile: templateFile,
		Metadata:     metaData,
	}

	// Process the MD content with the template
	// This will be used to process the full page from the template
	templateContent, err := b.getTemplateContent(pageData)
	if err != nil {
		return err
	}
	pageData.TemplateContent = templateContent

	// Execute the full page template with the built PageData
	var output bytes.Buffer
	if err := templates.ExecuteTemplate(&output, "fullpage.tmpl", pageData); err != nil {
		return err
	}

	// Use filesystem.Create to write the output to the specified path
	// Assuming filesystem.Create takes a string path and byte slice as content
	return filesystem.Create(outputPath, output.String())
}

func (b *Builder) buildIndexFiles(dirsMap map[string]DirectoryInfo) error {
	logger.Info("Placeholder for building index files")
	return nil
}

// Process the content in the pageData struct to generate tempalted contend
func (b *Builder) getTemplateContent(pageData PageData) (template.HTML, error) {
	// Process the template in the metsdata with the content in the metadata
	var tmplContent bytes.Buffer
	if err := templates.ExecuteTemplate(&tmplContent, pageData.TemplateFile, pageData); err != nil {
		return "", err
	}

	return template.HTML(tmplContent.String()), nil
}

// Parse the templates and store them in a global variable
func (b *Builder) initTemplates() error {
	var err error
	templateDir := command.templateDir
	templates, err = template.ParseGlob(filepath.Join(templateDir, "*.tmpl"))
	return err
}
