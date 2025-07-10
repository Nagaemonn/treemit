package main

import "fmt"

// 補完スクリプト出力用関数
func OutputCompletionScript() {
	fmt.Print(`# bash completion for treemit
__treemit() {
    local cur prev opts
    COMPREPLY=()
    cur="${COMP_WORDS[COMP_CWORD]}"
    prev="${COMP_WORDS[COMP_CWORD-1]}"
    opts="-a --all -d --dir -L --level -E --extension -I --ignore --help"

    case "${prev}" in
        --level|-L|--extension|-E|--ignore|-I)
            return 0
            ;;
    esac

    if [[ "$cur" == -* ]]; then
        COMPREPLY=( $(compgen -W "${opts}" -- "${cur}") )
        return 0
    else
        compopt -o filenames
        COMPREPLY=( $(compgen -f -- "$cur") )
    fi
}
complete -F __treemit -o bashdefault -o default treemit
`)
}

func OutputZshCompletionScript() {
	fmt.Print(`#compdef treemit

_arguments \
  '-a[All files are listed]' \
  '-d[List directories only]' \
  '-L+[Max display depth of the directory tree]:level' \
  '-E+[Max display files of the same extensions]:extension' \
  '-I+[List only those files that do not match the pattern given]:ignore pattern' \
  '--help[Print usage and this help message and exit]'
`)
}
