"use server";

import { getAccounts } from "./accounts";
import { getCategories } from "./categories";
import type { Account } from "../types/account";
import type { Category } from "../types/category";

export type SettingsData = {
  accounts: Account[];
  categories: Category[];
};

export async function getSettingsData(): Promise<SettingsData | null> {
  try {
    const [accounts, categories] = await Promise.all([getAccounts(), getCategories()]);
    return { accounts, categories };
  } catch {
    return null;
  }
}
