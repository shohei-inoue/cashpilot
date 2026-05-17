"use server";

import { getAccounts } from "./accounts";
import { getCategories } from "./categories";
import { getGoals } from "./goals";
import type { Account } from "../types/account";
import type { Category } from "../types/category";
import type { Goal } from "../types/goal";

export type SettingsData = {
  accounts: Account[];
  categories: Category[];
  goals: Goal[];
};

export async function getSettingsData(): Promise<SettingsData | null> {
  try {
    const [accounts, categories, goals] = await Promise.all([
      getAccounts(),
      getCategories(),
      getGoals(),
    ]);
    return { accounts, categories, goals };
  } catch {
    return null;
  }
}
