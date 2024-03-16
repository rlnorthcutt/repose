package filesystem

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/rlnorthcutt/repose/pkg/logger"
)

// Setup a temporary directory for testing
func setupTestDir(t *testing.T) (string, func()) {
	t.Helper()
	testDir, err := os.MkdirTemp("", "fs_test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}

	// Return the cleanup function to delete the test directory
	return testDir, func() {
		os.RemoveAll(testDir)
	}
}

func TestExists(t *testing.T) {
	testDir, cleanup := setupTestDir(t)
	defer cleanup()

	fs := New(logger.New(false)) // Assuming your logger has a New function that accepts a bool
	testFilePath := filepath.Join(testDir, "testfile.txt")

	// Test case for a non-existing file
	if fs.Exists(testFilePath) {
		t.Errorf("Expected file %s to not exist", testFilePath)
	}

	// Create a file to test the positive case
	err := os.WriteFile(testFilePath, []byte("test content"), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Test case for an existing file
	if !fs.Exists(testFilePath) {
		t.Errorf("Expected file %s to exist", testFilePath)
	}
}

func TestCreateAndDelete(t *testing.T) {
	testDir, cleanup := setupTestDir(t)
	defer cleanup()

	fs := New(logger.New(false))
	testFilePath := filepath.Join(testDir, "testfile.txt")
	testContent := "test content"

	// Test creating a file
	if err := fs.Create(testFilePath, testContent); err != nil {
		t.Fatalf("Failed to create file: %v", err)
	}

	// Verify the file exists
	if !fs.Exists(testFilePath) {
		t.Errorf("Expected file %s to exist after creation", testFilePath)
	}

	// Test deleting the file
	if err := fs.Delete(testFilePath); err != nil {
		t.Fatalf("Failed to delete file: %v", err)
	}

	// Verify the file no longer exists
	if fs.Exists(testFilePath) {
		t.Errorf("Expected file %s to not exist after deletion", testFilePath)
	}
}
