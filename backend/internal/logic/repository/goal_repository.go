package repository

import (
	"context"

	"backend/internal/logic/domain"

	"gorm.io/gorm"
)

// GoalRepository は目標の永続化インターフェース
type GoalRepository interface {
	ListByUserID(ctx context.Context, userID int) ([]*domain.Goal, error)
	FindByIDAndUserID(ctx context.Context, id, userID int) (*domain.Goal, error)
	Create(ctx context.Context, userID int, name string, targetAmount int, deadline *string) (*domain.Goal, error)
	Update(ctx context.Context, id, userID int, name string, targetAmount int, deadline *string) (*domain.Goal, error)
	Delete(ctx context.Context, id, userID int) error
}

var _ GoalRepository = (*GoalRepositoryImpl)(nil)

// GoalRepositoryImpl は GORM 用の GoalRepository 実装
type GoalRepositoryImpl struct {
	db *gorm.DB
}

// NewGoalRepository は GoalRepositoryImpl を生成する
func NewGoalRepository(db *gorm.DB) *GoalRepositoryImpl {
	return &GoalRepositoryImpl{db: db}
}

// ListByUserID は userID に紐づく目標一覧を取得する
func (r *GoalRepositoryImpl) ListByUserID(ctx context.Context, userID int) ([]*domain.Goal, error) {
	var list []*domain.Goal
	err := r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Order("id").
		Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

// FindByIDAndUserID は id と userID で目標を取得する
func (r *GoalRepositoryImpl) FindByIDAndUserID(ctx context.Context, id, userID int) (*domain.Goal, error) {
	var g domain.Goal
	err := r.db.WithContext(ctx).
		Where("id = ? AND user_id = ?", id, userID).
		First(&g).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &g, nil
}

// Create は新規目標を作成する
func (r *GoalRepositoryImpl) Create(ctx context.Context, userID int, name string, targetAmount int, deadline *string) (*domain.Goal, error) {
	g := domain.Goal{
		UserID:       userID,
		Name:         name,
		TargetAmount: targetAmount,
		Deadline:     deadline,
	}
	if err := r.db.WithContext(ctx).Create(&g).Error; err != nil {
		return nil, err
	}
	return r.FindByIDAndUserID(ctx, g.ID, userID)
}

// Update は目標を更新する
func (r *GoalRepositoryImpl) Update(ctx context.Context, id, userID int, name string, targetAmount int, deadline *string) (*domain.Goal, error) {
	updates := map[string]interface{}{
		"name":          name,
		"target_amount": targetAmount,
		"deadline":      deadline,
	}
	result := r.db.WithContext(ctx).Model(&domain.Goal{}).
		Where("id = ? AND user_id = ?", id, userID).
		Updates(updates)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return r.FindByIDAndUserID(ctx, id, userID)
}

// Delete は目標を削除する
func (r *GoalRepositoryImpl) Delete(ctx context.Context, id, userID int) error {
	result := r.db.WithContext(ctx).
		Where("id = ? AND user_id = ?", id, userID).
		Delete(&domain.Goal{})
	return result.Error
}
