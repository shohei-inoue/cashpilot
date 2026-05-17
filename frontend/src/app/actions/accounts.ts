"use server";

import { revalidatePath } from "next/cache";
import { createAuthClient } from "../libs/client";
import { handleResponse } from "../libs/response";
import type { Account, AccountsResponse } from "../types/account";

const REVALIDATE_PATHS = ["/settings", "/transactions", "/"];

async function revalidateAccountPaths() {
  for (const path of REVALIDATE_PATHS) {
    revalidatePath(path);
  }
}

export async function getAccounts(): Promise<Account[]> {
  const client = await createAuthClient();
  const res = await handleResponse(client.get<AccountsResponse>("/api/accounts"));
  return res.accounts ?? [];
}

export async function createAccount(type: string, name: string): Promise<Account> {
  const client = await createAuthClient();
  const account = await handleResponse(
    client.post<Account>("/api/accounts", { type, name })
  );
  await revalidateAccountPaths();
  return account;
}

export async function updateAccount(
  id: number,
  type: string,
  name: string
): Promise<Account> {
  const client = await createAuthClient();
  const account = await handleResponse(
    client.put<Account>(`/api/accounts/${id}`, { type, name })
  );
  await revalidateAccountPaths();
  return account;
}

export async function deleteAccount(id: number): Promise<void> {
  const client = await createAuthClient();
  await client.delete(`/api/accounts/${id}`);
  await revalidateAccountPaths();
}
