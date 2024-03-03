package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Defining a new public type 'Filesystem'
type Filesystem int

// Defining a global varaiable for Filesystem
var filesystem Filesystem

// **********  Public Filesystem methods  **********

// Check if the file (or directory) exists at the given path.
func (f *Filesystem) Exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// Check if a file with the given name exists in the given directory.
func (f *Filesystem) ExistsRecursive(fileName string, directory string) (bool, error) {
	var exists bool

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err // return any error encountered
		}

		// Check if the current file matches the fileName
		if filepath.Base(path) == fileName {
			exists = true
			return filepath.SkipDir // Stop walking as we've found the file
		}

		return nil
	})

	if err != nil {
		return false, err // return any error encountered during walking
	}

	return exists, nil
}

// Create a file or directory based on the given path and content.
func (f *Filesystem) Create(path string, content string) error {
	// Check if the path exists using checkPath
	if f.Exists(path) {
		errorMessage := path + " already exists"
		return errors.New(errorMessage)
	}

	// Determine if the path is for a file or directory and create accordingly
	if filepath.Ext(path) != "" {
		return f.createFile(path, content)
	}
	return f.createDirectory(path)
}

// Return the content of the file at the given path.
func (f *Filesystem) Read(path string) (string, error) {
	// Check if the path exists
	if !f.Exists(path) {
		warningMessage := path + " doesn't exist"
		logger.Warn(warningMessage)
		return "", errors.New(warningMessage)
	}

	// Read the file
	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// Remove a file or directory at the given path.
func (f *Filesystem) Delete(path string) error {
	// Check if the path exists
	if !f.Exists(path) {
		warningMessage := path + " doesn't exist"
		logger.Warn(warningMessage)
		return errors.New(warningMessage)
	}

	// Delete the file or directory
	return os.RemoveAll(path)
}

// Check if the given path is a directory.
func (f *Filesystem) IsDir(path string) (bool, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return false, err
	}

	absPath := filepath.Join(cwd, path)

	fileInfo, err := os.Stat(absPath)
	if err != nil {
		return false, err
	}
	return fileInfo.IsDir(), nil
}

// Basic YAML format parser that reads the content & returns a map of the data.
// We use this instead of loading the YAML modules to keep the size down
func (f *Filesystem) ParseYml(content string) (map[string]string, error) {
	// Check for empty content
	if content == "" {
		return nil, errors.New("parsing failed: empty yaml content")
	}

	scanner := bufio.NewScanner(strings.NewReader(content))
	ymlMap := make(map[string]string)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			// Log a warning for invalid lines
			logger.Warn("Invalid YAML line:", line)
			continue // Skip invalid lines
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		ymlMap[key] = value
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// Check for empty parsed map, indicating potential parsing issues
	if len(ymlMap) == 0 {
		return nil, errors.New("parsing failed: invalid yaml structure")
	}

	return ymlMap, nil
}

// Get information about the file at the given path.
func (f *Filesystem) GetFileInfo(rootDir string, path string) (relPath string, dir string, firstSubdir string, fileName string, extension string, err error) {
	// Check if the rootDir is a directory
	isRootDirDir, err := f.IsDir(rootDir)
	if err != nil {
		return "", "", "", "", "", fmt.Errorf("error checking rootDir: %s, error: %v", rootDir, err)
	}
	if !isRootDirDir {
		return "", "", "", "", "", fmt.Errorf("rootDir is not a directory: %s", rootDir)
	}

	// Check if the path exists using the Filesystem's Exists method
	if !f.Exists(path) {
		return "", "", "", "", "", fmt.Errorf("path does not exist: %s", path)
	}

	// Ensure the path is not a directory using IsDir
	isPathDir, err := f.IsDir(path)
	if err != nil {
		return "", "", "", "", "", fmt.Errorf("error checking path: %s, error: %v", path, err)
	}
	if isPathDir {
		return "", "", "", "", "", fmt.Errorf("path is a directory: %s", path)
	}

	// Get the relative path from the root directory using filepath.Rel
	relPath, err = filepath.Rel(rootDir, path)
	if err != nil {
		return "", "", "", "", "", fmt.Errorf("error getting relative path: %s, error: %v", path, err)
	}

	// Extract the directory using filepath.Dir
	dir = filepath.Dir(relPath)
	// If is is the current directory, set it to an empty string
	if dir == "." {
		dir = ""
	}

	// Split the relative path into components based on the file separator
	components := strings.Split(relPath, string(filepath.Separator))

	// Extract the first subdirectory (if it exists)
	if len(components) > 1 {
		firstSubdir = components[0]
	} else {
		firstSubdir = ""
	}

	// Get the filename from the relative path using filepath.Base
	fullfileName := filepath.Base(relPath)

	// Get the file extension using filepath.Ext and remove the leading dot
	extension = filepath.Ext(fullfileName)
	if extension != "" {
		extension = extension[1:]
	}

	// Get the filename without the extension
	fileName = strings.TrimSuffix(fullfileName, "."+extension)

	return relPath, dir, firstSubdir, fileName, extension, nil
}

// **********  Private Filesystem methods  **********

// Creates a file with the given path and writes content to it.
// It ensures that the parent directories exist before creating the file.
func (f *Filesystem) createFile(path string, content string) error {
	// Get the parent directory of the path
	parentDir := filepath.Dir(path)

	// Check if the parent directory exists, create it if not
	if !f.Exists(parentDir) {
		if err := f.createDirectory(parentDir); err != nil {
			return err
		}
	}

	// Create the file
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write content to the file
	_, err = file.WriteString(content)
	return err
}

// Creates a directory at the given path.
// It is recursive and can create the entire parent directory structure if needed.
func (f *Filesystem) createDirectory(path string) error {
	return os.MkdirAll(path, 0755)
}
