---
title: "Go Tutorial - Building a Static Site Generator - Part 7 testing"
description: "Exploring the beauty and simplicity of static websites in the modern web era."
tags: [static sites, web development, Repose]
image: /images/zen-static.jpg
index: true
publish: true
author: "Ron Northcutt"
publish_date: "2024-01-30"
template: "blog-post.tmpl"
---

# Go Tutorial - Building a Static Site Generator - Part  7 testing
## Filesystem.go testing
🧙🏾‍♂️: Writing tests for your `Filesystem` type is a great way to ensure its methods work as expected. In Go, you typically write tests in a separate file, conventionally named with `_test.go` suffix. Let's create a test file for your `Filesystem` methods.

### File: `filesystem_test.go`

```go
package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestFilesystem_Create(t *testing.T) {
	// Setup Filesystem instance
	fs := Filesystem{}

	// Define test directory and file
	testDir := "./testdir"
	testFile := filepath.Join(testDir, "testfile.txt")
	testContent := "Hello, Filesystem!"

	// Test creating a directory
	err := fs.Create(testDir, "")
	if err != nil {
		t.Errorf("Failed to create directory: %s", err)
	}

	// Test creating a file
	err = fs.Create(testFile, testContent)
	if err != nil {
		t.Errorf("Failed to create file: %s", err)
	}

	// Check if file content is correct
	content, _ := ioutil.ReadFile(testFile)
	if string(content) != testContent {
		t.Errorf("File content mismatch. Got: %s, Want: %s", string(content), testContent)
	}

	// Cleanup
	os.RemoveAll(testDir)
}

func TestFilesystem_Delete(t *testing.T) {
	// Setup Filesystem instance
	fs := Filesystem{}

	// Define test directory
	testDir := "./testdir"

	// Create a directory to test deletion
	os.MkdirAll(testDir, 0755)

	// Test deleting the directory
	err := fs.Delete(testDir)
	if err != nil {
		t.Errorf("Failed to delete directory: %s", err)
	}

	// Check if directory is deleted
	if _, err := os.Stat(testDir); !os.IsNotExist(err) {
		t.Errorf("Directory still exists after deletion")
	}
}
```

### Explanation

- `TestFilesystem_Create` tests the `Create` method. It attempts to create a directory and a file, then verifies if the file contains the expected content.
- `TestFilesystem_Delete` tests the `Delete` method. It creates a directory, then uses `Delete` to remove it and checks if the directory still exists.
- The tests use `t.Errorf` to report failures. This approach allows the test to continue running even after an assertion fails, which can be useful for identifying multiple issues in one test run.
- The `os.RemoveAll` call in the cleanup step ensures that the test directory is removed after each test, keeping your test environment clean.

### Running the Tests

To run these tests, use the `go test` command in the directory containing your `filesystem.go` and `filesystem_test.go` files:

```bash
go test
```

This command will compile your test code along with the code being tested and then run the tests.

### Notes

- These tests assume that the `logger` used in your `Filesystem` methods is properly set up. If `logger` is not defined, you'll need to either mock it or initialize it before running the tests.
- Depending on your environment or additional requirements, you might need to modify the tests, especially the file paths.

If you have any specific scenarios you want to test or need further assistance, feel free to ask! 🚀🧪