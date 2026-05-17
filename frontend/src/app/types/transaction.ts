export type Transaction = {
  id: number;
  account_id: number;
  account_name?: string;
  category_id?: number;
  category_name?: string;
  amount: number;
  memo?: string;
  occurred_at: string;
};

export type TransactionsResponse = {
  transactions: Transaction[];
};
