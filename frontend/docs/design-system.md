# デザインシステム

CashPilot フロントエンドのスタイル・変数・コンポーネントの設計指針。  
`src/app/styles/variables.scss` と連動する。

---

## 1. カラーパレット

100=薄い／900=濃い。

### 白 (white)

| 100 | 200 | 300 | 400 | 500 | 600 | 700 | 800 | 900 |
|-----|-----|-----|-----|-----|-----|-----|-----|-----|
| `#ffffff` | `#fafafa` | `#f5f5f5` | `#f0f0f0` | `#ebebeb` | `#e5e5e5` | `#d4d4d4` | `#a3a3a3` | `#737373` |

### 黒 (black)

| 100 | 200 | 300 | 400 | 500 | 600 | 700 | 800 | 900 |
|-----|-----|-----|-----|-----|-----|-----|-----|-----|
| `#f5f5f5` | `#e5e5e5` | `#d4d4d4` | `#a3a3a3` | `#737373` | `#525252` | `#404040` | `#262626` | `#000000` |

### グレイ (gray)

| 100 | 200 | 300 | 400 | 500 | 600 | 700 | 800 | 900 |
|-----|-----|-----|-----|-----|-----|-----|-----|-----|
| `#f9fafb` | `#f3f4f6` | `#e5e7eb` | `#d1d5db` | `#9ca3af` | `#6b7280` | `#4b5563` | `#374151` | `#111827` |

### 青 (blue)

| 100 | 200 | 300 | 400 | 500 | 600 | 700 | 800 | 900 |
|-----|-----|-----|-----|-----|-----|-----|-----|-----|
| `#dbeafe` | `#bfdbfe` | `#93c5fd` | `#60a5fa` | `#3b82f6` | `#2563eb` | `#1d4ed8` | `#1e40af` | `#1e3a8a` |

### 緑 (green)

| 100 | 200 | 300 | 400 | 500 | 600 | 700 | 800 | 900 |
|-----|-----|-----|-----|-----|-----|-----|-----|-----|
| `#dcfce7` | `#bbf7d0` | `#86efac` | `#4ade80` | `#22c55e` | `#16a34a` | `#15803d` | `#166534` | `#14532d` |

### 赤 (red)

| 100 | 200 | 300 | 400 | 500 | 600 | 700 | 800 | 900 |
|-----|-----|-----|-----|-----|-----|-----|-----|-----|
| `#fee2e2` | `#fecaca` | `#fca5a5` | `#f87171` | `#ef4444` | `#dc2626` | `#b91c1c` | `#991b1b` | `#7f1d1d` |

### セマンティック（上記パレット参照）

| 変数名 | 参照 |
|--------|------|
| `--color-primary` | blue-600 |
| `--color-primary-hover` | blue-700 |
| `--color-secondary` | gray-500 |
| `--color-secondary-hover` | gray-600 |
| `--color-income` | green-600 |
| `--color-expense` | red-600 |
| `--color-success` | green-600 |
| `--color-warning` | （オレンジ系は必要に応じて追加） |
| `--color-error` | red-700 |
| `--color-bg` | white-100（ライト）/ black-800（ダーク） |
| `--color-bg-subtle` | white-200 / black-700 |
| `--color-text` | black-900 / white-100 |
| `--color-text-muted` | gray-500 / gray-400 |
| `--color-border` | gray-300 / gray-700 |

---

## 2. タイポグラフィ

### フォントファミリ

| 変数名 | 用途 | 値 |
|--------|------|-----|
| `--font-sans` | 本文 | system-ui, sans-serif |
| `--font-mono` | 数値（金額等） | ui-monospace, monospace |

### フォントサイズ

| 変数名 | 用途 | 値 |
|--------|------|-----|
| `--font-size-xs` | キャプション | 12px |
| `--font-size-sm` | 補助 | 14px |
| `--font-size-base` | 本文 | 16px |
| `--font-size-lg` | 小見出し | 18px |
| `--font-size-xl` | 見出し | 20px |
| `--font-size-2xl` | 大見出し | 24px |
| `--font-size-3xl` | ヒーロー数値（残高等） | 30px |

### フォントウェイト

| 変数名 | 用途 | 値 |
|--------|------|-----|
| `--font-weight-normal` | 本文 | 400 |
| `--font-weight-medium` | 強調 | 500 |
| `--font-weight-semibold` | 小見出し | 600 |
| `--font-weight-bold` | 見出し・ヒーロー | 700 |

---

## 3. スペーシング

| 変数名 | 値 | 用途 |
|--------|-----|------|
| `--space-1` | 4px | 極小 |
| `--space-2` | 8px | 小 |
| `--space-3` | 12px | 中 |
| `--space-4` | 16px | 標準 |
| `--space-5` | 20px | 大 |
| `--space-6` | 24px | セクション間 |
| `--space-8` | 32px | ブロック間 |
| `--radius-sm` | 4px | ボタン等 |
| `--radius-md` | 8px | カード |
| `--radius-lg` | 12px | モーダル等 |

---

## 4. シャドウ・エレベーション

| 変数名 | 用途 |
|--------|------|
| `--shadow-sm` | カード・ボタン |
| `--shadow-md` | ドロップダウン・モーダル |
| `--shadow-lg` | オーバーレイ |

---

## 5. コンポーネント指針

### ボタン

- Primary: 背景 `--color-blue-600`、テキスト白、角丸 `--radius-sm`
- Secondary: 背景透明、ボーダー・テキスト `--color-gray-500`、ホバー時 `--color-gray-600`
- Danger: 背景 `--color-red-700`（削除等）

### カード

- 背景 `--color-white-200`（ライト）/ `--color-black-700`（ダーク）、パディング `--space-4`、角丸 `--radius-md`、シャドウ `--shadow-sm`

### 入力フィールド

- ボーダー `--color-gray-300`、フォーカス時 `--color-blue-600` でアウトライン

### 金額表示

- 収入: `--color-green-600`
- 支出: `--color-red-600`
- 残高: `--color-black-900` / `--color-white-100`

---

## 6. レスポンシブ

## ブレークポイント(variables.scssに記述 → mixin.scssでmediaquery対応)

| ブレークポイント | 幅 | 用途 |
|------------------|-----|------|
| `--bp-sm` | 640px | スマホ |
| `--bp-md` | 768px | タブレット |
| `--bp-lg` | 1024px | PC |
| `--bp-xl` | 1280px | 大画面 |

---

## 7. 参照

- `src/app/styles/variables.scss` - 実装
- `src/app/styles/globals.scss` - ベーススタイル
- [frontend.md](../frontend.md) - 全体方針
