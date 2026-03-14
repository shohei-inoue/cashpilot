package repository

import (
	"context"

	"backend/internal/logic/domain"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
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

// GoalRepositoryImpl は PostgreSQL 用の GoalRepository 実装
type GoalRepositoryImpl struct {
	pool *pgxpool.Pool
}

// NewGoalRepository は GoalRepositoryImpl を生成する
func NewGoalRepository(pool *pgxpool.Pool) *GoalRepositoryImpl {
	return &GoalRepositoryImpl{pool: pool}
}

// ListByUserID は userID に紐づく目標一覧を取得する
func (r *GoalRepositoryImpl) ListByUserID(ctx context.Context, userID int) ([]*domain.Goal, error) {
	rows, err := r.pool.Query(ctx,
		`SELECT id, user_id, name, target_amount, deadline::text, created_at::text, updated_at::text
		 FROM goals WHERE user_id = $1 ORDER BY id`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var goals []*domain.Goal
	for rows.Next() {
		var g domain.Goal
		if err := rows.Scan(&g.ID, &g.UserID, &g.Name, &g.TargetAmount, &g.Deadline, &g.CreatedAt, &g.UpdatedAt); err != nil {
			return nil, err
		}
		goals = append(goals, &g)
	}
	return goals, rows.Err()
}

// FindByIDAndUserID は id と userID で目標を取得する
func (r *GoalRepositoryImpl) FindByIDAndUserID(ctx context.Context, id, userID int) (*domain.Goal, error) {
	var g domain.Goal
	err := r.pool.QueryRow(ctx,
		`SELECT id, user_id, name, target_amount, deadline::text, created_at::text, updated_at::text
		 FROM goals WHERE id = $1 AND user_id = $2`,
		id, userID,
	).Scan(&g.ID, &g.UserID, &g.Name, &g.TargetAmount, &g.Deadline, &g.CreatedAt, &g.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &g, nil
}

// Create は新規目標を作成する
func (r *GoalRepositoryImpl) Create(ctx context.Context, userID int, name string, targetAmount int, deadline *string) (*domain.Goal, error) {
	var id int
	err := r.pool.QueryRow(ctx,
		`INSERT INTO goals (user_id, name, target_amount, deadline)
		 VALUES ($1, $2, $3, $4)
		 RETURNING id`,
		userID, name, targetAmount, deadline,
	).Scan(&id)
	if err != nil {
		return nil, err
	}
	return r.FindByIDAndUserID(ctx, id, userID)
}

// Update は目標を更新する
func (r *GoalRepositoryImpl) Update(ctx context.Context, id, userID int, name string, targetAmount int, deadline *string) (*domain.Goal, error) {
	result, err := r.pool.Exec(ctx,
		`UPDATE goals SET name = $1, target_amount = $2, deadline = $3, updated_at = now()
		 WHERE id = $4 AND user_id = $5`,
		name, targetAmount, deadline, id, userID,
	)
	if err != nil {
		return nil, err
	}
	if result.RowsAffected() == 0 {
		return nil, nil
	}
	return r.FindByIDAndUserID(ctx, id, userID)
}

// Delete は目標を削除する
func (r *GoalRepositoryImpl) Delete(ctx context.Context, id, userID int) error {
	result, err := r.pool.Exec(ctx,
		`DELETE FROM goals WHERE id = $1 AND user_id = $2`,
		id, userID,
	)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return nil
	}
	return nil
}
