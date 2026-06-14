# Cursor Automations セットアップガイド（cashpilot）

CashPilot で Cursor Automations を使う前に、Cursor Dashboard で行う初期設定の手順です。

## 前提条件


| 要件              | 内容                                           |
| --------------- | -------------------------------------------- |
| 有料プラン           | Cloud Agents の利用に必要                          |
| On-Demand Usage | Automation 作成時に **従量課金の有効化が必須**（$1 以上の上限でも可） |
| Git 接続          | コード変更系には GitHub / GitLab の read-write 接続     |
| Max Mode        | Automations は常に Max Mode（オフ不可）               |


## セットアップ手順

### 1. Cursor Dashboard にログイン

[https://cursor.com/dashboard](https://cursor.com/dashboard) にアクセスし、Cursor アカウントでログインします。

### 2. GitHub を接続

1. Dashboard → **Integrations** を開く
2. **GitHub** を選択し **Connect** をクリック
3. cashpilot リポジトリへのアクセス権を付与（**read-write** が必要）
4. 接続後、リポジトリ一覧に cashpilot が表示されることを確認

### 3. Cloud Agents 環境を設定

1. Dashboard → **Cloud Agents** を開く
2. デフォルトのベース環境を確認（Node.js / Go が利用可能な環境）
3. 必要に応じて **Personal** 環境を保存（リポジトリの `.cursor/environment.json` が最優先で適用される）

### 4. Secrets を登録

Dashboard → **Cloud Agents** → **Secrets** で、以下を登録します（`.env` はリポジトリに含めない）。


| Secret 名       | 用途          | 例                                                                         |
| -------------- | ----------- | ------------------------------------------------------------------------- |
| `DATABASE_URL` | API 結合テスト用  | `postgres://cashpilot:cashpilot@localhost:5432/cashpilot?sslmode=disable` |
| `JWT_SECRET`   | 認証 API テスト用 | 32 文字以上のランダム文字列                                                           |


ローカル開発用の値は `[.env.example](../../.env.example)` を参照してください。

### 5. On-Demand Usage を有効化

1. Dashboard → **Spending** を開く
2. **On-Demand Usage** を有効化
3. 月額上限を設定（例: $10）して意図しない高額課金を防ぐ

### 6. 任意の統合を接続

cashpilot の開発で使う場合の推奨:


| 統合         | 用途                      |
| ---------- | ----------------------- |
| **GitHub** | PR トリガー、コード変更、CI 連携（必須） |
| **Slack**  | 週次サマリー、バグ報告からの修正 PR（任意） |
| **Linear** | Issue 作成時の自動トリアージ（任意）   |


## セットアップ完了の確認

以下がすべて満たされていれば、次のステップ（[はじめの一歩](./getting-started.md)）に進めます。

- [x] GitHub に cashpilot リポジトリが接続されている
- [x] On-Demand Usage が有効化されている
- [x] Cloud Agents の Secrets に `JWT_SECRET` が登録されている
- [x] リポジトリに `AGENTS.md` と `.cursor/environment.json` が存在する（本リポジトリに含まれています）

## トラブルシューティング


| 症状                | 対処                                               |
| ----------------- | ------------------------------------------------ |
| Automation が起動しない | GitHub 接続、リポジトリ権限、有料プラン、on-demand 有効化を確認         |
| Secrets が使えない     | Dashboard → Cloud Agents → Secrets で名前のスペルミスを確認  |
| 環境の起動が遅い          | `.cursor/environment.json` の `install` スクリプトを最適化 |


## 参考リンク

- [Automations ドキュメント](https://cursor.com/docs/cloud-agent/automations)
- [Cloud Agent セットアップ](https://cursor.com/docs/cloud-agent/setup)
- [GitHub 統合](https://cursor.com/docs/integrations/github)

