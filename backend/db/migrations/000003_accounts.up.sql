-- accounts: お金の「場所」（銀行・クレカ・現金）
CREATE TABLE accounts (
  id serial PRIMARY KEY,                       -- 連番 PK
  user_id integer NOT NULL,                    -- 所有者（users.id）
  account_type_id integer NOT NULL,            -- 種別（account_types.id）
  name text NOT NULL,                          -- 口座名（例: 三井住友銀行、楽天カード）
  created_at timestamp DEFAULT now(),          -- 登録日時
  updated_at timestamp DEFAULT now(),          -- 更新日時
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (account_type_id) REFERENCES account_types(id)
);

CREATE INDEX idx_accounts_user ON accounts(user_id);
CREATE INDEX idx_accounts_account_type ON accounts(account_type_id);
