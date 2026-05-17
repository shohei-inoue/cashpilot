export type SimulationRunInput = {
  period_months: number;
  monthly_income: number;
  monthly_expense: number;
  hourly_rate?: number;
  hours_per_month?: number;
};

export type MonthlyBalance = {
  month: string;
  balance: number;
};

export type GoalProjection = {
  goal_id: number;
  name: string;
  target_amount: number;
  deadline?: string | null;
  projected_balance: number;
  achievable: boolean;
};

export type SimulationRunResponse = {
  start_balance: number;
  min_balance: number;
  end_balance: number;
  monthly_balances: MonthlyBalance[];
  goal_projections?: GoalProjection[];
};
