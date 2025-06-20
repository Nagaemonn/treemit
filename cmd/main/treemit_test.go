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
	root := walkTree(tempDir, opts, 0)

	if root == nil {
		t.Fatal("Expected non-nil root node")
	}

	if root.Name != filepath.Base(tempDir) {
		t.Errorf("Expected root name %q, got %q", filepath.Base(tempDir), root.Name)
	}
}

func TestMatchesPattern(t *testing.T) {
	tests := []struct {
		name     string
		pattern  string
		expected bool
	}{
		{"test.go", "*.go", true},
		{"test.txt", "*.go", false},
		{"test.go", "*.go|*.txt", true},
		{"test.txt", "*.go|*.txt", true},
		{"test.py", "*.go|*.txt", false},
		{"venv", "venv", true},
		{"venv", "venv|node_modules", true},
		{"node_modules", "venv|node_modules", true},
		{"src", "venv|node_modules", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := matchesPattern(tt.name, tt.pattern)
			if got != tt.expected {
				t.Errorf("matchesPattern(%q, %q) = %v, want %v", tt.name, tt.pattern, got, tt.expected)
			}
		})
	}
}

func TestWalkTreeWithIgnore(t *testing.T) {
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
		"venv/test1.py",
		"venv/test2.py",
		"node_modules/test1.js",
		"node_modules/test2.js",
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

	// venvとnode_modulesを除外
	opts := &options{ignorePattern: "venv|node_modules"}
	root := walkTree(tempDir, opts, 0)

	if root == nil {
		t.Fatal("Expected non-nil root node")
	}

	// 除外されたディレクトリが含まれていないことを確認
	for _, child := range root.Children {
		if child.Name == "venv" || child.Name == "node_modules" {
			t.Errorf("Expected %s to be ignored", child.Name)
		}
	}
}

func TestWalkTreeWithLevel(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "treemit_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	// 2階層のディレクトリ構造を作成
	files := []string{
		"a.txt",
		"dir1/b.txt",
		"dir1/dir2/c.txt",
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

	opts := &options{level: 2}
	root := walkTree(tempDir, opts, 0)
	if root == nil {
		t.Fatal("Expected non-nil root node")
	}
	// dir1/dir2/c.txt は level=2 ではdir2は含まれるが、そのChildrenは空
	foundDir2 := false
	for _, child := range root.Children {
		if child.Name == "dir1" {
			for _, sub := range child.Children {
				if sub.Name == "dir2" {
					foundDir2 = true
					if len(sub.Children) != 0 {
						t.Errorf("dir2's Children should be empty when level=2")
					}
				}
			}
		}
	}
	if !foundDir2 {
		t.Errorf("dir2 should be included when level=2")
	}
}
