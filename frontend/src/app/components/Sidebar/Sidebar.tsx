'use client';

import Link from 'next/link';
import { usePathname } from 'next/navigation';
import { SIDEBAR_ITEMS } from '../../constants/constants';
import styles from './Sidebar.module.scss';

const ICON_MAP: Record<string, string> = {
  home: 'home',
  transactions: 'receipt_long',
  simulation: 'trending_up',
  settings: 'settings',
};

type SidebarProps = {
  isOpen: boolean;
};

const Sidebar = ({ isOpen }: SidebarProps) => {
  const pathname = usePathname();

  return (
    <nav
      className={`${styles.sidebar} ${!isOpen ? styles.closed : ''}`}
      aria-label="メインナビゲーション"
      aria-hidden={!isOpen}
    >
      <ul className={styles.list}>
        {SIDEBAR_ITEMS.map((item) => {
          const isActive = item.activePath.some(
            (path) => pathname === path || pathname.startsWith(path + '/')
          );
          const iconName = ICON_MAP[item.icon] ?? item.icon;

          return (
            <li key={item.href}>
              <Link
                href={item.href}
                className={`${styles.link} ${isActive ? styles.active : ''}`}
                aria-current={isActive ? 'page' : undefined}
              >
                <span className={`material-symbols-rounded ${styles.icon}`}>{iconName}</span>
                <span className={styles.text}>{item.text}</span>
              </Link>
            </li>
          );
        })}
      </ul>
    </nav>
  );
};

export default Sidebar;
