# CashPilot インフラ

ローカル開発環境をまとめる。デプロイ環境は実装が進んでから検討する（AWS / Fly.io 等）。

---

## ローカル開発環境

Docker Compose で一括起動。

```
frontend (Next.js)  → localhost:3010
backend (Go API)    → localhost:8080
db (PostgreSQL)     → localhost:5432
migrate             → 起動時にマイグレーション自動適用
```

```bash
./start.sh
# または
docker compose up --build
```

`.env.example` をコピーして `.env` を作成（`start.sh` が自動作成する）。

---

## デプロイ（将来検討）

実装が進んだ段階で、以下のいずれかを検討する。

- **フロントエンド**: Vercel、Netlify 等
- **バックエンド**: AWS（ECS Fargate）、Fly.io、Render 等
- **データベース**: RDS、Supabase、Fly Postgres 等

---

## 参照

- [migration.md](./migration.md) - DB マイグレーション手順
- [database.md](./database.md) - DB 設計
