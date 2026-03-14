-- users にパスワードハッシュを追加（JWT + Cookie 認証用）
ALTER TABLE users ADD COLUMN password_hash text;
