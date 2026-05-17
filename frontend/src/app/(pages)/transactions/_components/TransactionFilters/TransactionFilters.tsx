'use client';

import { useRouter } from 'next/navigation';
import { useState } from 'react';
import type { TransactionsPageData } from '@/app/actions/transactions-page';
import Button from '@/app/components/Button/Button';
import Card from '@/app/components/Card/Card';
import Form from '@/app/components/Form/Form';
import Heading from '@/app/components/Heading/Heading';
import Input from '@/app/components/Input/Input';
import Select, { type SelectOption } from '@/app/components/Select/Select';
import {
  dateInputToISOEnd,
  dateInputToISOStart,
  toDateInputValue,
} from '@/app/libs/format';
import styles from './TransactionFilters.module.scss';

type TransactionFiltersProps = {
  accounts: TransactionsPageData['accounts'];
  categories: TransactionsPageData['categories'];
  filter: TransactionsPageData['filter'];
};

const TransactionFilters = ({
  accounts,
  categories,
  filter,
}: TransactionFiltersProps) => {
  const router = useRouter();
  const [fromDate, setFromDate] = useState(toDateInputValue(filter.from));
  const [toDate, setToDate] = useState(toDateInputValue(filter.to));
  const [accountId, setAccountId] = useState(
    filter.account_id ? String(filter.account_id) : ''
  );
  const [categoryId, setCategoryId] = useState(
    filter.category_id ? String(filter.category_id) : ''
  );

  const accountOptions: SelectOption[] = accounts.map((a) => ({
    value: String(a.id),
    label: a.name,
  }));

  const categoryOptions: SelectOption[] = categories.map((c) => ({
    value: String(c.id),
    label: `${c.name}（${c.type === 'income' ? '収入' : '支出'}）`,
  }));

  const applyFilters = (overrides?: {
    from?: string;
    to?: string;
    account_id?: string;
    category_id?: string;
  }) => {
    const params = new URLSearchParams();
    const from = overrides?.from ?? fromDate;
    const to = overrides?.to ?? toDate;
    const acc = overrides?.account_id ?? accountId;
    const cat = overrides?.category_id ?? categoryId;

    if (from) params.set('from', dateInputToISOStart(from));
    if (to) params.set('to', dateInputToISOEnd(to));
    if (acc) params.set('account_id', acc);
    if (cat) params.set('category_id', cat);

    const qs = params.toString();
    router.push(qs ? `/transactions?${qs}` : '/transactions');
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    applyFilters();
  };

  const handleReset = () => {
    router.push('/transactions');
  };

  return (
    <Card>
      <Heading level={2} className={styles.title}>
        絞り込み
      </Heading>
      <Form onSubmit={handleSubmit} className={styles.form}>
        <div className={styles.dateRow}>
          <Input
            label="開始日"
            type="date"
            name="from"
            value={fromDate}
            onChange={(e) => setFromDate(e.target.value)}
          />
          <Input
            label="終了日"
            type="date"
            name="to"
            value={toDate}
            onChange={(e) => setToDate(e.target.value)}
          />
        </div>
        <Select
          label="口座"
          name="filter_account_id"
          value={accountId}
          options={accountOptions}
          onChange={(e) => setAccountId(e.target.value)}
          placeholder="すべて"
        />
        <Select
          label="カテゴリ"
          name="filter_category_id"
          value={categoryId}
          options={categoryOptions}
          onChange={(e) => setCategoryId(e.target.value)}
          placeholder="すべて"
        />
        <div className={styles.actions}>
          <Button type="submit" variant="primary">
            適用
          </Button>
          <Button type="button" variant="secondary" onClick={handleReset}>
            リセット
          </Button>
        </div>
      </Form>
    </Card>
  );
};

export default TransactionFilters;
