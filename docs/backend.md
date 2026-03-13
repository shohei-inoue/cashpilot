## バックエンド仕様（MVP）

CashPilot のバックエンドは Go + Gin を用いた REST API とし、プロダクトの**目的**（[docs/README.md](./README.md)・[database.md](./database.md) に合わせる）のうち、次の4つを支える責務を持つ。

- **家計管理**: ユーザーごとの口座・カテゴリ・取引（transactions）のCRUDと一覧・集計APIを提供する。
- **キャッシュフロー分析**: 取引データに基づく月次・日次のキャッシュフロー集計、残高推移、平均生活費などのAPIを提供する。
- **シミュレーション**: 現在残高と予定収支をもとに将来残高や目標達成時期を計算し、結果を返すAPIを提供する。
- **将来マルチユーザー対応**: すべてのリソースを `user_id` でスコープし、認証済みユーザー以外のデータにアクセスできないようにする。

---

### 1. 全体方針（バックエンド）

- **認証**
  - Supabase Auth など外部認証基盤を利用し、バックエンドでは「発行済みトークンの検証」のみ行う。
  - すべてのアプリ本体 API は「認証必須」とする（MVP時点でパブリックAPIは想定しない）。
- **ドメイン**
  - ユーザーごとの口座（accounts）・カテゴリ（categories）・取引（transactions）・目標（goals）を管理し、家計管理・キャッシュフロー分析・シミュレーションに必要なデータと計算結果を提供する。
- **アーキテクチャ**
  - シンプルなレイヤード構成を想定:
    - handler（Gin のハンドラー / HTTP レイヤ）
    - service（ドメインロジック・ユースケース）
    - repository（DB アクセス）
- **DB設計**
  - **スキーマ・DDL・インデックス・代表クエリは [database.md](./database.md) を参照**。本ドキュメントでは API とドメインの対応のみ要約する。

---

### 2. ドメインモデル（MVP）— DB との対応

DB 設計は **Ledger 型 + User ownership**。正本は [database.md](./database.md)。

| エンティティ | 役割 | 備考 |
|-------------|------|------|
| **users** | 認証主体。Supabase `auth.users.id` と 1:1。 | 全データは `user_id` でスコープ。 |
| **accounts** | お金の「場所」（銀行・クレカ・現金）。 | `type`: cash / bank / credit。 |
| **categories** | 収入・支出のカテゴリ。 | `type`: income / expense。ユーザーごと。 |
| **transactions** | 1件の金銭移動。Ledger の中心。 | `amount`: 収入は正・支出は負。`occurred_at` で時系列。 |
| **goals** | 資金目標（引越し・購入・返済など）。 | `target_amount`, `deadline`。 |

**Transaction の設計ポイント**

- **amount**: 収入は正の整数、支出は負の整数で保存。`SUM(amount)` でキャッシュフローが得られる。
- **account_id / category_id**: 必ず口座と（任意で）カテゴリに紐づく。
- **occurred_at**: 取引が発生した日時。集計・シミュレーションはこの日付基準。

**シミュレーション（MVP）**

- 将来残高の計算は、`transactions` の集計と、ランタイムで渡す「予定収支」（時給×時間など）を組み合わせて **API 側で計算**。結果の永続化は将来 `simulation_runs` で対応可能。
- 目標達成時期の試算は、`goals` の `target_amount` / `deadline` と、上記キャッシュフローから算出する。

---

### 3. API設計（MVP）

ベースURL は `/api` を想定。

#### 3-1. 認証関連

バックエンドとしては、主に「トークン検証ミドルウェア」を提供するイメージ。

- **HTTP Header**
  - `Authorization: Bearer <access_token>`
- Gin のミドルウェアでトークンを検証し、`user_id` をコンテキストに埋め込む。

#### 3-2. ユーザー・アカウント・カテゴリ API

- **GET `/api/user`**  
  - 認証済みユーザー情報（`users` の id, email など）を返却。
- **GET `/api/accounts`**  
  - 当該ユーザーの `accounts` 一覧。
