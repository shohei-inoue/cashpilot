import { render, screen } from '@testing-library/react';
import { describe, expect, it } from 'vitest';
import type { CashflowTrendPoint } from '@/app/actions/dashboard';
import CashflowChart from './CashflowChart';

const samplePoints: CashflowTrendPoint[] = [
  { period: '2025-01', label: '2025年1月', balance: 10000 },
  { period: '2025-02', label: '2025年2月', balance: 25000 },
];

describe('CashflowChart', () => {
  it('shows empty message when no points', () => {
    render(<CashflowChart points={[]} />);
    expect(
      screen.getByText('取引データが増えると、残高推移が表示されます。')
    ).toBeInTheDocument();
  });

  it('renders chart with bars for each period', () => {
    render(<CashflowChart points={samplePoints} />);
    expect(screen.getByRole('img', { name: '月次残高推移' })).toBeInTheDocument();
    expect(screen.getByText('1月')).toBeInTheDocument();
    expect(screen.getByText('2月')).toBeInTheDocument();
  });

  it('sets title on bar with balance tooltip', () => {
    const { container } = render(<CashflowChart points={samplePoints} />);
    const bars = container.querySelectorAll('[title^="2025年1月"]');
    expect(bars.length).toBeGreaterThan(0);
    expect(bars[0]).toHaveAttribute('title', '2025年1月: ￥10,000');
  });
});
