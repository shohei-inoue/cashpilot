import { render, screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { describe, expect, it } from 'vitest';
import MainContainer from './MainContainer';

describe('MainContainer', () => {
  it('renders children and passes userEmail to Header', () => {
    render(
      <MainContainer userEmail="test@cashpilot.local">
        <p>ページ内容</p>
      </MainContainer>
    );
    expect(screen.getByText('ページ内容')).toBeInTheDocument();
    expect(screen.getByText('test@cashpilot.local')).toBeInTheDocument();
  });

  it('toggles sidebar aria-expanded on menu click', async () => {
    const user = userEvent.setup();
    render(
      <MainContainer userEmail="test@example.com">
        <p>content</p>
      </MainContainer>
    );
    const toggle = screen.getByRole('button', { name: 'メニューを閉じる' });
    expect(toggle).toHaveAttribute('aria-expanded', 'true');
    await user.click(toggle);
    expect(screen.getByRole('button', { name: 'メニューを開く' })).toHaveAttribute(
      'aria-expanded',
      'false'
    );
  });
});
