import Heading from '@/app/components/Heading/Heading';
import AppShell from '@/app/components/AppShell/AppShell';
import MainContent from '@/app/components/MainContent/MainContent';

export default function Transactions() {
  return (
    <AppShell>
      <MainContent>
        <Heading level={1}>Transactions</Heading>
      </MainContent>
    </AppShell>
  );
}
