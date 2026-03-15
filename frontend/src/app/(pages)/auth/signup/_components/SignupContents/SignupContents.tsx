'use client';

import { useState } from 'react';
import Link from 'next/link';
import { useRouter } from 'next/navigation';
import { signup } from '@/app/actions/auth';
import Card from '@/app/components/Card/Card';
import Form from '@/app/components/Form/Form';
import Input from '@/app/components/Input/Input';
import Button from '@/app/components/Button/Button';
import ErrorBlock from '@/app/components/ErrorBlock/ErrorBlock';
import styles from './SignupContents.module.scss';

const SignupContents = () => {
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
      await signup(email, password);
      router.push('/');
      router.refresh();
    } catch (err) {
      setError(err instanceof Error ? err.message : '新規登録に失敗しました');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className={styles.wrapper}>
      <Card>
        <h1 className={styles.title}>新規登録</h1>
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
            autoComplete="new-password"
            disabled={loading}
          />
          <Button type="submit" disabled={loading} variant="primary">
            {loading ? '登録中...' : '登録する'}
          </Button>
        </Form>
        <p className={styles.footer}>
          すでにアカウントをお持ちの方は <Link href="/auth/login">ログイン</Link>
        </p>
      </Card>
    </div>
  );
};

export default SignupContents;
