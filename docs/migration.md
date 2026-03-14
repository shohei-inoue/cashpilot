# DB マイグレーション

golang-migrate の [migrate/migrate](https://github.com/golang-migrate/migrate) イメージを使用。

---

## マイグレーションファイル

`backend/db/migrations/` 直下に配置。

| ファイル | 内容 |
|----------|------|
| `000001_users` | users テーブル |
| `000002_account_types` | 口座種別マスタ（cash, bank, credit） |
| `000003_accounts` | accounts テーブル |
| `000004_category_types` | カテゴリ種別マスタ（income, expense） |
| `000005_categories` | categories テーブル |
| `000006_transactions` | transactions テーブル + インデックス |
| `000007_goals` | goals テーブル |

---

## Docker 経由での実行

`docker compose up` 時に **migrate** サービスが自動で `migrate up` を実行する。  
バックエンドは migrate 完了後に起動する。

```bash
docker compose up --build
```

---

## 手動実行（ローカル）

PostgreSQL が起動している状態で:

```bash
# マイグレーション適用
docker run --rm -v $(pwd)/backend/db/migrations:/migrations \
  -e POSTGRES_USER=cashpilot \
  -e POSTGRES_PASSWORD=cashpilot \
  -e POSTGRES_DB=cashpilot \
  --network cashpilot_default \
  migrate/migrate \
  -path /migrations \
  -database "postgres://cashpilot:cashpilot@db:5432/cashpilot?sslmode=disable" \
  up
```

ローカルで migrate CLI をインストールしている場合:

```bash
migrate -path ./backend/db/migrations -database "postgres://cashpilot:cashpilot@localhost:5432/cashpilot?sslmode=disable" up
```

---

## ロールバック（down）

```bash
migrate -path ./backend/db/migrations -database "postgres://..." down 1
```

---

## 参照

- DB 設計: [database.md](./database.md)
