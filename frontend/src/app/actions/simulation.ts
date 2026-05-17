"use server";

import { createAuthClient } from "../libs/client";
import { handleResponse } from "../libs/response";
import type {
  SimulationRunInput,
  SimulationRunResponse,
} from "../types/simulation";

export async function runSimulation(
  input: SimulationRunInput
): Promise<SimulationRunResponse> {
  const client = await createAuthClient();
  const body: SimulationRunInput = {
    period_months: input.period_months,
    monthly_income: input.monthly_income,
    monthly_expense: input.monthly_expense,
  };
  if (input.hourly_rate != null && input.hours_per_month != null) {
    body.hourly_rate = input.hourly_rate;
    body.hours_per_month = input.hours_per_month;
  }
  return handleResponse(
    client.post<SimulationRunResponse>("/api/simulation/run", body)
  );
}
