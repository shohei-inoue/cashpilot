import { describe, expect, it } from 'vitest';
import {
  formatDateTime,
  formatExpenseTotal,
  formatYearMonth,
  formatYen,
  getCurrentMonthLabel,
  getCurrentMonthRange,
  getMonthsRange,
} from './format';

describe('formatYen', () => {
  it('formats amount as JPY', () => {
    expect(formatYen(30000)).toBe('￥30,000');
  });

  it('formats negative amount', () => {
    expect(formatYen(-20000)).toBe('-￥20,000');
  });
});

describe('formatExpenseTotal', () => {
  it('returns absolute value formatted as JPY', () => {
    expect(formatExpenseTotal(-20000)).toBe('￥20,000');
  });
});

describe('formatDateTime', () => {
  it('formats valid ISO string', () => {
    const result = formatDateTime('2025-01-15T12:00:00Z');
    expect(result).toContain('2025');
  });

  it('returns original string when invalid', () => {
    expect(formatDateTime('invalid')).toBe('invalid');
  });
});

describe('formatYearMonth', () => {
  it('formats YYYY-MM to Japanese label', () => {
    expect(formatYearMonth('2025-03')).toBe('2025年3月');
  });

  it('returns period as-is when malformed', () => {
    expect(formatYearMonth('invalid')).toBe('invalid');
  });
});

describe('getCurrentMonthLabel', () => {
  it('returns year and month in Japanese', () => {
    const now = new Date();
    const expected = `${now.getFullYear()}年${now.getMonth() + 1}月`;
    expect(getCurrentMonthLabel()).toBe(expected);
  });
});

describe('getCurrentMonthRange', () => {
  it('returns from at start of month and to at end of month', () => {
    const { from, to } = getCurrentMonthRange();
    expect(new Date(from).getTime()).toBeLessThanOrEqual(new Date(to).getTime());
  });
});

describe('getMonthsRange', () => {
  it('returns from before to', () => {
    const { from, to } = getMonthsRange(12);
    expect(new Date(from).getTime()).toBeLessThanOrEqual(new Date(to).getTime());
  });
});
