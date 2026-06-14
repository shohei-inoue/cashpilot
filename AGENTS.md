# CashPilot — Agent Instructions

Cloud Agent / Cursor Automations がこのリポジトリで作業するときの手順と規約です。

## プロジェクト概要

CashPilot は個人向けのキャッシュフローシミュレーターと収支トラッカーです。

| 領域 | 技術 |
|------|------|
| フロントエンド | Next.js 16, React 19, TypeScript, SCSS Modules |
| バックエンド | Go 1.25, Gin |
| データベース | PostgreSQL 16 |
| ローカル開発 | Docker Compose |

## 実装状況（重要）

- **完了**: Phase 0〜4（環境、DB、バックエンド API、認証 UI、レイアウト）
- **未実装**: Phase 5（ダッシュボード・取引・シミュレーションの API 連携、目標・設定画面）
- 詳細: [docs/implementation-flow.md](docs/implementation-flow.md)

Phase 5 の画面実装では、既存の Server Actions（`frontend/src/app/actions/`）と API クライアント（`frontend/src/app/libs/client.ts`）のパターンに従ってください。

## Cursor Cloud specific instructions

### 依存関係のインストール

```bash
cd frontend && pnpm install
cd ../backend && go mod download
```

### テスト（DB 不要）

```bash
# フロントエンド: lint + ユニットテスト
cd frontend && pnpm lint && pnpm test

# バックエンド: ユニットテストのみ
cd backend && go test ./tests/unit/...
```

### テスト（DB 必要）

API 結合テストは PostgreSQL が起動している必要があります。

```bash
# Docker で DB + マイグレーションを起動
docker compose up -d db migrate

# バックエンド API 結合テスト
cd backend && go test ./tests/api/...
```

### ローカル開発サーバー

Docker Compose で一括起動（推奨）:

```bash
./start.sh
# または
docker compose up --build
```

起動後:

- フロントエンド: http://localhost:3010
- バックエンド: http://localhost:8080
- PostgreSQL: localhost:5432

### 環境変数

`.env.example` を `.env` にコピーして使用します。`JWT_SECRET` は 32 文字以上が必要です。

Cloud Agent では Dashboard → Cloud Agents → Secrets に `JWT_SECRET` と `DATABASE_URL` を登録してください。

## コーディング規約

### フロントエンド

- App Router（`frontend/src/app/`）
- スタイルは SCSS Modules（`.module.scss`）
- サーバーアクションは `frontend/src/app/actions/` に配置
- 型定義は `frontend/src/app/types/` に配置
- 画面仕様: [docs/frontend.md](docs/frontend.md)

### バックエンド

- クリーンアーキテクチャ風: `domain` → `repository` → `usecase` → `controller` → `router`
- 認証済み API は `user_id` でスコープを徹底（他ユーザーのデータにアクセスしない）
- API 設計: [docs/backend.md](docs/backend.md)
- アーキテクチャ: [docs/backend-architecture.md](docs/backend-architecture.md)

### ブランチ・PR

- 機能ブランチは `cursor/<descriptive-name>-5869` 形式
- PR のベースブランチは `develop`
- 変更後は lint とテストを実行してから PR を作成

## 変更時のチェックリスト

1. `cd frontend && pnpm lint && pnpm test`
2. `cd backend && go test ./tests/unit/...`
3. API を変更した場合: `go test ./tests/api/...`（DB 起動時）
4. ドキュメントが古くなった場合は `docs/` を更新

## 参照ドキュメント

| ドキュメント | 内容 |
|-------------|------|
| [docs/implementation-flow.md](docs/implementation-flow.md) | 実装フェーズと優先順位 |
| [docs/backend.md](docs/backend.md) | API 設計 |
| [docs/frontend.md](docs/frontend.md) | 画面仕様 |
| [docs/database.md](docs/database.md) | DB 設計 |
| [docs/cursor-automations/](docs/cursor-automations/) | Automations 設定ガイド |
