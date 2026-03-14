# CashPilot DB

マイグレーションの実行方法と管理方法をまとめる。

## 概要

- **ツール**: [golang-migrate/migrate](https://github.com/golang-migrate/migrate)
- **DB**: PostgreSQL 16
- **マイグレーション配置**: `migrations/` 直下

## マイグレーションファイル

| ファイル | 内容 |
|----------|------|
| `000001_users` | users テーブル |
| `000002_account_types` | 口座種別マスタ（cash, bank, credit） |
| `000003_accounts` | accounts テーブル |
| `000004_category_types` | カテゴリ種別マスタ（income, expense） |
| `000005_categories` | categories テーブル |
| `000006_transactions` | transactions テーブル + インデックス |
| `000007_goals` | goals テーブル |
| `000008_add_password_hash` | users に password_hash 追加（認証用） |

命名規則: `{連番}_{名前}.up.sql` / `{連番}_{名前}.down.sql`

---

## 実行方法

### 1. Docker 経由（自動）

`docker compose up` 時に **migrate** サービスが自動で `migrate up` を実行する。

```bash
# プロジェクトルートで
docker compose up --build
```

バックエンドは migrate 完了後に起動する。

### 2. Docker 経由（手動）

PostgreSQL（db サービス）が起動している状態で実行。

```bash
# プロジェクトルートで
docker compose up -d db

# マイグレーション適用（up）
docker run --rm -v $(pwd)/backend/db/migrations:/migrations \
  -e POSTGRES_USER=${POSTGRES_USER:-cashpilot} \
  -e POSTGRES_PASSWORD=${POSTGRES_PASSWORD:-cashpilot} \
  -e POSTGRES_DB=${POSTGRES_DB:-cashpilot} \
  --network cashpilot_default \
  migrate/migrate \
  -path /migrations \
  -database "postgres://${POSTGRES_USER:-cashpilot}:${POSTGRES_PASSWORD:-cashpilot}@db:5432/${POSTGRES_DB:-cashpilot}?sslmode=disable" \
  up
```

### 3. ローカル CLI

`migrate` をインストールしている場合。

```bash
# .env の DATABASE_URL を使用する例
migrate -path ./backend/db/migrations \
  -database "postgres://cashpilot:cashpilot@localhost:5432/cashpilot?sslmode=disable" \
  up
```

---

## 管理コマンド

| コマンド | 説明 |
|----------|------|
| `up` | 未適用のマイグレーションをすべて適用 |
| `up 1` | 1 件だけ適用 |
| `down` | 最後に適用したマイグレーションを 1 件ロールバック |
| `down 1` | 同上 |
| `force N` | バージョン N を強制的に「適用済み」としてマーク（破損時など） |
| `version` | 現在のマイグレーションバージョンを表示 |

### ロールバック例

```bash
# 1 件ロールバック
migrate -path ./backend/db/migrations \
  -database "postgres://cashpilot:cashpilot@localhost:5432/cashpilot?sslmode=disable" \
  down 1
```

### 状態リセット（開発時）

既存の DB を削除して最初からやり直す場合。

```bash
docker compose down -v   # ボリューム削除
docker compose up --build
```

---

## 新規マイグレーションの作成

1. `migrations/` 直下にファイルを追加
2. 連番は既存の最大値 + 1（例: 次は `000008_xxx`）
3. `.up.sql` と `.down.sql` のペアを作成

```bash
# 例: migrate CLI で作成
migrate create -ext sql -dir ./backend/db/migrations -seq add_budgets
# -> 000008_add_budgets.up.sql, 000008_add_budgets.down.sql が生成される
```

---

## 環境変数

| 変数 | 説明 |
|------|------|
| `POSTGRES_USER` | DB ユーザー（既定: cashpilot） |
| `POSTGRES_PASSWORD` | DB パスワード（既定: cashpilot） |
| `POSTGRES_DB` | DB 名（既定: cashpilot） |

接続文字列: `postgres://{user}:{password}@{host}:5432/{dbname}?sslmode=disable`

---

## 参照

- DB 設計: [docs/database.md](../../docs/database.md)
- マイグレーション詳細: [docs/migration.md](../../docs/migration.md)
