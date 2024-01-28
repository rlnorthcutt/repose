package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFilesystem_Create(t *testing.T) {
	// Define test directory and file
	testDir := "./testdir"
	testFile := filepath.Join(testDir, "testfile.txt")
	testContent := "Hello, Filesystem!"

	// Test creating a directory
	err := filesystem.Create(testDir, "")
	if err != nil {
		t.Errorf("Failed to create directory: %s", err)
	}

	// Test creating a file
	err = filesystem.Create(testFile, testContent)
	if err != nil {
		t.Errorf("Failed to create file: %s", err)
	}

	// Check if file content is correct
	content, _ := os.ReadFile(testFile)
	if string(content) != testContent {
		t.Errorf("File content mismatch. Got: %s, Want: %s", string(content), testContent)
	}

	// Cleanup
	os.RemoveAll(testDir)
}

func TestFilesystem_Delete(t *testing.T) {
	// Define test directory and file
	testDir := "./testdir"
	testFile := filepath.Join(testDir, "testfile.txt")

	// Create a directory and a file to test deletion
	os.MkdirAll(testDir, 0755)
	os.WriteFile(testFile, []byte("test content"), 0644)

	// Test deleting the file
	err := filesystem.Delete(testFile)
	if err != nil {
		t.Errorf("Failed to delete file: %s", err)
	}

	// Check if file is deleted
	if _, err := os.Stat(testFile); !os.IsNotExist(err) {
		t.Errorf("File still exists after deletion")
	}

	// Test deleting the directory
	err = filesystem.Delete(testDir)
	if err != nil {
		t.Errorf("Failed to delete directory: %s", err)
	}

	// Check if directory is deleted
	if _, err := os.Stat(testDir); !os.IsNotExist(err) {
		t.Errorf("Directory still exists after deletion")
	}
}
