package repository

import (
	"context"

	"backend/internal/logic/domain"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// AccountRepository は口座の永続化インターフェース
type AccountRepository interface {
	ListByUserID(ctx context.Context, userID int) ([]*domain.Account, error)
	FindByIDAndUserID(ctx context.Context, id, userID int) (*domain.Account, error)
	Create(ctx context.Context, userID, accountTypeID int, name string) (*domain.Account, error)
	Update(ctx context.Context, id, userID, accountTypeID int, name string) (*domain.Account, error)
	Delete(ctx context.Context, id, userID int) error
	GetAccountTypeIDByName(ctx context.Context, name string) (int, error)
}

var _ AccountRepository = (*AccountRepositoryImpl)(nil)

// AccountRepositoryImpl は PostgreSQL 用の AccountRepository 実装
type AccountRepositoryImpl struct {
	pool *pgxpool.Pool
}

// NewAccountRepository は AccountRepositoryImpl を生成する
func NewAccountRepository(pool *pgxpool.Pool) *AccountRepositoryImpl {
	return &AccountRepositoryImpl{pool: pool}
}

// GetAccountTypeIDByName は account_types の name から id を取得する
func (r *AccountRepositoryImpl) GetAccountTypeIDByName(ctx context.Context, name string) (int, error) {
	var id int
	err := r.pool.QueryRow(ctx,
		`SELECT id FROM account_types WHERE name = $1`,
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

// ListByUserID は userID に紐づく口座一覧を取得する
func (r *AccountRepositoryImpl) ListByUserID(ctx context.Context, userID int) ([]*domain.Account, error) {
	rows, err := r.pool.Query(ctx,
		`SELECT a.id, a.user_id, a.account_type_id, at.name, a.name, a.created_at::text, a.updated_at::text
		 FROM accounts a
		 JOIN account_types at ON a.account_type_id = at.id
		 WHERE a.user_id = $1
		 ORDER BY a.id`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var accounts []*domain.Account
	for rows.Next() {
		var a domain.Account
		if err := rows.Scan(&a.ID, &a.UserID, &a.AccountTypeID, &a.AccountType, &a.Name, &a.CreatedAt, &a.UpdatedAt); err != nil {
			return nil, err
		}
		accounts = append(accounts, &a)
	}
	return accounts, rows.Err()
}

// FindByIDAndUserID は id と userID で口座を取得する
func (r *AccountRepositoryImpl) FindByIDAndUserID(ctx context.Context, id, userID int) (*domain.Account, error) {
	var a domain.Account
	err := r.pool.QueryRow(ctx,
		`SELECT a.id, a.user_id, a.account_type_id, at.name, a.name, a.created_at::text, a.updated_at::text
		 FROM accounts a
		 JOIN account_types at ON a.account_type_id = at.id
		 WHERE a.id = $1 AND a.user_id = $2`,
		id, userID,
	).Scan(&a.ID, &a.UserID, &a.AccountTypeID, &a.AccountType, &a.Name, &a.CreatedAt, &a.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &a, nil
}

// Create は新規口座を作成する
func (r *AccountRepositoryImpl) Create(ctx context.Context, userID, accountTypeID int, name string) (*domain.Account, error) {
	var id int
	err := r.pool.QueryRow(ctx,
		`INSERT INTO accounts (user_id, account_type_id, name)
		 VALUES ($1, $2, $3)
		 RETURNING id`,
		userID, accountTypeID, name,
	).Scan(&id)
	if err != nil {
		return nil, err
	}
	return r.FindByIDAndUserID(ctx, id, userID)
}

// Update は口座を更新する
func (r *AccountRepositoryImpl) Update(ctx context.Context, id, userID, accountTypeID int, name string) (*domain.Account, error) {
	result, err := r.pool.Exec(ctx,
		`UPDATE accounts SET account_type_id = $1, name = $2, updated_at = now()
		 WHERE id = $3 AND user_id = $4`,
		accountTypeID, name, id, userID,
	)
	if err != nil {
		return nil, err
	}
	if result.RowsAffected() == 0 {
		return nil, nil
	}
	return r.FindByIDAndUserID(ctx, id, userID)
}

// Delete は口座を削除する
func (r *AccountRepositoryImpl) Delete(ctx context.Context, id, userID int) error {
	result, err := r.pool.Exec(ctx,
		`DELETE FROM accounts WHERE id = $1 AND user_id = $2`,
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
