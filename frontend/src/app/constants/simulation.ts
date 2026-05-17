import type { SimulationRunInput } from '../types/simulation';

export const SIMULATION_PERIOD_PRESETS = [
  { value: 12, label: '1年' },
  { value: 36, label: '3年' },
  { value: 60, label: '5年' },
] as const;

export const DEFAULT_SIMULATION_INPUT: SimulationRunInput = {
  period_months: 12,
  monthly_income: 0,
  monthly_expense: 0,
};
