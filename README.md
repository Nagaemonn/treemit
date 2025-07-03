# :evergreen_tree:treemit
Enhanced "tree" command with file and depth limits.

[![Coverage Status](https://coveralls.io/repos/github/Nagaemonn/treemit/badge.svg?branch=main)](https://coveralls.io/github/Nagaemonn/treemit?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/Nagaemonn/treemit)](https://goreportcard.com/report/github.com/Nagaemonn/treemit)
[![DOI](https://zenodo.org/badge/967311302.svg)](https://doi.org/10.5281/zenodo.15751267)

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
    -I, --ignore       List only those files that do not match the pattern given.
                       Multiple patterns can be specified with '|'.
    --help             Print usage and this help message and exit.
```

## Examples

```sh
# Show all files including hidden ones
treemit -a

# Show only directories
treemit -d

# Limit display depth to 2 levels
treemit -L 2

# Show only 1 file per extension group
treemit -E 1

# Exclude .md and .go files
treemit -I "*.md|*.go"

# Combine multiple options
treemit -a -L 3 -E 2 -I "*.tmp|*.log"
```

## Features

- **Extension-based grouping**: Files with the same extension are grouped together for better readability
- **Streaming output**: Unlike traditional tree commands, treemit outputs results immediately without building the entire tree in memory
- **Silent error handling**: Access permission errors are handled gracefully without cluttering the output
- **Pattern exclusion**: Exclude files and directories using glob patterns