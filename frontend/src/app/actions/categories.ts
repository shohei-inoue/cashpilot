"use server";

import { revalidatePath } from "next/cache";
import { createAuthClient } from "../libs/client";
import { handleResponse } from "../libs/response";
import type { Category, CategoriesResponse } from "../types/category";

const REVALIDATE_PATHS = ["/settings", "/transactions", "/"];

async function revalidateCategoryPaths() {
  for (const path of REVALIDATE_PATHS) {
    revalidatePath(path);
  }
}

export async function getCategories(): Promise<Category[]> {
  const client = await createAuthClient();
  const res = await handleResponse(client.get<CategoriesResponse>("/api/categories"));
  return res.categories ?? [];
}

export async function createCategory(type: string, name: string): Promise<Category> {
  const client = await createAuthClient();
  const category = await handleResponse(
    client.post<Category>("/api/categories", { type, name })
  );
  await revalidateCategoryPaths();
  return category;
}

export async function updateCategory(
  id: number,
  type: string,
  name: string
): Promise<Category> {
  const client = await createAuthClient();
  const category = await handleResponse(
    client.put<Category>(`/api/categories/${id}`, { type, name })
  );
  await revalidateCategoryPaths();
  return category;
}

export async function deleteCategory(id: number): Promise<void> {
  const client = await createAuthClient();
  await client.delete(`/api/categories/${id}`);
  await revalidateCategoryPaths();
}
