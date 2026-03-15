package repository

import (
	"context"

	"backend/internal/logic/domain"

	"gorm.io/gorm"
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

// CategoryRepositoryImpl は GORM 用の CategoryRepository 実装
type CategoryRepositoryImpl struct {
	db *gorm.DB
}

// NewCategoryRepository は CategoryRepositoryImpl を生成する
func NewCategoryRepository(db *gorm.DB) *CategoryRepositoryImpl {
	return &CategoryRepositoryImpl{db: db}
}

func setCategoryTypeName(c *domain.Category) {
	if c.CategoryType != nil {
		c.Type = c.CategoryType.Name
	}
}

// GetCategoryTypeIDByName は category_types の name から id を取得する
func (r *CategoryRepositoryImpl) GetCategoryTypeIDByName(ctx context.Context, name string) (int, error) {
	var ct domain.CategoryType
	err := r.db.WithContext(ctx).Where("name = ?", name).First(&ct).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return 0, nil
		}
		return 0, err
	}
	return ct.ID, nil
}

// ListByUserID は userID に紐づくカテゴリ一覧を取得する
func (r *CategoryRepositoryImpl) ListByUserID(ctx context.Context, userID int) ([]*domain.Category, error) {
	var list []*domain.Category
	err := r.db.WithContext(ctx).
		Preload("CategoryType").
		Where("user_id = ?", userID).
		Order("id").
		Find(&list).Error
	if err != nil {
		return nil, err
	}
	for _, c := range list {
		setCategoryTypeName(c)
	}
	return list, nil
}

// FindByIDAndUserID は id と userID でカテゴリを取得する
func (r *CategoryRepositoryImpl) FindByIDAndUserID(ctx context.Context, id, userID int) (*domain.Category, error) {
	var c domain.Category
	err := r.db.WithContext(ctx).
		Preload("CategoryType").
		Where("id = ? AND user_id = ?", id, userID).
		First(&c).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	setCategoryTypeName(&c)
	return &c, nil
}

// Create は新規カテゴリを作成する
func (r *CategoryRepositoryImpl) Create(ctx context.Context, userID, categoryTypeID int, name string) (*domain.Category, error) {
	c := domain.Category{
		UserID:         userID,
		CategoryTypeID: categoryTypeID,
		Name:           name,
	}
	if err := r.db.WithContext(ctx).Create(&c).Error; err != nil {
		return nil, err
	}
	return r.FindByIDAndUserID(ctx, c.ID, userID)
}

// Update はカテゴリを更新する
func (r *CategoryRepositoryImpl) Update(ctx context.Context, id, userID, categoryTypeID int, name string) (*domain.Category, error) {
	result := r.db.WithContext(ctx).Model(&domain.Category{}).
		Where("id = ? AND user_id = ?", id, userID).
		Updates(map[string]interface{}{
			"category_type_id": categoryTypeID,
			"name":             name,
		})
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return r.FindByIDAndUserID(ctx, id, userID)
}

// Delete はカテゴリを削除する
func (r *CategoryRepositoryImpl) Delete(ctx context.Context, id, userID int) error {
	result := r.db.WithContext(ctx).
		Where("id = ? AND user_id = ?", id, userID).
		Delete(&domain.Category{})
	return result.Error
}
