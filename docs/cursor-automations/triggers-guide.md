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
| Issue opened | 手動で `@cursor plan` コメント |
| Issue label changed | 手動で `@cursor plan` コメント |
| Issue commented | `@cursor plan` / `@cursor ok` / `@cursor fix plan` / `@cursor fix` / `@cursor reject` |

## cashpilot 4 エージェントの正しい Trigger 設定

| Automation | UI で選ぶトリガー | 補足 |
|------------|------------------|------|
| `cashpilot-product-decision` | **Scheduled** | そのまま使える |
| `cashpilot-deliberation` | **不要（無効化推奨）** | GHA が `@cursor` コメントで起動 |
| `cashpilot-implementation` | **不要（無効化推奨）** | GHA が `@cursor` コメントで起動 |
| `cashpilot-review-merge` | **不要（GHA 経由）** | PR Open 時（非 Draft）に GHA が `@cursor` コメント |
| `cashpilot-auto-fix` | **Pull request commented** | `@cursor fix` フィルタ |

**重要**: 吟味・実装は **Webhook Automation と GHA の二重経路にしない**。Webhook 版は Issue への `gh` 実行が不安定だったため、GHA 経由の `@cursor` コメントに一本化する。

## 吟味・実装の起動経路（GHA → @cursor）

GitHub Actions が Issue イベントを検知し、Issue に `@cursor` コメントを投稿して Cloud Agent を起動します。

| ワークフロー | トリガー | 動作 |
|-------------|---------|------|
| `.github/workflows/agent-trigger-deliberation.yml` | Issue コメント `@cursor plan` | 難易度判定 + 【実装案】 |
| `.github/workflows/agent-trigger-issue-followup.yml` | `@cursor ok` / `@cursor fix plan` / `@cursor fix` / `@cursor reject` | 難易度別分岐 |
| `.github/workflows/agent-trigger-test-implementation.yml` | `agent:test-implementing` 付与 | hard のテスト実装 |
| `.github/workflows/agent-trigger-implementation.yml` | `agent:approved` 付与 | 本実装 + PR |

| `.github/workflows/agent-trigger-review-merge.yml` | PR `opened` / `ready_for_review`（非 Draft、`agent:needs-review`） | レビュー依頼の `@cursor` コメント |

**重要**: レビューは **Draft PR では走らない**。Open 時（Draft → Ready 含む）のみ起動。Push では起動しない。

### Cursor 側の設定

吟味・実装用の Webhook Automation（`cashpilot-deliberation`, `cashpilot-implementation`）は **無効化または削除** してください。残すと GHA と競合し、二重起動や未実行の原因になります。

有効にする Automation は次の 3 つ（+ 任意の auto-fix）:

- `cashpilot-product-decision`（Scheduled）
- `cashpilot-review-merge`（**PR トリガーは無効化**、GHA 経由）
- `cashpilot-auto-fix`（PR commented、`@cursor fix`）

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
