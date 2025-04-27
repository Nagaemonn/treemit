# :evergreen_tree:treemit
Enhanced "tree" command with file and depth limits.

[![Coverage Status](https://coveralls.io/repos/github/Nagaemonn/treemit/badge.svg?branch=main)](https://coveralls.io/github/Nagaemonn/treemit?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/Nagaemonn/treemit)](https://goreportcard.com/report/github.com/Nagaemonn/treemit)

## :bulb: Overview
This is part of the `tree` command implementation, with a new function to limit the file listing based on extensions, which is not supported by the current `tree`.

## Usage

```sh
treemit [DIRs] [OPTION]

OPTION
    -a                 All files are listed.
    -d                 List directories only.
    -L, --level        Max display depth of the directory tree.
    -E, --extension    Max display files of the same extensions.
    --help             Print usage and this help message and exit.

    W.I.P
```
