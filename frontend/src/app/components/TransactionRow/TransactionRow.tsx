import Amount from '@/app/components/Amount/Amount';
import { formatDateTime } from '@/app/libs/format';
import type { Transaction } from '@/app/types/transaction';
import styles from './TransactionRow.module.scss';

type TransactionRowProps = {
  transaction: Transaction;
};

const TransactionRow = ({ transaction }: TransactionRowProps) => {
  const categoryLabel =
    transaction.category_name ?? (transaction.amount >= 0 ? '収入' : '支出');

  return (
    <li className={styles.row}>
      <div className={styles.main}>
        <span className={styles.category}>{categoryLabel}</span>
        <span className={styles.meta}>
          {transaction.account_name && (
            <span className={styles.account}>{transaction.account_name}</span>
          )}
          <time className={styles.date} dateTime={transaction.occurred_at}>
            {formatDateTime(transaction.occurred_at)}
          </time>
        </span>
        {transaction.memo && <span className={styles.memo}>{transaction.memo}</span>}
      </div>
      <Amount amount={transaction.amount} />
    </li>
  );
};

export default TransactionRow;
