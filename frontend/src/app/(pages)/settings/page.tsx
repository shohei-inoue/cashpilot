import { redirect } from 'next/navigation';
import { getSettingsData } from '@/app/actions/settings';
import AppShell from '@/app/components/AppShell/AppShell';
import SettingsContents from './_components/SettingsContents/SettingsContents';

export default async function SettingsPage() {
  const data = await getSettingsData();
  if (!data) {
    redirect('/auth/login');
  }

  return (
    <AppShell>
      <SettingsContents data={data} />
    </AppShell>
  );
}
