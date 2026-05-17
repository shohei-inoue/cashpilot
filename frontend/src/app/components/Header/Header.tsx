import styles from './Header.module.scss';
import Link from 'next/link';
import Image from 'next/image';

type HeaderProps = {
  isSidebarOpen: boolean;
  onToggleSidebar: () => void;
  userEmail: string;
};

const Header = ({ isSidebarOpen, onToggleSidebar, userEmail }: HeaderProps) => {
  return (
    <header className={styles.header}>
      <div className={styles.left}>
        <button
          type="button"
          className={styles.sidebarButton}
          onClick={onToggleSidebar}
          aria-label={isSidebarOpen ? 'メニューを閉じる' : 'メニューを開く'}
          aria-expanded={isSidebarOpen}
        >
          <span className={styles.sidebarIcon} aria-hidden />
        </button>
        <Link href="/" className={styles.logoLink}>
          <Image src="/cashpilot_logo.svg" alt="CashPilot" width={120} height={40} priority />
        </Link>
      </div>
      <div className={styles.right}>
        <span className={styles.userEmail} title={userEmail}>
          {userEmail}
        </span>
      </div>
    </header>
  );
};

export default Header;
