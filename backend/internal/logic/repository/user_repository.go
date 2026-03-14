package repository

import (
	"context"

	"backend/internal/logic/domain"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// UserRepository はユーザーの永続化インターフェース
type UserRepository interface {
	FindByID(ctx context.Context, id int) (*domain.User, error)
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	// FindByEmailWithPassword はログイン用。user と password_hash を返す
	FindByEmailWithPassword(ctx context.Context, email string) (*domain.User, string, error)
	Create(ctx context.Context, email, passwordHash string) (*domain.User, error)
}

var _ UserRepository = (*UserRepositoryImpl)(nil)

// UserRepositoryImpl は PostgreSQL 用の UserRepository 実装
type UserRepositoryImpl struct {
	pool *pgxpool.Pool
}

// NewUserRepository は UserRepositoryImpl を生成する
func NewUserRepository(pool *pgxpool.Pool) *UserRepositoryImpl {
	return &UserRepositoryImpl{pool: pool}
}

// FindByID は id でユーザーを取得する
func (r *UserRepositoryImpl) FindByID(ctx context.Context, id int) (*domain.User, error) {
	var u domain.User
	err := r.pool.QueryRow(ctx,
		`SELECT id, uuid::text, email, COALESCE(password_hash, ''), created_at::text FROM users WHERE id = $1`,
		id,
	).Scan(&u.ID, &u.UUID, &u.Email, &u.PasswordHash, &u.CreatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}

// FindByEmail は email でユーザーを取得する
func (r *UserRepositoryImpl) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	u, _, err := r.FindByEmailWithPassword(ctx, email)
	return u, err
}

// FindByEmailWithPassword はログイン用。user と password_hash を返す
func (r *UserRepositoryImpl) FindByEmailWithPassword(ctx context.Context, email string) (*domain.User, string, error) {
	var u domain.User
	err := r.pool.QueryRow(ctx,
		`SELECT id, uuid::text, email, COALESCE(password_hash, ''), created_at::text FROM users WHERE email = $1`,
		email,
	).Scan(&u.ID, &u.UUID, &u.Email, &u.PasswordHash, &u.CreatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, "", nil
		}
		return nil, "", err
	}
	return &u, u.PasswordHash, nil
}

// Create は新規ユーザーを作成する
func (r *UserRepositoryImpl) Create(ctx context.Context, email, passwordHash string) (*domain.User, error) {
	var u domain.User
	err := r.pool.QueryRow(ctx,
		`INSERT INTO users (email, password_hash) VALUES ($1, $2) RETURNING id, uuid::text, email, created_at::text`,
		email, passwordHash,
	).Scan(&u.ID, &u.UUID, &u.Email, &u.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
