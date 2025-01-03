# PokeMatch

PokeMatchは、ユーザーが身長と体重を入力すると、近しいポケモンを表示するWebアプリケーションです。

## 主な機能
- 身長と体重を入力して近しいポケモンを表示

## 技術要素
- **バックエンド**
  - Go 1.23.2
  - Gin 1.10.0
- **データベース**
  - Postgres 16.6
- **デプロイ**
  - Dockerを使用して、コンテナ化されたアプリケーションをデプロイ

## 実装内容
- Controller, Service, Repositoryに分割
- [google/wire](https://github.com/google/wire)
- ↑DIコンテナを使った実装
- Gormを使ったマイグレーション
- バリデーションチェック
- リクエストごとのログ出力
- Dockerでの開発
- テーブル定義書、画面遷移図作成

## 次の目標
[次回の開発リポジトリ（Giter）](https://github.com/develop-suda/Giter)
- JWTを使ったログイン機能の実装
- GitHubログイン機能
- GitHub GraphQL APIを使ったデータ取得
- より良いエラーハンドリング
- テストコードの実装
- セキュリティを考慮したコード
- コメント書く
