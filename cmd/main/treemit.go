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
	all       bool
	dirOnly   bool
	level     int
	extension int
	help      bool
}

// ファイルツリーのノード構造体（雛形）
type Node struct {
	Name     string
	Children []*Node
	// 必要に応じて他の情報も追加
}

// ファイル構造を探索する関数（ルートノードを返す形に変更）
func walkTree(root string, opts *options) *Node {
	info, err := os.Stat(root)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading directory %s: %v\n", root, err)
		return nil
	}
	rootNode := &Node{Name: info.Name()}
	if info.IsDir() {
		entries, err := os.ReadDir(root)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading directory %s: %v\n", root, err)
			return rootNode
		}
		for _, entry := range entries {
			if !opts.all && strings.HasPrefix(entry.Name(), ".") {
				continue
			}
			childPath := filepath.Join(root, entry.Name())
			childNode := walkTree(childPath, opts)
			if childNode != nil {
				rootNode.Children = append(rootNode.Children, childNode)
			}
		}
	}
	return rootNode
}

// ツリーを表示する関数（ルート名も表示）
func printTree(root *Node, opts *options) {
	fmt.Println(root.Name)
	printTreeFancy(root.Children, "", false)
}

func printTreeFancy(nodes []*Node, prefix string, isRoot bool) {
	for i, node := range nodes {
		isLast := i == len(nodes)-1
		var branch string
		if isRoot {
			branch = ""
		} else if isLast {
			branch = "└── "
		} else {
			branch = "├── "
		}
		fmt.Printf("%s%s%s\n", prefix, branch, node.Name)
		// 次の階層のprefixを決定
		nextPrefix := prefix
		if !isRoot {
			if isLast {
				nextPrefix += "    "
			} else {
				nextPrefix += "│   "
			}
		}
		if len(node.Children) > 0 {
			printTreeFancy(node.Children, nextPrefix, false)
		}
	}
}

func helpMessage() string {
	return `Usage: treemit [DIRs] [OPTION]

` +
		`OPTION
` +
		`    -a                 All files are listed.
` +
		`    -d                 List directories only.
` +
		`    -L, --level        Max display depth of the directory tree.
` +
		`    -E, --extension    Max display files of the same extensions.
` +
		`    --help             Print usage and this help message and exit.
`
	// -hオプションも追加したかったが，本来のtreeコマンドに表示形式用-hオプション(humanize)があったため不採用
}

func buildFlagSet() (*flag.FlagSet, *options) {
	opts := &options{}
	flags := flag.NewFlagSet("treemit", flag.ContinueOnError)
	flags.Usage = func() { fmt.Fprint(os.Stderr, helpMessage()) }

	// レイアウトの縦揃えを試したが，gofmtによる自動整形で詰められるので，それを受け入れる
	flags.BoolVarP(&opts.all, "all", "a", false, "All files are listed.")
	flags.BoolVarP(&opts.dirOnly, "dir", "d", false, "List directories only.")
	flags.IntVarP(&opts.level, "level", "L", 0, "Max display depth of the directory tree.")
	flags.IntVarP(&opts.extension, "extension", "E", 0, "Max display files of the same extensions.")
	flags.BoolVar(&opts.help, "help", false, "Print usage and this help message and exit.")

	return flags, opts
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
		printTree(walkTree(".", opts), opts)
	} else {
		for i := 0; i < flags.NArg(); i++ {
			dir := flags.Arg(i)
			printTree(walkTree(dir, opts), opts)
		}
	}
	return 0
}

func main() {
	status := goMain(os.Args)
	os.Exit(status)
}
