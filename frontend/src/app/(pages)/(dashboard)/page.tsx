import { redirect } from 'next/navigation';
import { getDashboardData } from '@/app/actions/dashboard';
import AppShell from '@/app/components/AppShell/AppShell';
import DashboardContents from './_components/DashboardContents/DashboardContents';
import DashboardErrorState from './_components/DashboardErrorState/DashboardErrorState';

export default async function DashboardPage() {
  const result = await getDashboardData();

  if (result.status === 'unauthorized') {
    redirect('/auth/login');
  }

  return (
    <AppShell>
      {result.status === 'ok' ? (
        <DashboardContents data={result.data} />
      ) : (
        <DashboardErrorState message={result.message} />
      )}
    </AppShell>
  );
}
