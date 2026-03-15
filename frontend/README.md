# CashPilot フロントエンド

CashPilot のフロントエンドアプリ。Next.js (App Router) + React + TypeScript + SCSS Modules で構築。

## 技術スタック

- **Next.js** 16（App Router）
- **React** 19
- **TypeScript**
- **SCSS Modules**（`src/app/styles/` に globals.scss, variables.scss, mixin.scss）

## 開発

パッケージマネージャーは **pnpm** を利用。

```bash
pnpm install
pnpm dev
```

開発サーバー起動後、[http://localhost:3010](http://localhost:3010) でアプリを確認できる。

## スクリプト

| コマンド | 説明 |
|----------|------|
| `pnpm dev` | 開発サーバー起動（ホットリロード） |
| `pnpm build` | 本番ビルド |
| `pnpm start` | 本番サーバー起動 |
| `pnpm lint` | ESLint 実行 |

## ディレクトリ構成

```
src/
└── app/
    ├── (pages)/
    │   ├── (dashboard)/   # ダッシュボード（プレースホルダー）
    │   ├── auth/          # サインアップ・ログイン（API 連携済み）
    │   ├── transactions/  # 取引一覧・登録（プレースホルダー）
    │   └── simulation/    # シミュレーション（プレースホルダー）
    ├── components/        # Header, Sidebar, MainContainer, Form など
    ├── constants/
    ├── styles/            # グローバル SCSS（globals, variables, mixin）
    └── layout.tsx
```

コンポーネント単位のスタイルは `.module.scss` で定義する（SCSS Modules）。

## 実装状況

| 画面・機能 | 状態 | 備考 |
|------------|------|------|
| 認証（サインアップ・ログイン） | 実装済み | `/api/auth/*` と連携 |
| ダッシュボード | 骨組みのみ | 集計・直近取引の表示は未実装 |
| 取引一覧・登録 | 骨組みのみ | API 連携・フォーム未実装 |
| シミュレーション | 骨組みのみ | 入力・結果表示・API 連携未実装 |
| 目標・設定（口座・カテゴリ） | 未実装 | 画面・ナビ未追加 |

次にやること: ダッシュボード・取引・シミュレーションの API 連携と実データ表示、目標・設定画面の追加。全体の優先順位は [docs/implementation-flow.md](../docs/implementation-flow.md) を参照。

## 参照

- 画面仕様・UI 方針: [docs/frontend.md](../docs/frontend.md)
- 全体方針: [docs/README.md](../docs/README.md)

## デプロイ

Vercel へのデプロイを想定。詳細は [Next.js デプロイメント](https://nextjs.org/docs/app/building-your-application/deploying) を参照。
