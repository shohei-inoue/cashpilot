package repository

import (
	"context"
	"fmt"
	"time"

	"backend/internal/logic/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

// CashflowGroupBy はキャッシュフローの集計単位
type CashflowGroupBy string

const (
	CashflowGroupByAll     CashflowGroupBy = "all"
	CashflowGroupByDaily   CashflowGroupBy = "daily"
	CashflowGroupByWeekly  CashflowGroupBy = "weekly"
	CashflowGroupByMonthly CashflowGroupBy = "monthly"
	CashflowGroupByYearly  CashflowGroupBy = "yearly"
)

// AnalyticsRepository は集計用のリポジトリ
type AnalyticsRepository interface {
	GetSummary(ctx context.Context, userID int, from, to *time.Time, accountID, categoryID *int) (*domain.TransactionSummary, error)
	GetCashflow(ctx context.Context, userID int, groupBy CashflowGroupBy, from, to *time.Time, accountID, categoryID *int) ([]*domain.CashflowByPeriod, error)
}

var _ AnalyticsRepository = (*AnalyticsRepositoryImpl)(nil)

// AnalyticsRepositoryImpl は PostgreSQL 用の AnalyticsRepository 実装
type AnalyticsRepositoryImpl struct {
	pool *pgxpool.Pool
}

// NewAnalyticsRepository は AnalyticsRepositoryImpl を生成する
func NewAnalyticsRepository(pool *pgxpool.Pool) *AnalyticsRepositoryImpl {
	return &AnalyticsRepositoryImpl{pool: pool}
}

// GetSummary は期間内の取引サマリーを取得する
func (r *AnalyticsRepositoryImpl) GetSummary(ctx context.Context, userID int, from, to *time.Time, accountID, categoryID *int) (*domain.TransactionSummary, error) {
	query := `SELECT
		COALESCE(SUM(CASE WHEN amount > 0 THEN amount ELSE 0 END), 0) AS total_income,
		COALESCE(SUM(CASE WHEN amount < 0 THEN amount ELSE 0 END), 0) AS total_expense,
		COALESCE(SUM(amount), 0) AS net_cashflow,
		COUNT(*)::int AS transaction_count
		FROM transactions
		WHERE user_id = $1`
	args := []interface{}{userID}
	argIdx := 2

	if from != nil {
		query += fmt.Sprintf(` AND occurred_at >= $%d`, argIdx)
		args = append(args, from)
		argIdx++
	}
	if to != nil {
		query += fmt.Sprintf(` AND occurred_at <= $%d`, argIdx)
		args = append(args, to)
		argIdx++
	}
	if accountID != nil {
		query += fmt.Sprintf(` AND account_id = $%d`, argIdx)
		args = append(args, *accountID)
		argIdx++
	}
	if categoryID != nil {
		query += fmt.Sprintf(` AND category_id = $%d`, argIdx)
		args = append(args, *categoryID)
		argIdx++
	}

	var s domain.TransactionSummary
	err := r.pool.QueryRow(ctx, query, args...).Scan(
		&s.TotalIncome,
		&s.TotalExpense,
		&s.NetCashflow,
		&s.TransactionCount,
	)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

// GetCashflow は指定単位でキャッシュフローを取得する
func (r *AnalyticsRepositoryImpl) GetCashflow(ctx context.Context, userID int, groupBy CashflowGroupBy, from, to *time.Time, accountID, categoryID *int) ([]*domain.CashflowByPeriod, error) {
	if groupBy == CashflowGroupByAll {
		return r.getCashflowAll(ctx, userID, from, to, accountID, categoryID)
	}

	var periodExpr, orderExpr string
	switch groupBy {
	case CashflowGroupByDaily:
		periodExpr = `TO_CHAR(occurred_at::date, 'YYYY-MM-DD')`
		orderExpr = periodExpr
	case CashflowGroupByWeekly:
		periodExpr = `TO_CHAR(date_trunc('week', occurred_at)::date, 'YYYY-MM-DD')`
		orderExpr = periodExpr
	case CashflowGroupByMonthly:
		periodExpr = `TO_CHAR(occurred_at, 'YYYY-MM')`
		orderExpr = periodExpr
	case CashflowGroupByYearly:
		periodExpr = `TO_CHAR(occurred_at, 'YYYY')`
		orderExpr = periodExpr
	default:
		periodExpr = `TO_CHAR(occurred_at, 'YYYY-MM')`
		orderExpr = periodExpr
	}

	query := fmt.Sprintf(`SELECT
		%s AS period,
		COALESCE(SUM(CASE WHEN amount > 0 THEN amount ELSE 0 END), 0)::int AS total_income,
		COALESCE(SUM(CASE WHEN amount < 0 THEN amount ELSE 0 END), 0)::int AS total_expense,
		COALESCE(SUM(amount), 0)::int AS net_cashflow,
		COUNT(*)::int AS transaction_count
		FROM transactions
		WHERE user_id = $1`, periodExpr)
	args := []interface{}{userID}
	argIdx := 2

	if from != nil {
		query += fmt.Sprintf(` AND occurred_at >= $%d`, argIdx)
		args = append(args, from)
		argIdx++
	}
	if to != nil {
		query += fmt.Sprintf(` AND occurred_at <= $%d`, argIdx)
		args = append(args, to)
		argIdx++
	}
	if accountID != nil {
		query += fmt.Sprintf(` AND account_id = $%d`, argIdx)
		args = append(args, *accountID)
		argIdx++
	}
	if categoryID != nil {
		query += fmt.Sprintf(` AND category_id = $%d`, argIdx)
		args = append(args, *categoryID)
		argIdx++
	}

	query += fmt.Sprintf(` GROUP BY %s ORDER BY %s`, periodExpr, orderExpr)

	rows, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []*domain.CashflowByPeriod
	for rows.Next() {
		var m domain.CashflowByPeriod
		if err := rows.Scan(&m.Period, &m.TotalIncome, &m.TotalExpense, &m.NetCashflow, &m.TransactionCount); err != nil {
			return nil, err
		}
		items = append(items, &m)
	}
	return items, rows.Err()
}

// getCashflowAll は期間全体を1件で返す
func (r *AnalyticsRepositoryImpl) getCashflowAll(ctx context.Context, userID int, from, to *time.Time, accountID, categoryID *int) ([]*domain.CashflowByPeriod, error) {
	query := `SELECT
		COALESCE(SUM(CASE WHEN amount > 0 THEN amount ELSE 0 END), 0)::int AS total_income,
		COALESCE(SUM(CASE WHEN amount < 0 THEN amount ELSE 0 END), 0)::int AS total_expense,
		COALESCE(SUM(amount), 0)::int AS net_cashflow,
		COUNT(*)::int AS transaction_count
		FROM transactions
		WHERE user_id = $1`
	args := []interface{}{userID}
	argIdx := 2

	if from != nil {
		query += fmt.Sprintf(` AND occurred_at >= $%d`, argIdx)
		args = append(args, from)
		argIdx++
	}
	if to != nil {
		query += fmt.Sprintf(` AND occurred_at <= $%d`, argIdx)
		args = append(args, to)
		argIdx++
	}
	if accountID != nil {
		query += fmt.Sprintf(` AND account_id = $%d`, argIdx)
		args = append(args, *accountID)
		argIdx++
	}
	if categoryID != nil {
		query += fmt.Sprintf(` AND category_id = $%d`, argIdx)
		args = append(args, *categoryID)
		argIdx++
	}

	var m domain.CashflowByPeriod
	m.Period = "all"
	err := r.pool.QueryRow(ctx, query, args...).Scan(
		&m.TotalIncome,
		&m.TotalExpense,
		&m.NetCashflow,
		&m.TransactionCount,
	)
	if err != nil {
		return nil, err
	}
	return []*domain.CashflowByPeriod{&m}, nil
}
