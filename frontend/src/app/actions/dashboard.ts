"use server";

import { createAuthClient } from "../libs/client";
import {
  getCurrentMonthLabel,
  getCurrentMonthRange,
  getMonthsRange,
  formatYearMonth,
} from "../libs/format";
import { handleResponse } from "../libs/response";
import type {
  CashflowResponse,
  TransactionSummaryResponse,
} from "../types/analytics";
import type { TransactionsResponse } from "../types/transaction";

export type CashflowTrendPoint = {
  period: string;
  label: string;
  balance: number;
};

export type DashboardData = {
  monthLabel: string;
  balance: number;
  monthSummary: TransactionSummaryResponse["summary"];
  recentTransactions: TransactionsResponse["transactions"];
  cashflowTrend: CashflowTrendPoint[];
};

export type DashboardResult =
  | {
      status: "ok";
      data: DashboardData;
    }
  | {
      status: "unauthorized";
    }
  | {
      status: "error";
      message: string;
    };

function buildCashflowTrend(items: CashflowResponse["items"]): CashflowTrendPoint[] {
  const sorted = [...items].sort((a, b) => a.period.localeCompare(b.period));
  let balance = 0;
  return sorted.map((item) => {
    balance += item.net_cashflow;
    return {
      period: item.period,
      label: formatYearMonth(item.period),
      balance,
    };
  });
}

/** ダッシュボード表示用データを取得 */
export async function getDashboardData(): Promise<DashboardResult> {
  const client = await createAuthClient();
  const monthRange = getCurrentMonthRange();
  const trendRange = getMonthsRange(12);

  try {
    const [allTimeRes, monthRes, transactionsRes, cashflowRes] = await Promise.all([
      handleResponse(
        client.get<TransactionSummaryResponse>("/api/transactions/summary")
      ),
      handleResponse(
        client.get<TransactionSummaryResponse>("/api/transactions/summary", {
          params: monthRange,
        })
      ),
      handleResponse(
        client.get<TransactionsResponse>("/api/transactions", {
          params: { limit: 10 },
        })
      ),
      handleResponse(
        client.get<CashflowResponse>("/api/analytics/cashflow", {
          params: { group_by: "monthly", ...trendRange },
        })
      ),
    ]);

    return {
      status: "ok",
      data: {
        monthLabel: getCurrentMonthLabel(),
        balance: allTimeRes.summary.net_cashflow,
        monthSummary: monthRes.summary,
        recentTransactions: transactionsRes.transactions ?? [],
        cashflowTrend: buildCashflowTrend(cashflowRes.items ?? []),
      },
    };
  } catch (error) {
    if (error instanceof Error && error.message === "Unauthorized") {
      return { status: "unauthorized" };
    }

    return {
      status: "error",
      message: "ダッシュボード情報の取得に失敗しました。時間をおいて再試行してください。",
    };
  }
}
