import { render, screen } from '@testing-library/react';
import { describe, expect, it } from 'vitest';
import BalanceDisplay from './BalanceDisplay';

describe('BalanceDisplay', () => {
  it('renders label and formatted amount', () => {
    render(<BalanceDisplay label="現在の残高" amount={30000} />);
    expect(screen.getByText('現在の残高')).toBeInTheDocument();
    expect(screen.getByText('￥30,000')).toBeInTheDocument();
  });
});
