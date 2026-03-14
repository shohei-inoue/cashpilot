# バックエンド アーキテクチャ設計（クリーンアーキテクチャ）

CashPilot バックエンドは **クリーンアーキテクチャ** に沿って設計する。  
レイヤーは Domain / Usecase / Controller / Repository に分け、`internal` ディレクトリ直下で構成する。

---

## 1. ディレクトリ構成

```
backend/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── domain/         # ドメインエンティティ
│   ├── usecase/        # ユースケース（ビジネスロジック）
│   ├── controller/     # HTTP ハンドラー（入力アダプタ）
│   ├── repository/     # リポジトリインターフェース
│   └── router/         # 各 API のルート定義
├── db/                 # DB 関連（リポジトリ実装・PostgreSQL 等）
├── pkg/                # 外部公開ユーティリティ（任意）
└── go.mod
```

### ルーティングの扱い

- **router/**  
  各 API リソースごとのルート定義を記述する。
  - 例: `transaction_router.go`, `user_router.go`, `goal_router.go`
  - `*RouterGroup` に対して `GET`, `POST` などを登録
- **controller/**  
  ルートから呼ばれる HTTP ハンドラー実装。リクエストのパース・レスポンス生成・ユースケース呼び出しを行う。
- **まとめ方**  
  `main.go` または `router/router.go` で、`controller` を生成し、`router` の各関数に渡してルートを一括で登録する。

---

## 2. 各レイヤーの役割と依存関係

| レイヤー | ディレクトリ | 役割 | 依存先 |
|----------|--------------|------|--------|
| Domain | `internal/domain/` | エンティティ定義。他レイヤーに依存しない | なし |
| Repository（I/F） | `internal/repository/` | データ取得・永続化のインターフェース | Domain |
| Usecase | `internal/usecase/` | アプリケーション固有のビジネスロジック | Domain, Repository（I/F） |
| Controller | `internal/controller/` | HTTP リクエストの受け取り、Usecase 呼び出し、レスポンス返却 | Usecase |
| Router | `internal/router/` | ルート定義と Controller へのマッピング | Controller |
| DB | `db/` | リポジトリ実装（PostgreSQL 等） | Repository（I/F）, Domain |

依存の向き: Domain ← Repository(I/F) ← Usecase ← Controller ← Router。db は Repository を実装する。

---

## 3. ファイル命名ルール

| 種類 | 命名 | 例 |
|------|------|-----|
| エンティティ | 単数形 | `user.go`, `account.go`, `transaction.go`, `goal.go`, `category.go` |
| ユースケース | `_usecase.go` | `transaction_usecase.go`, `simulation_usecase.go` |
| コントローラー | `_controller.go` | `transaction_controller.go`, `goal_controller.go` |
| リポジトリ（I/F） | `_repository.go` | `transaction_repository.go`, `account_repository.go` |
| リポジトリ実装 | `_repository_impl.go` または `postgres_*.go` | `postgres_transaction_repository.go` |
| ルーター | `_router.go` | `transaction_router.go`, `goal_router.go` |

パッケージ名はディレクトリ名に合わせる（`domain`, `usecase`, `controller`, `repository`, `router`, `db`）。

---

## 4. ディレクトリ・ファイル一覧（MVP 想定）

```
backend/
├── internal/
│   ├── domain/
│   │   ├── user.go
│   │   ├── account.go
│   │   ├── category.go
│   │   ├── transaction.go
│   │   └── goal.go
│   ├── usecase/
│   │   ├── account_usecase.go
│   │   ├── category_usecase.go
│   │   ├── transaction_usecase.go
│   │   ├── goal_usecase.go
│   │   └── simulation_usecase.go
│   ├── controller/
│   │   ├── account_controller.go
│   │   ├── category_controller.go
│   │   ├── transaction_controller.go
│   │   ├── goal_controller.go
│   │   └── simulation_controller.go
│   ├── repository/
│   │   ├── account_repository.go
│   │   ├── category_repository.go
│   │   ├── transaction_repository.go
│   │   └── goal_repository.go
│   └── router/
│       ├── router.go           # 全ルートの集約・登録
│       ├── account_router.go
│       ├── category_router.go
│       ├── transaction_router.go
│       ├── goal_router.go
│       └── simulation_router.go
└── db/
    ├── migrations/              # DB マイグレーション（001〜007）
    │   ├── 000001_users.up.sql
    │   ├── 000002_account_types.up.sql
    │   ├── 000003_accounts.up.sql
    │   └── ...
    ├── account_repository_impl.go
    ├── category_repository_impl.go
    ├── transaction_repository_impl.go
    └── goal_repository_impl.go
```

---

## 5. ルーティングの流れ

1. `main.go` で Gin エンジンを作成し、認証ミドルウェアを設定。
2. `/api` をプレフィックスとする `RouterGroup` を作成。
3. `router/router.go` の `Setup(r *gin.RouterGroup, c *Controllers)` で、各 `*_router.go` を呼び出し、ルートを登録。
4. 各 `*_router.go` は、受け取った `Controller` のメソッドを `r.GET`, `r.POST` などにバインド。

```go
// router/router.go のイメージ
func Setup(r *gin.RouterGroup, c *Controllers) {
    AccountRouter(r.Group("/accounts"), c.Account)
    CategoryRouter(r.Group("/categories"), c.Category)
    TransactionRouter(r.Group("/transactions"), c.Transaction)
    GoalRouter(r.Group("/goals"), c.Goal)
    SimulationRouter(r.Group("/simulation"), c.Simulation)
}
```

---

## 6. 参照

- API 設計・エンドポイント詳細: [backend.md](./backend.md)
- DB 設計: [database.md](./database.md)
