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
| インフラストラクチャ | Docker |
| CI/CD | GitHub Actions |
| デプロイ環境 | Vercel |
| マイグレーション | golang-migrate |
| 認証 | Supabase Auth |
| ストレージ | Supabase |
| インフラ(将来実装) | Terraform |
| パッケージ管理 | pnpm(Node.js) Go Modules|

## プロジェクト構造
```
cashpilot
├ frontend
├ backend
├ docs
├ docker-compose.yml
├ start.sh
└ .github
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