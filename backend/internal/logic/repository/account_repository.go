package repository

import (
	"context"

	"backend/internal/logic/domain"

	"gorm.io/gorm"
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

// AccountRepositoryImpl は GORM 用の AccountRepository 実装
type AccountRepositoryImpl struct {
	db *gorm.DB
}

// NewAccountRepository は AccountRepositoryImpl を生成する
func NewAccountRepository(db *gorm.DB) *AccountRepositoryImpl {
	return &AccountRepositoryImpl{db: db}
}

func setAccountTypeName(a *domain.Account) {
	if a.AccountType != nil {
		a.Type = a.AccountType.Name
	}
}

// GetAccountTypeIDByName は account_types の name から id を取得する
func (r *AccountRepositoryImpl) GetAccountTypeIDByName(ctx context.Context, name string) (int, error) {
	var at domain.AccountType
	err := r.db.WithContext(ctx).Where("name = ?", name).First(&at).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return 0, nil
		}
		return 0, err
	}
	return at.ID, nil
}

// ListByUserID は userID に紐づく口座一覧を取得する
func (r *AccountRepositoryImpl) ListByUserID(ctx context.Context, userID int) ([]*domain.Account, error) {
	var list []*domain.Account
	err := r.db.WithContext(ctx).
		Preload("AccountType").
		Where("user_id = ?", userID).
		Order("id").
		Find(&list).Error
	if err != nil {
		return nil, err
	}
	for _, a := range list {
		setAccountTypeName(a)
	}
	return list, nil
}

// FindByIDAndUserID は id と userID で口座を取得する
func (r *AccountRepositoryImpl) FindByIDAndUserID(ctx context.Context, id, userID int) (*domain.Account, error) {
	var a domain.Account
	err := r.db.WithContext(ctx).
		Preload("AccountType").
		Where("id = ? AND user_id = ?", id, userID).
		First(&a).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	setAccountTypeName(&a)
	return &a, nil
}

// Create は新規口座を作成する
func (r *AccountRepositoryImpl) Create(ctx context.Context, userID, accountTypeID int, name string) (*domain.Account, error) {
	a := domain.Account{
		UserID:        userID,
		AccountTypeID: accountTypeID,
		Name:          name,
	}
	if err := r.db.WithContext(ctx).Create(&a).Error; err != nil {
		return nil, err
	}
	return r.FindByIDAndUserID(ctx, a.ID, userID)
}

// Update は口座を更新する
func (r *AccountRepositoryImpl) Update(ctx context.Context, id, userID, accountTypeID int, name string) (*domain.Account, error) {
	result := r.db.WithContext(ctx).Model(&domain.Account{}).
		Where("id = ? AND user_id = ?", id, userID).
		Updates(map[string]interface{}{
			"account_type_id": accountTypeID,
			"name":            name,
		})
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return r.FindByIDAndUserID(ctx, id, userID)
}

// Delete は口座を削除する
func (r *AccountRepositoryImpl) Delete(ctx context.Context, id, userID int) error {
	result := r.db.WithContext(ctx).
		Where("id = ? AND user_id = ?", id, userID).
		Delete(&domain.Account{})
	return result.Error
}
