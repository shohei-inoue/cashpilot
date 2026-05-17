import { render, screen } from '@testing-library/react';
import { describe, expect, it } from 'vitest';
import Heading from './Heading';

describe('Heading', () => {
  it.each([1, 2, 3, 4, 5, 6] as const)(
    'renders h%i element for level %i',
    (level) => {
      render(<Heading level={level}>見出し</Heading>);
      expect(screen.getByRole('heading', { level, name: '見出し' })).toBeInTheDocument();
    }
  );

  it('applies custom className', () => {
    render(
      <Heading level={2} className="custom">
        セクション
      </Heading>
    );
    expect(screen.getByRole('heading', { name: 'セクション' })).toHaveClass('custom');
  });

  it('sets id when provided', () => {
    render(
      <Heading level={1} id="page-title">
        タイトル
      </Heading>
    );
    expect(screen.getByRole('heading', { name: 'タイトル' })).toHaveAttribute('id', 'page-title');
  });
});
