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

/** Date を datetime-local 入力用の文字列に変換 */
export function toDatetimeLocalValue(date: Date = new Date()): string {
  const pad = (n: number) => String(n).padStart(2, '0');
  return `${date.getFullYear()}-${pad(date.getMonth() + 1)}-${pad(date.getDate())}T${pad(date.getHours())}:${pad(date.getMinutes())}`;
}

/** datetime-local の値を ISO 文字列に変換 */
export function datetimeLocalToISO(local: string): string {
  return new Date(local).toISOString();
}

/** ISO を date 入力用（YYYY-MM-DD）に変換 */
export function toDateInputValue(iso: string): string {
  const d = new Date(iso);
  if (Number.isNaN(d.getTime())) return '';
  const pad = (n: number) => String(n).padStart(2, '0');
  return `${d.getFullYear()}-${pad(d.getMonth() + 1)}-${pad(d.getDate())}`;
}

/** date 入力の開始日を ISO（その日 00:00:00 ローカル） */
export function dateInputToISOStart(date: string): string {
  const [y, m, d] = date.split('-').map(Number);
  return new Date(y, m - 1, d, 0, 0, 0, 0).toISOString();
}

/** date 入力の終了日を ISO（その日 23:59:59 ローカル） */
export function dateInputToISOEnd(date: string): string {
  const [y, m, d] = date.split('-').map(Number);
  return new Date(y, m - 1, d, 23, 59, 59, 999).toISOString();
}

/** YYYY-MM-DD または ISO を表示用日付に */
export function formatDateOnly(value: string): string {
  const normalized = value.includes('T') ? value : `${value}T00:00:00`;
  const d = new Date(normalized);
  if (Number.isNaN(d.getTime())) return value;
  return new Intl.DateTimeFormat('ja-JP', {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
  }).format(d);
}
