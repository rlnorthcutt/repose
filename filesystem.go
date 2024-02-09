package main

import (
	"bufio"
	"errors"
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

// Basic YAML format parser that reads the content & returns a map of the data.
// We use this instead of loading the YAML modules to keep the size down
func (f *Filesystem) ParseYml(content string) (map[string]string, error) {
	scanner := bufio.NewScanner(strings.NewReader(content))
	ymlMap := make(map[string]string)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue // Skip invalid lines
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		ymlMap[key] = value
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return ymlMap, nil
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
