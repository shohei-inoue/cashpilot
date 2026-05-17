'use client';

import {
  createAccount,
  deleteAccount,
  updateAccount,
} from '@/app/actions/accounts';
import {
  createCategory,
  deleteCategory,
  updateCategory,
} from '@/app/actions/categories';
import type { SettingsData } from '@/app/actions/settings';
import {
  ACCOUNT_TYPE_OPTIONS,
  CATEGORY_TYPE_OPTIONS,
  getAccountTypeLabel,
  getCategoryTypeLabel,
} from '@/app/constants/settings';
import Heading from '@/app/components/Heading/Heading';
import MainContent from '@/app/components/MainContent/MainContent';
import GoalSection from '../GoalSection/GoalSection';
import ResourceSection from '../ResourceSection/ResourceSection';
import styles from './SettingsContents.module.scss';

type SettingsContentsProps = {
  data: SettingsData;
};

const SettingsContents = ({ data }: SettingsContentsProps) => {
  return (
    <MainContent>
      <header className={styles.header}>
        <Heading level={1}>設定</Heading>
        <p className={styles.lead}>
          口座・カテゴリ・目標を管理します。取引登録やシミュレーションで使用します。
        </p>
      </header>

      <ResourceSection
        title="口座"
        description="お金の置き場所（銀行・クレジットカード・現金など）を登録します。"
        items={data.accounts}
        typeOptions={ACCOUNT_TYPE_OPTIONS}
        getTypeLabel={getAccountTypeLabel}
        emptyMessage="口座がまだありません。下のフォームから追加してください。"
        nameLabel="口座名"
        typeLabel="種別"
        createLabel="口座を追加"
        onCreate={createAccount}
        onUpdate={updateAccount}
        onDelete={deleteAccount}
      />

      <ResourceSection
        title="カテゴリ"
        description="収入・支出の分類を登録します。"
        items={data.categories}
        typeOptions={CATEGORY_TYPE_OPTIONS}
        getTypeLabel={getCategoryTypeLabel}
        emptyMessage="カテゴリがまだありません。下のフォームから追加してください。"
        nameLabel="カテゴリ名"
        typeLabel="種別"
        createLabel="カテゴリを追加"
        onCreate={createCategory}
        onUpdate={updateCategory}
        onDelete={deleteCategory}
      />

      <GoalSection goals={data.goals} />
    </MainContent>
  );
};

export default SettingsContents;
