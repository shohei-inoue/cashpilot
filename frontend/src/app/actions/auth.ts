"use server";

import { cookies } from "next/headers";
import type { User } from "../types/user";
import { createAuthClient } from "../libs/client";
import { handleResponse } from "../libs/response";

export type { User };

/** サインアップ */
export async function signup(
  email: string,
  password: string
): Promise<User> {
  const client = await createAuthClient();
  return handleResponse(client.post<User>("/api/auth/signup", { email, password }));
}

/** ログイン */
export async function login(
  email: string,
  password: string
): Promise<User> {
  const client = await createAuthClient();
  return handleResponse(client.post<User>("/api/auth/login", { email, password }));
}

/** ログアウト */
export async function logout(): Promise<void> {
  const client = await createAuthClient();
  await handleResponse(client.post("/api/auth/logout")).catch(() => {});
  const cookieStore = await cookies();
  cookieStore.delete("token");
}
