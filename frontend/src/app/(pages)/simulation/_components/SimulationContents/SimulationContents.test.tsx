import { render, screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { describe, expect, it, vi } from 'vitest';
import SimulationContents from './SimulationContents';

vi.mock('@/app/actions/simulation', () => ({
  runSimulation: vi.fn().mockResolvedValue({
    start_balance: 50000,
    min_balance: 50000,
    end_balance: 170000,
    monthly_balances: [
      { month: '2025-05', balance: 120000 },
      { month: '2025-06', balance: 170000 },
    ],
    goal_projections: [],
  }),
}));

const sampleResult = {
  start_balance: 100000,
  min_balance: 80000,
  end_balance: 200000,
  monthly_balances: [{ month: '2025-05', balance: 150000 }],
  goal_projections: [],
};

describe('SimulationContents', () => {
  it('renders form and initial result', () => {
    render(<SimulationContents initialResult={sampleResult} />);
    expect(screen.getByRole('heading', { name: 'シミュレーション' })).toBeInTheDocument();
    expect(screen.getByText('開始残高')).toBeInTheDocument();
    expect(screen.getByText('￥100,000')).toBeInTheDocument();
  });

  it('shows initial error', () => {
    render(
      <SimulationContents initialResult={null} initialError="計算に失敗しました" />
    );
    expect(screen.getByText('計算に失敗しました')).toBeInTheDocument();
  });

  it('runs simulation on submit', async () => {
    const { runSimulation } = await import('@/app/actions/simulation');
    const user = userEvent.setup();
    render(<SimulationContents initialResult={sampleResult} />);

    await user.clear(screen.getByLabelText(/毎月の収入/));
    await user.type(screen.getByLabelText(/毎月の収入/), '300000');
    await user.click(screen.getByRole('button', { name: 'シミュレーション実行' }));

    expect(runSimulation).toHaveBeenCalledWith(
      expect.objectContaining({
        period_months: 12,
        monthly_income: 300000,
      })
    );
  });
});
