# 認証設計（JWT + httpOnly Cookie）

メール＋パスワード認証を JWT と httpOnly Cookie で実現する。

---

## 1. 全体フロー

```
サインアップ: POST /api/auth/signup (email, password)
    → users に登録
    → JWT 発行 → Cookie にセット
    → 201 + user

ログイン: POST /api/auth/login (email, password)
    → パスワード検証
    → JWT 発行 → Cookie にセット
    → 200 + user

API 呼び出し: GET /api/accounts 等
    → リクエストに Cookie が自動付与
    → ミドルウェアで JWT 検証 → user_id をコンテキストに格納
    → ハンドラーで user_id を利用

ログアウト: POST /api/auth/logout
    → Cookie を削除（空文字・過去の expiry で上書き）
    → 200
```

---

## 2. API エンドポイント

| メソッド | パス | 認証 | 説明 |
|----------|------|------|------|
| POST | `/api/auth/signup` | 不要 | サインアップ。成功時 Cookie セット。 |
| POST | `/api/auth/login` | 不要 | ログイン。成功時 Cookie セット。 |
| POST | `/api/auth/logout` | 不要 | ログアウト。Cookie 削除。 |
| GET | `/api/user` | 必須 | 認証済みユーザー情報を返却。 |

### 2-1. POST /api/auth/signup

**Request**
```json
{
  "email": "user@example.com",
  "password": "min8chars"
}
```

**Response 201**
```json
{
  "id": 1,
  "email": "user@example.com",
  "created_at": "2025-03-13T12:00:00Z"
}
```
+ `Set-Cookie`: `token=...` を付与

**エラー**
- 400: バリデーションエラー（email 形式、password 8文字以上）
- 409: メールアドレス重複

### 2-2. POST /api/auth/login

**Request**
```json
{
  "email": "user@example.com",
  "password": "password"
}
```

**Response 200**
```json
{
  "id": 1,
  "email": "user@example.com",
  "created_at": "2025-03-13T12:00:00Z"
}
```
+ `Set-Cookie`: `token=...` を付与

**エラー**
- 400: バリデーションエラー
- 401: メールまたはパスワードが不正

### 2-3. POST /api/auth/logout

**Response 200**
+ `Set-Cookie`: `token=; Max-Age=0` で Cookie 削除

### 2-4. GET /api/user

認証必須。Cookie の JWT を検証し、ユーザー情報を返す。

**Response 200**
```json
{
  "id": 1,
  "email": "user@example.com",
  "created_at": "2025-03-13T12:00:00Z"
}
```

**エラー**
- 401: トークンなし or 無効 or 期限切れ

---

## 3. JWT 仕様

### ペイロード（Claims）

| キー | 型 | 説明 |
|------|-----|------|
| sub | string | ユーザー ID（users.id の文字列） |
| email | string | メールアドレス |
| exp | number | 有効期限（Unix timestamp） |
| iat | number | 発行日時（Unix timestamp） |

### 有効期限

- **Access Token**: 24 時間（86400 秒）
- リフレッシュトークンは MVP では未実装。期限切れ時は再ログイン。

### 署名

- アルゴリズム: HS256
- 秘密鍵: 環境変数 `JWT_SECRET`（32 文字以上推奨）

---

## 4. Cookie 仕様

| 項目 | 値 |
|------|-----|
| 名前 | `token` |
| 値 | JWT 文字列 |
| Path | `/` |
| HttpOnly | `true` |
| Secure | 本番時 `true`（HTTPS のみ） |
| SameSite | `Strict` |
| Max-Age | 86400（24 時間、JWT の exp に合わせる） |

### ドメイン・CORS

- ローカル開発: FE `localhost:3010`、BE `localhost:8080` の場合、異なるオリジンなので Cookie はデフォルトで送信されない。
- **対応**: API 呼び出し時に `credentials: 'include'` を指定し、BE で `Access-Control-Allow-Credentials: true` と `Access-Control-Allow-Origin` をフロントのオリジンに固定する（`*` は不可）。

---

## 5. パスワード

- **ハッシュ**: bcrypt（cost 10〜12）
- **保存**: `users.password_hash`（カラム追加が必要）
- **バリデーション**: 8 文字以上

---

## 6. DB 変更

`users` テーブルに `password_hash` を追加するマイグレーションが必要。

```sql
-- 000008_add_password_hash.up.sql
ALTER TABLE users ADD COLUMN password_hash text;
```

- 既存ユーザーがいない前提なら `NOT NULL` も付与可能
- サインアップ時は必ず `password_hash` を設定する

---

## 7. セキュリティ

| 対策 | 内容 |
|------|------|
| XSS | `HttpOnly` で JS から Cookie を参照不可 |
| CSRF | `SameSite=Strict` で他サイトからのリクエストで Cookie 送信しない |
| パスワード | bcrypt でハッシュ保存。平文は保存しない。 |
| JWT 秘密鍵 | `JWT_SECRET` を環境変数で管理。本番では強力な乱数を使用。 |

---

## 8. フロントエンド側の実装方針

- `fetch` / Axios で API 呼び出し時に `credentials: 'include'` を指定
- ログイン・サインアップ成功後は Cookie が自動で付与されるため、トークンを JS で保持する必要はない
- 401 時はログイン画面へリダイレクト

---

## 9. 参照

- [backend.md](./backend.md) - API 設計
- [database.md](./database.md) - users テーブル
