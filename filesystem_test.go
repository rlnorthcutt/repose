package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func relPath(t *testing.T, target string) string {
	t.Helper()

	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get working directory: %v", err)
	}

	rel, err := filepath.Rel(cwd, target)
	if err != nil {
		t.Fatalf("failed to compute relative path: %v", err)
	}

	return rel
}

func TestFilesystemExists(t *testing.T) {
	tempDir := t.TempDir()
	existingFile := filepath.Join(tempDir, "exists.txt")
	if err := os.WriteFile(existingFile, []byte("content"), 0o644); err != nil {
		t.Fatalf("failed to write test file: %v", err)
	}

	if !filesystem.Exists(existingFile) {
		t.Fatalf("expected file %s to exist", existingFile)
	}

	missingFile := filepath.Join(tempDir, "missing.txt")
	if filesystem.Exists(missingFile) {
		t.Fatalf("expected file %s to be missing", missingFile)
	}
}

func TestFilesystemExistsRecursive(t *testing.T) {
	tempDir := t.TempDir()
	nestedDir := filepath.Join(tempDir, "a", "b")
	if err := os.MkdirAll(nestedDir, 0o755); err != nil {
		t.Fatalf("failed to create nested directory: %v", err)
	}

	targetFile := filepath.Join(nestedDir, "target.txt")
	if err := os.WriteFile(targetFile, []byte("hello"), 0o644); err != nil {
		t.Fatalf("failed to write target file: %v", err)
	}

	exists, err := filesystem.ExistsRecursive("target.txt", tempDir)
	if err != nil {
		t.Fatalf("unexpected error searching for existing file: %v", err)
	}
	if !exists {
		t.Fatalf("expected target file to be found recursively")
	}

	exists, err = filesystem.ExistsRecursive("missing.txt", tempDir)
	if err != nil {
		t.Fatalf("unexpected error searching for missing file: %v", err)
	}
	if exists {
		t.Fatalf("did not expect missing file to be reported as existing")
	}

	_, err = filesystem.ExistsRecursive("whatever.txt", filepath.Join(tempDir, "does-not-exist"))
	if err == nil {
		t.Fatalf("expected an error when walking a directory that does not exist")
	}
}

func TestFilesystemCreate(t *testing.T) {
	tempDir := t.TempDir()

	filePath := filepath.Join(tempDir, "nested", "file.txt")
	if err := filesystem.Create(filePath, "Hello, Filesystem!"); err != nil {
		t.Fatalf("failed to create file: %v", err)
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("failed to read created file: %v", err)
	}
	if string(data) != "Hello, Filesystem!" {
		t.Fatalf("unexpected file contents: %s", string(data))
	}

	dirPath := filepath.Join(tempDir, "anotherDir")
	if err := filesystem.Create(dirPath, ""); err != nil {
		t.Fatalf("failed to create directory: %v", err)
	}
	if info, err := os.Stat(dirPath); err != nil || !info.IsDir() {
		t.Fatalf("expected %s to be a directory", dirPath)
	}

	if err := filesystem.Create(filePath, "new content"); err == nil {
		t.Fatalf("expected error when attempting to recreate existing file")
	}
}

func TestFilesystemRead(t *testing.T) {
	tempDir := t.TempDir()
	filePath := filepath.Join(tempDir, "readme.txt")

	if err := os.WriteFile(filePath, []byte("Read me"), 0o644); err != nil {
		t.Fatalf("failed to write file for reading: %v", err)
	}

	content, err := filesystem.Read(filePath)
	if err != nil {
		t.Fatalf("failed to read existing file: %v", err)
	}
	if content != "Read me" {
		t.Fatalf("unexpected content from Read: %s", content)
	}

	_, err = filesystem.Read(filepath.Join(tempDir, "missing.txt"))
	if err == nil {
		t.Fatalf("expected error when reading missing file")
	}
}

func TestFilesystemDelete(t *testing.T) {
	tempDir := t.TempDir()
	filePath := filepath.Join(tempDir, "deleteme.txt")

	if err := os.WriteFile(filePath, []byte("delete"), 0o644); err != nil {
		t.Fatalf("failed to write file for deletion: %v", err)
	}

	if err := filesystem.Delete(filePath); err != nil {
		t.Fatalf("failed to delete existing file: %v", err)
	}
	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		t.Fatalf("expected file to be deleted")
	}

	if err := filesystem.Delete(filepath.Join(tempDir, "missing.txt")); err == nil {
		t.Fatalf("expected error when deleting a missing file")
	}
}

