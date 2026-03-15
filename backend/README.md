# CashPilot バックエンド

CashPilot のバックエンド API。Go + Gin で構築。

## 現状

Phase 0〜3 まで実装済み。認証・CRUD・集計・シミュレーション API が利用可能。

- **ルート**: `GET /` … 疎通確認（`{"message": "CashPilot API", "status": "ok"}`）
- **ヘルス**: `GET /api/health` … 起動・DB 接続確認
- **認証**: `POST /api/auth/signup`, `POST /api/auth/login`, `POST /api/auth/logout`
- **ユーザー**: `GET /api/user`（認証必須）
- **口座**: `GET/POST/PUT/DELETE /api/accounts`（認証必須）
- **カテゴリ**: `GET/POST/PUT/DELETE /api/categories`（認証必須）
- **取引**: `GET/POST/PUT/DELETE /api/transactions`, `GET /api/transactions/:id`（認証必須）
- **集計**: `GET /api/transactions/summary`, `GET /api/analytics/cashflow`（認証必須）
- **目標**: `GET/POST/PUT/DELETE /api/goals`, `GET /api/goals/:id`（認証必須）
- **シミュレーション**: `POST /api/simulation/run`（認証必須）

詳細なリクエスト・レスポンスは [docs/backend.md](../docs/backend.md) を参照。

## 開発

### ローカルで実行

```bash
go run ./cmd/api
```

### Docker でホットリロード（Air）

プロジェクトルートで `./start.sh` または `docker compose up --build` を実行すると、バックエンドは Air で起動し、Go ファイルの編集で自動ビルド・再起動されます。

Docker 関連の設定は `docker/` にまとまっています（`Dockerfile`, `.air.toml`）。

サーバー起動後、[http://localhost:8080](http://localhost:8080) にアクセスして動作を確認できる。

## 技術スタック

- **Go** 1.25
- **Gin** … HTTP フレームワーク

## 参照

- API 設計・ドメイン・シミュレーション仕様: [docs/backend.md](../docs/backend.md)
- DB 設計: [docs/database.md](../docs/database.md)
- 実装フロー: [docs/implementation-flow.md](../docs/implementation-flow.md)
