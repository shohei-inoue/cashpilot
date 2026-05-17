import { render, screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { beforeEach, describe, expect, it, vi } from 'vitest';
import GoalSection from './GoalSection';

const mockRefresh = vi.fn();

vi.mock('next/navigation', () => ({
  useRouter: () => ({ refresh: mockRefresh }),
}));

vi.mock('@/app/actions/goals', () => ({
  createGoal: vi.fn().mockResolvedValue({ id: 1 }),
  updateGoal: vi.fn().mockResolvedValue({ id: 1 }),
  deleteGoal: vi.fn().mockResolvedValue(undefined),
}));

const goals = [
  {
    id: 1,
    name: 'MacBook購入',
    target_amount: 300000,
    deadline: '2025-12-31',
  },
];

describe('GoalSection', () => {
  beforeEach(() => {
    vi.clearAllMocks();
  });

  it('renders goal with amount and deadline', () => {
    render(<GoalSection goals={goals} />);
    expect(screen.getByText('MacBook購入')).toBeInTheDocument();
    expect(screen.getByText('￥300,000')).toBeInTheDocument();
    expect(screen.getByText(/期限:/)).toBeInTheDocument();
  });

  it('creates a new goal on submit', async () => {
    const { createGoal } = await import('@/app/actions/goals');
    const user = userEvent.setup();
    render(<GoalSection goals={[]} />);

    await user.type(screen.getByLabelText(/目標名/), '旅行資金');
    await user.type(screen.getByLabelText(/目標金額/), '100000');
    await user.click(screen.getByRole('button', { name: '追加する' }));

    expect(createGoal).toHaveBeenCalledWith({
      name: '旅行資金',
      target_amount: 100000,
      deadline: undefined,
    });
    expect(mockRefresh).toHaveBeenCalled();
  });

  it('enters edit mode when edit is clicked', async () => {
    const user = userEvent.setup();
    render(<GoalSection goals={goals} />);
    await user.click(screen.getByRole('button', { name: '編集' }));
    expect(screen.getByRole('heading', { name: '目標を編集' })).toBeInTheDocument();
    expect(screen.getByDisplayValue('MacBook購入')).toBeInTheDocument();
  });
});
