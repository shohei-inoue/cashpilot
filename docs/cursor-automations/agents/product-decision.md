# 意思決定エージェント: cashpilot-product-decision

プロダクトの次タスクを分析し、GitHub Issue として提案する Automation です。

## 設定

| 項目 | 値 |
|------|-----|
| 名前 | `cashpilot-product-decision` |
| トリガー | cron: 毎週月曜 9:00（JST に合わせて調整） |
| リポジトリモード | Single repository |
| ツール | なし（Issue 作成は gh CLI または GitHub MCP） |

**注意**: Comment on PR ツールは不要。Issue 作成にはリポジトリへの書き込み権限が必要です。Cloud Agent は `gh issue create` を実行できます。

## プロンプト

```markdown
## Goal

CashPilot のプロダクトバックログを分析し、次に実装すべきタスクを 1 件 GitHub Issue として提案する。

## Process

1. `product-planner` subagent を使い、バックログ分析を行う
2. 次タスクを 1 件に絞る
3. GitHub Issue を作成:
   - タイトル: `[Agent] Phase 5-X: <画面名> の実装`
   - 本文: product-planner の出力（提案・理由・受け入れ条件）
   - ラベル: `agent:proposed`
4. 既に `agent:proposed` または `agent:approved` の未完了 Issue がある場合は新規作成をスキップし、既存 Issue へのコメントで状況を報告

## Constraints

- コード変更・PR 作成は行わない
- Phase 5 のスコープ内のみ提案
- 1 回の実行で Issue は最大 1 件

## Reference

- docs/implementation-flow.md
- docs/frontend.md
- AGENTS.md
```

## 動作確認

- [ ] 週次実行後、GitHub Issues に `agent:proposed` ラベル付き Issue ができる
- [ ] Issue 本文に受け入れ条件が含まれる
- [ ] 未完了 Issue があるときは重複作成しない

## 次のエージェント

→ [吟味エージェント](./deliberation.md) が `agent:proposed` Issue を処理
