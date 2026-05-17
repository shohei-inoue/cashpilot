import { redirect } from 'next/navigation';
import { runSimulation } from '@/app/actions/simulation';
import AppShell from '@/app/components/AppShell/AppShell';
import { DEFAULT_SIMULATION_INPUT } from '@/app/constants/simulation';
import SimulationContents from './_components/SimulationContents/SimulationContents';

export default async function SimulationPage() {
  let initialResult = null;
  let initialError: string | null = null;

  try {
    initialResult = await runSimulation(DEFAULT_SIMULATION_INPUT);
  } catch (err) {
    if (err instanceof Error && err.message === 'Unauthorized') {
      redirect('/auth/login');
    }
    initialError =
      err instanceof Error ? err.message : 'シミュレーションの初期読み込みに失敗しました';
  }

  if (!initialResult && !initialError) {
    redirect('/auth/login');
  }

  return (
    <AppShell>
      <SimulationContents initialResult={initialResult} initialError={initialError} />
    </AppShell>
  );
}
