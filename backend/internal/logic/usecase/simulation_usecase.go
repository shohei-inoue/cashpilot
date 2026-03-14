package usecase

import (
	"context"
	"time"

	"backend/internal/apperrors"
	"backend/internal/logic/domain"
	"backend/internal/logic/repository"
)

// SimulationUsecase はシミュレーションのユースケース
type SimulationUsecase interface {
	Run(ctx context.Context, userID int, req *domain.SimulationRunRequest) (*domain.SimulationRunResponse, error)
}

var _ SimulationUsecase = (*SimulationUsecaseImpl)(nil)

// SimulationUsecaseImpl は SimulationUsecase の実装
type SimulationUsecaseImpl struct {
	analyticsRepo repository.AnalyticsRepository
	goalRepo      repository.GoalRepository
}

// NewSimulationUsecase は SimulationUsecaseImpl を生成する
func NewSimulationUsecase(analyticsRepo repository.AnalyticsRepository, goalRepo repository.GoalRepository) *SimulationUsecaseImpl {
	return &SimulationUsecaseImpl{analyticsRepo: analyticsRepo, goalRepo: goalRepo}
}

// Run はシミュレーションを実行する
func (u *SimulationUsecaseImpl) Run(ctx context.Context, userID int, req *domain.SimulationRunRequest) (*domain.SimulationRunResponse, error) {
	if req.PeriodMonths < 1 || req.PeriodMonths > 120 {
		return nil, apperrors.ErrInvalidSimulationPeriod
	}

	// 現在残高（取引の集計）
	summary, err := u.analyticsRepo.GetSummary(ctx, userID, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	startBalance := summary.NetCashflow

	// 月次収支
	monthlyNet := req.MonthlyIncome - req.MonthlyExpense
	if req.HourlyRate != nil && req.HoursPerMonth != nil && *req.HourlyRate > 0 && *req.HoursPerMonth > 0 {
		monthlyNet += *req.HourlyRate * *req.HoursPerMonth
	}

	// 月次残高を計算
	now := time.Now()
	monthlyBalances := make([]domain.MonthlyBalance, 0, req.PeriodMonths)
	balance := startBalance
	minBalance := balance

	for i := 0; i < req.PeriodMonths; i++ {
		m := now.AddDate(0, i, 0)
		month := m.Format("2006-01")
		balance += monthlyNet
		if balance < minBalance {
			minBalance = balance
		}
		monthlyBalances = append(monthlyBalances, domain.MonthlyBalance{Month: month, Balance: balance})
	}

	endBalance := balance

	// 目標の達成見込み
	goals, err := u.goalRepo.ListByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	goalProjections := make([]domain.GoalProjection, 0, len(goals))
	for _, g := range goals {
		if g.Deadline == nil || *g.Deadline == "" {
			continue
		}
		deadline, err := time.Parse("2006-01-02", *g.Deadline)
		if err != nil {
			continue
		}
		deadlineMonth := deadline.Format("2006-01")

		var projectedBalance int
		achievable := false
		for _, mb := range monthlyBalances {
			if mb.Month == deadlineMonth {
				projectedBalance = mb.Balance
				achievable = mb.Balance >= g.TargetAmount
				break
			}
		}
		goalProjections = append(goalProjections, domain.GoalProjection{
			GoalID:           g.ID,
			Name:             g.Name,
			TargetAmount:     g.TargetAmount,
			Deadline:         g.Deadline,
			ProjectedBalance: projectedBalance,
			Achievable:       achievable,
		})
	}

	return &domain.SimulationRunResponse{
		StartBalance:     startBalance,
		MinBalance:       minBalance,
		EndBalance:       endBalance,
		MonthlyBalances:  monthlyBalances,
		GoalProjections:  goalProjections,
	}, nil
}
