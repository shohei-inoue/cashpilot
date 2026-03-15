# CashPilot 実装フロー（MVP）

実装の順序と依存関係を整理する。**DB → バックエンド → フロントエンド** の流れで、各フェーズで「動く単位」を小さく作る。

---

## 現状（実装状況）

| フェーズ | 状態 | 備考 |
|----------|------|------|
| Phase 0 | 完了 | Docker Compose、`.env`、モノレポ構成 |
| Phase 1 | 完了 | マイグレーション 001〜008（users, account_types, accounts, category_types, categories, transactions, goals, password_hash） |
| Phase 2 | 完了 | Gin, CORS, GORM 接続, JWT 認証ミドルウェア, `GET /api/health` |
| Phase 3 | 完了 | 認証・user・accounts・categories・transactions・goals・analytics・simulation API 一式 |
| Phase 4 | 完了 | Next.js, 認証 UI（サインアップ・ログイン）, レイアウト（Header/Sidebar）, ナビ（Dashboard / Transactions / Simulation） |
| Phase 5 | 未実装 | 各画面はプレースホルダーのみ。API 連携・ダッシュボード実データ・目標・設定画面が未実装 |
| Phase 6 | 未着手 | ローカル結合確認・エラーハンドリング統一 |

**次にやること（MVP 達成に向けて）**: Phase 5 のフロントエンド画面実装（ダッシュボードに集計・取引表示、取引一覧・登録、シミュレーション入力・結果表示、目標・口座・カテゴリの設定画面の追加と API 連携）。

---

## 全体の流れ（概要）

```
Phase 0: 環境構築
    ↓
Phase 1: DB マイグレーション
    ↓
Phase 2: バックエンド基盤（認証・接続）
    ↓
Phase 3: バックエンド API（CRUD → 集計 → シミュレーション）
    ↓
Phase 4: フロントエンド基盤（Next.js・認証・レイアウト）
    ↓
Phase 5: フロントエンド画面（ダッシュボード → 取引 → 目標・設定 → シミュレーション）
    ↓
Phase 6: 結合・デプロイ
```

---

## Phase 0: 環境構築

| 項目 | 内容 |
|------|------|
| リポジトリ | `frontend` / `backend` / `docs` / `.github` の構成を確定。モノレポ or 別リポジトリは方針に応じて。 |
| 認証 | メール＋パスワード or 外部 ID プロバイダ。認証基盤は別途決定。 |
| 環境変数 | バックエンド: DB URL、認証用の secret 等。フロント: API ベース URL。 |
| ローカル DB | Docker Compose で PostgreSQL をローカル起動（[infra.md](./infra.md) 参照）。 |

**成果物**: ローカル開発環境、`.env.example` の整備。

---

## Phase 1: DB マイグレーション

[database.md](./database.md) の「9. 初期マイグレーション順序」に従う。

| 順番 | マイグレーション | 内容 |
|------|-----------------|------|
| 1 | `001_users` | `users`（id: serial PK、uuid: 認証連携用、email）。 |
| 2 | `002_account_types` | 口座種別マスタ（cash, bank, credit）。タグ選択用。 |
| 3 | `003_accounts` | `accounts`（user_id, account_type_id, name）。 |
| 4 | `004_category_types` | カテゴリ種別マスタ（income, expense）。タグ選択用。 |
| 5 | `005_categories` | `categories`（user_id, category_type_id, name）。 |
| 6 | `006_transactions` | `transactions`（user_id, account_id, category_id, amount, memo, occurred_at）。 |
| 7 | `007_goals` | `goals`（user_id, name, target_amount, deadline）。 |

**補足**
- マイグレーションツール: golang-migrate を使用（[migration.md](./migration.md) 参照）。
- 初回利用時: サインアップ時に `users` に 1 件挿入するか、認証基盤の webhook で作成するか決める。

**成果物**: 上記 7 本のマイグレーションが適用されたスキーマ。必要ならシード（デフォルトカテゴリなど）を追加。

---

## Phase 2: バックエンド基盤

| 順番 | タスク | 内容 |
|------|--------|------|
| 2-1 | プロジェクト初期化 | Go モジュール、Gin ルーター、`/api` プレフィックス。CORS 設定。 |
| 2-2 | DB 接続 | PostgreSQL 接続プール。環境変数から URL を読む。 |
| 2-3 | 認証ミドルウェア | Cookie の JWT を検証し、`user_id` をコンテキストに格納。未認証は 401。詳細は [auth.md](./auth.md)。 |
| 2-4 | ヘルスチェック | `GET /api/health` など、起動・DB 接続確認用。 |

**成果物**: 認証必須の `/api/*` にアクセスできる骨組み。この時点ではまだ CRUD は未実装でよい。

---

## Phase 3: バックエンド API

[backend.md](./backend.md) の API 設計に沿って、**依存の少ないものから** 実装する。

| 順番 | API 群 | 内容 | 依存 |
|------|--------|------|------|
| 3-1 | ユーザー・口座・カテゴリ | `GET /api/user`, `GET/POST/PUT/DELETE /api/accounts`, `GET/POST/PUT /api/categories`。 | users, accounts, categories テーブル |
| 3-2 | 取引 | `GET/POST/PUT/DELETE /api/transactions`。一覧は `from`/`to`/`account_id`/`category_id` でフィルタ。 | accounts, categories が存在すること |
| 3-3 | 目標 | `GET/POST/PUT/DELETE /api/goals`。 | goals テーブル |
| 3-4 | 集計（キャッシュフロー分析） | 月次・日次のキャッシュフロー、残高推移、平均生活費など。`GET /api/transactions/summary` や `GET /api/analytics/cashflow` など、設計に合わせてエンドポイントを 1 本用意。 | transactions |
| 3-5 | シミュレーション | `POST /api/simulation/run`。現在残高（取引の集計）＋リクエストの前提（月収・副業・固定費など）で将来残高を計算。目標（goals）との突き合わせは任意。 | transactions, goals（任意） |

