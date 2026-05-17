'use client';

import { useRouter } from 'next/navigation';
import { useState } from 'react';
import { deleteTransaction } from '@/app/actions/transactions';
import type { Transaction } from '@/app/types/transaction';
import Button from '@/app/components/Button/Button';
import Card from '@/app/components/Card/Card';
import EmptyState from '@/app/components/EmptyState/EmptyState';
import Heading from '@/app/components/Heading/Heading';
import TransactionRow from '@/app/components/TransactionRow/TransactionRow';
import styles from './TransactionList.module.scss';

type TransactionListProps = {
  transactions: Transaction[];
  onEdit: (transaction: Transaction) => void;
};

const TransactionList = ({ transactions, onEdit }: TransactionListProps) => {
  const router = useRouter();
  const [deletingId, setDeletingId] = useState<number | null>(null);

  const handleDelete = async (tx: Transaction) => {
    if (!window.confirm(`「${tx.category_name ?? (tx.amount >= 0 ? '収入' : '支出')}」の取引を削除しますか？`)) {
      return;
    }
    setDeletingId(tx.id);
    try {
      await deleteTransaction(tx.id);
      router.refresh();
    } catch {
      window.alert('削除に失敗しました');
    } finally {
      setDeletingId(null);
    }
  };

  return (
    <Card>
      <Heading level={2} className={styles.title}>
        取引一覧
      </Heading>
      {transactions.length === 0 ? (
        <EmptyState message="この期間の取引はありません。" />
      ) : (
        <ul className={styles.list}>
          {transactions.map((tx) => (
            <li key={tx.id} className={styles.item}>
              <TransactionRow transaction={tx} />
              <div className={styles.actions}>
                <Button
                  type="button"
                  variant="secondary"
                  disabled={deletingId === tx.id}
                  onClick={() => onEdit(tx)}
                >
                  編集
                </Button>
                <Button
                  type="button"
                  variant="danger"
                  disabled={deletingId !== null}
                  onClick={() => handleDelete(tx)}
                >
                  削除
                </Button>
              </div>
            </li>
          ))}
        </ul>
      )}
    </Card>
  );
};

export default TransactionList;
