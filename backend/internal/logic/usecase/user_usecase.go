package usecase

import (
	"context"

	"backend/internal/apperrors"
	"backend/internal/logic/domain"
	"backend/internal/logic/repository"
)

// UserUsecase はユーザー取得のユースケース
type UserUsecase interface {
	GetUser(ctx context.Context, userID int) (*domain.User, error)
}

var _ UserUsecase = (*UserUsecaseImpl)(nil)

// UserUsecaseImpl は UserUsecase の実装
type UserUsecaseImpl struct {
	userRepo repository.UserRepository
}

// NewUserUsecase は UserUsecaseImpl を生成する
func NewUserUsecase(userRepo repository.UserRepository) *UserUsecaseImpl {
	return &UserUsecaseImpl{userRepo: userRepo}
}

// GetUser は userID でユーザーを取得する
func (u *UserUsecaseImpl) GetUser(ctx context.Context, userID int) (*domain.User, error) {
	user, err := u.userRepo.FindByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, apperrors.ErrUserNotFound
	}
	return user, nil
}
