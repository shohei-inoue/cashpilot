# レビュー・マージエージェント: cashpilot-review-merge

PR をレビューし、基準を満たせば承認してマージ準備を整える Automation です。

## 設定

| 項目 | 値 |
|------|-----|
| 名前 | `cashpilot-review-merge` |
| トリガー | **Pull request opened**, **Pull request pushed**（+ **CI completed** が UI にあれば追加） |
| フィルタ | PR ラベル `agent:needs-review`（あれば） |
| リポジトリモード | Single repository |
| ツール | Comment on pull request（承認を有効化） |

**注意**: **CI completed** はプランや連携状態によって UI に表示されない場合があります。なければ PR pushed のみで運用。[triggers-guide.md](../triggers-guide.md) 参照。

**既存 Automation との関係**

- `cashpilot-pr-review` は **無効化** し、この Automation に統合する
- `cashpilot-auto-fix` は **継続利用**（修正が必要なとき PR に `@cursor fix`）

## プロンプト

```markdown
## Goal

`agent:needs-review` ラベル付き PR をレビューし、マージ可能なら承認して `agent:merge-ready` ラベルを付与する。

## Process

1. PR の diff と関連 Issue の受け入れ条件を読む
2. `verifier` subagent で検証:
   - `cd frontend && pnpm lint && pnpm test`
   - `cd backend && go test ./tests/unit/...`
3. レビュー結果を PR コメントに投稿
4. 判定:

### マージ可の条件（すべて満たす）
- lint / test が通過
- 受け入れ条件を充足
- セキュリティ上の問題なし
- diff が Phase 5 スコープ内

### マージ可の場合
- PR を承認（Approve）
- ラベル `agent:merge-ready` を追加
- PR コメント: 「マージ準備完了。GitHub Actions が自動マージします。」

### 修正必要の場合
- 変更要求コメントを投稿
- PR コメント: `@cursor fix` で修正を依頼（cashpilot-auto-fix が起動）
- `agent:needs-review` を維持

### 却下の場合
- 理由を詳述し、Issue を再オープンするよう提案
- `agent:rejected` を PR に付与

## Constraints

- この Automation ではコード変更・PR 作成は行わない（修正は cashpilot-auto-fix へ）
- **merge ボタンは押さない**（GitHub Actions が `agent:merge-ready` + CI 成功で実行）

## Reference

- AGENTS.md
- .cursor/approval-policies/ROUTING.md
```

## 動作確認

- [ ] `agent:needs-review` PR にレビューコメントが付く
- [ ] 基準充足時に Approve と `agent:merge-ready` ラベルが付く
- [ ] 修正必要時に `@cursor fix` コメントが付く

## マージの実行

→ [.github/workflows/agent-auto-merge.yml](../../../.github/workflows/agent-auto-merge.yml) が `agent:merge-ready` + CI 成功で `develop` へマージ

## 次のサイクル

マージ後、次週の [意思決定エージェント](./product-decision.md) が新タスクを提案
