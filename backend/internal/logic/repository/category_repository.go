package repository

import (
	"context"

	"backend/internal/logic/domain"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// CategoryRepository はカテゴリの永続化インターフェース
type CategoryRepository interface {
	ListByUserID(ctx context.Context, userID int) ([]*domain.Category, error)
	FindByIDAndUserID(ctx context.Context, id, userID int) (*domain.Category, error)
	Create(ctx context.Context, userID, categoryTypeID int, name string) (*domain.Category, error)
	Update(ctx context.Context, id, userID, categoryTypeID int, name string) (*domain.Category, error)
	Delete(ctx context.Context, id, userID int) error
	GetCategoryTypeIDByName(ctx context.Context, name string) (int, error)
}

var _ CategoryRepository = (*CategoryRepositoryImpl)(nil)

// CategoryRepositoryImpl は PostgreSQL 用の CategoryRepository 実装
type CategoryRepositoryImpl struct {
	pool *pgxpool.Pool
}

// NewCategoryRepository は CategoryRepositoryImpl を生成する
func NewCategoryRepository(pool *pgxpool.Pool) *CategoryRepositoryImpl {
	return &CategoryRepositoryImpl{pool: pool}
}

// GetCategoryTypeIDByName は category_types の name から id を取得する
func (r *CategoryRepositoryImpl) GetCategoryTypeIDByName(ctx context.Context, name string) (int, error) {
	var id int
	err := r.pool.QueryRow(ctx,
		`SELECT id FROM category_types WHERE name = $1`,
		name,
	).Scan(&id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}
	return id, nil
}

// ListByUserID は userID に紐づくカテゴリ一覧を取得する
func (r *CategoryRepositoryImpl) ListByUserID(ctx context.Context, userID int) ([]*domain.Category, error) {
	rows, err := r.pool.Query(ctx,
		`SELECT c.id, c.user_id, c.category_type_id, ct.name, c.name, c.created_at::text, c.updated_at::text
		 FROM categories c
		 JOIN category_types ct ON c.category_type_id = ct.id
		 WHERE c.user_id = $1
		 ORDER BY c.id`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*domain.Category
	for rows.Next() {
		var cat domain.Category
		if err := rows.Scan(&cat.ID, &cat.UserID, &cat.CategoryTypeID, &cat.CategoryType, &cat.Name, &cat.CreatedAt, &cat.UpdatedAt); err != nil {
			return nil, err
		}
		categories = append(categories, &cat)
	}
	return categories, rows.Err()
}

// FindByIDAndUserID は id と userID でカテゴリを取得する
func (r *CategoryRepositoryImpl) FindByIDAndUserID(ctx context.Context, id, userID int) (*domain.Category, error) {
	var cat domain.Category
	err := r.pool.QueryRow(ctx,
		`SELECT c.id, c.user_id, c.category_type_id, ct.name, c.name, c.created_at::text, c.updated_at::text
		 FROM categories c
		 JOIN category_types ct ON c.category_type_id = ct.id
		 WHERE c.id = $1 AND c.user_id = $2`,
		id, userID,
	).Scan(&cat.ID, &cat.UserID, &cat.CategoryTypeID, &cat.CategoryType, &cat.Name, &cat.CreatedAt, &cat.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &cat, nil
}

// Create は新規カテゴリを作成する
func (r *CategoryRepositoryImpl) Create(ctx context.Context, userID, categoryTypeID int, name string) (*domain.Category, error) {
	var id int
	err := r.pool.QueryRow(ctx,
		`INSERT INTO categories (user_id, category_type_id, name)
		 VALUES ($1, $2, $3)
		 RETURNING id`,
		userID, categoryTypeID, name,
	).Scan(&id)
	if err != nil {
		return nil, err
	}
	return r.FindByIDAndUserID(ctx, id, userID)
}

// Update はカテゴリを更新する
func (r *CategoryRepositoryImpl) Update(ctx context.Context, id, userID, categoryTypeID int, name string) (*domain.Category, error) {
	result, err := r.pool.Exec(ctx,
		`UPDATE categories SET category_type_id = $1, name = $2, updated_at = now()
		 WHERE id = $3 AND user_id = $4`,
		categoryTypeID, name, id, userID,
	)
	if err != nil {
		return nil, err
	}
	if result.RowsAffected() == 0 {
		return nil, nil
	}
	return r.FindByIDAndUserID(ctx, id, userID)
}

// Delete はカテゴリを削除する
func (r *CategoryRepositoryImpl) Delete(ctx context.Context, id, userID int) error {
	result, err := r.pool.Exec(ctx,
		`DELETE FROM categories WHERE id = $1 AND user_id = $2`,
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
