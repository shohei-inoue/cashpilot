package usecase

import (
	"context"
	"strings"

	"backend/internal/apperrors"
	"backend/internal/logic/domain"
	"backend/internal/logic/repository"
)

// GoalUsecase は目標のユースケース
type GoalUsecase interface {
	List(ctx context.Context, userID int) ([]*domain.Goal, error)
	Get(ctx context.Context, id, userID int) (*domain.Goal, error)
	Create(ctx context.Context, userID int, name string, targetAmount int, deadline *string) (*domain.Goal, error)
	Update(ctx context.Context, id, userID int, name string, targetAmount int, deadline *string) (*domain.Goal, error)
	Delete(ctx context.Context, id, userID int) error
}

var _ GoalUsecase = (*GoalUsecaseImpl)(nil)

// GoalUsecaseImpl は GoalUsecase の実装
type GoalUsecaseImpl struct {
	goalRepo repository.GoalRepository
}

// NewGoalUsecase は GoalUsecaseImpl を生成する
func NewGoalUsecase(goalRepo repository.GoalRepository) *GoalUsecaseImpl {
	return &GoalUsecaseImpl{goalRepo: goalRepo}
}

// List は userID に紐づく目標一覧を取得する
func (u *GoalUsecaseImpl) List(ctx context.Context, userID int) ([]*domain.Goal, error) {
	return u.goalRepo.ListByUserID(ctx, userID)
}

// Get は id と userID で目標を取得する
func (u *GoalUsecaseImpl) Get(ctx context.Context, id, userID int) (*domain.Goal, error) {
	g, err := u.goalRepo.FindByIDAndUserID(ctx, id, userID)
	if err != nil {
		return nil, err
	}
	if g == nil {
		return nil, apperrors.ErrGoalNotFound
	}
	return g, nil
}

// Create は新規目標を作成する
func (u *GoalUsecaseImpl) Create(ctx context.Context, userID int, name string, targetAmount int, deadline *string) (*domain.Goal, error) {
	if strings.TrimSpace(name) == "" {
		return nil, apperrors.ErrInvalidGoalName
	}
	if targetAmount <= 0 {
		return nil, apperrors.ErrInvalidTargetAmount
	}
	return u.goalRepo.Create(ctx, userID, strings.TrimSpace(name), targetAmount, deadline)
}

// Update は目標を更新する
func (u *GoalUsecaseImpl) Update(ctx context.Context, id, userID int, name string, targetAmount int, deadline *string) (*domain.Goal, error) {
	g, err := u.goalRepo.FindByIDAndUserID(ctx, id, userID)
	if err != nil {
		return nil, err
	}
	if g == nil {
		return nil, apperrors.ErrGoalNotFound
	}
	if strings.TrimSpace(name) == "" {
		return nil, apperrors.ErrInvalidGoalName
	}
	if targetAmount <= 0 {
		return nil, apperrors.ErrInvalidTargetAmount
	}
	return u.goalRepo.Update(ctx, id, userID, strings.TrimSpace(name), targetAmount, deadline)
}

// Delete は目標を削除する
func (u *GoalUsecaseImpl) Delete(ctx context.Context, id, userID int) error {
	g, err := u.goalRepo.FindByIDAndUserID(ctx, id, userID)
	if err != nil {
		return err
	}
	if g == nil {
		return apperrors.ErrGoalNotFound
	}
	return u.goalRepo.Delete(ctx, id, userID)
}
