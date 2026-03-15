"use server";

import type { User } from "../types/user";
import { createAuthClient } from "../libs/client";
import { handleResponse } from "../libs/response";

/** 現在のユーザー取得 */
export async function getCurrentUser(): Promise<User | null> {
  const client = await createAuthClient();
  return handleResponse(client.get<User>("/api/user")).catch(() => null);
}
