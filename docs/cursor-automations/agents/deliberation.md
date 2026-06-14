# 吟味エージェント: cashpilot-deliberation

意思決定エージェントの提案を評価し、承認または却下する Automation です。

## 設定

| 項目 | 値 |
|------|-----|
| 名前 | `cashpilot-deliberation` |
| トリガー | **Webhook**（GitHub Actions が Issue イベント時に POST） |
| リポジトリモード | Single repository |
| ツール | なし（Issue コメント・ラベルは gh CLI） |

**注意**: UI に **Issue opened** トリガーはありません。[triggers-guide.md](../triggers-guide.md) 参照。

## プロンプト

```markdown
## Goal

Webhook ペイロードの Issue（`agent:proposed` ラベル付き）を吟味し、実装の承認または却下を判定する。

## Webhook context

ペイロードに `issue_number`, `issue_url`, `issue_title` が含まれる。該当 Issue を `gh issue view` で読む。

## Process

1. Issue 本文と `docs/implementation-flow.md`, `docs/frontend.md`, `docs/backend.md` を読む
2. `architect` subagent を使い、設計・スコープ・リスクを評価する
3. Issue に吟味結果をコメントする（判定・理由・懸念点・条件）
4. ラベルを更新:
   - 承認: `agent:approved` を追加、`agent:proposed` を削除
   - 却下: `agent:rejected` を追加、`agent:proposed` を削除
5. 条件付き承認の場合は条件を Issue コメントに明記し、`agent:approved` を付与

## 判定基準

### 承認
- Phase 5 スコープ内
- バックエンド API が既に存在
- MVP の目的（家計管理・分析・シミュレーション）に貢献

### 却下
- スコープ外（Phase 6 以降、新 API 開発が必要）
- 既に同等の Issue / PR が存在
- リスクが高く受け入れ条件が曖昧

## Constraints

- コード変更・PR 作成は行わない
- 判定は「承認」「却下」を必ず明示

## Reference

- AGENTS.md
- docs/cursor-automations/autonomous-org.md
```

## 動作確認

- [ ] `agent:proposed` Issue 作成後、自動でコメントが付く
- [ ] 承認時に `agent:approved` ラベルが付く
- [ ] 却下時に `agent:rejected` ラベルが付く

## 次のエージェント

→ [実装エージェント](./implementation.md) が `agent:approved` を検知して実装開始
