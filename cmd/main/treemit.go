package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	flag "github.com/spf13/pflag"
)

// オプションを保持する構造体
type options struct {
	all           bool
	dirOnly       bool
	level         int
	extension     int
	help          bool
	ignorePattern string // 除外パターン
}

// ファイルツリーのノード構造体（雛形）
type Node struct {
	Name     string
	Children []*Node
	IsDir    bool
	Ext      string // ファイルの拡張子
	// 必要に応じて他の情報も追加
}

// ファイルの拡張子を取得する関数
func getFileExtension(name string) string {
	ext := filepath.Ext(name)
	if ext == "" {
		return ""
	}
	return ext[1:] // 先頭のドットを除去
}

// ファイル構造を探索する関数（ルートノードを返す形に変更）
// depth: 現在の深さ（ルートは0）
func walkTree(root string, opts *options, depth int) *Node {
	if opts.level > 0 && depth >= opts.level {
		return nil
	}
	info, err := os.Stat(root)
	if err != nil {
		return nil
	}
	rootNode := &Node{Name: info.Name(), IsDir: info.IsDir(), Ext: getFileExtension(info.Name())}
	if info.IsDir() {
		entries, err := os.ReadDir(root)
		if err != nil {
			return rootNode
		}
		for _, entry := range entries {
			if !opts.all && strings.HasPrefix(entry.Name(), ".") {
				continue
			}
			// 除外パターンにマッチする場合はスキップ
			if opts.ignorePattern != "" && matchesPattern(entry.Name(), opts.ignorePattern) {
				continue
			}
			childPath := filepath.Join(root, entry.Name())
			childNode := walkTree(childPath, opts, depth+1)
			if childNode != nil {
				rootNode.Children = append(rootNode.Children, childNode)
			}
		}
	}
	return rootNode
}

// ツリーを即時出力しながら探索する関数（本家tree風）
func printTreeStreaming(root string, opts *options, prefix string, isRoot bool, depth int) {
	if opts.level > 0 && depth >= opts.level {
		return
	}
	info, err := os.Stat(root)
	if err != nil {
		return
	}
	if isRoot {
		fmt.Println(info.Name())
	}
	if info.IsDir() {
		entries, err := os.ReadDir(root)
		if err != nil {
			return
		}
		// フィルタリング
		var dirs []os.DirEntry
		var files []os.DirEntry
		for _, entry := range entries {
			if !opts.all && strings.HasPrefix(entry.Name(), ".") {
				continue
			}
			if opts.ignorePattern != "" && matchesPattern(entry.Name(), opts.ignorePattern) {
				continue
			}
			if entry.IsDir() {
				dirs = append(dirs, entry)
			} else if !opts.dirOnly {
				files = append(files, entry)
			}
		}
		// ディレクトリを先に出力
		for i, entry := range dirs {
			isLastDir := i == len(dirs)-1 && len(files) == 0
			branch := "├── "
			if isLastDir {
				branch = "└── "
			}
			fmt.Printf("%s%s%s\n", prefix, branch, entry.Name())
			nextPrefix := prefix
			if isLastDir {
				nextPrefix += "    "
			} else {
				nextPrefix += "│   "
			}
			childPath := filepath.Join(root, entry.Name())
			printTreeStreaming(childPath, opts, nextPrefix, false, depth+1)
		}
		// ファイルを拡張子ごとにグループ化して出力
		if len(files) > 0 && !opts.dirOnly {
			// グループ化
			groups := make(map[string][]os.DirEntry)
			var extOrder []string
			for _, entry := range files {
				ext := getFileExtension(entry.Name())
				if _, ok := groups[ext]; !ok {
					extOrder = append(extOrder, ext)
				}
				groups[ext] = append(groups[ext], entry)
			}
			groupIndex := 0
			for _, ext := range extOrder {
				groupFiles := groups[ext]
				displayCount := len(groupFiles)
				if opts.extension > 0 && displayCount > opts.extension {
					displayCount = opts.extension
				}
				for i, entry := range groupFiles[:displayCount] {
					isLastFile := (groupIndex == len(extOrder)-1) && (i == displayCount-1)
					branch := "├── "
					if isLastFile {
						branch = "└── "
					}
					fmt.Printf("%s%s%s\n", prefix, branch, entry.Name())
				}
				// 残りのファイル数を表示
				if opts.extension > 0 && len(groupFiles) > opts.extension {
					remaining := len(groupFiles) - opts.extension
					branch := "└── "
					if groupIndex != len(extOrder)-1 {
						branch = "├── "
					}
					fmt.Printf("%s%s... +%d\n", prefix, branch, remaining)
				}
				groupIndex++
			}
		}
	}
}

func helpMessage() string {
	return `Usage: treemit [DIRs] [OPTION]

OPTION
    -a                 All files are listed.
    -d                 List directories only.
    -L, --level        Max display depth of the directory tree.
    -E, --extension    Max display files of the same extensions.
    -I, --ignore       List only those files that do not match the pattern given.
                      Multiple patterns can be specified with '|'.
    --help             Print usage and this help message and exit.
`
}

func buildFlagSet() (*flag.FlagSet, *options) {
	opts := &options{}
	flags := flag.NewFlagSet("treemit", flag.ContinueOnError)
	flags.Usage = func() { fmt.Fprint(os.Stderr, helpMessage()) }

	flags.BoolVarP(&opts.all, "all", "a", false, "All files are listed.")
	flags.BoolVarP(&opts.dirOnly, "dir", "d", false, "List directories only.")
	flags.IntVarP(&opts.level, "level", "L", 0, "Max display depth of the directory tree.")
	flags.IntVarP(&opts.extension, "extension", "E", 0, "Max display files of the same extensions.")
	flags.StringVarP(&opts.ignorePattern, "ignore", "I", "", "List only those files that do not match the pattern given. Multiple patterns can be specified with '|'.")
	flags.BoolVar(&opts.help, "help", false, "Print usage and this help message and exit.")

	return flags, opts
}

// パターンにマッチするかチェックする関数
func matchesPattern(name string, pattern string) bool {
	if pattern == "" {
		return false
	}
	patterns := strings.Split(pattern, "|")
	for _, p := range patterns {
		matched, err := filepath.Match(strings.TrimSpace(p), name)
		if err != nil {
			continue
		}
		if matched {
			return true
		}
	}
	return false
}

func goMain(args []string) int {
	flags, opts := buildFlagSet()
	if err := flags.Parse(args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		flags.Usage()
		return 2
	}
	if opts.help {
		flags.Usage()
		return 0
	}
	if flags.NArg() == 0 {
		// 引数がなければカレントディレクトリのみ
		printTreeStreaming(".", opts, "", true, 0)
	} else {
		for i := 0; i < flags.NArg(); i++ {
			dir := flags.Arg(i)
			printTreeStreaming(dir, opts, "", true, 0)
		}
	}
	return 0
}

func main() {
	status := goMain(os.Args)
	os.Exit(status)
}
