import { render, screen, waitFor } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { beforeEach, describe, expect, it, vi } from 'vitest';
import Sidebar from './Sidebar';

const mockReplace = vi.fn();
const mockRefresh = vi.fn();
const mockPathname = vi.fn(() => '/');
const mockLogout = vi.fn();

vi.mock('next/navigation', () => ({
  usePathname: () => mockPathname(),
  useRouter: () => ({
    replace: mockReplace,
    refresh: mockRefresh,
  }),
}));

vi.mock('../../actions/auth', () => ({
  logout: mockLogout,
}));

describe('Sidebar', () => {
  beforeEach(() => {
    vi.clearAllMocks();
    mockPathname.mockReturnValue('/');
    mockLogout.mockResolvedValue(undefined);
  });

  it('renders navigation links and logout button', () => {
    render(<Sidebar isOpen={true} />);

    expect(screen.getByRole('link', { name: /ダッシュボード/ })).toBeInTheDocument();
    expect(screen.getByRole('link', { name: /収支入力/ })).toBeInTheDocument();
    expect(screen.getByRole('button', { name: 'ログアウト' })).toBeInTheDocument();
  });

  it('calls logout and redirects to login page', async () => {
    const user = userEvent.setup();
    render(<Sidebar isOpen={true} />);

    await user.click(screen.getByRole('button', { name: 'ログアウト' }));

    await waitFor(() => {
      expect(mockLogout).toHaveBeenCalledTimes(1);
    });
    expect(mockReplace).toHaveBeenCalledWith('/auth/login');
    expect(mockRefresh).toHaveBeenCalledTimes(1);
  });

  it('disables logout button while processing to prevent double click', async () => {
    const user = userEvent.setup();
    let resolveLogout: (() => void) | undefined;
    const pendingLogout = new Promise<void>((resolve) => {
      resolveLogout = resolve;
    });
    mockLogout.mockReturnValueOnce(pendingLogout);

    render(<Sidebar isOpen={true} />);
    const logoutButton = screen.getByRole('button', { name: 'ログアウト' });

    await user.click(logoutButton);
    await waitFor(() => {
      expect(logoutButton).toBeDisabled();
    });

    await user.click(logoutButton);
    expect(mockLogout).toHaveBeenCalledTimes(1);

    resolveLogout?.();
    await waitFor(() => {
      expect(mockReplace).toHaveBeenCalledWith('/auth/login');
    });
  });
});
