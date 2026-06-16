import Link from 'next/link';
import Card from '@/app/components/Card/Card';
import EmptyState from '@/app/components/EmptyState/EmptyState';
import Heading from '@/app/components/Heading/Heading';
import MainContent from '@/app/components/MainContent/MainContent';
import styles from './DashboardErrorState.module.scss';

type DashboardErrorStateProps = {
  message: string;
};

const DashboardErrorState = ({ message }: DashboardErrorStateProps) => {
  return (
    <MainContent>
      <header className={styles.header}>
        <Heading level={1}>ダッシュボード</Heading>
      </header>
      <Card>
        <EmptyState
          message={message}
          action={
            <div className={styles.actions}>
              <Link href="/" className={styles.primaryButton}>
                再試行する
              </Link>
              <div className={styles.links}>
                <Link href="/transactions">収支を入力</Link>
                <Link href="/simulation">シミュレーション</Link>
                <Link href="/settings">設定を開く</Link>
              </div>
            </div>
          }
        />
      </Card>
    </MainContent>
  );
};

export default DashboardErrorState;
