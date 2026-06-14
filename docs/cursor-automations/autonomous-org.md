# CashPilot 自立型エージェント組織

CashPilot を「組織」として運用し、意思決定からマージまでをエージェントが担う設計です。
Cursor Automations の制約内で **可能な限り自立** させ、人間は最終オーバーライドのみに限定します。

## 組織図

```mermaid
flowchart TB
    subgraph decision [意思決定]
        PD["cashpilot-product-decision"]
        PD --> Issue["GitHub_Issue_作成"]
    end

    subgraph deliberation [吟味]
        DL["cashpilot-deliberation"]
        Issue --> DL
        DL --> LabelApproved["label_agent-approved"]
        DL --> LabelRejected["label_agent-rejected"]
    end

    subgraph implementation [実装]
        IM["cashpilot-implementation"]
        LabelApproved --> IM
        IM --> PR["実装PR作成"]
    end

    subgraph reviewMerge [レビューとマージ]
        RM["cashpilot-review-merge"]
        AF["cashpilot-auto-fix"]
        PR --> RM
        RM --> AF
        AF --> RM
        RM --> MergeReady["label_agent-merge-ready"]
        MergeReady --> GHA["GitHub_Actions_自動マージ"]
        GHA --> Develop["develop_へマージ"]
    end
```

## 4 つのエージェント

| エージェント | Automation 名 | 役割 | コード変更 |
|-------------|---------------|------|-----------|
| **意思決定** | `cashpilot-product-decision` | バックログ分析、次タスク提案、Issue 作成 | なし |
| **吟味** | `cashpilot-deliberation` | 提案の妥当性評価、承認/却下ラベル付与 | なし |
| **実装** | `cashpilot-implementation` | 承認済み Issue を実装し PR 作成 | あり |
| **レビュー・マージ** | `cashpilot-review-merge` | PR レビュー、承認、マージ準備ラベル | なし（修正は auto-fix へ委譲） |

既存の `cashpilot-pr-review` と `cashpilot-auto-fix` は **レビュー・マージ** フェーズに統合または併用します。

## 状態管理: GitHub ラベル

Automations 同士は直接呼び出せないため、**GitHub ラベル** で状態を受け渡します。

| ラベル | 意味 | 付与者 |
|--------|------|--------|
| `agent:proposed` | 意思決定エージェントが提案した Issue | product-decision |
| `agent:deliberating` | 吟味中 | deliberation |
| `agent:approved` | 実装承認済み | deliberation |
| `agent:rejected` | 却下 | deliberation |
| `agent:implementing` | 実装中 | implementation |
| `agent:needs-review` | PR 作成済み、レビュー待ち | implementation |
| `agent:merge-ready` | マージ可能 | review-merge |

### ラベルの作成

GitHub リポジトリ → **Issues** → **Labels** で上記ラベルを作成してください。

## 1 サイクルの流れ

```
1. [cron 週1] product-decision
   → docs/implementation-flow.md を読み、次タスクを Issue に提案
   → ラベル: agent:proposed

2. [Issue opened] deliberation
   → 提案を吟味（readonly subagent: architect）
   → 承認: agent:approved / 却下: agent:rejected

3. [Label: agent:approved] implementation
   → subagent: planner → implementer → verifier
   → PR 作成、ラベル: agent:needs-review

4. [PR opened + CI completed] review-merge
   → レビュー、問題あれば @cursor fix（auto-fix）
   → 基準を満たせば承認 + agent:merge-ready

5. [Label: agent:merge-ready + CI green] GitHub Actions
   → develop へ自動マージ
```

## 完全自立の限界（正直な説明）

Cursor Automations には次の制約があります。

| できること | できないこと |
|-----------|-------------|
| Issue 作成、ラベル付与、PR 作成、PR 承認コメント | Automation から Automation を直接起動 |
| `@cursor fix` による修正 PR | LLM が任意のタイミングで merge ボタンを押す |
| GitHub Actions による条件付きマージ | プロダクト判断の 100% 自動化（人間の価値観は docs に依存） |

**マージ** は Cursor の Automation ツールにはないため、`.github/workflows/agent-auto-merge.yml` で **ラベル + CI 成功** を条件に機械的に実行します。

## セットアップ順序

| 順番 | 作業 | ドキュメント |
|------|------|-------------|
| 1 | GitHub ラベル 7 個を作成 | 上記表 |
| 2 | Subagent 4 個をリポジトリに追加 | [.cursor/agents/](../../.cursor/agents/) |
| 3 | 意思決定 Automation 作成 | [product-decision.md](./agents/product-decision.md) |
| 4 | 吟味 Automation 作成 | [deliberation.md](./agents/deliberation.md) |
| 5 | 実装 Automation 作成 | [implementation.md](./agents/implementation.md) |
| 6 | レビュー・マージ Automation 作成 | [review-merge.md](./agents/review-merge.md) |
| 7 | 既存 auto-fix を接続 | [auto-fix-automation.md](./auto-fix-automation.md) |
| 8 | GitHub Actions マージワークフロー有効化 | [.github/workflows/agent-auto-merge.yml](../../.github/workflows/agent-auto-merge.yml) |
| 9 | Approval Policy 確認 | [.cursor/approval-policies/ROUTING.md](../../.cursor/approval-policies/ROUTING.md) |

## Subagent（1 実行内の役割分担）

Automation 1 回の実行内では `.cursor/agents/` の Subagent に委譲します。

| Subagent | 役割 | readonly |
|----------|------|----------|
| `product-planner` | バックログ分析、タスク提案 | はい |
| `architect` | 設計の吟味、トレードオフ評価 | はい |
| `implementer` | コード実装 | いいえ |
| `verifier` | テスト・lint 検証 | はい |

## 人間の関与ポイント（最小化）

| タイミング | 人間がやること |
|-----------|---------------|
| 初回 | ラベル作成、Automation 5 個の設定、On-Demand 有効化 |
| 運用中 | `agent:rejected` の Issue を確認（任意） |
| 緊急時 | Automation を無効化、手動マージ |

## 既存構成との関係

| 既存 | 自立型組織での位置づけ |
|------|---------------------|
| `cashpilot-pr-review` | `cashpilot-review-merge` に統合推奨（重複レビューを避ける） |
| `cashpilot-auto-fix` | review-merge フェーズで `@cursor fix` として継続利用 |
| `phase5-automations.md` | implementation エージェントのプロンプトに吸収 |

## 参考

- [Automations ドキュメント](https://cursor.com/docs/cloud-agent/automations)
- [Subagents](https://cursor.com/docs/subagents)
- [Approval Agents](https://cursor.com/docs/approval-agents)
- [Merge Queue](https://cursor.com/docs/cursor-review/merge-queue)
