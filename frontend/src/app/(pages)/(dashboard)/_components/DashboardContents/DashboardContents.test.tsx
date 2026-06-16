import { render, screen } from '@testing-library/react';
import { describe, expect, it } from 'vitest';
import type { DashboardData } from '@/app/actions/dashboard';
import DashboardContents from './DashboardContents';

const baseData: DashboardData = {
  monthLabel: '2026年6月',
  balance: 120000,
  monthSummary: {
    total_income: 300000,
    total_expense: -180000,
    net_cashflow: 120000,
  },
  recentTransactions: [
    {
      id: 1,
      account_id: 1,
      account_name: 'メイン口座',
      category_name: '給与',
      amount: 300000,
      occurred_at: '2026-06-10T12:00:00Z',
    },
  ],
  cashflowTrend: [
    { period: '2026-05', label: '2026年5月', balance: 100000 },
    { period: '2026-06', label: '2026年6月', balance: 120000 },
  ],
};

describe('DashboardContents', () => {
  it('現状把握カードと次アクションCTAを表示する', () => {
    render(<DashboardContents data={baseData} />);

    expect(screen.getByRole('heading', { name: '現状把握' })).toBeInTheDocument();
    expect(screen.getByText('現在の残高')).toBeInTheDocument();
    expect(screen.getByText('今月の収支差分')).toBeInTheDocument();
    expect(
      screen.getByRole('link', { name: '収支を入力する' })
    ).toHaveAttribute('href', '/transactions');
    expect(screen.getByRole('link', { name: 'シミュレーション' })).toHaveAttribute(
      'href',
      '/simulation'
    );
    expect(screen.getByRole('link', { name: '設定を開く' })).toHaveAttribute(
      'href',
      '/settings'
    );
  });

  it('取引が0件のとき空状態を表示する', () => {
    render(<DashboardContents data={{ ...baseData, recentTransactions: [] }} />);

    expect(
      screen.getByText('まだ取引がありません。最初の収支を登録しましょう。')
    ).toBeInTheDocument();
    const transactionLinks = screen.getAllByRole('link', { name: '収支を入力する' });
    expect(transactionLinks.length).toBeGreaterThan(0);
    expect(transactionLinks.some((link) => link.getAttribute('href') === '/transactions')).toBe(
      true
    );
  });
});
