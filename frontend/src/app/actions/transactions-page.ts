"use server";

import { getAccounts } from "./accounts";
import { getCategories } from "./categories";
import { getTransactions, type TransactionFilter } from "./transactions";
import { getCurrentMonthRange } from "../libs/format";
import type { Account } from "../types/account";
import type { Category } from "../types/category";
import type { Transaction } from "../types/transaction";

export type TransactionsPageData = {
  accounts: Account[];
  categories: Category[];
  transactions: Transaction[];
  filter: {
    from: string;
    to: string;
    account_id?: number;
    category_id?: number;
  };
};

export type TransactionsPageSearchParams = {
  from?: string;
  to?: string;
  account_id?: string;
  category_id?: string;
};

function parseOptionalId(value?: string): number | undefined {
  if (!value) return undefined;
  const n = Number(value);
  return Number.isInteger(n) && n > 0 ? n : undefined;
}

export async function getTransactionsPageData(
  searchParams: TransactionsPageSearchParams = {}
): Promise<TransactionsPageData | null> {
  const monthRange = getCurrentMonthRange();
  const from = searchParams.from || monthRange.from;
  const to = searchParams.to || monthRange.to;
  const account_id = parseOptionalId(searchParams.account_id);
  const category_id = parseOptionalId(searchParams.category_id);

  const filter: TransactionFilter = { from, to, account_id, category_id, limit: 50 };

  try {
    const [accounts, categories, transactions] = await Promise.all([
      getAccounts(),
      getCategories(),
      getTransactions(filter),
    ]);
    return {
      accounts,
      categories,
      transactions,
      filter: { from, to, account_id, category_id },
    };
  } catch {
    return null;
  }
}
