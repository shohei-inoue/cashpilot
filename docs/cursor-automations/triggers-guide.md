# Cursor Automations トリガー一覧（実際の UI）

Cursor Automations の GitHub トリガーは **PR 中心** です。Issue 関連のトリガーは UI にありません（2026年6月時点）。

## UI にあるトリガー（GitHub / GitLab）

| UI の名前 | 説明 |
|-----------|------|
| **Scheduled** | cron / 定期実行 |
| **Draft opened** | ドラフト PR 作成時 |
| **Pull request opened** | PR オープン時 |
| **Pull request pushed** | PR に push 時 |
| **Pull request label changed** | **PR** のラベル変更時（Issue ではない） |
| **Pull request merged** | PR マージ時 |
| **Pull request commented** | PR にコメント時 |
| **Push to branch** | 特定ブランチへの push |
| **CI completed** | CI 完了時（表示されない場合は PR pushed で代替） |
| **Webhook** | 外部から HTTP POST |

## UI にないトリガー（よく探されるが存在しない）

| 探している名前 | 代替手段 |
|---------------|----------|
| Issue opened | **Webhook** + GitHub Actions |
| Issue label changed | **Webhook** + GitHub Actions |
| Issue commented | **Webhook** + GitHub Actions |

## cashpilot 4 エージェントの正しい Trigger 設定

| Automation | UI で選ぶトリガー | 補足 |
|------------|------------------|------|
| `cashpilot-product-decision` | **Scheduled** | そのまま使える |
| `cashpilot-deliberation` | **Webhook** | GHA が Issue opened 時に POST |
| `cashpilot-implementation` | **Webhook** | GHA が Issue labeled 時に POST |
| `cashpilot-review-merge` | **Pull request opened** + **Pull request pushed** | CI completed があれば追加 |
| `cashpilot-auto-fix` | **Pull request commented** | `@cursor fix` フィルタ |

## Webhook の設定手順

### 1. Cursor 側

1. Automation 作成時にトリガー **Webhook** を選択
2. 保存後に **Webhook URL** と **API キー** が表示される
3. GitHub リポジトリの Secrets に登録:
   - `CURSOR_WEBHOOK_DELIBERATION_URL`
   - `CURSOR_WEBHOOK_IMPLEMENTATION_URL`
   - `CURSOR_AUTOMATION_TOKEN`

### 2. GitHub 側

次のワークフローが Issue イベントを Cursor Webhook に転送します。

- `.github/workflows/agent-trigger-deliberation.yml` — Issue 作成時
- `.github/workflows/agent-trigger-implementation.yml` — `agent:approved` ラベル付与時

## CI completed が表示されない場合

プランや GitHub 連携の状態によっては **CI completed** が UI に出ないことがあります。代替:

| 代替 | 設定 |
|------|------|
| **Pull request pushed** | PR 更新のたびに review-merge が起動 |
| **Scheduled**（5分おき等） | ポーリングで未レビュー PR を処理（コスト高） |

PR push で十分なことが多いです。

## 参考

- [Automations ドキュメント](https://cursor.com/docs/cloud-agent/automations)
- [autonomous-org.md](./autonomous-org.md)
