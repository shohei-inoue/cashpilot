-- categories: 支出・収入のカテゴリ
CREATE TABLE categories (
  id serial PRIMARY KEY,                       -- 連番 PK
  user_id integer NOT NULL,                    -- 所有者（users.id）
  category_type_id integer NOT NULL,           -- 種別（category_types.id: income | expense）
  name text NOT NULL,                          -- カテゴリ名（例: salary, rent, food）
  created_at timestamp DEFAULT now(),          -- 登録日時
  updated_at timestamp DEFAULT now(),          -- 更新日時
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (category_type_id) REFERENCES category_types(id)
);

CREATE INDEX idx_categories_user ON categories(user_id);
CREATE INDEX idx_categories_category_type ON categories(category_type_id);
