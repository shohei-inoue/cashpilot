export const ACCOUNT_TYPE_OPTIONS = [
  { value: 'cash', label: '現金' },
  { value: 'bank', label: '銀行' },
  { value: 'credit', label: 'クレジット' },
] as const;

export const CATEGORY_TYPE_OPTIONS = [
  { value: 'income', label: '収入' },
  { value: 'expense', label: '支出' },
] as const;

const ACCOUNT_TYPE_LABELS: Record<string, string> = Object.fromEntries(
  ACCOUNT_TYPE_OPTIONS.map((o) => [o.value, o.label])
);

const CATEGORY_TYPE_LABELS: Record<string, string> = Object.fromEntries(
  CATEGORY_TYPE_OPTIONS.map((o) => [o.value, o.label])
);

export function getAccountTypeLabel(type: string): string {
  return ACCOUNT_TYPE_LABELS[type] ?? type;
}

export function getCategoryTypeLabel(type: string): string {
  return CATEGORY_TYPE_LABELS[type] ?? type;
}
