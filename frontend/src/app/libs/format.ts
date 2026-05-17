/** 金額を日本円表示（整数） */
export function formatYen(amount: number): string {
  return new Intl.NumberFormat('ja-JP', {
    style: 'currency',
    currency: 'JPY',
    maximumFractionDigits: 0,
  }).format(amount);
}

/** 支出合計（負の値）を表示用の正の金額に */
export function formatExpenseTotal(expense: number): string {
  return formatYen(Math.abs(expense));
}

/** ISO 日時を表示用（日付 + 時刻） */
export function formatDateTime(iso: string): string {
  const d = new Date(iso);
  if (Number.isNaN(d.getTime())) return iso;
  return new Intl.DateTimeFormat('ja-JP', {
    year: 'numeric',
    month: 'short',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit',
  }).format(d);
}

/** YYYY-MM を「YYYY年M月」 */
export function formatYearMonth(period: string): string {
  const [y, m] = period.split('-');
  if (!y || !m) return period;
  return `${y}年${Number(m)}月`;
}

/** 今月の表示ラベル */
export function getCurrentMonthLabel(): string {
  const now = new Date();
  return `${now.getFullYear()}年${now.getMonth() + 1}月`;
}

/** 今月の from / to（ISO） */
export function getCurrentMonthRange(): { from: string; to: string } {
  const now = new Date();
  const from = new Date(now.getFullYear(), now.getMonth(), 1);
  const to = new Date(now.getFullYear(), now.getMonth() + 1, 0, 23, 59, 59, 999);
  return { from: from.toISOString(), to: to.toISOString() };
}

/** 過去 N ヶ月の from / to（ISO） */
export function getMonthsRange(months: number): { from: string; to: string } {
  const now = new Date();
  const from = new Date(now.getFullYear(), now.getMonth() - (months - 1), 1);
  return { from: from.toISOString(), to: now.toISOString() };
}
