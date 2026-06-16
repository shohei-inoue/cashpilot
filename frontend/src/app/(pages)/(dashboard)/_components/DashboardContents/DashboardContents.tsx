import Link from 'next/link';
import type { DashboardData } from '@/app/actions/dashboard';
import BalanceDisplay from '@/app/components/BalanceDisplay/BalanceDisplay';
import Card from '@/app/components/Card/Card';
import EmptyState from '@/app/components/EmptyState/EmptyState';
import Heading from '@/app/components/Heading/Heading';
import MainContent from '@/app/components/MainContent/MainContent';
import TransactionRow from '@/app/components/TransactionRow/TransactionRow';
import { formatExpenseTotal, formatYen } from '@/app/libs/format';
import CashflowChart from '../CashflowChart/CashflowChart';
import styles from './DashboardContents.module.scss';

type DashboardContentsProps = {
  data: DashboardData;
};

const DashboardContents = ({ data }: DashboardContentsProps) => {
  const { balance, monthSummary, monthLabel, recentTransactions, cashflowTrend } =
    data;

  return (
    <MainContent>
      <header className={styles.header}>
        <div>
          <Heading level={1}>ダッシュボード</Heading>
          <p className={styles.subtitle}>現状を把握して、次のアクションにつなげましょう。</p>
        </div>
      </header>

      <section className={styles.overviewSection}>
        <Card>
          <div className={styles.sectionHeader}>
            <Heading level={2} className={styles.sectionTitle}>
              現状把握
            </Heading>
            <p className={styles.sectionHint}>
              残高と{monthLabel}の収支を1つのカードで確認できます。
            </p>
          </div>
          <div className={styles.balanceWrap}>
            <BalanceDisplay label="現在の残高" amount={balance} />
          </div>
          <dl className={styles.summaryGrid}>
            <div className={styles.summaryItem}>
              <dt>今月の収入</dt>
              <dd className={styles.income}>{formatYen(monthSummary.total_income)}</dd>
            </div>
            <div className={styles.summaryItem}>
              <dt>今月の支出</dt>
              <dd className={styles.expense}>
                {formatExpenseTotal(monthSummary.total_expense)}
              </dd>
            </div>
            <div className={styles.summaryItem}>
              <dt>今月の収支差分</dt>
              <dd
                className={
                  monthSummary.net_cashflow >= 0 ? styles.income : styles.expense
                }
              >
                {formatYen(monthSummary.net_cashflow)}
              </dd>
            </div>
          </dl>
        </Card>
      </section>

      <section className={styles.actionsSection}>
        <Card>
          <Heading level={2} className={styles.sectionTitle}>
            次のアクション
          </Heading>
          <p className={styles.sectionHint}>
            収支入力・将来シミュレーション・設定をすぐに開けます。
          </p>
          <div className={styles.actions}>
            <Link href="/transactions" className={styles.linkButton}>
              収支を入力する
            </Link>
            <Link href="/simulation" className={styles.linkButton}>
              シミュレーション
            </Link>
            <Link href="/settings" className={`${styles.linkButton} ${styles.secondary}`}>
              設定を開く
            </Link>
          </div>
        </Card>
      </section>

      <section className={styles.chartSection}>
        <Card>
          <Heading level={2} className={styles.sectionTitle}>
            残高推移（直近12ヶ月）
          </Heading>
          <p className={styles.sectionHint}>
            取引の累計から算出しています。将来予測はシミュレーション画面で確認できます。
          </p>
          <CashflowChart points={cashflowTrend} />
        </Card>
      </section>

      <section className={styles.transactionsSection}>
        <Card>
          <div className={styles.transactionsHeader}>
            <Heading level={2} className={styles.sectionTitle}>
              直近の取引
            </Heading>
            <Link href="/transactions" className={styles.viewAll}>
              すべて見る
            </Link>
          </div>
          {recentTransactions.length === 0 ? (
            <EmptyState
              message="まだ取引がありません。最初の収支を登録しましょう。"
              action={
                <Link href="/transactions" className={styles.linkButton}>
                  収支を入力する
                </Link>
              }
            />
          ) : (
            <ul className={styles.transactionList}>
              {recentTransactions.map((tx) => (
                <TransactionRow key={tx.id} transaction={tx} />
              ))}
            </ul>
          )}
        </Card>
      </section>
    </MainContent>
  );
};

export default DashboardContents;
