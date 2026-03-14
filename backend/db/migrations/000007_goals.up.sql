-- goals: 資金目標（引越し・購入・返済など）
CREATE TABLE goals (
  id serial PRIMARY KEY,                    -- 連番 PK
  user_id integer NOT NULL,                 -- 所有者（users.id）
  name text NOT NULL,                       -- 目標名（例: 引越し、MacBook）
  target_amount integer NOT NULL,           -- 目標金額（円）
  deadline date,                            -- 目標期限（任意）
  created_at timestamp DEFAULT now(),       -- 登録日時
  updated_at timestamp DEFAULT now(),       -- 更新日時
  FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE INDEX idx_goals_user ON goals(user_id);
