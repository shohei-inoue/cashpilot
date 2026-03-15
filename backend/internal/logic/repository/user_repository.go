package repository

import (
	"context"

	"backend/internal/logic/domain"

	"gorm.io/gorm"
)

// UserRepository はユーザーの永続化インターフェース
type UserRepository interface {
	FindByID(ctx context.Context, id int) (*domain.User, error)
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	FindByEmailWithPassword(ctx context.Context, email string) (*domain.User, string, error)
	Create(ctx context.Context, email, passwordHash string) (*domain.User, error)
}

var _ UserRepository = (*UserRepositoryImpl)(nil)

// UserRepositoryImpl は GORM 用の UserRepository 実装
type UserRepositoryImpl struct {
	db *gorm.DB
}

// NewUserRepository は UserRepositoryImpl を生成する
func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

// FindByID は id でユーザーを取得する
func (r *UserRepositoryImpl) FindByID(ctx context.Context, id int) (*domain.User, error) {
	var u domain.User
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&u).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
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
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&u).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, "", nil
		}
		return nil, "", err
	}
	return &u, u.PasswordHash, nil
}

// Create は新規ユーザーを作成する
func (r *UserRepositoryImpl) Create(ctx context.Context, email, passwordHash string) (*domain.User, error) {
	u := domain.User{
		Email:        email,
		PasswordHash: passwordHash,
	}
	if err := r.db.WithContext(ctx).Create(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}
