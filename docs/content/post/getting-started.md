---
title: "treemit: 拡張子でフィルタリングできるツリーコマンド"
description: "ディレクトリ構造を可視化する「tree」コマンドを拡張し、拡張子によるフィルタリング機能を追加したCLIツール"
date: 2025-06-19
image: "img/featured-getting-started.png"
categories:
    - Tutorial
tags:
    - CLI
    - Go
    - Productivity
    - Development Tools
---

## treemitとは？

treemitは、Unixの`tree`コマンドを拡張した、より柔軟なディレクトリ構造可視化ツールです。
従来の`tree`コマンドでは難しかった「特定の拡張子のファイルだけを表示する」といった操作が簡単にできます。

## なぜtreemitを使うのか？

開発プロジェクトで以下のような場面に遭遇したことはありませんか？

- プロジェクトの構造を把握したいが、不要なファイルが多すぎて見づらい
- `.go`ファイルだけのディレクトリ構造を確認したい
- ディレクトリの深さを制限しつつ、特定の拡張子のファイルだけを表示したい

treemitは、このような開発者の日常的なニーズに応えるために開発されました。

## インストール方法

### Goを使用してインストール

```bash
go install github.com/Nagaemonn/treemit@latest
```

### Homebrewを使用してインストール（予定）

```bash
brew tap Nagaemonn/treemit
brew install treemit
```

## 基本的な使い方

### 基本的なディレクトリツリーの表示

```bash
treemit [ディレクトリパス]
```

### 特定の拡張子のファイルだけを表示

```bash
treemit -E go  # .goファイルだけを表示
treemit -E md  # .mdファイルだけを表示
```

### ディレクトリの深さを制限して表示

```bash
treemit -L 2  # 深さ2階層まで表示
```

### すべてのファイルを表示（隠しファイルを含む）

```bash
treemit -a
```

### ディレクトリのみを表示

```bash
treemit -d
```

## オプション一覧

| オプション | 説明 |
|------------|------|
| -a | すべてのファイルを表示（隠しファイルを含む） |
| -d | ディレクトリのみを表示 |
| -L, --level | ディレクトリツリーの最大深さを指定 |
| -E, --extension | 表示する拡張子を指定 |
| --help | ヘルプメッセージを表示 |

## 使用例

### Goプロジェクトのソースファイル構造を確認

```bash
treemit -E go ./src
```

### ドキュメントファイルの構造を2階層まで表示

```bash
treemit -E md -L 2 ./docs
```

## 開発背景

既存の`tree`コマンドは優れたツールですが、特定の拡張子のファイルだけを表示する機能が不足していました。
treemitは、この機能を追加することで、より効率的なプロジェクト構造の把握を可能にします。

## 今後の展望

- より多くのフィルタリングオプションの追加
- 出力形式のカスタマイズ機能
- パフォーマンスの最適化

## フィードバック

バグ報告や機能リクエストは、[GitHubのIssue](https://github.com/Nagaemonn/treemit/issues)にお願いします。
プルリクエストも大歓迎です！

## ライセンス

このプロジェクトはMITライセンスの下で公開されています。 