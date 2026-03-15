'use client';

import { useState } from 'react';
import Link from 'next/link';
import { useRouter } from 'next/navigation';
import { login } from '@/app/actions/auth';
import Card from '@/app/components/Card/Card';
import Form from '@/app/components/Form/Form';
import Input from '@/app/components/Input/Input';
import Button from '@/app/components/Button/Button';
import ErrorBlock from '@/app/components/ErrorBlock/ErrorBlock';
import styles from './LoginContents.module.scss';

const LoginContents = () => {
  const router = useRouter();
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError(null);
    setLoading(true);
    try {
      await login(email, password);
      router.push('/');
      router.refresh();
    } catch (err) {
      setError(err instanceof Error ? err.message : 'ログインに失敗しました');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className={styles.wrapper}>
      <Card>
        <h1 className={styles.title}>ログイン</h1>
        <Form onSubmit={handleSubmit}>
          {error && <ErrorBlock>{error}</ErrorBlock>}
          <Input
            label="メールアドレス"
            type="email"
            name="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            required
            autoComplete="email"
            disabled={loading}
          />
          <Input
            label="パスワード"
            type="password"
            name="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
            autoComplete="current-password"
            disabled={loading}
          />
          <Button
            type="submit"
            disabled={loading}
            variant="primary"
          >
            {loading ? 'ログイン中...' : 'ログイン'}
          </Button>
        </Form>
        <p className={styles.footer}>
          アカウントをお持ちでない方は <Link href="/auth/signup">新規登録</Link>
        </p>
      </Card>
    </div>
  );
};

export default LoginContents;
