-- account_types: 口座種別マスタ（タグのように選択する用）
CREATE TABLE account_types (
  id serial PRIMARY KEY,       -- 連番 PK
  name text UNIQUE NOT NULL,   -- 種別名: cash | bank | credit
  created_at timestamp DEFAULT now(),
  updated_at timestamp DEFAULT now()
);

-- 初期データ
INSERT INTO account_types (name) VALUES ('cash'), ('bank'), ('credit');
