package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGetFileExtension(t *testing.T) {
	tests := []struct {
		name     string
		expected string
	}{
		{"test.go", "go"},
		{"test.txt", "txt"},
		{"test", ""},
		{"test.", ""},
		{".gitignore", "gitignore"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getFileExtension(tt.name)
			if got != tt.expected {
				t.Errorf("getFileExtension(%q) = %q, want %q", tt.name, got, tt.expected)
			}
		})
	}
}

func TestGroupFilesByExtension(t *testing.T) {
	nodes := []*Node{
		{Name: "test1.go", IsDir: false, Ext: "go"},
		{Name: "test2.go", IsDir: false, Ext: "go"},
		{Name: "test3.txt", IsDir: false, Ext: "txt"},
		{Name: "dir1", IsDir: true, Ext: ""},
	}

	groups := groupFilesByExtension(nodes)

	if len(groups) != 2 {
		t.Errorf("Expected 2 groups, got %d", len(groups))
	}

	if len(groups["go"]) != 2 {
		t.Errorf("Expected 2 .go files, got %d", len(groups["go"]))
	}

	if len(groups["txt"]) != 1 {
		t.Errorf("Expected 1 .txt file, got %d", len(groups["txt"]))
	}
}

func TestWalkTree(t *testing.T) {
	// テスト用の一時ディレクトリを作成
	tempDir, err := os.MkdirTemp("", "treemit_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	// テスト用のファイルとディレクトリを作成
	files := []string{
		"test1.go",
		"test2.go",
		"test3.go",
		"test1.txt",
		"test2.txt",
		"dir1/test1.go",
		"dir1/test2.go",
	}

	for _, file := range files {
		path := filepath.Join(tempDir, file)
		dir := filepath.Dir(path)
		if err := os.MkdirAll(dir, 0755); err != nil {
			t.Fatal(err)
		}
		if _, err := os.Create(path); err != nil {
			t.Fatal(err)
		}
	}

	opts := &options{extension: 2}
	root := walkTree(tempDir, opts)

	if root == nil {
		t.Fatal("Expected non-nil root node")
	}

	if root.Name != filepath.Base(tempDir) {
		t.Errorf("Expected root name %q, got %q", filepath.Base(tempDir), root.Name)
	}
}
