import { redirect } from 'next/navigation';
import { getDashboardData } from '@/app/actions/dashboard';
import AppShell from '@/app/components/AppShell/AppShell';
import DashboardContents from './_components/DashboardContents/DashboardContents';

export default async function DashboardPage() {
  const data = await getDashboardData();
  if (!data) {
    redirect('/auth/login');
  }

  return (
    <AppShell>
      <DashboardContents data={data} />
    </AppShell>
  );
}
