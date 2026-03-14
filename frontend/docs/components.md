# 共通コンポーネント TODO

各画面で再利用する想定のコンポーネント一覧。実装時はチェックを付けていく。

---

## レイアウト

- [ ] **Header** - ロゴ、ナビリンク（ダッシュボード/取引/シミュレーション/目標）、ログアウト
- [ ] **AuthLayout** - 認証必須ページ用ラッパー（未認証時は /login へリダイレクト）
- [ ] **PageLayout** - 共通ページレイアウト（Header + main）

---

## ボタン・リンク

- [ ] **Button** - variant: primary / secondary / danger、disabled、loading 対応
- [ ] **IconButton** - アイコン付きボタン（編集・削除等）
- [ ] **Link** - ナビ用リンク

---

## カード・コンテナ

- [ ] **Card** - 汎用カード（見出しオプション）
- [ ] **EmptyState** - データなし時（アイコン + メッセージ + アクションリンク）
- [ ] **ErrorMessage** - エラー表示 + リトライボタン

---

## フォーム

- [ ] **Input** - text / number / date / datetime-local 対応
- [ ] **Select** - 単一選択（口座、カテゴリ等）
- [ ] **Textarea** - 複数行（メモ）
- [ ] **FormField** - label + 入力要素 + エラーメッセージのラッパー
- [ ] **DateRangePicker** - 期間 from / to（オプション）

---

## フィードバック

- [ ] **Spinner** - ローディング表示
- [ ] **Skeleton** - スケルトンローディング（カード・一覧行）
- [ ] **Toast** - 成功・エラー通知（「登録しました」等）

---

## モーダル・ダイアログ

- [ ] **Modal** - 汎用オーバーレイ（目標追加・編集フォーム等）
- [ ] **ConfirmDialog** - 削除確認（「○○ を削除しますか？」）

---

## データ表示

- [ ] **Amount** - 金額表示（3桁区切り、収入=緑/支出=赤）
- [ ] **BalanceDisplay** - 残高表示（大きめフォント）
- [ ] **TransactionRow** - 取引1行（日時、口座、カテゴリ、金額、アクション）
- [ ] **GoalCard** - 目標カード（名前、金額、期限、編集・削除）
- [ ] **GoalProjectionItem** - 目標達成見込み1件（達成/未達バッジ）

---

## グラフ

- [ ] **CashflowLineChart** - 残高推移折れ線グラフ（Recharts 等）

---

## その他

- [ ] **PageTitle** - ページ見出し + アクションボタン（例: 「目標」+ 「目標を追加」）

---

## 参照

- [design-system.md](./design-system.md)
- [pages/](./pages/)
