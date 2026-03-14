# フロントエンド仕様

CashPilot フロントエンドの詳細仕様。  
ルートは `(pages)` 配下とし、トップをダッシュボードとする。

---

## ルーティング

| パス | 画面 | 認証 |
|------|------|------|
| `/` | ダッシュボード | 必須 |
| `/goals` | 目標 | 必須 |
| `/transactions` | 取引（収支入力・一覧） | 必須 |
| `/simulation` | シミュレーション | 必須 |
| `/login` | サインアップ／ログイン | 不要 |

---

## ドキュメント一覧

| ファイル | 内容 |
|----------|------|
| [design-system.md](./design-system.md) | デザインシステム（カラー・タイポ・スペーシング） |
| [components.md](./components.md) | 共通コンポーネント TODO |
| [pages/dashboard.md](./pages/dashboard.md) | ダッシュボード画面仕様 |
| [pages/goals.md](./pages/goals.md) | 目標画面仕様 |
| [pages/transactions.md](./pages/transactions.md) | 取引画面仕様 |
| [pages/simulation.md](./pages/simulation.md) | シミュレーション画面仕様 |

---

## 参照

- [frontend.md](../frontend.md) - 全体方針・画面一覧
- [auth.md](../auth.md) - 認証（JWT + Cookie）
