# CashPilot バックエンド

CashPilot のバックエンド API。Go + Gin で構築。

## 現状

開発初期段階。現時点では起動確認用の簡単なエンドポイントのみ実装済み。

- `GET /` … JSON レスポンス（`{"message": "Hello, World!", "status": "ok"}`）

## 開発

```bash
go run main.go
```

サーバー起動後、[http://localhost:8080](http://localhost:8080) にアクセスして動作を確認できる。

## 技術スタック

- **Go** 1.25
- **Gin** … HTTP フレームワーク

## 参照

- API 設計・ドメイン・シミュレーション仕様: [docs/backend.md](../docs/backend.md)
- DB 設計: [docs/database.md](../docs/database.md)
- 実装フロー: [docs/implementation-flow.md](../docs/implementation-flow.md)
