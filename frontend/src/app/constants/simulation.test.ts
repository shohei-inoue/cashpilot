import { describe, expect, it } from 'vitest';
import {
  DEFAULT_SIMULATION_INPUT,
  SIMULATION_PERIOD_PRESETS,
} from './simulation';

describe('simulation constants', () => {
  it('has default period of 12 months', () => {
    expect(DEFAULT_SIMULATION_INPUT.period_months).toBe(12);
  });

  it('includes common period presets', () => {
    const values = SIMULATION_PERIOD_PRESETS.map((p) => p.value);
    expect(values).toContain(12);
    expect(values).toContain(36);
  });
});
