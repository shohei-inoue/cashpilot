# はじめの一歩: Marketplace テンプレートで PR レビュー Automation

セットアップ完了後、最初に試す Automation の作成手順です。**コメントのみ** モードで運用します。自動修正は別 Automation（[auto-fix-automation.md](./auto-fix-automation.md)）で行います。

## 目標

cashpilot の PR がオープンされたとき、Cloud Agent が自動でコードレビューを行い、GitHub にコメントを残す Automation を 1 つ作成します。

## 作成手順

### 1. Automations 画面を開く

次のいずれかから開きます。

- Web: [https://cursor.com/automations](https://cursor.com/automations)
- IDE: **Agents Window** → **Automations** タブ

### 2. Marketplace テンプレートを選択

1. **Create Automation** または **Browse Marketplace** をクリック
2. Automations セクションから **Find vulnerabilities** または **Add test coverage** を選択
3. **Use template** でテンプレートを読み込む

### 3. トリガーを設定

| 設定項目 | 値 |
|----------|-----|
| トリガー | **Pull request opened** |
| リポジトリ | cashpilot（対象の GitHub リポジトリ） |
| ブランチ | `develop` または `main`（マージ先ブランチ） |

初回は **Pull request opened** のみにし、push や merged は後から追加します。

### 4. ツールを設定

初回は **安全モード** で以下のみ有効化します。

| ツール | 設定 |
|--------|------|
| **Comment on pull request** | 有効 |
| **Pull request creation** | 無効 |
| **Request reviewers** | 無効 |

自動修正はこの Automation では行いません。[auto-fix-automation.md](./auto-fix-automation.md) で別途設定します。

### 5. リポジトリモードを選択

**Single repository** を選択し、cashpilot リポジトリを指定します。

### 6. プロンプトをカスタマイズ

テンプレートのプロンプトを、cashpilot 向けに次の内容に置き換えます。

```markdown
## Goal

cashpilot リポジトリの Pull Request をレビューし、品質・セキュリティ・テストの観点でフィードバックを PR コメントとして残す。

## Review scope

- フロントエンド（Next.js / React / TypeScript）: `frontend/`
- バックエンド（Go / Gin）: `backend/`
- 変更にテストが含まれているか（`frontend/` は Vitest、`backend/` は Go test）

## Review checklist

1. **セキュリティ**: JWT / Cookie 認証の漏洩、SQL インジェクション、他ユーザーデータへのアクセス
2. **API 設計**: `docs/backend.md` に沿ったエンドポイント・レスポンス形式
3. **フロントエンド**: Server Actions / API クライアントのエラーハンドリング
4. **テスト**: 変更ファイルに対応するテストの有無

## Verification

レビュー前に以下を実行し、結果をコメントに含める:

- `cd frontend && pnpm lint && pnpm test`
- `cd backend && go test ./tests/unit/...`

## Output policy

- **コメントのみ**: この Automation では PR を自動作成しない
- Critical な問題が 3 件以上ある場合: **Changes requested** 相当のコメント
- 問題が少ない場合: 改善提案をインラインコメントで残す
- 問題がない場合: 簡潔な承認コメント

## Project context

- MVP の未実装領域は `docs/implementation-flow.md` Phase 5 を参照
- エージェント向け手順は `AGENTS.md` を参照
```

### 7. 作成して保存

1. Automation に名前を付ける（例: `cashpilot-pr-review`）
2. **Create** をクリック
3. 権限スコープは **Private**（個人利用）または **Team Visible**（チーム共有）を選択

## 動作確認

### テスト用 PR を作成

1. 小さな変更（コメント追加など）を含むブランチを作成
2. `develop`（または `main`）向けに PR をオープン
3. 数分以内に Cloud Agent が起動し、PR にコメントが付くことを確認

### 確認ポイント

- [ ] Automation の実行ログが [cursor.com/automations](https://cursor.com/automations) に表示される
- [ ] PR にレビューコメントが投稿される
- [ ] lint / test の実行結果がコメントに含まれる
- [ ] 意図しない PR が自動作成されていない（初回設定どおり）

## 次のステップ

動作確認後:

1. **Pull request pushed** トリガーを追加（更新時の再レビュー）
2. [自動修正 Automation（分離構成）](./auto-fix-automation.md) を作成
3. [Phase 5 向け Automation 設計](./phase5-automations.md) を適用

## 参考リンク

- [Marketplace](https://cursor.com/marketplace)
- [Automations ドキュメント](https://cursor.com/docs/cloud-agent/automations)
- [Cloud Agent ベストプラクティス](https://cursor.com/docs/cloud-agent/best-practices)
