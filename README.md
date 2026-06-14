# CashPilot

CashPilotは、個人のキャッシュフローシミュレーターと収支トラッカーです。
収入、支出、副業が将来の財政状況にどのように影響するかを理解するのに役立ちます。

## 機能

- 収入管理
- 支出追跡
- キャッシュフローシミュレーション
- 貯蓄予測

## ユーザー
- 個人
- 家計

## 利用シーン
- 月次見直し
- 金銭の関係する目標設定
- 副業検討

## 技術スタック

| カテゴリ | 技術 |
|----------|------|
| フロントエンド | TypeScript, Next.js, React, SCSS Modules |
| バックエンド | Go, Gin |
| データベース | PostgreSQL |
| インフラストラクチャ | Docker（ローカル開発） |
| CI/CD | GitHub Actions |
| デプロイ | 将来検討（AWS / Fly.io 等） |
| マイグレーション | golang-migrate |
| 認証 | JWT + httpOnly Cookie（メール＋パスワード） |
| パッケージ管理 | pnpm（Node.js）、Go Modules |

## 現在の進捗（MVP）

- **バックエンド**: Phase 0〜3 完了（環境・マイグレーション・認証・CRUD・集計・シミュレーション API）
- **フロントエンド**: Phase 4 完了（認証 UI・レイアウト・ナビ）。Phase 5 は各画面の API 連携・目標・設定画面が未実装

詳細なマイルストーンと次にやることは [docs/implementation-flow.md](docs/implementation-flow.md) を参照。

## Cursor Automations

Cloud Agent を使った自動レビュー・Phase 5 実装支援の設定は [docs/cursor-automations/](docs/cursor-automations/) を参照。エージェント向け手順は [AGENTS.md](AGENTS.md)、VM 環境は [.cursor/environment.json](.cursor/environment.json) に定義しています。

## プロジェクト構造

```
cashpilot/
├── frontend/                    # Next.js フロントエンド（App Router, TypeScript, SCSS Modules）
│   ├── src/
│   │   └── app/
│   │       ├── (pages)/         # ルートグループ
│   │       │   ├── (dashboard)/ # ダッシュボード /
│   │       │   ├── auth/        # 認証（login, signup）
│   │       │   ├── transactions/ # 取引一覧・登録
│   │       │   └── simulation/  # シミュレーション
│   │       ├── components/      # Header, Sidebar, MainContainer, Form, Button, Card など
│   │       ├── actions/         # サーバーアクション（auth, user）
│   │       ├── libs/            # API クライアント（client.ts）, response 型
│   │       ├── constants/       # ナビ項目など
│   │       ├── types/           # 型定義（user, ui）
│   │       └── styles/           # グローバル SCSS（globals, variables, mixin）
│   ├── public/
│   ├── docs/                    # フロント用ドキュメント（design-system, pages）
│   ├── package.json
│   └── Dockerfile
│
├── backend/                     # Go + Gin API
│   ├── cmd/
│   │   └── api/
│   │       └── main.go          # エントリポイント・ルート登録
│   ├── db/
│   │   ├── postgres.go          # GORM / pgx 接続
│   │   ├── migrations/          # golang-migrate 用 SQL
│   │   └── README.md
│   ├── internal/
│   │   ├── config/              # 環境変数（DB URL, JWT secret）
│   │   ├── jwt/                 # JWT 発行・検証
│   │   ├── middleware/          # CORS, 認証（Cookie JWT → user_id）
│   │   ├── response/            # 共通レスポンス・エラーコード
│   │   ├── apperrors/           # アプリケーションエラー
│   │   └── logic/               # クリーンアーキテクチャ風
│   │       ├── domain/          # エンティティ（user, account, category, transaction, goal, analytics, simulation）
│   │       ├── repository/      # DB アクセス
│   │       ├── usecase/         # ビジネスロジック
│   │       ├── controller/      # HTTP ハンドラー
│   │       └── router/          # ルート定義（health, auth, user, accounts, categories, transactions, goals, analytics, simulation）
│   ├── tests/
│   │   ├── api/                 # 統合テスト（各 API）
│   │   └── unit/
│   ├── docker/                  # Dockerfile, .air.toml（ホットリロード）
│   ├── go.mod
│   └── README.md
│
├── docs/                        # 仕様・設計ドキュメント
│   ├── README.md                # 全体方針・MVP スコープ
│   ├── implementation-flow.md  # 実装順序・マイルストーン
│   ├── database.md              # DB 設計
│   ├── backend.md               # API 設計・ドメイン
│   ├── backend-architecture.md
│   ├── frontend.md              # 画面一覧・UI
│   ├── auth.md                  # 認証設計
│   ├── migration.md             # マイグレーション手順
│   └── infra.md                 # ローカル環境・デプロイ方針
│
├── docker-compose.yml           # FE, BE, PostgreSQL, migrate
├── start.sh                     # 一括起動スクリプト
├── .env.example
└── .github/                     # CI/CD（GitHub Actions）
```

## ローカル開発（Docker）

FE・BE・PostgreSQL を一括で起動する場合:

```bash
./start.sh
# または
docker compose up --build
```

起動後:
- フロントエンド: http://localhost:3010
- バックエンド: http://localhost:8080
- PostgreSQL: localhost:5432

初回は `.env.example` をコピーして `.env` を作成する（`start.sh` が自動作成する）。