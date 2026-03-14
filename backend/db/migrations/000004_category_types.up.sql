-- category_types: カテゴリ種別マスタ（タグのように選択する用）
CREATE TABLE category_types (
  id serial PRIMARY KEY,       -- 連番 PK
  name text UNIQUE NOT NULL,   -- 種別名: income | expense
  created_at timestamp DEFAULT now(),
  updated_at timestamp DEFAULT now()
);

-- 初期データ
INSERT INTO category_types (name) VALUES ('income'), ('expense');
