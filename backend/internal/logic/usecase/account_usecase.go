package usecase

import (
	"context"
	"strings"

	"backend/internal/apperrors"
	"backend/internal/logic/domain"
	"backend/internal/logic/repository"
)

var validAccountTypes = map[string]bool{"cash": true, "bank": true, "credit": true}

// AccountUsecase は口座のユースケース
type AccountUsecase interface {
	List(ctx context.Context, userID int) ([]*domain.Account, error)
	Get(ctx context.Context, id, userID int) (*domain.Account, error)
	Create(ctx context.Context, userID int, accountType, name string) (*domain.Account, error)
	Update(ctx context.Context, id, userID int, accountType, name string) (*domain.Account, error)
	Delete(ctx context.Context, id, userID int) error
}

var _ AccountUsecase = (*AccountUsecaseImpl)(nil)

// AccountUsecaseImpl は AccountUsecase の実装
type AccountUsecaseImpl struct {
	accountRepo repository.AccountRepository
}

// NewAccountUsecase は AccountUsecaseImpl を生成する
func NewAccountUsecase(accountRepo repository.AccountRepository) *AccountUsecaseImpl {
	return &AccountUsecaseImpl{accountRepo: accountRepo}
}

func validateAccountType(t string) bool {
	return validAccountTypes[strings.ToLower(t)]
}

// List は userID に紐づく口座一覧を取得する
func (u *AccountUsecaseImpl) List(ctx context.Context, userID int) ([]*domain.Account, error) {
	return u.accountRepo.ListByUserID(ctx, userID)
}

// Get は id と userID で口座を取得する
func (u *AccountUsecaseImpl) Get(ctx context.Context, id, userID int) (*domain.Account, error) {
	acc, err := u.accountRepo.FindByIDAndUserID(ctx, id, userID)
	if err != nil {
		return nil, err
	}
	if acc == nil {
		return nil, apperrors.ErrAccountNotFound
	}
	return acc, nil
}

// Create は新規口座を作成する
func (u *AccountUsecaseImpl) Create(ctx context.Context, userID int, accountType, name string) (*domain.Account, error) {
	if !validateAccountType(accountType) {
		return nil, apperrors.ErrInvalidAccountType
	}
	if strings.TrimSpace(name) == "" {
		return nil, apperrors.ErrInvalidAccountName
	}

	typeID, err := u.accountRepo.GetAccountTypeIDByName(ctx, strings.ToLower(accountType))
	if err != nil || typeID == 0 {
		return nil, apperrors.ErrInvalidAccountType
	}

	return u.accountRepo.Create(ctx, userID, typeID, strings.TrimSpace(name))
}

// Update は口座を更新する
func (u *AccountUsecaseImpl) Update(ctx context.Context, id, userID int, accountType, name string) (*domain.Account, error) {
	acc, err := u.accountRepo.FindByIDAndUserID(ctx, id, userID)
	if err != nil {
		return nil, err
	}
	if acc == nil {
		return nil, apperrors.ErrAccountNotFound
	}

	if !validateAccountType(accountType) {
		return nil, apperrors.ErrInvalidAccountType
	}
	if strings.TrimSpace(name) == "" {
		return nil, apperrors.ErrInvalidAccountName
	}

	typeID, err := u.accountRepo.GetAccountTypeIDByName(ctx, strings.ToLower(accountType))
	if err != nil || typeID == 0 {
		return nil, apperrors.ErrInvalidAccountType
	}

	return u.accountRepo.Update(ctx, id, userID, typeID, strings.TrimSpace(name))
}

// Delete は口座を削除する
func (u *AccountUsecaseImpl) Delete(ctx context.Context, id, userID int) error {
	acc, err := u.accountRepo.FindByIDAndUserID(ctx, id, userID)
	if err != nil {
		return err
	}
	if acc == nil {
		return apperrors.ErrAccountNotFound
	}
	return u.accountRepo.Delete(ctx, id, userID)
}
