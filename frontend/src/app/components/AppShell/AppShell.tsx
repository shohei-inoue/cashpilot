import { redirect } from 'next/navigation';
import { getCurrentUser } from '@/app/actions/user';
import MainContainer from '@/app/components/MainContainer/MainContainer';

type AppShellProps = {
  children: React.ReactNode;
};

/** 認証済みページ用レイアウト（ユーザー情報を Header に渡す） */
const AppShell = async ({ children }: AppShellProps) => {
  const user = await getCurrentUser();
  if (!user) {
    redirect('/auth/login');
  }

  return <MainContainer userEmail={user.email}>{children}</MainContainer>;
};

export default AppShell;
