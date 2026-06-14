# 吟味エージェント: cashpilot-deliberation

Issue の要望（一言でも可）を読み、**難易度（easy / normal / hard）** を判定し、段階に応じた提案を行います。

## 難易度とフロー

| 難易度 | 初回提案 | ユーザーの OK 後 | 本実装 |
|--------|---------|-----------------|--------|
| **easy** | 概要 + 変更ファイル概略 | 即本実装開始 | 1 回で PR |
| **normal** | 具体的な実装案（ステップ付き） | 即本実装開始 | 1 回で PR |
| **hard** | 具体的な実装案 + テスト実装の検証ポイント | テスト実装 → 確認 → 本実装 | 修正を加えながら PR |

## 設定

| 項目 | 値 |
|------|-----|
| 名前 | `cashpilot-deliberation` |
| トリガー | **不要**（GitHub Actions が `@cursor` コメントで起動） |
| Automation（Webhook） | オプション。**無効化推奨** |

**推奨経路**:

- 初回判定: Issue に `@cursor plan` コメント → `.github/workflows/agent-trigger-deliberation.yml`
- 追加入力: `.github/workflows/agent-trigger-issue-followup.yml`
- hard テスト実装: `.github/workflows/agent-trigger-test-implementation.yml`

## 対話フロー

```
1. Issue に要望を書く
2. `@cursor plan` とコメント
3. 【実装案】難易度: easy|normal|hard + 提案内容
   ラベル: agent:plan-proposed + agent:difficulty-*
3. ユーザー入力:
   easy/normal + OK  → agent:approved → 本実装
   hard + OK          → agent:test-implementing → 【テスト実装】
   hard + テスト確認後 OK → agent:approved → 本実装（修正しながら）
   修正案:            → 【修正案】更新
   却下              → agent:rejected
```

### ユーザー向けコマンド

| 入力 | easy / normal | hard（plan 段階） | hard（test 段階） |
|------|--------------|------------------|------------------|
| `OK` / `/ok` | 本実装開始 | テスト実装開始 | 本実装開始 |
| `修正案: ...` / `/revise` | 実装案を更新 | 実装案を更新 | — |
| `修正: ...` / `/revise-test` | — | — | テスト実装を修正 |
| `却下` / `/reject` | 却下 | 却下 | 却下 |

## 難易度の判定基準

### easy
- 1〜2 ファイルの局所変更
- 既存パターンの踏襲（スタイル修正、文言、単純な表示追加）
- 新 API・新アーキテクチャ不要

### normal
- 複数ファイルの変更
- 軽い設計判断（コンポーネント分割、既存 API の組み合わせ）
- Phase 5 スコープ内、バックエンド API 既存

### hard
- 広範囲の変更または不確実なアプローチ
- 新パターンの導入、複数画面・API の統合
- テスト実装で検証すべき技術的リスクがある

## プロンプト（初回・難易度判定）

```markdown
## Goal

難易度を判定し、easy / normal / hard に応じた【実装案】を投稿する。

## Process

1. architect subagent で評価
2. 【実装案】難易度: X を投稿（段階に応じた詳細度）
3. agent:plan-proposed + agent:difficulty-X を付与

## Constraints

- コード変更・PR・agent:approved は行わない
```

## プロンプト（追加入力）

```markdown
## Goal

OK / 修正案 / 却下 を難易度と状態に応じて処理する。

## Process

- easy/normal + OK → agent:approved
- hard + OK（plan 段階）→ agent:test-implementing
- hard + OK（test 段階）→ agent:approved
- 修正案 → 【修正案】、ラベル維持
- テスト修正 → agent:test-implementing 再付与
```

## 動作確認

- [ ] easy Issue → 簡潔な【実装案】+ agent:difficulty-easy
- [ ] normal Issue → 詳細な【実装案】+ agent:difficulty-normal
- [ ] hard Issue → 詳細案 + OK で【テスト実装】→ 再 OK で PR
- [ ] 「修正案:」で【修正案】が更新される

## 次のエージェント

→ [実装エージェント](./implementation.md) が `agent:approved` を検知して本実装開始
