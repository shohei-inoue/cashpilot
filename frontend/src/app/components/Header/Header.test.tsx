import { render, screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { describe, expect, it, vi } from 'vitest';
import Header from './Header';

describe('Header', () => {
  it('displays user email', () => {
    render(
      <Header
        isSidebarOpen={true}
        onToggleSidebar={vi.fn()}
        userEmail="user@example.com"
      />
    );
    const email = screen.getByText('user@example.com');
    expect(email).toBeInTheDocument();
    expect(email).toHaveAttribute('title', 'user@example.com');
  });

  it('calls onToggleSidebar when menu button is clicked', async () => {
    const user = userEvent.setup();
    const onToggleSidebar = vi.fn();
    render(
      <Header
        isSidebarOpen={false}
        onToggleSidebar={onToggleSidebar}
        userEmail="user@example.com"
      />
    );
    await user.click(screen.getByRole('button', { name: 'メニューを開く' }));
    expect(onToggleSidebar).toHaveBeenCalledOnce();
  });

  it('shows close label when sidebar is open', () => {
    render(
      <Header
        isSidebarOpen={true}
        onToggleSidebar={vi.fn()}
        userEmail="user@example.com"
      />
    );
    expect(screen.getByRole('button', { name: 'メニューを閉じる' })).toHaveAttribute(
      'aria-expanded',
      'true'
    );
  });
});
