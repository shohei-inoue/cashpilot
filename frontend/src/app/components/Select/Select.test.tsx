import { render, screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { describe, expect, it, vi } from 'vitest';
import Select from './Select';

const OPTIONS = [
  { value: 'bank', label: '銀行' },
  { value: 'cash', label: '現金' },
] as const;

describe('Select', () => {
  it('renders label and options', () => {
    render(
      <Select
        label="種別"
        value="bank"
        options={OPTIONS}
        onChange={vi.fn()}
      />
    );
    expect(screen.getByLabelText(/種別/)).toBeInTheDocument();
    expect(screen.getByRole('option', { name: '銀行' })).toBeInTheDocument();
    expect(screen.getByRole('option', { name: '現金' })).toBeInTheDocument();
  });

  it('calls onChange when selection changes', async () => {
    const user = userEvent.setup();
    const onChange = vi.fn();
    render(
      <Select label="種別" value="bank" options={OPTIONS} onChange={onChange} />
    );
    await user.selectOptions(screen.getByLabelText(/種別/), 'cash');
    expect(onChange).toHaveBeenCalled();
  });

  it('shows error message', () => {
    render(
      <Select
        label="種別"
        value="bank"
        options={OPTIONS}
        onChange={vi.fn()}
        error="必須です"
      />
    );
    expect(screen.getByRole('alert')).toHaveTextContent('必須です');
  });
});
