package command

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/rlnorthcutt/repose/internal/config"
	"github.com/rlnorthcutt/repose/internal/filesystem"
	"github.com/rlnorthcutt/repose/pkg/logger"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Defining a new public type 'Command'
type Command struct {
	args     []string       // arguments passed to the command
	config   *config.Config // loaded configuration struct
	logger   *logger.Logger // Loaded logger struct
	rootPath string         // rootPath for the project
}

// Defining a global varaiable for Command
var command Command

// **********  Public Command Methods  **********

// dispatchCommand will take the command name and dispatch it to the correct function
func (c *Command) Dispatch(args []string, logger *logger.Logger, rootPath string) {

	// Load the config file
	configurator := config.New(c.logger)
	logger.Detail("Loading the config file")
	config, err := configurator.Load(rootPath)
	if err != nil {
		logger.Error("Error loading the config file: %s", err)
		logger.Detail("Check your config file for errors.")
		os.Exit(1)
	}

	// Set up the command struct
	c.args = args
	c.config = &config
	c.logger = logger
	c.rootPath = rootPath

	// Get the command name
	commandName := args[0]

	// Dispatch the command
	switch commandName {
	case "init":
		c.Init()
	case "new":
		c.New()
	case "demo":
		c.Demo()
	case "build":
		c.Build()
	case "preview":
		c.Preview()
	case "update":
		c.Update()
	case "help":
		c.Help()
	default:
		c.logger.Error("Unknown command: %s\n", os.Args[1])
	}
}

// Initializes a new Repose project.
// It creates the proper folder structure and starter files.
func (c *Command) Init() string {
	configFile := c.config.FileName()
	if c.rootPath != "" {
		configFile = filepath.Join(c.rootPath, configFile)
	}

	// Check if the config.yml file already exists
	fs := filesystem.New(c.logger)
	if fs.Exists(configFile) {
		c.logger.Fatal("Warning: The config file exists at %s. Please choose a new root directory.", configFile)
	}

	// Create the project files
	if err := initCommand.CreateNewProjectFiles(c.rootPath); err != nil {
		c.logger.Fatal("Error creating site structure: ", err)
	}
	return ""
}

// Creates new content based on the provided content type and filename.
// It requires two arguments: content type and filename.
// The content type defines the path, so it can also include a subfolder
func (c *Command) New() {
	if len(c.args) < 3 {
		c.logger.Warn("Missing arguments. Usage: \033[1;96m`repose new [CONTENTTYPE] [FILENAME]`\033[0m ")
		c.logger.Detail("CONTENTTYPE can be 'post', 'page', or a custom type. This will create a new directory in the content folder.")
		c.logger.Detail("FILENAME cannot contain spaces and will be used as the title of the content and the filename.")
		c.logger.Detail("Example: `repose new post my-first-post`")
		return
	} else if len(c.args) > 3 {
		c.logger.Warn("File name cannot contain spaces. Usage: \033[1;96m`repose new [CONTENTTYPE] [FILENAME]`\033[0m ")
		c.logger.Detail("FILENAME cannot contain spaces and will be used as the title of the content and the filename.")
		return
	}

	typeDirectory := c.args[1]
	fileNameParam := c.args[2]

	if err := c.createNewContent(typeDirectory, fileNameParam); err != nil {
		c.logger.Error(err.Error())
	}
}

// Generates a new project with demo content and templates to create a new site.
// @TODO: create demo content so this works
func (c *Command) Demo() string {
	c.logger.Info("Generating demo content")
	return ""
}

// Builds the Repose site based on the current project default values.
// It uses command-line flags to modify the root directory and config file.
// If there is an error parsing the command flags, it prints an error message.
func (c *Command) Build() {
	c.logger.Info("Building site from %s with %s", c.rootPath, c.config.ConfigFile)
	if err := buildCommand.BuildSite(); err != nil {
		c.logger.Fatal("Error building site:", err)
	}
	c.logger.Success("Site built successfully")
}

// Starts serving the Repose site for local preview.
func (c *Command) Preview() {
	c.logger.Info("Setting up the local preview server")

	// Define the directory to serve.
	if c.rootPath == "" {
		c.rootPath = "."
	}
	webDir := filepath.Join(c.rootPath, c.config.OutputDirectory)

	// Setup the HTTP server.
	fs := http.FileServer(http.Dir(webDir))
	http.Handle("/", fs)

	// Start the server in a new goroutine so it doesn't block opening the browser.
	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			c.logger.Error("Error starting server:", err)
			panic(err)
		}
	}()

	c.logger.Info("Preview server ready at %s/index.html", c.config.PreviewURL)
	c.logger.Detail("Press Ctrl+C to stop the server")

	// Give the server a moment to start.
	time.Sleep(500 * time.Millisecond)

	// Open the browser.
	c.openBrowser(c.config.PreviewURL + "/index.html")

	// Keep the server running.
	select {}
}

