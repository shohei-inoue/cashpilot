import { render, screen } from '@testing-library/react';
import { describe, expect, it } from 'vitest';
import DashboardErrorState from './DashboardErrorState';

describe('DashboardErrorState', () => {
  it('エラーメッセージと再試行導線を表示する', () => {
    render(
      <DashboardErrorState message="ダッシュボード情報の取得に失敗しました。時間をおいて再試行してください。" />
    );

    expect(
      screen.getByText('ダッシュボード情報の取得に失敗しました。時間をおいて再試行してください。')
    ).toBeInTheDocument();
    expect(screen.getByRole('link', { name: '再試行する' })).toHaveAttribute('href', '/');
    expect(screen.getByRole('link', { name: '収支を入力' })).toHaveAttribute(
      'href',
      '/transactions'
    );
    expect(screen.getByRole('link', { name: 'シミュレーション' })).toHaveAttribute(
      'href',
      '/simulation'
    );
    expect(screen.getByRole('link', { name: '設定を開く' })).toHaveAttribute(
      'href',
      '/settings'
    );
  });
});
