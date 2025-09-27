package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCommandProcessFileName(t *testing.T) {
	var c Command

	fileName, title := c.processFileName("my-first-post")
	if fileName != "my-first-post.md" {
		t.Fatalf("expected markdown extension, got %s", fileName)
	}
	if title != "My First Post" {
		t.Fatalf("unexpected title: %s", title)
	}

	fileName, title = c.processFileName("already.html")
	if fileName != "already.html" {
		t.Fatalf("expected filename to remain unchanged, got %s", fileName)
	}
	if title != "Already" {
		t.Fatalf("unexpected title for html file: %s", title)
	}
}

func TestCommandDefaultContent(t *testing.T) {
	originalConfig := config
	config = Config{Author: "Test Author"}
	t.Cleanup(func() { config = originalConfig })

	var c Command
	content := c.defaultContent("post", "My Great Post")

	if !strings.Contains(content, "title: My Great Post") {
		t.Fatalf("expected title to be inserted into content: %s", content)
	}
	if !strings.Contains(content, "description: post about My Great Post") {
		t.Fatalf("expected description to include content type and title: %s", content)
	}
	if !strings.Contains(content, "author: Test Author") {
		t.Fatalf("expected author to be pulled from config: %s", content)
	}
	if !strings.Contains(content, "template: post.tmpl") {
		t.Fatalf("expected template name to be based on content type: %s", content)
	}
}

func TestCommandCreateNewContent(t *testing.T) {
	originalWD, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get working directory: %v", err)
	}

	tempRoot := t.TempDir()
	if err := os.Chdir(tempRoot); err != nil {
		t.Fatalf("failed to change working directory: %v", err)
	}
	t.Cleanup(func() {
		_ = os.Chdir(originalWD)
	})

	if err := os.MkdirAll("template", 0o755); err != nil {
		t.Fatalf("failed to create template directory: %v", err)
	}

	originalConfig := config
	config = Config{Author: "Author", ContentDirectory: "content"}
	t.Cleanup(func() { config = originalConfig })

	cfg := config

	var c Command
	if err := c.createNewContent(cfg, "posts", "new-entry"); err != nil {
		t.Fatalf("createNewContent returned error: %v", err)
	}

	expectedPath := filepath.Join("content", "posts", "new-entry.md")
	data, err := os.ReadFile(expectedPath)
	if err != nil {
		t.Fatalf("expected file to be created at %s: %v", expectedPath, err)
	}

	if !strings.Contains(string(data), "title: New Entry") {
		t.Fatalf("expected generated file to contain formatted title, got: %s", string(data))
	}
	if !strings.Contains(string(data), "author: Author") {
		t.Fatalf("expected generated file to contain author from config, got: %s", string(data))
	}
}

func TestConfigLoad(t *testing.T) {
	tempRoot := t.TempDir()
	configContent := "sitename: Example\nauthor: Alice\neditor: vim\ncontentDirectory: content\noutputDirectory: public\nurl: https://example.com\npreviewUrl: http://localhost:9000\ntheme: pico\n"

	configPath := filepath.Join(tempRoot, ConfigFile)
	if err := os.WriteFile(configPath, []byte(configContent), 0o644); err != nil {
		t.Fatalf("failed to write config file: %v", err)
	}

	originalRoot := buildCommand.rootPath
	buildCommand.rootPath = tempRoot
	t.Cleanup(func() { buildCommand.rootPath = originalRoot })

	var cfg Config
	loaded, err := cfg.Load()
	if err != nil {
		t.Fatalf("unexpected error loading config: %v", err)
	}

	if loaded.Sitename != "Example" {
		t.Fatalf("expected sitename Example, got %s", loaded.Sitename)
	}
	if loaded.Author != "Alice" {
		t.Fatalf("expected author Alice, got %s", loaded.Author)
	}
	if loaded.Editor != "vim" {
		t.Fatalf("expected editor vim, got %s", loaded.Editor)
	}
	if loaded.ContentDirectory != "content" {
		t.Fatalf("expected content directory content, got %s", loaded.ContentDirectory)
	}
	if loaded.OutputDirectory != "public" {
		t.Fatalf("expected output directory public, got %s", loaded.OutputDirectory)
	}
	if loaded.URL != "https://example.com" {
		t.Fatalf("expected URL https://example.com, got %s", loaded.URL)
	}
	if loaded.PreviewURL != "http://localhost:9000" {
		t.Fatalf("expected preview URL http://localhost:9000, got %s", loaded.PreviewURL)
	}
	if loaded.Theme != "pico" {
		t.Fatalf("expected theme pico, got %s", loaded.Theme)
	}
}
