'use client';

import { useRouter } from 'next/navigation';
import { useEffect, useState } from 'react';
import {
  createGoal,
  deleteGoal,
  updateGoal,
} from '@/app/actions/goals';
import type { Goal } from '@/app/types/goal';
import Amount from '@/app/components/Amount/Amount';
import Button from '@/app/components/Button/Button';
import Card from '@/app/components/Card/Card';
import EmptyState from '@/app/components/EmptyState/EmptyState';
import ErrorBlock from '@/app/components/ErrorBlock/ErrorBlock';
import Form from '@/app/components/Form/Form';
import Heading from '@/app/components/Heading/Heading';
import Input from '@/app/components/Input/Input';
import { formatDateOnly } from '@/app/libs/format';
import styles from './GoalSection.module.scss';

type GoalSectionProps = {
  goals: Goal[];
};

const GoalSection = ({ goals }: GoalSectionProps) => {
  const router = useRouter();
  const [editingId, setEditingId] = useState<number | null>(null);
  const [name, setName] = useState('');
  const [targetAmount, setTargetAmount] = useState('');
  const [deadline, setDeadline] = useState('');
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);

  const resetForm = () => {
    setEditingId(null);
    setName('');
    setTargetAmount('');
    setDeadline('');
    setError(null);
  };

  const startEdit = (goal: Goal) => {
    setEditingId(goal.id);
    setName(goal.name);
    setTargetAmount(String(goal.target_amount));
    setDeadline(goal.deadline ?? '');
    setError(null);
  };

  useEffect(() => {
    if (editingId === null) return;
    const stillExists = goals.some((g) => g.id === editingId);
    if (!stillExists) resetForm();
  }, [editingId, goals]);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError(null);

    const amount = parseInt(targetAmount, 10);
    if (!name.trim()) {
      setError('目標名を入力してください');
      return;
    }
    if (!amount || Number.isNaN(amount) || amount <= 0) {
      setError('目標金額を正の数で入力してください');
      return;
    }

    const input = {
      name: name.trim(),
      target_amount: amount,
      deadline: deadline || undefined,
    };

    setLoading(true);
    try {
      if (editingId !== null) {
        await updateGoal(editingId, input);
      } else {
        await createGoal(input);
      }
      resetForm();
      router.refresh();
    } catch (err) {
      setError(err instanceof Error ? err.message : '保存に失敗しました');
    } finally {
      setLoading(false);
    }
  };

  const handleDelete = async (goal: Goal) => {
    if (!window.confirm(`「${goal.name}」を削除しますか？`)) return;
    setLoading(true);
    setError(null);
    try {
      await deleteGoal(goal.id);
      if (editingId === goal.id) resetForm();
      router.refresh();
    } catch (err) {
      setError(err instanceof Error ? err.message : '削除に失敗しました');
    } finally {
      setLoading(false);
    }
  };

  return (
    <section className={styles.section}>
      <Card>
        <Heading level={2} className={styles.title}>
          目標
        </Heading>
        <p className={styles.description}>
          資金目標（引っ越し・購入など）を登録します。シミュレーションで達成見込みを確認できます。
        </p>

        {goals.length === 0 ? (
          <EmptyState message="目標がまだありません。下のフォームから追加してください。" />
        ) : (
          <ul className={styles.list}>
            {goals.map((goal) => (
              <li key={goal.id} className={styles.listItem}>
                <div className={styles.itemMain}>
                  <span className={styles.itemName}>{goal.name}</span>
                  <div className={styles.itemMeta}>
                    <Amount amount={goal.target_amount} variant="income" />
                    {goal.deadline && (
                      <span className={styles.deadline}>
                        期限: {formatDateOnly(goal.deadline)}
                      </span>
                    )}
                  </div>
                </div>
                <div className={styles.itemActions}>
                  <Button
                    type="button"
                    variant="secondary"
                    disabled={loading}
                    onClick={() => startEdit(goal)}
                  >
                    編集
                  </Button>
                  <Button
                    type="button"
                    variant="danger"
                    disabled={loading}
                    onClick={() => handleDelete(goal)}
                  >
                    削除
                  </Button>
                </div>
              </li>
            ))}
          </ul>
        )}

        <Form onSubmit={handleSubmit} className={styles.form}>
          <Heading level={3} className={styles.formTitle}>
            {editingId !== null ? '目標を編集' : '目標を追加'}
          </Heading>
          {error && <ErrorBlock>{error}</ErrorBlock>}
          <Input
            label="目標名"
            name="name"
            value={name}
            onChange={(e) => setName(e.target.value)}
            required
            disabled={loading}
            placeholder="例: マイホーム購入"
          />
          <Input
            label="目標金額"
            type="number"
            name="target_amount"
            value={targetAmount}
            onChange={(e) => setTargetAmount(e.target.value)}
            required
            disabled={loading}
            placeholder="300000"
          />
          <Input
            label="期限（任意）"
            type="date"
            name="deadline"
            value={deadline}
            onChange={(e) => setDeadline(e.target.value)}
            disabled={loading}
          />
          <div className={styles.formActions}>
            <Button
              type="submit"
              variant="primary"
              disabled={loading || !name.trim() || !targetAmount}
            >
              {loading ? '保存中...' : editingId !== null ? '更新する' : '追加する'}
            </Button>
            {editingId !== null && (
              <Button type="button" variant="secondary" disabled={loading} onClick={resetForm}>
                キャンセル
              </Button>
            )}
          </div>
        </Form>
      </Card>
    </section>
  );
};

export default GoalSection;
