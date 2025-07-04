---
title: "treemit: 拡張子でフィルタリングできるツリーコマンド"
description: "「tree」コマンドを拡張し，拡張子によるフィルタリング機能を追加したCLIツール"
date: 2025-07-03
image: "img/featured-getting-started.png"
categories:
    - Tutorial
tags:
    - CLI
    - Go
    - Development Tools
---

## treemitとは？

treemitは，Unixの`tree`コマンドを拡張したディレクトリ構造可視化ツールです．
従来の`tree`コマンドでは扱えなかった「特定の拡張子のファイルだけを表示する」といった操作が行えます．

## treemitが便利な場面
- プロジェクトのファイル構造を一覧したいが，treeでは冗長なファイルが多すぎて見づらい時
- findやawkを組み合わせたシェルスクリプトも出力がイマイチな時
- プロジェクトの構造をLLMに伝えたいが，ファイル数が多すぎて上限に引っかかる時

treemitは，このようなAI時代における開発者の日常的なニーズに応えるために開発されました．

## インストール（Homebrew）

```bash
brew tap Nagaemonn/treemit
brew install treemit
```

## 基本的な使い方

### 基本的なディレクトリツリーの表示

```bash
treemit [ディレクトリパス]
```

### 特定の拡張子ごとの表示数を制限

```bash
treemit -E 2  # 各拡張子ごとに2ファイルまで表示
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
| -E, --extension | 各拡張子ごとの表示数上限を指定 |
| --help | ヘルプメッセージを表示 |

## 使用例

### Goプロジェクトのソースファイル構造を確認

```bash
treemit -E 2 ./src  # 各拡張子ごとに2ファイルまで表示
```

### ドキュメントファイルの構造を2階層まで表示

```bash
treemit -E 1 -L 2 ./docs  # 各拡張子ごとに1ファイルまで，深さ2階層まで表示
```

## 開発背景

既存の`tree`コマンドはディレクトリの構造を一覧するのに優れたツールですが，拡張子ベースで表示するファイル数を制限する機能が不足していました．
treemitは，この機能を実装することで，より効率的なプロジェクト構造の把握を可能にします．
さらに，コンパクトに構造を表示することでLLMに伝える際のトークン数削減に非常に効果的です．

## 今後の展望

- より多くのフィルタリングオプションの追加
- 出力形式のカスタマイズ機能
- パフォーマンスの最適化

## フィードバック

バグ報告や機能リクエストは，[GitHubのIssue](https://github.com/Nagaemonn/treemit/issues)にお願いします．

## ライセンス

このプロジェクトはMITライセンスの下で公開されています． 