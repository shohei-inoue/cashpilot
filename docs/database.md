# CashPilot DB設計 v1

Ledger型設計 + User ownership を採用し、将来拡張も見据えたスキーマとする。

## 目的（他ドキュメントと共通）

本 DB が支えるプロダクトの目的は次の 4 つ（[docs/README.md](./README.md) と同一）。

- 家計管理
- キャッシュフロー分析
- シミュレーション
- 将来マルチユーザー対応

---

## ER構造

```
users
  │
  ├── account_types（口座種別マスタ: cash, bank, credit）
  ├── accounts      ← account_type_id で account_types を参照
  │
  ├── category_types（カテゴリ種別マスタ: income, expense）
  ├── categories    ← category_type_id で category_types を参照
  │
  ├── transactions
  │
  └── goals
```

**将来追加予定**

- `budgets`
- `recurring_transactions`
- `simulation_runs`

---

## 1. users

- **id**: serial（連番）を主キーとする。内部の参照・結合には `id` を使用。
- **uuid**: 認証連携用。外部認証基盤のユーザー ID と紐付ける場合はこのカラムを使用。

```sql
CREATE TABLE users (
  id serial PRIMARY KEY,
  uuid uuid UNIQUE NOT NULL DEFAULT gen_random_uuid(),
  email text UNIQUE NOT NULL,
  created_at timestamp DEFAULT now()
);
```

---

## 2. account_types（口座種別マスタ）

タグのように選択する用。初期データ: cash, bank, credit。

```sql
CREATE TABLE account_types (
  id serial PRIMARY KEY,
  name text UNIQUE NOT NULL
);
```

---

## 3. accounts

お金の「場所」。`account_type_id` で種別（account_types）を参照する。

```sql
CREATE TABLE accounts (
  id serial PRIMARY KEY,
  user_id integer NOT NULL,
  account_type_id integer NOT NULL,
  name text NOT NULL,
  created_at timestamp DEFAULT now(),
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (account_type_id) REFERENCES account_types(id)
);

CREATE INDEX idx_accounts_user ON accounts(user_id);
```

---

## 4. category_types（カテゴリ種別マスタ）

タグのように選択する用。初期データ: income, expense。

```sql
CREATE TABLE category_types (
  id serial PRIMARY KEY,
  name text UNIQUE NOT NULL
);
```

---

## 5. categories

支出・収入のカテゴリ。`category_type_id` で種別（category_types）を参照する。

```sql
CREATE TABLE categories (
  id serial PRIMARY KEY,
  user_id integer NOT NULL,
  category_type_id integer NOT NULL,
  name text NOT NULL,
  created_at timestamp DEFAULT now(),
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (category_type_id) REFERENCES category_types(id)
);

CREATE INDEX idx_categories_user ON categories(user_id);
```

---

## 6. transactions（最重要）

すべての金銭移動を記録する。**Ledger の中心**。

```sql
CREATE TABLE transactions (
  id serial PRIMARY KEY,
  user_id integer NOT NULL,
  account_id integer NOT NULL,
  category_id integer,
  amount integer NOT NULL,
  memo text,
  occurred_at timestamp NOT NULL,
  created_at timestamp DEFAULT now(),

  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (account_id) REFERENCES accounts(id),
  FOREIGN KEY (category_id) REFERENCES categories(id)
);
```

### amount ルール（推奨）

- **収入**: 正の値（`+value`）
- **支出**: 負の値（`-value`）

例:

- 給与: `+300000`
- 食費: `-1200`
- 家賃: `-80000`
- 副業収入: `+50000`

このルールにすると、**キャッシュフロー** は次の 1 クエリで得られる。

```sql
SELECT SUM(amount) FROM transactions WHERE user_id = $1;
```

---

## 7. goals

資金目標（引越し・購入・返済など）。

```sql
CREATE TABLE goals (
  id serial PRIMARY KEY,
  user_id integer NOT NULL,
  name text NOT NULL,
  target_amount integer NOT NULL,
  deadline date,
  created_at timestamp DEFAULT now(),

  FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE INDEX idx_goals_user ON goals(user_id);
```

例: 引越し、MacBook、借金返済

---

## 8. 重要インデックス

家計アプリは **期間クエリ** が頻出するため、以下を必須とする。

```sql
-- ユーザー×日付（一覧・集計の主力）
CREATE INDEX idx_transactions_user_date
ON transactions(user_id, occurred_at);

-- カテゴリ別集計
CREATE INDEX idx_transactions_category
ON transactions(category_id);
```

---

## 9. 代表的なクエリ

### 月次キャッシュフロー

```sql
SELECT
  date_trunc('month', occurred_at) AS month,
  SUM(amount) AS total
FROM transactions
WHERE user_id = $1
GROUP BY month
ORDER BY month;
```

### 日別支出

```sql
SELECT
  date_trunc('day', occurred_at) AS day,
  SUM(amount) AS daily_total
FROM transactions
WHERE user_id = $1 AND amount < 0
GROUP BY day
ORDER BY day;
```

### 平均生活費（日次支出の平均）

```sql
SELECT AVG(daily_total) AS avg_daily_expense
FROM (
  SELECT
    date_trunc('day', occurred_at) AS day,
    SUM(amount) AS daily_total
  FROM transactions
  WHERE user_id = $1 AND amount < 0
  GROUP BY day
) t;
```

---

## 10. CashPilot 固有の機能（シミュレーション）

**副業シミュレーション** などは、入力（`hourly_rate`, `hours`）に基づき  
`side_income = hourly_rate × hours` を **ランタイム計算** で行う。  
DB にシミュレーション結果を保存しない形でよい（保存する場合は将来 `simulation_runs` を利用）。

---

## 11. 初期マイグレーション順序

依存関係に従い、以下の順で実行する。

1. `001_users`
2. `002_account_types`
3. `003_accounts`
4. `004_category_types`
5. `005_categories`
6. `006_transactions`
7. `007_goals`

---

## 12. この設計のメリット

- 家計管理
- 資産管理（口座別残高）
- クレカ管理（account type = credit）
- 借金管理（負の残高 or 別テーブルで拡張）
- 副業分析（category + シミュレーション）

を **同一の Ledger モデル** で扱える。

---

## 13. 将来追加テーブル（v2 以降）

| テーブル | 用途 |
|----------|------|
| `budgets` | 予算（月別・カテゴリ別など） |
| `recurring_transactions` | 定期収支（家賃・給与・サブスク等）のテンプレート。**シミュレーション・予測の精度向上に直結** |
| `debts` | 借入・返済計画（任意） |
| `simulation_runs` | シミュレーション結果の保存・比較 |

**recurring_transactions** を入れると、「毎月の固定収支」を定義でき、  
「Xヶ月後の残高」「目標達成に必要な月数」などのシミュレーションが安定して設計しやすくなる。

---

## 14. 参照

- バックエンド API 設計・ドメインロジック: [backend.md](./backend.md)
- 全体方針: [README.md](./README.md)
