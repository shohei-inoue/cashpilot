"use server";

import { revalidatePath } from "next/cache";
import { createAuthClient } from "../libs/client";
import { handleResponse } from "../libs/response";
import type { Transaction, TransactionsResponse } from "../types/transaction";

const REVALIDATE_PATHS = ["/transactions", "/", "/settings"];

async function revalidateTransactionPaths() {
  for (const path of REVALIDATE_PATHS) {
    revalidatePath(path);
  }
}

export type TransactionFilter = {
  from?: string;
  to?: string;
  account_id?: number;
  category_id?: number;
  limit?: number;
};

export type TransactionInput = {
  account_id: number;
  category_id?: number;
  amount: number;
  memo?: string;
  occurred_at: string;
};

export async function getTransactions(
  filter: TransactionFilter = {}
): Promise<Transaction[]> {
  const client = await createAuthClient();
  const params: Record<string, string | number> = { limit: filter.limit ?? 50 };
  if (filter.from) params.from = filter.from;
  if (filter.to) params.to = filter.to;
  if (filter.account_id) params.account_id = filter.account_id;
  if (filter.category_id) params.category_id = filter.category_id;

  const res = await handleResponse(
    client.get<TransactionsResponse>("/api/transactions", { params })
  );
  return res.transactions ?? [];
}

export async function createTransaction(input: TransactionInput): Promise<Transaction> {
  const client = await createAuthClient();
  const body = {
    account_id: input.account_id,
    category_id: input.category_id ?? undefined,
    amount: input.amount,
    memo: input.memo || undefined,
    occurred_at: input.occurred_at,
  };
  const tx = await handleResponse(client.post<Transaction>("/api/transactions", body));
  await revalidateTransactionPaths();
  return tx;
}

export async function updateTransaction(
  id: number,
  input: TransactionInput
): Promise<Transaction> {
  const client = await createAuthClient();
  const body = {
    account_id: input.account_id,
    category_id: input.category_id ?? undefined,
    amount: input.amount,
    memo: input.memo || undefined,
    occurred_at: input.occurred_at,
  };
  const tx = await handleResponse(
    client.put<Transaction>(`/api/transactions/${id}`, body)
  );
  await revalidateTransactionPaths();
  return tx;
}

export async function deleteTransaction(id: number): Promise<void> {
  const client = await createAuthClient();
  await client.delete(`/api/transactions/${id}`);
  await revalidateTransactionPaths();
}
