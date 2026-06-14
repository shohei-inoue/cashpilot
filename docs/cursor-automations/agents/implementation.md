# 実装エージェント: cashpilot-implementation

承認済み Issue を実装し、PR を作成する Automation です。

## 設定

| 項目 | 値 |
|------|-----|
| 名前 | `cashpilot-implementation` |
| トリガー | GitHub Issue label changed → `agent:approved` が追加されたとき |
| リポジトリモード | Single repository |
| ベースブランチ | `develop` |
| ツール | Pull request creation, Comment on pull request |

## プロンプト

```markdown
## Goal

`agent:approved` ラベル付き Issue を実装し、PR を作成する。

## Process

1. Issue 本文の受け入れ条件を読む
2. `implementer` subagent で実装する
3. 検証:
   - `cd frontend && pnpm lint && pnpm test`
   - `cd backend && go test ./tests/unit/...`
4. `verifier` subagent で受け入れ条件とテスト結果を確認
5. verifier が「マージ可」でなければ修正して再検証
6. PR を作成:
   - ブランチ: `cursor/<task-id>-5869`
   - タイトル: `feat(frontend): <Issue タイトル>`
   - 本文: 変更概要、受け入れ条件の充足状況、テスト結果、Closes #<Issue番号>
   - PR ラベル: `agent:needs-review`
7. Issue に PR リンクをコメントし、`agent:implementing` を `agent:needs-review` に更新（Issue 側はクローズしない）

## Implementation rules

- 既存パターン: Server Actions, API クライアント, SCSS Modules
- 1 PR = 1 Issue
- バックエンド API の変更は行わない（必要なら Issue にコメントして停止）

## Constraints

- Phase 5 スコープ外の機能は実装しない
- テストが通らない PR は作成しない

## Reference

- AGENTS.md
- docs/frontend.md
- docs/implementation-flow.md
```

## 動作確認

- [ ] `agent:approved` Issue から PR が自動作成される
- [ ] PR に `agent:needs-review` ラベルが付く
- [ ] lint / test が通った状態で PR が作られる

## 次のエージェント

→ [レビュー・マージエージェント](./review-merge.md) が PR をレビュー
