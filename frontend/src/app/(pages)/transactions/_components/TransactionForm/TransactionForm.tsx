'use client';

import { useRouter } from 'next/navigation';
import { useEffect, useMemo, useState } from 'react';
import {
  createTransaction,
  updateTransaction,
  type TransactionInput,
} from '@/app/actions/transactions';
import type { Account } from '@/app/types/account';
import type { Category } from '@/app/types/category';
import type { Transaction } from '@/app/types/transaction';
import Button from '@/app/components/Button/Button';
import Card from '@/app/components/Card/Card';
import ErrorBlock from '@/app/components/ErrorBlock/ErrorBlock';
import Form from '@/app/components/Form/Form';
import Heading from '@/app/components/Heading/Heading';
import Input from '@/app/components/Input/Input';
import Select, { type SelectOption } from '@/app/components/Select/Select';
import {
  datetimeLocalToISO,
  toDatetimeLocalValue,
} from '@/app/libs/format';
import styles from './TransactionForm.module.scss';

type FlowType = 'income' | 'expense';

type TransactionFormProps = {
  accounts: Account[];
  categories: Category[];
  editingTransaction?: Transaction | null;
  onCancelEdit?: () => void;
};

function buildAmount(flowType: FlowType, amountStr: string): number {
  const n = Math.abs(parseInt(amountStr, 10));
  if (!n || Number.isNaN(n)) return 0;
  return flowType === 'expense' ? -n : n;
}

function flowTypeFromAmount(amount: number): FlowType {
  return amount < 0 ? 'expense' : 'income';
}

const TransactionForm = ({
  accounts,
  categories,
  editingTransaction,
  onCancelEdit,
}: TransactionFormProps) => {
  const router = useRouter();
  const isEditing = editingTransaction != null;

  const [flowType, setFlowType] = useState<FlowType>('expense');
  const [amount, setAmount] = useState('');
  const [accountId, setAccountId] = useState('');
  const [categoryId, setCategoryId] = useState('');
  const [occurredAt, setOccurredAt] = useState(toDatetimeLocalValue());
  const [memo, setMemo] = useState('');
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);

  const accountOptions: SelectOption[] = useMemo(
    () => accounts.map((a) => ({ value: String(a.id), label: a.name })),
    [accounts]
  );

  const categoryOptions: SelectOption[] = useMemo(
    () =>
      categories
        .filter((c) => c.type === flowType)
        .map((c) => ({ value: String(c.id), label: c.name })),
    [categories, flowType]
  );

  useEffect(() => {
    if (!editingTransaction) {
      setFlowType('expense');
      setAmount('');
      setAccountId(accounts[0] ? String(accounts[0].id) : '');
      setCategoryId('');
      setOccurredAt(toDatetimeLocalValue());
      setMemo('');
      setError(null);
      return;
    }
    setFlowType(flowTypeFromAmount(editingTransaction.amount));
    setAmount(String(Math.abs(editingTransaction.amount)));
    setAccountId(String(editingTransaction.account_id));
    setCategoryId(
      editingTransaction.category_id ? String(editingTransaction.category_id) : ''
    );
    setOccurredAt(toDatetimeLocalValue(new Date(editingTransaction.occurred_at)));
    setMemo(editingTransaction.memo ?? '');
    setError(null);
  }, [editingTransaction, accounts]);

  useEffect(() => {
    if (!accountId && accounts.length > 0) {
      setAccountId(String(accounts[0].id));
    }
  }, [accounts, accountId]);

  useEffect(() => {
    if (categoryId && !categoryOptions.some((o) => o.value === categoryId)) {
      setCategoryId('');
    }
  }, [categoryId, categoryOptions]);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError(null);

    const parsedAmount = buildAmount(flowType, amount);
    if (parsedAmount === 0) {
      setError('金額を入力してください');
      return;
    }
    if (!accountId) {
      setError('口座を選択してください');
      return;
    }

    const input: TransactionInput = {
      account_id: Number(accountId),
      category_id: categoryId ? Number(categoryId) : undefined,
      amount: parsedAmount,
      memo: memo.trim() || undefined,
      occurred_at: datetimeLocalToISO(occurredAt),
    };

    setLoading(true);
    try {
      if (isEditing && editingTransaction) {
        await updateTransaction(editingTransaction.id, input);
        onCancelEdit?.();
      } else {
        await createTransaction(input);
        setAmount('');
        setCategoryId('');
        setMemo('');
        setOccurredAt(toDatetimeLocalValue());
      }
      router.refresh();
    } catch (err) {
      setError(err instanceof Error ? err.message : '保存に失敗しました');
    } finally {
      setLoading(false);
    }
  };

  if (accounts.length === 0) {
    return (
      <Card>
        <Heading level={2} className={styles.title}>
          収支を登録
        </Heading>
        <p className={styles.hint}>
          取引を登録するには、先に設定画面で口座を追加してください。
        </p>
      </Card>
    );
  }

  return (
    <Card>
      <Heading level={2} className={styles.title}>
        {isEditing ? '取引を編集' : '収支を登録'}
      </Heading>
      <Form onSubmit={handleSubmit} className={styles.form}>
        {error && <ErrorBlock>{error}</ErrorBlock>}

        <div className={styles.flowToggle} role="group" aria-label="収支の種類">
          <button
            type="button"
            className={`${styles.flowButton} ${flowType === 'expense' ? styles.flowActive : ''}`}
            onClick={() => setFlowType('expense')}
            disabled={loading}
          >
            支出
          </button>
          <button
            type="button"
            className={`${styles.flowButton} ${flowType === 'income' ? styles.flowActiveIncome : ''}`}
            onClick={() => setFlowType('income')}
            disabled={loading}
          >
            収入
          </button>
        </div>

        <Input
          label="金額"
          type="number"
          name="amount"
          value={amount}
          onChange={(e) => setAmount(e.target.value)}
          required
          disabled={loading}
          placeholder="1000"
        />

        <Select
          label="口座"
          name="account_id"
          value={accountId}
          options={accountOptions}
          onChange={(e) => setAccountId(e.target.value)}
          required
          disabled={loading}
        />

        <Select
          label="カテゴリ"
          name="category_id"
          value={categoryId}
          options={categoryOptions}
          onChange={(e) => setCategoryId(e.target.value)}
          disabled={loading || categoryOptions.length === 0}
          placeholder="（未選択）"
        />

        <Input
          label="発生日時"
          type="datetime-local"
          name="occurred_at"
          value={occurredAt}
          onChange={(e) => setOccurredAt(e.target.value)}
          required
          disabled={loading}
        />

        <Input
          label="メモ"
          name="memo"
          value={memo}
          onChange={(e) => setMemo(e.target.value)}
          disabled={loading}
          placeholder="任意"
        />

        <div className={styles.actions}>
          <Button type="submit" variant="primary" disabled={loading || !amount}>
            {loading ? '保存中...' : isEditing ? '更新する' : '登録する'}
          </Button>
          {isEditing && onCancelEdit && (
            <Button type="button" variant="secondary" disabled={loading} onClick={onCancelEdit}>
              キャンセル
            </Button>
          )}
        </div>
      </Form>
    </Card>
  );
};

export default TransactionForm;
