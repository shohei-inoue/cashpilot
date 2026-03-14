package domain

// SimulationRunRequest はシミュレーション実行のリクエスト
type SimulationRunRequest struct {
	PeriodMonths    int  `json:"period_months"`    // シミュレーション期間（月）、1〜120
	MonthlyIncome   int  `json:"monthly_income"`   // 月収（正の整数、任意）
	MonthlyExpense  int  `json:"monthly_expense"`  // 月間固定費（正の整数、任意）
	HourlyRate      *int `json:"hourly_rate,omitempty"`      // 時給（副業、任意）
	HoursPerMonth   *int `json:"hours_per_month,omitempty"`  // 月間労働時間（任意）
}

// MonthlyBalance は月ごとの残高
type MonthlyBalance struct {
	Month   string `json:"month"`   // YYYY-MM
	Balance int    `json:"balance"`
}

// GoalProjection は目標の達成見込み
type GoalProjection struct {
	GoalID           int     `json:"goal_id"`
	Name             string  `json:"name"`
	TargetAmount     int     `json:"target_amount"`
	Deadline         *string `json:"deadline,omitempty"`
	ProjectedBalance int     `json:"projected_balance"` // その時点の予測残高
	Achievable       bool    `json:"achievable"`        // 目標達成可否
}

// SimulationRunResponse はシミュレーション実行のレスポンス
type SimulationRunResponse struct {
	StartBalance    int               `json:"start_balance"`
	MinBalance      int               `json:"min_balance"`
	EndBalance      int               `json:"end_balance"`
	MonthlyBalances []MonthlyBalance  `json:"monthly_balances"`
	GoalProjections []GoalProjection  `json:"goal_projections,omitempty"`
}
