package domain

// TransactionSummary は期間内の取引サマリー
type TransactionSummary struct {
	TotalIncome     int `json:"total_income"`
	TotalExpense    int `json:"total_expense"`
	NetCashflow     int `json:"net_cashflow"`
	TransactionCount int `json:"transaction_count"`
}

// CashflowByPeriod は期間単位のキャッシュフロー（daily: YYYY-MM-DD, weekly: YYYY-MM-DD週初, monthly: YYYY-MM, yearly: YYYY）
type CashflowByPeriod struct {
	Period          string `json:"period"`
	TotalIncome     int    `json:"total_income"`
	TotalExpense    int    `json:"total_expense"`
	NetCashflow     int    `json:"net_cashflow"`
	TransactionCount int   `json:"transaction_count"`
}
