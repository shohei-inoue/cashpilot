import { redirect } from 'next/navigation';
import { getTransactionsPageData } from '@/app/actions/transactions-page';
import AppShell from '@/app/components/AppShell/AppShell';
import TransactionsContents from './_components/TransactionsContents/TransactionsContents';

type PageProps = {
  searchParams: Promise<{
    from?: string;
    to?: string;
    account_id?: string;
    category_id?: string;
  }>;
};

export default async function TransactionsPage({ searchParams }: PageProps) {
  const params = await searchParams;
  const data = await getTransactionsPageData(params);
  if (!data) {
    redirect('/auth/login');
  }

  return (
    <AppShell>
      <TransactionsContents data={data} />
    </AppShell>
  );
}
