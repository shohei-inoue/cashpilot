import { render, screen } from '@testing-library/react';
import { describe, expect, it } from 'vitest';
import GoalProjectionList from './GoalProjectionList';

const projections = [
  {
    goal_id: 1,
    name: '引っ越し',
    target_amount: 500000,
    deadline: '2025-12-31',
    projected_balance: 600000,
    achievable: true,
  },
  {
    goal_id: 2,
    name: 'MacBook',
    target_amount: 300000,
    deadline: '2026-03-15',
    projected_balance: 200000,
    achievable: false,
  },
];

describe('GoalProjectionList', () => {
  it('shows empty message when no projections', () => {
    render(<GoalProjectionList projections={[]} />);
    expect(screen.getByText(/期限付きの目標/)).toBeInTheDocument();
  });

  it('renders achievable and unachievable badges', () => {
    render(<GoalProjectionList projections={projections} />);
    expect(screen.getByText('引っ越し')).toBeInTheDocument();
    expect(screen.getByText('達成見込み')).toBeInTheDocument();
    expect(screen.getByText('未達の見込み')).toBeInTheDocument();
  });
});
