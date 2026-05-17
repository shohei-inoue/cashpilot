'use client';

import Link from 'next/link';
import { useCallback, useState } from 'react';
import { runSimulation } from '@/app/actions/simulation';
import {
  DEFAULT_SIMULATION_INPUT,
  SIMULATION_PERIOD_PRESETS,
} from '@/app/constants/simulation';
import type {
  SimulationRunInput,
  SimulationRunResponse,
} from '@/app/types/simulation';
import BalanceDisplay from '@/app/components/BalanceDisplay/BalanceDisplay';
import Button from '@/app/components/Button/Button';
import Card from '@/app/components/Card/Card';
import ErrorBlock from '@/app/components/ErrorBlock/ErrorBlock';
import Form from '@/app/components/Form/Form';
import Heading from '@/app/components/Heading/Heading';
import Input from '@/app/components/Input/Input';
import MainContent from '@/app/components/MainContent/MainContent';
import { formatYearMonth } from '@/app/libs/format';
import CashflowChart from '@/app/(pages)/(dashboard)/_components/CashflowChart/CashflowChart';
import GoalProjectionList from '../GoalProjectionList/GoalProjectionList';
import styles from './SimulationContents.module.scss';

type SimulationContentsProps = {
  initialResult: SimulationRunResponse | null;
  initialError?: string | null;
};

const SimulationContents = ({
  initialResult,
  initialError = null,
}: SimulationContentsProps) => {
  const [periodMonths, setPeriodMonths] = useState(
    String(DEFAULT_SIMULATION_INPUT.period_months)
  );
  const [monthlyIncome, setMonthlyIncome] = useState(
    String(DEFAULT_SIMULATION_INPUT.monthly_income)
  );
  const [monthlyExpense, setMonthlyExpense] = useState(
    String(DEFAULT_SIMULATION_INPUT.monthly_expense)
  );
  const [hourlyRate, setHourlyRate] = useState('');
  const [hoursPerMonth, setHoursPerMonth] = useState('');
  const [result, setResult] = useState<SimulationRunResponse | null>(initialResult);
  const [error, setError] = useState<string | null>(initialError);
  const [loading, setLoading] = useState(false);

  const buildInput = useCallback((): SimulationRunInput | null => {
    const period = parseInt(periodMonths, 10);
    if (!period || period < 1 || period > 120) {
      setError('期間は1〜120ヶ月で指定してください');
      return null;
    }
    const income = parseInt(monthlyIncome, 10) || 0;
    const expense = parseInt(monthlyExpense, 10) || 0;
    const input: SimulationRunInput = {
      period_months: period,
      monthly_income: Math.max(0, income),
      monthly_expense: Math.max(0, expense),
    };
    const rate = parseInt(hourlyRate, 10);
    const hours = parseInt(hoursPerMonth, 10);
    if (rate > 0 && hours > 0) {
      input.hourly_rate = rate;
      input.hours_per_month = hours;
    }
    return input;
  }, [periodMonths, monthlyIncome, monthlyExpense, hourlyRate, hoursPerMonth]);

  const executeSimulation = useCallback(
    async (input: SimulationRunInput) => {
      setLoading(true);
      setError(null);
      try {
        const res = await runSimulation(input);
        setResult(res);
      } catch (err) {
        setResult(null);
        setError(err instanceof Error ? err.message : 'シミュレーションに失敗しました');
      } finally {
        setLoading(false);
      }
    },
    []
  );

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    const input = buildInput();
    if (!input) return;
    await executeSimulation(input);
  };

  const chartPoints =
    result?.monthly_balances.map((mb) => ({
      period: mb.month,
      label: formatYearMonth(mb.month),
      balance: mb.balance,
    })) ?? [];

  return (
    <MainContent>
      <header className={styles.header}>
        <Heading level={1}>シミュレーション</Heading>
        <p className={styles.lead}>
          現在の残高と将来の収支前提から残高推移を試算します。目標は
          <Link href="/settings" className={styles.inlineLink}>
            設定画面
          </Link>
          で管理できます。
        </p>
      </header>

      <section className={styles.section}>
        <Card>
          <Heading level={2} className={styles.sectionTitle}>
            前提条件
          </Heading>
          <Form onSubmit={handleSubmit} className={styles.form}>
            {error && <ErrorBlock>{error}</ErrorBlock>}

            <div className={styles.periodRow}>
              <Input
                label="シミュレーション期間（ヶ月）"
                type="number"
                name="period_months"
                value={periodMonths}
                onChange={(e) => setPeriodMonths(e.target.value)}
                required
                disabled={loading}
              />
              <div className={styles.presets} role="group" aria-label="期間プリセット">
                {SIMULATION_PERIOD_PRESETS.map((preset) => (
                  <button
                    key={preset.value}
                    type="button"
                    className={styles.presetButton}
                    disabled={loading}
                    onClick={() => setPeriodMonths(String(preset.value))}
                  >
                    {preset.label}
                  </button>
                ))}
              </div>
            </div>

            <Input
              label="毎月の収入（円）"
              type="number"
              name="monthly_income"
              value={monthlyIncome}
              onChange={(e) => setMonthlyIncome(e.target.value)}
              disabled={loading}
              placeholder="0"
            />
            <Input
              label="毎月の固定費（円）"
              type="number"
              name="monthly_expense"
              value={monthlyExpense}
              onChange={(e) => setMonthlyExpense(e.target.value)}
              disabled={loading}
              placeholder="0"
            />

            <div className={styles.sideJobRow}>
              <Input
                label="副業: 時給（円）"
                type="number"
                name="hourly_rate"
                value={hourlyRate}
                onChange={(e) => setHourlyRate(e.target.value)}
                disabled={loading}
                placeholder="任意"
              />
              <Input
                label="副業: 月間時間"
                type="number"
                name="hours_per_month"
                value={hoursPerMonth}
                onChange={(e) => setHoursPerMonth(e.target.value)}
                disabled={loading}
                placeholder="任意"
              />
            </div>

            <Button type="submit" variant="primary" disabled={loading}>
              {loading ? '計算中...' : 'シミュレーション実行'}
            </Button>
          </Form>
        </Card>
      </section>

      {result && (
        <section className={styles.section}>
          <Card>
            <Heading level={2} className={styles.sectionTitle}>
              試算結果
            </Heading>
            <div className={styles.summary}>
              <BalanceDisplay label="開始残高" amount={result.start_balance} />
              <BalanceDisplay label="期間中最小残高" amount={result.min_balance} />
              <BalanceDisplay label="終了時残高" amount={result.end_balance} />
            </div>

            <Heading level={3} className={styles.chartTitle}>
              残高推移
            </Heading>
            <CashflowChart points={chartPoints} />

            <Heading level={3} className={styles.goalsTitle}>
              目標達成見込み
            </Heading>
            <GoalProjectionList projections={result.goal_projections ?? []} />
          </Card>
        </section>
      )}
    </MainContent>
  );
};

export default SimulationContents;
