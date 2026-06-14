# CashPilot 承認・マージルーティング

レビュー・マージエージェント（`cashpilot-review-merge`）と GitHub Actions が参照するポリシーです。

## 自動承認（agent:merge-ready）の条件

以下を **すべて** 満たす PR のみ `agent:merge-ready` を付与してよい。

| # | 条件 |
|---|------|
| 1 | PR ラベルに `agent:needs-review` がある |
| 2 | 関連 Issue の受け入れ条件を充足 |
| 3 | `pnpm lint && pnpm test` が通過 |
| 4 | `go test ./tests/unit/...` が通過 |
| 5 | セキュリティ: 認証スコープ違反、SQL インジェクション、秘密情報の混入がない |
| 6 | diff が Phase 5 スコープ内（フロントエンド API 連携中心） |
| 7 | 1 PR = 1 Issue（スコープクリープなし） |

## 人間レビューが必要なケース

以下に該当する PR は `agent:merge-ready` を付与しない。

- バックエンド API の変更を含む
- DB マイグレーションを含む
- 認証・セキュリティ設計の変更
- 500 行を超える大規模 diff
- `agent:rejected` ラベルが付いている

## マージ実行

`agent:merge-ready` + CI 成功 → `.github/workflows/agent-auto-merge.yml` が squash merge

## エスカレーション

自動マージが 3 回連続で失敗した場合、Automation を一時停止し人間が確認する。
