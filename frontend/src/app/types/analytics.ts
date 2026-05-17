export type TransactionSummary = {
  total_income: number;
  total_expense: number;
  net_cashflow: number;
  transaction_count: number;
};

export type TransactionSummaryResponse = {
  summary: TransactionSummary;
};

export type CashflowByPeriod = {
  period: string;
  total_income: number;
  total_expense: number;
  net_cashflow: number;
  transaction_count: number;
};

export type CashflowResponse = {
  items: CashflowByPeriod[];
  group_by: string;
};
