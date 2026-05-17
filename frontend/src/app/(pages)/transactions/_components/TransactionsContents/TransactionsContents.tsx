'use client';

import { useState } from 'react';
import type { TransactionsPageData } from '@/app/actions/transactions-page';
import type { Transaction } from '@/app/types/transaction';
import Heading from '@/app/components/Heading/Heading';
import MainContent from '@/app/components/MainContent/MainContent';
import TransactionFilters from '../TransactionFilters/TransactionFilters';
import TransactionForm from '../TransactionForm/TransactionForm';
import TransactionList from '../TransactionList/TransactionList';
import styles from './TransactionsContents.module.scss';

type TransactionsContentsProps = {
  data: TransactionsPageData;
};

const TransactionsContents = ({ data }: TransactionsContentsProps) => {
  const [editingTransaction, setEditingTransaction] = useState<Transaction | null>(
    null
  );

  return (
    <MainContent>
      <header className={styles.header}>
        <Heading level={1}>収支入力</Heading>
        <p className={styles.lead}>
          取引の登録・一覧・編集ができます。口座やカテゴリは設定画面で管理します。
        </p>
      </header>

      <section className={styles.section}>
        <TransactionForm
          accounts={data.accounts}
          categories={data.categories}
          editingTransaction={editingTransaction}
          onCancelEdit={() => setEditingTransaction(null)}
        />
      </section>

      <section className={styles.section}>
        <TransactionFilters
          accounts={data.accounts}
          categories={data.categories}
          filter={data.filter}
        />
      </section>

      <section className={styles.section}>
        <TransactionList
          transactions={data.transactions}
          onEdit={setEditingTransaction}
        />
      </section>
    </MainContent>
  );
};

export default TransactionsContents;
