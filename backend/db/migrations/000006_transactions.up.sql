-- transactions: すべての金銭移動（Ledger の中心）
CREATE TABLE transactions (
  id serial PRIMARY KEY,                    -- 連番 PK
  user_id integer NOT NULL,                 -- 所有者（users.id）
  account_id integer NOT NULL,              -- 口座（accounts.id）
  category_id integer,                      -- カテゴリ（categories.id。任意）
  amount integer NOT NULL CHECK (amount != 0), -- 金額（収入: 正、支出: 負。0円禁止）
  memo text,                                -- メモ（任意）
  occurred_at timestamp NOT NULL,           -- 取引発生日時
  created_at timestamp DEFAULT now(),       -- 登録日時
  updated_at timestamp DEFAULT now(),       -- 更新日時（入力ミス修正等）
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (account_id) REFERENCES accounts(id),
  FOREIGN KEY (category_id) REFERENCES categories(id)
);

CREATE INDEX idx_transactions_user_date ON transactions(user_id, occurred_at);
CREATE INDEX idx_transactions_category ON transactions(category_id);
