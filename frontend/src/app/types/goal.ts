export type Goal = {
  id: number;
  name: string;
  target_amount: number;
  deadline?: string | null;
  created_at?: string;
  updated_at?: string;
};

export type GoalsResponse = {
  goals: Goal[];
};

export type GoalInput = {
  name: string;
  target_amount: number;
  deadline?: string;
};
