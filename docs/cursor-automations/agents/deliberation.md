# 吟味エージェント: cashpilot-deliberation

Issue の要望（一言でも可）を読み、**実装案を提案**するエージェントです。ユーザーの追加入力後に実装・修正案・却下へ進みます。

## 設定

| 項目 | 値 |
|------|-----|
| 名前 | `cashpilot-deliberation` |
| トリガー | **不要**（GitHub Actions が `@cursor` コメントで起動） |
| Automation（Webhook） | オプション。**無効化推奨**（GHA 経由と競合する） |
| リポジトリモード | Single repository（Automation を使う場合） |
| ツール | なし |

**推奨経路**:

- 初回: `.github/workflows/agent-trigger-deliberation.yml`（Issue 作成 / `agent:proposed`）
- 追加入力: `.github/workflows/agent-trigger-issue-followup.yml`（Issue コメント）

## 対話フロー

```
1. ユーザーが Issue に一言で要望を書く（ラベル任意）
2. GHA → @cursor → 【実装案】をコメント、agent:plan-proposed
3. ユーザーが追加入力:
   - 「実装して」→ agent:approved → 実装フェーズ
   - 「修正案: ...」→ 【修正案】を更新、agent:plan-proposed 維持
   - 「却下」→ agent:rejected
```

### ユーザー向けコマンド例

| 入力 | 動作 |
|------|------|
| `実装して` / `進めて` / `LGTM` / `/implement` | 実装案に基づき実装開始 |
| `修正案: グラフは円グラフに` / `/revise` | 実装案を更新して再提案 |
| `却下` / `やめて` / `/reject` | 却下 |

## プロンプト（初回・実装案）

```markdown
## Goal

Issue の要望を読み、実装案を提案する。即時の承認・実装は行わない。

## Context

GitHub Actions が Issue に投稿した `@cursor` コメントから起動する。

## Process

1. Issue 本文と `docs/implementation-flow.md`, `docs/frontend.md`, `docs/backend.md` を読む
2. `architect` subagent で設計・スコープ・リスクを評価する
3. Issue に【実装案】をコメント（要望の解釈、変更箇所、ステップ、受け入れ条件、リスク）
4. ラベル: `agent:plan-proposed` を追加、`agent:proposed` を削除
5. ユーザーへの次アクション案内を末尾に含める

## Constraints

- コード変更・PR 作成・`agent:approved` 付与は行わない
- スコープ外の場合も【実装案】内で却下理由と代替案を提示
```

## プロンプト（追加入力・followup）

```markdown
## Goal

ユーザーの Issue コメントを解釈し、実装開始・修正案・却下のいずれかに進める。

## Process

1. Issue とコメント履歴、【実装案】/【修正案】を読む
2. 意図判定:
   - 実装依頼 → `agent:approved` 付与（実装ワークフローへ委譲）
   - 修正依頼 → 【修正案】を投稿、`agent:plan-proposed` 維持
   - 却下 → `agent:rejected` 付与
   - 不明 → 【エージェント】で質問
3. エージェントコメントは先頭に【エージェント】を付ける

## Constraints

- 修正依頼時は PR を作らない
- 実装は `agent:approved` 付与のみ（実装エージェントに委譲）
```

## 動作確認

- [ ] 一言 Issue 作成後、【実装案】コメントと `agent:plan-proposed` が付く
- [ ] 「実装して」コメント後、`agent:approved` になる
- [ ] 「修正案:」コメント後、【修正案】が更新される
- [ ] 「却下」コメント後、`agent:rejected` になる

## 次のエージェント

→ [実装エージェント](./implementation.md) が `agent:approved` を検知して実装開始
