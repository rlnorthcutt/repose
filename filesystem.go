package main

import (
	"errors"
	"os"
	"path/filepath"
)

// Defining a new public type 'Filesystem'
type Filesystem int

// Defining a global varaiable for Filesystem
var filesystem Filesystem

// **********  Public Filesystem methods  **********

// Create a file or directory based on the given path and content.
// If the path already exists, it returns an error.
func (f *Filesystem) Create(path string, content string) error {
	// Check if the path exists using checkPath
	if f.pathExists(path) {
		errorMessage := path + " already exists"
		logger.Error(errorMessage)
		return errors.New(errorMessage)
	}

	// Determine if the path is for a file or directory and create accordingly
	if filepath.Ext(path) != "" {
		return f.createFile(path, content)
	}
	return f.createDirectory(path)
}

// Delete removes a file or directory at the given path.
// If the path does not exist, it returns an error.
func (f *Filesystem) Delete(path string) error {
	// Check if the path exists
	if !f.pathExists(path) {
		warningMessage := path + " doesn't exist"
		logger.Warn(warningMessage)
		return errors.New(warningMessage)
	}

	// Delete the file or directory
	return os.RemoveAll(path)
}

// **********  Private Filesystem methods  **********

// pathExists checks if the given path exists.
// It returns true if the path exists, and false if it does not.
func (f *Filesystem) pathExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// createFile creates a file with the given path and writes content to it.
// It ensures that the parent directories exist before creating the file.
func (f *Filesystem) createFile(path string, content string) error {
	// Get the parent directory of the path
	parentDir := filepath.Dir(path)

	// Check if the parent directory exists, create it if not
	if !f.pathExists(parentDir) {
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

// createDirectory creates a directory at the given path.
// It is recursive and can create the entire parent directory structure if needed.
func (f *Filesystem) createDirectory(path string) error {
	return os.MkdirAll(path, 0755)
}
