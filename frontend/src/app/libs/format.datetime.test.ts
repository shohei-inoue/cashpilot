import { describe, expect, it } from 'vitest';
import {
  dateInputToISOEnd,
  dateInputToISOStart,
  datetimeLocalToISO,
  formatDateOnly,
  toDateInputValue,
  toDatetimeLocalValue,
} from './format';

describe('datetime helpers', () => {
  it('toDatetimeLocalValue returns YYYY-MM-DDTHH:mm format', () => {
    const value = toDatetimeLocalValue(new Date('2025-03-15T14:30:00'));
    expect(value).toMatch(/^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}$/);
  });

  it('datetimeLocalToISO converts to ISO string', () => {
    const iso = datetimeLocalToISO('2025-03-15T14:30');
    expect(new Date(iso).toISOString()).toBe(iso);
  });

  it('toDateInputValue extracts date part', () => {
    expect(toDateInputValue('2025-03-15T14:30:00Z')).toBe('2025-03-15');
  });

  it('dateInputToISOStart is before dateInputToISOEnd', () => {
    const start = dateInputToISOStart('2025-03-01');
    const end = dateInputToISOEnd('2025-03-31');
    expect(new Date(start).getTime()).toBeLessThan(new Date(end).getTime());
  });

  it('formatDateOnly formats YYYY-MM-DD', () => {
    const formatted = formatDateOnly('2025-12-31');
    expect(formatted).toContain('2025');
    expect(formatted).toContain('31');
  });
});
