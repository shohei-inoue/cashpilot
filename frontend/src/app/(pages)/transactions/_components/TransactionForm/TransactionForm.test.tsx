import { render, screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { beforeEach, describe, expect, it, vi } from 'vitest';
import TransactionForm from './TransactionForm';

const mockRefresh = vi.fn();

vi.mock('next/navigation', () => ({
  useRouter: () => ({ refresh: mockRefresh }),
}));

vi.mock('@/app/actions/transactions', () => ({
  createTransaction: vi.fn().mockResolvedValue({ id: 1 }),
  updateTransaction: vi.fn().mockResolvedValue({ id: 1 }),
}));

const accounts = [{ id: 1, type: 'bank' as const, name: 'メイン口座' }];
const categories = [
  { id: 10, type: 'expense' as const, name: '食費' },
  { id: 11, type: 'income' as const, name: '給与' },
];

describe('TransactionForm', () => {
  beforeEach(() => {
    vi.clearAllMocks();
  });

  it('shows hint when no accounts exist', () => {
    render(<TransactionForm accounts={[]} categories={categories} />);
    expect(screen.getByText(/口座を追加/)).toBeInTheDocument();
  });

  it('renders income and expense toggle', () => {
    render(<TransactionForm accounts={accounts} categories={categories} />);
    expect(screen.getByRole('button', { name: '支出' })).toBeInTheDocument();
    expect(screen.getByRole('button', { name: '収入' })).toBeInTheDocument();
  });

  it('submits expense transaction', async () => {
    const { createTransaction } = await import('@/app/actions/transactions');
    const user = userEvent.setup();
    render(<TransactionForm accounts={accounts} categories={categories} />);

    await user.type(screen.getByLabelText(/金額/), '1500');
    await user.click(screen.getByRole('button', { name: '登録する' }));

    expect(createTransaction).toHaveBeenCalledWith(
      expect.objectContaining({
        account_id: 1,
        amount: -1500,
      })
    );
    expect(mockRefresh).toHaveBeenCalled();
  });
});
