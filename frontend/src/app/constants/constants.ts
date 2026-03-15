import { SidebarItem } from '../types/ui';

export const SIDEBAR_ITEMS: SidebarItem[] = [
  {
    href: '/',
    icon: 'home',
    text: 'Dashboard',
    activePath: ['/'],
  },
  {
    href: '/transactions',
    icon: 'transactions',
    text: 'Transactions',
    activePath: ['/transactions'],
  },
  {
    href: '/simulation',
    icon: 'simulation',
    text: 'Simulation',
    activePath: ['/simulation'],
  },
];