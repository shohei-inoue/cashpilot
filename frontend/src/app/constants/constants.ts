import { SidebarItem } from '../types/ui';

export const SIDEBAR_ITEMS: SidebarItem[] = [
  {
    href: '/',
    icon: 'home',
    text: 'ダッシュボード',
    activePath: ['/'],
  },
  {
    href: '/transactions',
    icon: 'transactions',
    text: '収支入力',
    activePath: ['/transactions'],
  },
  {
    href: '/simulation',
    icon: 'simulation',
    text: 'シミュレーション',
    activePath: ['/simulation'],
  },
  {
    href: '/settings',
    icon: 'settings',
    text: '設定',
    activePath: ['/settings'],
  },
];