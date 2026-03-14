package usecase

import (
	"context"
	"time"

	"backend/internal/apperrors"
	"backend/internal/logic/domain"
	"backend/internal/logic/repository"
)

// TransactionUsecase は取引のユースケース
type TransactionUsecase interface {
	List(ctx context.Context, userID int, filter repository.ListFilter) ([]*domain.Transaction, error)
	Get(ctx context.Context, id, userID int) (*domain.Transaction, error)
	Create(ctx context.Context, userID, accountID int, categoryID *int, amount int, memo *string, occurredAt time.Time) (*domain.Transaction, error)
	Update(ctx context.Context, id, userID, accountID int, categoryID *int, amount int, memo *string, occurredAt time.Time) (*domain.Transaction, error)
	Delete(ctx context.Context, id, userID int) error
}

var _ TransactionUsecase = (*TransactionUsecaseImpl)(nil)

// TransactionUsecaseImpl は TransactionUsecase の実装
type TransactionUsecaseImpl struct {
	transactionRepo repository.TransactionRepository
}

// NewTransactionUsecase は TransactionUsecaseImpl を生成する
func NewTransactionUsecase(transactionRepo repository.TransactionRepository) *TransactionUsecaseImpl {
	return &TransactionUsecaseImpl{transactionRepo: transactionRepo}
}

// List は userID に紐づく取引一覧を取得する
func (u *TransactionUsecaseImpl) List(ctx context.Context, userID int, filter repository.ListFilter) ([]*domain.Transaction, error) {
	return u.transactionRepo.List(ctx, userID, filter)
}

// Get は id と userID で取引を取得する
func (u *TransactionUsecaseImpl) Get(ctx context.Context, id, userID int) (*domain.Transaction, error) {
	tx, err := u.transactionRepo.FindByIDAndUserID(ctx, id, userID)
	if err != nil {
		return nil, err
	}
	if tx == nil {
		return nil, apperrors.ErrTransactionNotFound
	}
	return tx, nil
}

// Create は新規取引を作成する。amount の検証のみ usecase で行い、以降は 1 クエリで完結する。
func (u *TransactionUsecaseImpl) Create(ctx context.Context, userID, accountID int, categoryID *int, amount int, memo *string, occurredAt time.Time) (*domain.Transaction, error) {
	if amount == 0 {
		return nil, apperrors.ErrInvalidAmount
	}
	tx, err := u.transactionRepo.Create(ctx, userID, accountID, categoryID, amount, memo, occurredAt)
	if err != nil {
		return nil, err
	}
	if tx == nil {
		return nil, apperrors.ErrInvalidAccountRef
	}
	return tx, nil
}

// Update は取引を更新する。amount の検証のみ usecase で行い、以降は 1 クエリで完結する。
func (u *TransactionUsecaseImpl) Update(ctx context.Context, id, userID, accountID int, categoryID *int, amount int, memo *string, occurredAt time.Time) (*domain.Transaction, error) {
	if amount == 0 {
		return nil, apperrors.ErrInvalidAmount
	}
	tx, err := u.transactionRepo.Update(ctx, id, userID, accountID, categoryID, amount, memo, occurredAt)
	if err != nil {
		return nil, err
	}
	if tx == nil {
		return nil, apperrors.ErrTransactionNotFound
	}
	return tx, nil
}

// Delete は取引を削除する
func (u *TransactionUsecaseImpl) Delete(ctx context.Context, id, userID int) error {
	tx, err := u.transactionRepo.FindByIDAndUserID(ctx, id, userID)
	if err != nil {
		return err
	}
	if tx == nil {
		return apperrors.ErrTransactionNotFound
	}
	return u.transactionRepo.Delete(ctx, id, userID)
}
