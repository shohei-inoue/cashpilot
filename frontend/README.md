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
    │   └── (dashboard)/   # ダッシュボード関連
    ├── styles/            # グローバル SCSS（globals, variables, mixin）
    └── layout.tsx
```

コンポーネント単位のスタイルは `.module.scss` で定義する（SCSS Modules）。

## 参照

- 画面仕様・UI 方針: [docs/frontend.md](../docs/frontend.md)
- 全体方針: [docs/README.md](../docs/README.md)

## デプロイ

Vercel へのデプロイを想定。詳細は [Next.js デプロイメント](https://nextjs.org/docs/app/building-your-application/deploying) を参照。
