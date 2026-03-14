package usecase

import (
	"context"
	"time"

	"backend/internal/logic/domain"
	"backend/internal/logic/repository"
)

// AnalyticsUsecase は集計のユースケース
type AnalyticsUsecase interface {
	GetSummary(ctx context.Context, userID int, from, to *time.Time, accountID, categoryID *int) (*domain.TransactionSummary, error)
	GetCashflow(ctx context.Context, userID int, groupBy repository.CashflowGroupBy, from, to *time.Time, accountID, categoryID *int) ([]*domain.CashflowByPeriod, error)
}

var _ AnalyticsUsecase = (*AnalyticsUsecaseImpl)(nil)

// AnalyticsUsecaseImpl は AnalyticsUsecase の実装
type AnalyticsUsecaseImpl struct {
	analyticsRepo repository.AnalyticsRepository
}

// NewAnalyticsUsecase は AnalyticsUsecaseImpl を生成する
func NewAnalyticsUsecase(analyticsRepo repository.AnalyticsRepository) *AnalyticsUsecaseImpl {
	return &AnalyticsUsecaseImpl{analyticsRepo: analyticsRepo}
}

// GetSummary は期間内の取引サマリーを取得する
func (u *AnalyticsUsecaseImpl) GetSummary(ctx context.Context, userID int, from, to *time.Time, accountID, categoryID *int) (*domain.TransactionSummary, error) {
	return u.analyticsRepo.GetSummary(ctx, userID, from, to, accountID, categoryID)
}

// GetCashflow は指定単位でキャッシュフローを取得する
func (u *AnalyticsUsecaseImpl) GetCashflow(ctx context.Context, userID int, groupBy repository.CashflowGroupBy, from, to *time.Time, accountID, categoryID *int) ([]*domain.CashflowByPeriod, error) {
	return u.analyticsRepo.GetCashflow(ctx, userID, groupBy, from, to, accountID, categoryID)
}
