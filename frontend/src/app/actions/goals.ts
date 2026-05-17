"use server";

import { revalidatePath } from "next/cache";
import { createAuthClient } from "../libs/client";
import { handleResponse } from "../libs/response";
import type { Goal, GoalInput, GoalsResponse } from "../types/goal";

const REVALIDATE_PATHS = ["/settings", "/simulation", "/"];

async function revalidateGoalPaths() {
  for (const path of REVALIDATE_PATHS) {
    revalidatePath(path);
  }
}

export async function getGoals(): Promise<Goal[]> {
  const client = await createAuthClient();
  const res = await handleResponse(client.get<GoalsResponse>("/api/goals"));
  return res.goals ?? [];
}

export async function createGoal(input: GoalInput): Promise<Goal> {
  const client = await createAuthClient();
  const body = {
    name: input.name,
    target_amount: input.target_amount,
    deadline: input.deadline || undefined,
  };
  const goal = await handleResponse(client.post<Goal>("/api/goals", body));
  await revalidateGoalPaths();
  return goal;
}

export async function updateGoal(id: number, input: GoalInput): Promise<Goal> {
  const client = await createAuthClient();
  const body = {
    name: input.name,
    target_amount: input.target_amount,
    deadline: input.deadline || undefined,
  };
  const goal = await handleResponse(client.put<Goal>(`/api/goals/${id}`, body));
  await revalidateGoalPaths();
  return goal;
}

export async function deleteGoal(id: number): Promise<void> {
  const client = await createAuthClient();
  await client.delete(`/api/goals/${id}`);
  await revalidateGoalPaths();
}
