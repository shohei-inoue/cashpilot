package repository

import (
	"context"
	"fmt"
	"time"

	"backend/internal/logic/domain"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// ListFilter は取引一覧のフィルタ条件
type ListFilter struct {
	From       *time.Time
	To         *time.Time
	AccountID  *int
	CategoryID *int
	Limit      int
	Offset     int
}

// TransactionRepository は取引の永続化インターフェース
type TransactionRepository interface {
	List(ctx context.Context, userID int, filter ListFilter) ([]*domain.Transaction, error)
	FindByIDAndUserID(ctx context.Context, id, userID int) (*domain.Transaction, error)
	Create(ctx context.Context, userID, accountID int, categoryID *int, amount int, memo *string, occurredAt time.Time) (*domain.Transaction, error)
	Update(ctx context.Context, id, userID, accountID int, categoryID *int, amount int, memo *string, occurredAt time.Time) (*domain.Transaction, error)
	Delete(ctx context.Context, id, userID int) error
}

var _ TransactionRepository = (*TransactionRepositoryImpl)(nil)

// TransactionRepositoryImpl は PostgreSQL 用の TransactionRepository 実装
type TransactionRepositoryImpl struct {
	pool *pgxpool.Pool
}

// NewTransactionRepository は TransactionRepositoryImpl を生成する
func NewTransactionRepository(pool *pgxpool.Pool) *TransactionRepositoryImpl {
	return &TransactionRepositoryImpl{pool: pool}
}

// List は userID に紐づく取引一覧を取得する
func (r *TransactionRepositoryImpl) List(ctx context.Context, userID int, filter ListFilter) ([]*domain.Transaction, error) {
	query := `SELECT t.id, t.user_id, t.account_id, a.name, t.category_id, c.name, t.amount, t.memo, t.occurred_at::text, t.created_at::text, t.updated_at::text
		FROM transactions t
		JOIN accounts a ON t.account_id = a.id
		LEFT JOIN categories c ON t.category_id = c.id
		WHERE t.user_id = $1`
	args := []interface{}{userID}
	argIdx := 2

	if filter.From != nil {
		query += fmt.Sprintf(` AND t.occurred_at >= $%d`, argIdx)
		args = append(args, filter.From)
		argIdx++
	}
	if filter.To != nil {
		query += fmt.Sprintf(` AND t.occurred_at <= $%d`, argIdx)
		args = append(args, filter.To)
		argIdx++
	}
	if filter.AccountID != nil {
		query += fmt.Sprintf(` AND t.account_id = $%d`, argIdx)
		args = append(args, *filter.AccountID)
		argIdx++
	}
	if filter.CategoryID != nil {
		query += fmt.Sprintf(` AND t.category_id = $%d`, argIdx)
		args = append(args, *filter.CategoryID)
		argIdx++
	}

	query += ` ORDER BY t.occurred_at DESC, t.id DESC`
	if filter.Limit > 0 {
		query += fmt.Sprintf(` LIMIT $%d`, argIdx)
		args = append(args, filter.Limit)
		argIdx++
	}
	if filter.Offset > 0 {
		query += fmt.Sprintf(` OFFSET $%d`, argIdx)
		args = append(args, filter.Offset)
	}

	rows, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []*domain.Transaction
	for rows.Next() {
		var t domain.Transaction
		var catName *string
		if err := rows.Scan(&t.ID, &t.UserID, &t.AccountID, &t.AccountName, &t.CategoryID, &catName, &t.Amount, &t.Memo, &t.OccurredAt, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, err
		}
		t.CategoryName = catName
		transactions = append(transactions, &t)
	}
	return transactions, rows.Err()
}

// FindByIDAndUserID は id と userID で取引を取得する
func (r *TransactionRepositoryImpl) FindByIDAndUserID(ctx context.Context, id, userID int) (*domain.Transaction, error) {
	var t domain.Transaction
	var catName *string
	err := r.pool.QueryRow(ctx,
		`SELECT t.id, t.user_id, t.account_id, a.name, t.category_id, c.name, t.amount, t.memo, t.occurred_at::text, t.created_at::text, t.updated_at::text
		 FROM transactions t
		 JOIN accounts a ON t.account_id = a.id
		 LEFT JOIN categories c ON t.category_id = c.id
		 WHERE t.id = $1 AND t.user_id = $2`,
		id, userID,
	).Scan(&t.ID, &t.UserID, &t.AccountID, &t.AccountName, &t.CategoryID, &catName, &t.Amount, &t.Memo, &t.OccurredAt, &t.CreatedAt, &t.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	t.CategoryName = catName
	return &t, nil
}

// Create は新規取引を作成する。account/category の検証を含め 1 クエリで完結する。
// 検証失敗（account が存在しない or user に属さない、category が無効）時は (nil, nil) を返す。
func (r *TransactionRepositoryImpl) Create(ctx context.Context, userID, accountID int, categoryID *int, amount int, memo *string, occurredAt time.Time) (*domain.Transaction, error) {
	var t domain.Transaction
	var catName *string
	err := r.pool.QueryRow(ctx,
		`WITH ins AS (
			INSERT INTO transactions (user_id, account_id, category_id, amount, memo, occurred_at)
			SELECT $1, $2, $3, $4, $5, $6
			FROM accounts a
			WHERE a.id = $2 AND a.user_id = $1
			AND ($3::integer IS NULL OR EXISTS (SELECT 1 FROM categories c WHERE c.id = $3 AND c.user_id = $1))
			RETURNING id, user_id, account_id, category_id, amount, memo, occurred_at::text, created_at::text, updated_at::text
		)
		SELECT i.id, i.user_id, i.account_id, a.name, i.category_id, c.name, i.amount, i.memo, i.occurred_at, i.created_at, i.updated_at
		FROM ins i
		JOIN accounts a ON a.id = i.account_id
		LEFT JOIN categories c ON c.id = i.category_id`,
		userID, accountID, categoryID, amount, memo, occurredAt,
	).Scan(&t.ID, &t.UserID, &t.AccountID, &t.AccountName, &t.CategoryID, &catName, &t.Amount, &t.Memo, &t.OccurredAt, &t.CreatedAt, &t.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	t.CategoryName = catName
	return &t, nil
}

// Update は取引を更新する。存在チェック・account/category 検証・更新・取得を 1 クエリで完結する。
// 取引が存在しない、または検証失敗時は (nil, nil) を返す。
func (r *TransactionRepositoryImpl) Update(ctx context.Context, id, userID, accountID int, categoryID *int, amount int, memo *string, occurredAt time.Time) (*domain.Transaction, error) {
	var t domain.Transaction
	var catName *string
	err := r.pool.QueryRow(ctx,
		`WITH upd AS (
			UPDATE transactions t
			SET account_id = $1, category_id = $2, amount = $3, memo = $4, occurred_at = $5, updated_at = now()
			FROM accounts a
			WHERE t.id = $6 AND t.user_id = $7
			AND a.id = $1 AND a.user_id = $7
			AND ($2::integer IS NULL OR EXISTS (SELECT 1 FROM categories c WHERE c.id = $2 AND c.user_id = $7))
			RETURNING t.id, t.user_id, t.account_id, t.category_id, t.amount, t.memo, t.occurred_at, t.created_at, t.updated_at
		)
		SELECT u.id, u.user_id, u.account_id, a.name, u.category_id, c.name, u.amount, u.memo, u.occurred_at::text, u.created_at::text, u.updated_at::text
		FROM upd u
		JOIN accounts a ON a.id = u.account_id
		LEFT JOIN categories c ON c.id = u.category_id`,
		accountID, categoryID, amount, memo, occurredAt, id, userID,
	).Scan(&t.ID, &t.UserID, &t.AccountID, &t.AccountName, &t.CategoryID, &catName, &t.Amount, &t.Memo, &t.OccurredAt, &t.CreatedAt, &t.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	t.CategoryName = catName
	return &t, nil
}

// Delete は取引を削除する
func (r *TransactionRepositoryImpl) Delete(ctx context.Context, id, userID int) error {
	result, err := r.pool.Exec(ctx,
		`DELETE FROM transactions WHERE id = $1 AND user_id = $2`,
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
