# Backend テスト

バックエンド API のテストディレクトリ。

## 構成

```
tests/
├── api/           # API / 結合テスト
│                 # ルーター経由でエンドポイントを叩くテスト
├── unit/
│   ├── controller/   # コントローラーのユニットテスト
│   ├── repository/   # リポジトリのユニットテスト
│   └── usecase/      # ユースケースのユニットテスト
└── README.md
```

## 実行方法

```bash
# 全テスト
go test ./tests/...

# api のみ
go test ./tests/api/...

# unit のみ
go test ./tests/unit/...
```

## 注意（API テスト）

`tests/api/auth_test.go`、`tests/api/account_test.go`、`tests/api/category_test.go`、`tests/api/transaction_test.go`、`tests/api/goal_test.go` は DB 接続が必要。PostgreSQL が起動している状態で実行する（例: `./start.sh` で Docker 起動後）。

DB に接続できない場合は `t.Skip` でスキップされる。