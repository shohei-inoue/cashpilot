import { render, screen } from '@testing-library/react';
import { describe, expect, it } from 'vitest';
import EmptyState from './EmptyState';

describe('EmptyState', () => {
  it('renders message', () => {
    render(<EmptyState message="データがありません" />);
    expect(screen.getByText('データがありません')).toBeInTheDocument();
  });

  it('renders action when provided', () => {
    render(
      <EmptyState
        message="空です"
        action={<button type="button">追加する</button>}
      />
    );
    expect(screen.getByRole('button', { name: '追加する' })).toBeInTheDocument();
  });

  it('does not render action area when action is omitted', () => {
    const { container } = render(<EmptyState message="空です" />);
    expect(container.querySelector('.action')).not.toBeInTheDocument();
  });
});