**実装のコツ**
- 各 API は「認証済みユーザーの `user_id` でスコープ」を徹底。他ユーザーのデータは返さない・更新しない。
- 集計・シミュレーションは service レイヤにロジックをまとめ、handler は薄く保つ。

**成果物**: フロントから呼べる CRUD ＋ 集計 ＋ シミュレーション API。

---

## Phase 4: フロントエンド基盤

| 順番 | タスク | 内容 |
|------|--------|------|
| 4-1 | プロジェクト初期化 | Next.js (App Router), TypeScript, SCSS Modules。環境変数で API ベース URL を設定。 |
| 4-2 | 認証 UI | サインアップ・ログイン画面。認証基盤に合わせて実装。ログイン後にダッシュボードへリダイレクト。 |
| 4-3 | 認証状態の保持 | セッション・トークンの保持。API 呼び出し時に `Authorization: Bearer <token>` を付与する共通クライアント（fetch または Axios）を用意。 |
| 4-4 | レイアウト・ナビ | 認証済みユーザー向けのレイアウト（ヘッダー、ナビ: ダッシュボード／収支入力／シミュレーション／目標・設定、ログアウト）。未認証時はログインへリダイレクト。 |

**成果物**: ログインしてナビだけある「空のダッシュボード」まで表示できる状態。

---

## Phase 5: フロントエンド画面

[frontend.md](./frontend.md) の画面一覧に沿って、**データの流れが分かりやすい順** で実装する。

| 順番 | 画面・機能 | 内容 | 依存 API |
|------|------------|------|----------|
| 5-1 | ダッシュボード | 現在残高（集計 API または取引一覧から計算）、今月の収支サマリー、直近取引一覧。将来残高グラフは 5-4 で実装後、簡易表示してもよい。 | user, accounts, transactions, 集計 |
| 5-2 | 口座・カテゴリの設定 | 設定画面の一部として、口座一覧（追加・編集・削除）、カテゴリ一覧（追加・編集）。 | accounts, categories |
| 5-3 | 取引登録・一覧 | 収支入力画面。口座・カテゴリ・金額・日時・メモで 1 件登録。一覧は期間・口座・カテゴリでフィルタ可能に。 | accounts, categories, transactions |
| 5-4 | 目標（goals） | 目標一覧・追加・編集・削除。設定画面または専用画面。 | goals |
| 5-5 | シミュレーション | 期間・予定収支（月収・副業・固定費など）を入力し、`POST /api/simulation/run` を呼ぶ。結果をグラフ＋数値で表示。目標との突き合わせがあれば表示。 | simulation/run, goals（任意） |
| 5-6 | プロフィール・設定の統合 | 口座・カテゴリ・目標を 1 つの「設定」または「目標・設定」画面にまとめ、ナビと 5-2 / 5-4 を対応させる。 | 上記すべて |

**実装のコツ**
- ダッシュボードと取引一覧で「残高」「今月の収支」を同じ集計 API で揃える。
- シミュレーションは入力フォームと結果表示を同じページにまとめ、初回表示時はデフォルト期間だけ渡して結果を出してもよい。

**成果物**: 認証から家計管理・分析・シミュレーションまで一通り操作できる UI。

---

## Phase 6: 結合・デプロイ

ローカルでの結合確認を優先。本番デプロイは実装が進んでから検討（AWS / Fly.io 等、[infra.md](./infra.md) 参照）。

| 順番 | タスク | 内容 |
|------|--------|------|
| 6-1 | ローカル結合 | バックエンドとフロントを同時に起動し、サインアップ → 口座・カテゴリ作成 → 取引登録 → ダッシュボード・シミュレーション確認。 |
| 6-2 | エラーハンドリングの統一 | API エラー時のメッセージ表示、ネットワークエラー時のリトライやメッセージ。 |
| 6-3 | 本番デプロイ（将来） | 実装完了後にデプロイ環境（AWS / Fly.io 等）を選定し、環境変数・CI/CD を整備。 |

**成果物**: ローカルで一連の操作ができる状態。ドキュメント（README、環境構築手順）の更新。

---

## マイルストーン（チェックリスト）

- [x] **M1** Phase 0〜1 完了: ローカル環境構築とマイグレーション適用
- [x] **M2** Phase 2〜3-3 完了: 認証付きで CRUD API がすべて動く
- [x] **M3** Phase 3-4〜3-5 完了: 集計・シミュレーション API が動く
- [x] **M4** Phase 4 完了: ログイン〜ダッシュボード骨組みまで表示
- [ ] **M5** Phase 5 完了: 全画面がつながり、一通り操作できる
- [ ] **M6** Phase 6 完了: ローカル結合確認が完了（本番デプロイは将来検討）

---

## 参照

- [README.md](./README.md) - 全体方針・目的
- [database.md](./database.md) - DB 設計・マイグレーション順
- [backend.md](./backend.md) - API 設計・ドメイン
- [frontend.md](./frontend.md) - 画面・UI 仕様
- [infra.md](./infra.md) - ローカル開発環境。デプロイは将来検討。
