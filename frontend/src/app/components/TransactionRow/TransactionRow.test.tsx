import { render, screen } from '@testing-library/react';
import { describe, expect, it } from 'vitest';
import type { Transaction } from '@/app/types/transaction';
import TransactionRow from './TransactionRow';

const baseTransaction: Transaction = {
  id: 1,
  account_id: 1,
  account_name: 'メイン口座',
  category_name: '給与',
  amount: 50000,
  occurred_at: '2025-01-10T12:00:00Z',
};

describe('TransactionRow', () => {
  it('renders category, account, amount and datetime', () => {
    render(<TransactionRow transaction={baseTransaction} />);
    expect(screen.getByText('給与')).toBeInTheDocument();
    expect(screen.getByText('メイン口座')).toBeInTheDocument();
    expect(screen.getByText('￥50,000')).toBeInTheDocument();
    expect(screen.getByRole('time')).toHaveAttribute('dateTime', baseTransaction.occurred_at);
  });

  it('shows memo when present', () => {
    render(
      <TransactionRow
        transaction={{ ...baseTransaction, memo: 'ボーナス' }}
      />
    );
    expect(screen.getByText('ボーナス')).toBeInTheDocument();
  });

  it('falls back to 収入 label when category is missing and amount is positive', () => {
    render(
      <TransactionRow
        transaction={{ ...baseTransaction, category_name: undefined }}
      />
    );
    expect(screen.getByText('収入')).toBeInTheDocument();
  });

  it('falls back to 支出 label when category is missing and amount is negative', () => {
    render(
      <TransactionRow
        transaction={{
          ...baseTransaction,
          category_name: undefined,
          amount: -1000,
        }}
      />
    );
    expect(screen.getByText('支出')).toBeInTheDocument();
  });
});
