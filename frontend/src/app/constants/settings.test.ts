import { describe, expect, it } from 'vitest';
import {
  getAccountTypeLabel,
  getCategoryTypeLabel,
} from './settings';

describe('settings constants', () => {
  it('returns Japanese label for account types', () => {
    expect(getAccountTypeLabel('bank')).toBe('銀行');
    expect(getAccountTypeLabel('cash')).toBe('現金');
    expect(getAccountTypeLabel('credit')).toBe('クレジット');
  });

  it('returns Japanese label for category types', () => {
    expect(getCategoryTypeLabel('income')).toBe('収入');
    expect(getCategoryTypeLabel('expense')).toBe('支出');
  });

  it('returns raw type for unknown values', () => {
    expect(getAccountTypeLabel('unknown')).toBe('unknown');
  });
});
