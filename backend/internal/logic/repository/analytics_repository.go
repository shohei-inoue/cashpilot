package repository

import (
	"context"
	"fmt"
	"time"

	"backend/internal/logic/domain"

	"gorm.io/gorm"
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

// AnalyticsRepositoryImpl は GORM 用の AnalyticsRepository 実装
type AnalyticsRepositoryImpl struct {
	db *gorm.DB
}

// NewAnalyticsRepository は AnalyticsRepositoryImpl を生成する
func NewAnalyticsRepository(db *gorm.DB) *AnalyticsRepositoryImpl {
	return &AnalyticsRepositoryImpl{db: db}
}

func (r *AnalyticsRepositoryImpl) summaryWhere(db *gorm.DB, userID int, from, to *time.Time, accountID, categoryID *int) *gorm.DB {
	db = db.Where("user_id = ?", userID)
	if from != nil {
		db = db.Where("occurred_at >= ?", from)
	}
	if to != nil {
		db = db.Where("occurred_at <= ?", to)
	}
	if accountID != nil {
		db = db.Where("account_id = ?", *accountID)
	}
	if categoryID != nil {
		db = db.Where("category_id = ?", *categoryID)
	}
	return db
}

// GetSummary は期間内の取引サマリーを取得する
func (r *AnalyticsRepositoryImpl) GetSummary(ctx context.Context, userID int, from, to *time.Time, accountID, categoryID *int) (*domain.TransactionSummary, error) {
	var s domain.TransactionSummary
	q := r.db.WithContext(ctx).Table("transactions").
		Select("COALESCE(SUM(CASE WHEN amount > 0 THEN amount ELSE 0 END), 0) AS total_income, COALESCE(SUM(CASE WHEN amount < 0 THEN amount ELSE 0 END), 0) AS total_expense, COALESCE(SUM(amount), 0) AS net_cashflow, COUNT(*)::int AS transaction_count")
	q = r.summaryWhere(q, userID, from, to, accountID, categoryID)
	if err := q.Scan(&s).Error; err != nil {
		return nil, err
	}
	return &s, nil
}

// GetCashflow は指定単位でキャッシュフローを取得する
func (r *AnalyticsRepositoryImpl) GetCashflow(ctx context.Context, userID int, groupBy CashflowGroupBy, from, to *time.Time, accountID, categoryID *int) ([]*domain.CashflowByPeriod, error) {
	if groupBy == CashflowGroupByAll {
		return r.getCashflowAll(ctx, userID, from, to, accountID, categoryID)
	}

	var periodExpr string
	switch groupBy {
	case CashflowGroupByDaily:
		periodExpr = `TO_CHAR(occurred_at::date, 'YYYY-MM-DD')`
	case CashflowGroupByWeekly:
		periodExpr = `TO_CHAR(date_trunc('week', occurred_at)::date, 'YYYY-MM-DD')`
	case CashflowGroupByMonthly:
		periodExpr = `TO_CHAR(occurred_at, 'YYYY-MM')`
	case CashflowGroupByYearly:
		periodExpr = `TO_CHAR(occurred_at, 'YYYY')`
	default:
		periodExpr = `TO_CHAR(occurred_at, 'YYYY-MM')`
	}

	query := fmt.Sprintf(`SELECT %s AS period,
		COALESCE(SUM(CASE WHEN amount > 0 THEN amount ELSE 0 END), 0)::int AS total_income,
		COALESCE(SUM(CASE WHEN amount < 0 THEN amount ELSE 0 END), 0)::int AS total_expense,
		COALESCE(SUM(amount), 0)::int AS net_cashflow,
		COUNT(*)::int AS transaction_count
		FROM transactions WHERE user_id = ?`, periodExpr)
	args := []interface{}{userID}
	if from != nil {
		query += " AND occurred_at >= ?"
		args = append(args, from)
	}
	if to != nil {
		query += " AND occurred_at <= ?"
		args = append(args, to)
	}
	if accountID != nil {
		query += " AND account_id = ?"
		args = append(args, *accountID)
	}
	if categoryID != nil {
		query += " AND category_id = ?"
		args = append(args, *categoryID)
	}
	query += fmt.Sprintf(" GROUP BY %s ORDER BY %s", periodExpr, periodExpr)

	var items []*domain.CashflowByPeriod
	if err := r.db.WithContext(ctx).Raw(query, args...).Scan(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (r *AnalyticsRepositoryImpl) getCashflowAll(ctx context.Context, userID int, from, to *time.Time, accountID, categoryID *int) ([]*domain.CashflowByPeriod, error) {
	var m domain.CashflowByPeriod
	q := r.db.WithContext(ctx).Table("transactions").
		Select("COALESCE(SUM(CASE WHEN amount > 0 THEN amount ELSE 0 END), 0)::int AS total_income, COALESCE(SUM(CASE WHEN amount < 0 THEN amount ELSE 0 END), 0)::int AS total_expense, COALESCE(SUM(amount), 0)::int AS net_cashflow, COUNT(*)::int AS transaction_count")
	q = r.summaryWhere(q, userID, from, to, accountID, categoryID)
	if err := q.Scan(&m).Error; err != nil {
		return nil, err
	}
	m.Period = "all"
	return []*domain.CashflowByPeriod{&m}, nil
}
