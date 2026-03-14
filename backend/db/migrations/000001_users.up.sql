-- users: 認証主体
CREATE TABLE users (
  id serial PRIMARY KEY,                              -- 連番 PK（内部参照用）
  uuid uuid UNIQUE NOT NULL DEFAULT gen_random_uuid(), -- Supabase Auth 連携用（auth.users.id と紐付ける）
  email text UNIQUE NOT NULL,                         -- メールアドレス
  created_at timestamp DEFAULT now()                  -- 登録日時
);
