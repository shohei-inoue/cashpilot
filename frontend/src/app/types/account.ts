export type AccountType = 'cash' | 'bank' | 'credit';

export type Account = {
  id: number;
  type: AccountType;
  name: string;
  created_at?: string;
  updated_at?: string;
};

export type AccountsResponse = {
  accounts: Account[];
};
