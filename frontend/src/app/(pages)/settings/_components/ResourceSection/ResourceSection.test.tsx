import { render, screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { beforeEach, describe, expect, it, vi } from 'vitest';
import ResourceSection from './ResourceSection';

const TYPE_OPTIONS = [
  { value: 'bank', label: '銀行' },
  { value: 'cash', label: '現金' },
] as const;

const mockRefresh = vi.fn();

vi.mock('next/navigation', () => ({
  useRouter: () => ({ refresh: mockRefresh }),
}));

describe('ResourceSection', () => {
  const onCreate = vi.fn().mockResolvedValue(undefined);
  const onUpdate = vi.fn().mockResolvedValue(undefined);
  const onDelete = vi.fn().mockResolvedValue(undefined);

  const defaultProps = {
    title: '口座',
    description: '口座の説明',
    items: [{ id: 1, name: 'メイン銀行', type: 'bank' }],
    typeOptions: TYPE_OPTIONS,
    getTypeLabel: (type: string) => (type === 'bank' ? '銀行' : type),
    emptyMessage: '口座がありません',
    nameLabel: '口座名',
    typeLabel: '種別',
    createLabel: '口座を追加',
    onCreate,
    onUpdate,
    onDelete,
  };

  beforeEach(() => {
    vi.clearAllMocks();
  });

  it('renders items with type badge', () => {
    const { container } = render(<ResourceSection {...defaultProps} />);
    expect(screen.getByText('メイン銀行')).toBeInTheDocument();
    expect(container.querySelector('.typeBadge')).toHaveTextContent('銀行');
  });

  it('creates a new item on submit', async () => {
    const user = userEvent.setup();
    render(<ResourceSection {...defaultProps} items={[]} />);
    await user.type(screen.getByLabelText(/口座名/), '新規口座');
    await user.click(screen.getByRole('button', { name: '追加する' }));
    expect(onCreate).toHaveBeenCalledWith('bank', '新規口座');
    expect(mockRefresh).toHaveBeenCalled();
  });

  it('enters edit mode when edit is clicked', async () => {
    const user = userEvent.setup();
    render(<ResourceSection {...defaultProps} />);
    await user.click(screen.getByRole('button', { name: '編集' }));
    expect(screen.getByRole('heading', { name: '編集' })).toBeInTheDocument();
    expect(screen.getByDisplayValue('メイン銀行')).toBeInTheDocument();
  });
});
