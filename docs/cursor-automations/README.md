# Cursor Automations（cashpilot）

CashPilot プロジェクトで Cursor Automations を活用するためのドキュメント集です。

## ドキュメント一覧

| ドキュメント | 内容 |
|-------------|------|
| [autonomous-org.md](./autonomous-org.md) | **自立型 4 エージェント組織**（意思決定→吟味→実装→マージ） |
| [agents/](./agents/) | 各エージェントの Automation プロンプト |
| [setup-guide.md](./setup-guide.md) | Dashboard での初期設定（GitHub 接続、Secrets、On-Demand） |
| [triggers-guide.md](./triggers-guide.md) | **実際の UI トリガー一覧**（Issue トリガーがない問題と代替） |
| [getting-started.md](./getting-started.md) | PR レビュー Automation（コメントのみ） |
| [auto-fix-automation.md](./auto-fix-automation.md) | 自動修正 Automation（レビューと分離、`@cursor fix`） |
| [phase5-automations.md](./phase5-automations.md) | Phase 5 向け Automation（implementation に統合） |

## Automation 設定との関係（Agent Instructions とは）

Cursor Automations には **2 層** の指示があります。両方使うとエージェントの精度が上がります。

| 層 | 場所 | 役割 |
|----|------|------|
| **リポジトリの Agent Instructions** | この `AGENTS.md` | プロジェクト共通の手順・規約。Cloud Agent が自動で読む |
| **Automation のプロンプト** | [cursor.com/automations](https://cursor.com/automations) の各 Automation 設定画面 | タスク固有の指示（レビュー / 自動修正 / Phase 5 実装など） |

`cashpilot-auto-fix` を作成するとき:

1. **Automation プロンプト** — [auto-fix-automation.md](./auto-fix-automation.md) の全文を貼る（トリガー・修正範囲・PR 形式）
2. **Agent Instructions（リポジトリ）** — 追加作業不要。リポジトリの `AGENTS.md` が自動適用される

Automation 設定画面に「Agent Instructions」フィールドがある場合は空のままでよい。リポジトリの `AGENTS.md` が優先されます。

## リポジトリ設定

Cloud Agent / Automation が参照する設定ファイル:

| ファイル | 用途 |
|---------|------|
| [AGENTS.md](../../AGENTS.md) | テスト手順、コーディング規約、Phase 5 の注意事項 |
| [.cursor/environment.json](../../.cursor/environment.json) | VM 起動時の依存関係インストールと dev サーバー |

## 学習パス

### 段階 1: 基本（完了済みであればスキップ可）

1. [セットアップガイド](./setup-guide.md) で Dashboard を整備
2. [はじめの一歩](./getting-started.md) で PR レビュー Automation を作成
3. [自動修正 Automation](./auto-fix-automation.md) で `cashpilot-auto-fix` を追加

### 段階 2: 自立型組織（推奨）

4. [自立型エージェント組織](./autonomous-org.md) の設計を読む
5. GitHub ラベル 7 個を作成
6. [agents/](./agents/) の 4 Automation を順に設定
7. `cashpilot-pr-review` を無効化し `cashpilot-review-merge` に統合
8. GitHub Actions 自動マージ（`.github/workflows/agent-auto-merge.yml`）を有効化

## 公式ドキュメント

- [Automations](https://cursor.com/docs/cloud-agent/automations)
- [Cloud Agent セットアップ](https://cursor.com/docs/cloud-agent/setup)
- [Marketplace](https://cursor.com/marketplace)