func TestFilesystemIsDir(t *testing.T) {
	tempDir := t.TempDir()
	dirRel := relPath(t, tempDir)

	isDir, err := filesystem.IsDir(dirRel)
	if err != nil {
		t.Fatalf("unexpected error checking directory: %v", err)
	}
	if !isDir {
		t.Fatalf("expected %s to be a directory", dirRel)
	}

	filePath := filepath.Join(tempDir, "file.txt")
	if err := os.WriteFile(filePath, []byte("content"), 0o644); err != nil {
		t.Fatalf("failed to write file: %v", err)
	}
	fileRel := relPath(t, filePath)

	isDir, err = filesystem.IsDir(fileRel)
	if err != nil {
		t.Fatalf("unexpected error checking file: %v", err)
	}
	if isDir {
		t.Fatalf("expected %s to be reported as a file", fileRel)
	}

	_, err = filesystem.IsDir(filepath.Join(dirRel, "missing"))
	if err == nil {
		t.Fatalf("expected error when checking a path that does not exist")
	}
}

func TestFilesystemParseYml(t *testing.T) {
	content := "sitename: Example\nurl: https://example.com\n"

	result, err := filesystem.ParseYml(content)
	if err != nil {
		t.Fatalf("unexpected error parsing valid yaml: %v", err)
	}

	if result["sitename"] != "Example" {
		t.Fatalf("expected sitename to be parsed, got %s", result["sitename"])
	}
	if result["url"] != "https://example.com" {
		t.Fatalf("expected url to be parsed, got %s", result["url"])
	}

	if _, err := filesystem.ParseYml(""); err == nil {
		t.Fatalf("expected error when parsing empty yaml")
	}

	if _, err := filesystem.ParseYml("invalid-line-without-colon"); err == nil {
		t.Fatalf("expected error when parsing invalid yaml content")
	}
}

func TestFilesystemGetFileInfo(t *testing.T) {
	tempDir := t.TempDir()
	contentDir := filepath.Join(tempDir, "content")
	if err := os.MkdirAll(filepath.Join(contentDir, "blog"), 0o755); err != nil {
		t.Fatalf("failed to create content directory: %v", err)
	}

	filePath := filepath.Join(contentDir, "blog", "post.md")
	if err := os.WriteFile(filePath, []byte("content"), 0o644); err != nil {
		t.Fatalf("failed to write content file: %v", err)
	}

	rootRel := relPath(t, contentDir)
	fileRel := relPath(t, filePath)

	relPathResult, dir, firstSubdir, fileName, extension, err := filesystem.GetFileInfo(rootRel, fileRel)
	if err != nil {
		t.Fatalf("unexpected error getting file info: %v", err)
	}

	expectedRel := filepath.Join("blog", "post.md")
	if relPathResult != expectedRel {
		t.Fatalf("expected relative path %s, got %s", expectedRel, relPathResult)
	}
	if dir != "blog" {
		t.Fatalf("expected directory 'blog', got %s", dir)
	}
	if firstSubdir != "blog" {
		t.Fatalf("expected first subdir 'blog', got %s", firstSubdir)
	}
	if fileName != "post" {
		t.Fatalf("expected file name 'post', got %s", fileName)
	}
	if extension != "md" {
		t.Fatalf("expected extension 'md', got %s", extension)
	}
}

func TestFilesystemGetFileInfoErrors(t *testing.T) {
	tempDir := t.TempDir()
	contentDir := filepath.Join(tempDir, "content")
	if err := os.MkdirAll(contentDir, 0o755); err != nil {
		t.Fatalf("failed to create content directory: %v", err)
	}

	rootRel := relPath(t, contentDir)

	_, _, _, _, _, err := filesystem.GetFileInfo(rootRel, filepath.Join(rootRel, "missing.md"))
	if err == nil || !strings.Contains(err.Error(), "path does not exist") {
		t.Fatalf("expected missing path error, got %v", err)
	}

	_, _, _, _, _, err = filesystem.GetFileInfo(rootRel, rootRel)
	if err == nil || !strings.Contains(err.Error(), "path is a directory") {
		t.Fatalf("expected directory path error, got %v", err)
	}

	filePath := filepath.Join(tempDir, "notadir.txt")
	if err := os.WriteFile(filePath, []byte("content"), 0o644); err != nil {
		t.Fatalf("failed to create file: %v", err)
	}
	fileRel := relPath(t, filePath)

	_, _, _, _, _, err = filesystem.GetFileInfo(fileRel, fileRel)
	if err == nil || !strings.Contains(err.Error(), "rootDir is not a directory") {
		t.Fatalf("expected rootDir error, got %v", err)
	}
}
