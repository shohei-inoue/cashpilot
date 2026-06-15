'use client';

import { useState } from 'react';
import Link from 'next/link';
import { usePathname, useRouter } from 'next/navigation';
import { logout } from '../../actions/auth';
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
  const router = useRouter();
  const [isLoggingOut, setIsLoggingOut] = useState(false);

  const handleLogout = async () => {
    if (isLoggingOut) {
      return;
    }

    setIsLoggingOut(true);
    try {
      await logout();
    } finally {
      router.replace('/auth/login');
      router.refresh();
      setIsLoggingOut(false);
    }
  };

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
      <div className={styles.logoutSection}>
        <button
          type="button"
          className={styles.logoutButton}
          onClick={handleLogout}
          disabled={isLoggingOut}
        >
          <span className={`material-symbols-rounded ${styles.icon}`}>logout</span>
          <span className={styles.text}>{isLoggingOut ? 'ログアウト中...' : 'ログアウト'}</span>
        </button>
      </div>
    </nav>
  );
};

export default Sidebar;