// Updates the Repose binary in the current directory
func (c *Command) Update() string {
	fmt.Printf("Repose update placeholder")
	return ""
}

// Displays the help text for the command-line tool
func (c *Command) Help() {
	response := c.coloredLogo(AsciiLogo) + "\n\n" + HelpText
	fmt.Print(response)
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
	content = strings.Replace(content, "{author}", c.config.Author, -1)

	return content
}

// createNewContent creates a new content file in the specified directory.
// @TODO refactor this and break into 2-3 methods
func (c *Command) createNewContent(typeDirectory, fileNameParam string) error {
	fileName, title := c.processFileName(fileNameParam)
	filesystem := filesystem.New(c.logger)

	// Determine content type
	contentType := typeDirectory
	if strings.Contains(typeDirectory, "/") {
		contentType = strings.Split(typeDirectory, "/")[0]
	}

	// Construct the path
	c.logger.Info("Creating new %s in %s", contentType, c.config.ContentDirectory)
	path := filepath.Join(c.config.ContentDirectory, typeDirectory, fileName)

	// Get default content
	// @TODO: change thsi to use a yml file for the metadata like default.yml or post.yml
	content := c.defaultContent(contentType, title)

	// Create the file or directory
	if err := filesystem.Create(path, content); err != nil {
		return fmt.Errorf("failed to create %s: %v", path, err)
	}

	c.logger.Success("Successfully created new %s: %s", contentType, path)

	// Check if the template exists
	templateName := contentType + ".tmpl"
	found, err := filesystem.ExistsRecursive(templateName, "template")
	if err != nil {
		fmt.Println("Error searching for template:", err)
		return nil
	}

	// Ask the user to create the template file if it doesn't exist
	if !found {
		c.logger.Warn("Template file not found: %s", templateName)
		c.logger.Detail("Do you want to create this template? (Yes/no)")

		// Read the user's response
		reader := bufio.NewReader(os.Stdin)
		response, err := reader.ReadString('\n')
		if err != nil {
			c.logger.Error("Error reading response:", err)
			return nil
		}

		// Trim whitespace and newline character
		response = strings.ToLower(strings.TrimSpace(response))
		// If yes, then create the template file
		if response == "yes" || response == "" {
			c.logger.Info("Creating template file: %s", templateName)
			path := "template/" + templateName
			// @TODO: Change this to use the default.tmpl file
			template := DefaultTemplate_none
			// Replace "default.tmpl" with the value of templateName
			labeledTemplate := strings.Replace(template, "default.tmpl", templateName, -1)
			if err := filesystem.Create(path, labeledTemplate); err != nil {
				c.logger.Error("Error creating template:", err)
				return nil
			}
			c.logger.Success("Template created successfully.")
		}
	}

	// Check if the editor is set and not empty, then open the file with it
	if c.config.Editor != "" && c.config.Editor != "none" {
		if err := c.openFileInEditor(c.config.Editor, path); err != nil {
			// Log the error but do not fail the entire operation
			c.logger.Error("Failed to open file in editor: %v", err)
		}
	}

	return nil
}

// openFileInEditor opens the specified file in the given editor.
func (c *Command) openFileInEditor(editor, filePath string) error {
	c.logger.Info("Opening file in editor: %s", editor)
	// Pause for a moment before opening the editor
	time.Sleep(500 * time.Millisecond)

	cmd := exec.Command(editor, filePath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// openBrowser tries to open the browser with a given URL.
func (c *Command) openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	if err != nil {
		c.logger.Error("Failed to open the browser: %v\n", err)
	}
}

func (c *Command) coloredLogo(text string) string {
	colors := []string{
		"\033[34m", // Blue
		"\033[36m", // Cyan
		"\033[33m", // Yellow
		"\033[37m", // Light grey
	}

	lines := strings.Split(text, "\n")
	coloredLines := make([]string, len(lines))

	for i, line := range lines {
		// Calculate color index based on cycle and line position within the cycle
		colorIndex := (i / 2) % len(colors)
		coloredLines[i] = colors[colorIndex] + line + "\033[0m"
	}

	return strings.Join(coloredLines, "\n")
}
