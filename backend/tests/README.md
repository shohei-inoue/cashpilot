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
