import { render, screen } from '@testing-library/react';
import { describe, expect, it } from 'vitest';
import Amount from './Amount';

describe('Amount', () => {
  it('renders positive amount with income class when variant is neutral', () => {
    const { container } = render(<Amount amount={5000} />);
    expect(screen.getByText('￥5,000')).toBeInTheDocument();
    expect(container.firstChild).toHaveClass('income');
  });

  it('renders negative amount with expense class when variant is neutral', () => {
    const { container } = render(<Amount amount={-3000} />);
    expect(screen.getByText('-￥3,000')).toBeInTheDocument();
    expect(container.firstChild).toHaveClass('expense');
  });

  it('uses explicit income variant', () => {
    const { container } = render(<Amount amount={1000} variant="income" />);
    expect(container.firstChild).toHaveClass('income');
  });

  it('uses explicit expense variant', () => {
    const { container } = render(<Amount amount={-1000} variant="expense" />);
    expect(container.firstChild).toHaveClass('expense');
  });

  it('applies custom className', () => {
    const { container } = render(<Amount amount={0} className="extra" />);
    expect(container.firstChild).toHaveClass('extra');
    expect(container.firstChild).toHaveClass('neutral');
  });
});
