# Cursor Automations（cashpilot）

CashPilot プロジェクトで Cursor Automations を活用するためのドキュメント集です。

## ドキュメント一覧

| ドキュメント | 内容 |
|-------------|------|
| [setup-guide.md](./setup-guide.md) | Dashboard での初期設定（GitHub 接続、Secrets、On-Demand） |
| [getting-started.md](./getting-started.md) | PR レビュー Automation（コメントのみ）の作成 |
| [auto-fix-automation.md](./auto-fix-automation.md) | 自動修正 Automation（レビューと分離、`@cursor fix`） |
| [phase5-automations.md](./phase5-automations.md) | Phase 5（画面 API 連携）向け Automation のプロンプト設計 |

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

1. [セットアップガイド](./setup-guide.md) で Dashboard を整備
2. [はじめの一歩](./getting-started.md) で PR レビュー Automation（`cashpilot-pr-review`）を作成
3. [自動修正 Automation](./auto-fix-automation.md) で `cashpilot-auto-fix` を追加
4. Phase 5 開発開始時に [phase5-automations.md](./phase5-automations.md) の Automation を追加

## 公式ドキュメント

- [Automations](https://cursor.com/docs/cloud-agent/automations)
- [Cloud Agent セットアップ](https://cursor.com/docs/cloud-agent/setup)
- [Marketplace](https://cursor.com/marketplace)
