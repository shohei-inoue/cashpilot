'use client';

import { useRouter } from 'next/navigation';
import { useState } from 'react';
import Button from '@/app/components/Button/Button';
import Card from '@/app/components/Card/Card';
import EmptyState from '@/app/components/EmptyState/EmptyState';
import ErrorBlock from '@/app/components/ErrorBlock/ErrorBlock';
import Form from '@/app/components/Form/Form';
import Heading from '@/app/components/Heading/Heading';
import Input from '@/app/components/Input/Input';
import Select, { type SelectOption } from '@/app/components/Select/Select';
import styles from './ResourceSection.module.scss';

export type ResourceItem = {
  id: number;
  name: string;
  type: string;
};

type ResourceSectionProps = {
  title: string;
  description: string;
  items: ResourceItem[];
  typeOptions: readonly SelectOption[];
  getTypeLabel: (type: string) => string;
  emptyMessage: string;
  nameLabel: string;
  typeLabel: string;
  createLabel: string;
  onCreate: (type: string, name: string) => Promise<unknown>;
  onUpdate: (id: number, type: string, name: string) => Promise<unknown>;
  onDelete: (id: number) => Promise<unknown>;
};

const ResourceSection = ({
  title,
  description,
  items,
  typeOptions,
  getTypeLabel,
  emptyMessage,
  nameLabel,
  typeLabel,
  createLabel,
  onCreate,
  onUpdate,
  onDelete,
}: ResourceSectionProps) => {
  const router = useRouter();
  const [editingId, setEditingId] = useState<number | null>(null);
  const [type, setType] = useState(typeOptions[0]?.value ?? '');
  const [name, setName] = useState('');
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);

  const resetForm = () => {
    setEditingId(null);
    setType(typeOptions[0]?.value ?? '');
    setName('');
    setError(null);
  };

  const startEdit = (item: ResourceItem) => {
    setEditingId(item.id);
    setType(item.type);
    setName(item.name);
    setError(null);
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError(null);
    setLoading(true);
    try {
      if (editingId !== null) {
        await onUpdate(editingId, type, name.trim());
      } else {
        await onCreate(type, name.trim());
      }
      resetForm();
      router.refresh();
    } catch (err) {
      setError(err instanceof Error ? err.message : '保存に失敗しました');
    } finally {
      setLoading(false);
    }
  };

  const handleDelete = async (item: ResourceItem) => {
    if (!window.confirm(`「${item.name}」を削除しますか？`)) return;
    setError(null);
    setLoading(true);
    try {
      await onDelete(item.id);
      if (editingId === item.id) resetForm();
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
          {title}
        </Heading>
        <p className={styles.description}>{description}</p>

        {items.length === 0 ? (
          <EmptyState message={emptyMessage} />
        ) : (
          <ul className={styles.list}>
            {items.map((item) => (
              <li key={item.id} className={styles.listItem}>
                <div className={styles.itemMain}>
                  <span className={styles.itemName}>{item.name}</span>
                  <span className={styles.typeBadge}>{getTypeLabel(item.type)}</span>
                </div>
                <div className={styles.itemActions}>
                  <Button
                    type="button"
                    variant="secondary"
                    disabled={loading}
                    onClick={() => startEdit(item)}
                  >
                    編集
                  </Button>
                  <Button
                    type="button"
                    variant="danger"
                    disabled={loading}
                    onClick={() => handleDelete(item)}
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
            {editingId !== null ? '編集' : createLabel}
          </Heading>
          {error && <ErrorBlock>{error}</ErrorBlock>}
          <Select
            label={typeLabel}
            name="type"
            value={type}
            options={typeOptions}
            onChange={(e) => setType(e.target.value)}
            required
            disabled={loading}
          />
          <Input
            label={nameLabel}
            name="name"
            value={name}
            onChange={(e) => setName(e.target.value)}
            required
            disabled={loading}
            placeholder="名前を入力"
          />
          <div className={styles.formActions}>
            <Button type="submit" variant="primary" disabled={loading || !name.trim()}>
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

export default ResourceSection;
