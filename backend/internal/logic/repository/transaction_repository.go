package repository

import (
	"context"
	"time"

	"backend/internal/logic/domain"

	"gorm.io/gorm"
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

// TransactionRepositoryImpl は GORM 用の TransactionRepository 実装
type TransactionRepositoryImpl struct {
	db *gorm.DB
}

// NewTransactionRepository は TransactionRepositoryImpl を生成する
func NewTransactionRepository(db *gorm.DB) *TransactionRepositoryImpl {
	return &TransactionRepositoryImpl{db: db}
}

func (r *TransactionRepositoryImpl) listQuery(db *gorm.DB, userID int, filter ListFilter) *gorm.DB {
	q := db.Table("transactions").
		Select("transactions.id, transactions.user_id, transactions.account_id, a.name AS account_name, transactions.category_id, c.name AS category_name, transactions.amount, transactions.memo, transactions.occurred_at, transactions.created_at, transactions.updated_at").
		Joins("JOIN accounts a ON a.id = transactions.account_id").
		Joins("LEFT JOIN categories c ON c.id = transactions.category_id").
		Where("transactions.user_id = ?", userID)

	if filter.From != nil {
		q = q.Where("transactions.occurred_at >= ?", filter.From)
	}
	if filter.To != nil {
		q = q.Where("transactions.occurred_at <= ?", filter.To)
	}
	if filter.AccountID != nil {
		q = q.Where("transactions.account_id = ?", *filter.AccountID)
	}
	if filter.CategoryID != nil {
		q = q.Where("transactions.category_id = ?", *filter.CategoryID)
	}
	q = q.Order("transactions.occurred_at DESC, transactions.id DESC")
	if filter.Limit > 0 {
		q = q.Limit(filter.Limit)
	}
	if filter.Offset > 0 {
		q = q.Offset(filter.Offset)
	}
	return q
}

// List は userID に紐づく取引一覧を取得する
func (r *TransactionRepositoryImpl) List(ctx context.Context, userID int, filter ListFilter) ([]*domain.Transaction, error) {
	var list []*domain.Transaction
	err := r.listQuery(r.db.WithContext(ctx), userID, filter).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

// FindByIDAndUserID は id と userID で取引を取得する
func (r *TransactionRepositoryImpl) FindByIDAndUserID(ctx context.Context, id, userID int) (*domain.Transaction, error) {
	var t domain.Transaction
	err := r.db.WithContext(ctx).Table("transactions").
		Select("transactions.id, transactions.user_id, transactions.account_id, a.name AS account_name, transactions.category_id, c.name AS category_name, transactions.amount, transactions.memo, transactions.occurred_at, transactions.created_at, transactions.updated_at").
		Joins("JOIN accounts a ON a.id = transactions.account_id").
		Joins("LEFT JOIN categories c ON c.id = transactions.category_id").
		Where("transactions.id = ? AND transactions.user_id = ?", id, userID).
		First(&t).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &t, nil
}

// Create は新規取引を作成する。account/category の検証を行う。
func (r *TransactionRepositoryImpl) Create(ctx context.Context, userID, accountID int, categoryID *int, amount int, memo *string, occurredAt time.Time) (*domain.Transaction, error) {
	var acc domain.Account
	if err := r.db.WithContext(ctx).Where("id = ? AND user_id = ?", accountID, userID).First(&acc).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	if categoryID != nil {
		var cat domain.Category
		if err := r.db.WithContext(ctx).Where("id = ? AND user_id = ?", *categoryID, userID).First(&cat).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, nil
			}
			return nil, err
		}
	}
	t := domain.Transaction{
		UserID:     userID,
		AccountID:  accountID,
		CategoryID: categoryID,
		Amount:     amount,
		Memo:       memo,
		OccurredAt: occurredAt,
	}
	if err := r.db.WithContext(ctx).Create(&t).Error; err != nil {
		return nil, err
	}
	return r.FindByIDAndUserID(ctx, t.ID, userID)
}

// Update は取引を更新する。
func (r *TransactionRepositoryImpl) Update(ctx context.Context, id, userID, accountID int, categoryID *int, amount int, memo *string, occurredAt time.Time) (*domain.Transaction, error) {
	var existing domain.Transaction
	if err := r.db.WithContext(ctx).Where("id = ? AND user_id = ?", id, userID).First(&existing).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	var acc domain.Account
	if err := r.db.WithContext(ctx).Where("id = ? AND user_id = ?", accountID, userID).First(&acc).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	if categoryID != nil {
		var cat domain.Category
		if err := r.db.WithContext(ctx).Where("id = ? AND user_id = ?", *categoryID, userID).First(&cat).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, nil
			}
			return nil, err
		}
	}
	updates := map[string]interface{}{
		"account_id":  accountID,
		"category_id": categoryID,
		"amount":      amount,
		"memo":        memo,
		"occurred_at": occurredAt,
	}
	if err := r.db.WithContext(ctx).Model(&domain.Transaction{}).Where("id = ? AND user_id = ?", id, userID).Updates(updates).Error; err != nil {
		return nil, err
	}
	return r.FindByIDAndUserID(ctx, id, userID)
}

// Delete は取引を削除する
func (r *TransactionRepositoryImpl) Delete(ctx context.Context, id, userID int) error {
	result := r.db.WithContext(ctx).Where("id = ? AND user_id = ?", id, userID).Delete(&domain.Transaction{})
	return result.Error
}