- **POST `/api/accounts`**, **PUT `/api/accounts/:id`**, **DELETE `/api/accounts/:id`**  
  - 口座の作成・更新・削除。`type`: cash / bank / credit。
- **GET `/api/categories`**  
  - 当該ユーザーの `categories` 一覧（`type`: income / expense）。
- **POST `/api/categories`**, **PUT `/api/categories/:id`**  
  - カテゴリの作成・更新（必要に応じて）。

#### 3-3. トランザクション API

- **GET `/api/transactions`**
  - クエリ: `from`, `to`（日付）, `account_id`, `category_id`, `limit`, `offset`。
  - レスポンス: トランザクション一覧（`occurred_at` 降順）+ 期間内の `SUM(amount)` など集計値を含めてもよい。
- **POST `/api/transactions`**
  - ボディ: `account_id`, `category_id`（任意）, `amount`（符号付き）, `memo`, `occurred_at`。
  - 1件登録。`user_id` は認証から補完。
- **PUT `/api/transactions/:id`**, **DELETE `/api/transactions/:id`**  
  - 更新・削除。常に当該ユーザーに属するもののみ操作可能。

#### 3-4. 目標（goals）API

- **GET `/api/goals`**  
  - 当該ユーザーの `goals` 一覧。
- **POST `/api/goals`**  
  - ボディ: `name`, `target_amount`, `deadline`（任意）。
- **PUT `/api/goals/:id`**, **DELETE `/api/goals/:id`**  
  - 目標の更新・削除。

#### 3-5. シミュレーション実行 API

- **POST `/api/simulation/run`**
  - ボディ例: `period_months`, 任意で `hourly_rate` / `hours`（副業シミュレーション）など。
  - 処理:
    1. 当該ユーザーの `transactions` を `occurred_at` で集計し、現在時点までの残高（または口座別残高）を算出。
    2. 将来分は、ランタイムで渡した前提（例: 月収・副業見込み・固定費）で月次キャッシュフローを計算。
    3. `goals` の `deadline` / `target_amount` と照らして、目標達成可否・不足額などを算出してもよい。
  - レスポンス:
    - `monthly_balances`: `[{ month, balance }, ...]`
    - `min_balance`, `end_balance`
    - 必要に応じて `goals` ごとの達成見込み

---

### 4. シミュレーションロジック（MVP 概要）

- **現在残高**: 各口座（または全体）について、`transactions` の `SUM(amount)` でキャッシュフローを計算。口座別の「初期残高」をどう扱うかは実装方針（例: 口座マスタに `initial_balance` を持たせる、または最初の残高調整用トランザクションで表現する）。
- **将来シミュレーション**: リクエストで渡す「月収・副業（hourly_rate × hours）・固定費」などで月次を積み上げ、`period_months` 分の残高時系列を返す。DB にシミュレーション結果を保存するかは任意（将来 `simulation_runs` で保存可）。
- **目標（goals）**: `deadline` と `target_amount` から、その時点での予測残高との差（不足額／達成可否）をランタイム計算で返す。

---

### 5. エラーハンドリング・バリデーション（MVP）

- **エラー形式**
  - JSON `{ "error": { "code": "bad_request", "message": "..." } }` のようなシンプルな形式。
- **主なバリデーション**
  - `amount` は整数（収入は正・支出は負）。0 は不可とするか、許容するかは方針次第。
  - `account_id` / `category_id` は当該ユーザーに属するもののみ。
  - 日付／期間は `from <= to`。シミュレーション期間は上限（例: 10年 = 120ヶ月）を設ける。

---

### 6. セキュリティ・認可（MVP）

- すべてのリソースは `user_id` でスコープされる。
- 認証ミドルウェアで抽出した `user_id` 以外のデータにはアクセスできないようにクエリレベルで制限。
- 管理者機能などは MVP では考慮しない。

---

### 7. 今後の拡張余地（MVP以降）

- **DB**: `budgets`, `recurring_transactions`, `simulation_runs`（[database.md](./database.md) 参照）
- 複数シミュレーション・シナリオ比較（A/B）
- カテゴリ別・口座別レポート API
- 集計キャッシュ・バッチ
- 外部API連携（銀行・クレカ等）

