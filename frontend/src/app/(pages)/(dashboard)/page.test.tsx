import { render, screen } from '@testing-library/react';
import { beforeEach, describe, expect, it, vi } from 'vitest';
import type { DashboardData } from '@/app/actions/dashboard';
import DashboardPage from './page';

const {
  mockGetDashboardData,
  mockRedirect,
  mockDashboardContents,
  mockDashboardErrorState,
} = vi.hoisted(() => ({
  mockGetDashboardData: vi.fn(),
  mockRedirect: vi.fn(),
  mockDashboardContents: vi.fn(({ data }: { data: DashboardData }) => (
    <div data-testid="dashboard-contents">{data.monthLabel}</div>
  )),
  mockDashboardErrorState: vi.fn(({ message }: { message: string }) => (
    <div data-testid="dashboard-error">{message}</div>
  )),
}));

vi.mock('next/navigation', () => ({
  redirect: mockRedirect,
}));

vi.mock('@/app/actions/dashboard', () => ({
  getDashboardData: mockGetDashboardData,
}));

vi.mock('@/app/components/AppShell/AppShell', () => ({
  default: ({ children }: { children: React.ReactNode }) => (
    <div data-testid="app-shell">{children}</div>
  ),
}));

vi.mock('./_components/DashboardContents/DashboardContents', () => ({
  default: mockDashboardContents,
}));

vi.mock('./_components/DashboardErrorState/DashboardErrorState', () => ({
  default: mockDashboardErrorState,
}));

describe('DashboardPage', () => {
  beforeEach(() => {
    vi.clearAllMocks();
  });

  it('未認証時はログインへリダイレクトする', async () => {
    mockGetDashboardData.mockResolvedValue({ status: 'unauthorized' });
    mockRedirect.mockImplementation(() => {
      throw new Error('redirected');
    });

    await expect(DashboardPage()).rejects.toThrow('redirected');
    expect(mockRedirect).toHaveBeenCalledWith('/auth/login');
  });

  it('取得成功時はダッシュボード内容を表示する', async () => {
    mockGetDashboardData.mockResolvedValue({
      status: 'ok',
      data: {
        monthLabel: '2026年6月',
        balance: 1000,
        monthSummary: {
          total_income: 2000,
          total_expense: -1000,
          net_cashflow: 1000,
        },
        recentTransactions: [],
        cashflowTrend: [],
      },
    });

    const ui = await DashboardPage();
    render(ui);

    expect(screen.getByTestId('app-shell')).toBeInTheDocument();
    expect(screen.getByTestId('dashboard-contents')).toHaveTextContent('2026年6月');
  });

  it('APIエラー時はエラー表示を出す', async () => {
    mockGetDashboardData.mockResolvedValue({
      status: 'error',
      message: '取得失敗',
    });

    const ui = await DashboardPage();
    render(ui);

    expect(screen.getByTestId('dashboard-error')).toHaveTextContent('取得失敗');
  });
});
