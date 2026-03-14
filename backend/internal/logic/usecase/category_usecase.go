package usecase

import (
	"context"
	"strings"

	"backend/internal/apperrors"
	"backend/internal/logic/domain"
	"backend/internal/logic/repository"
)

var validCategoryTypes = map[string]bool{"income": true, "expense": true}

// CategoryUsecase はカテゴリのユースケース
type CategoryUsecase interface {
	List(ctx context.Context, userID int) ([]*domain.Category, error)
	Get(ctx context.Context, id, userID int) (*domain.Category, error)
	Create(ctx context.Context, userID int, categoryType, name string) (*domain.Category, error)
	Update(ctx context.Context, id, userID int, categoryType, name string) (*domain.Category, error)
	Delete(ctx context.Context, id, userID int) error
}

var _ CategoryUsecase = (*CategoryUsecaseImpl)(nil)

// CategoryUsecaseImpl は CategoryUsecase の実装
type CategoryUsecaseImpl struct {
	categoryRepo repository.CategoryRepository
}

// NewCategoryUsecase は CategoryUsecaseImpl を生成する
func NewCategoryUsecase(categoryRepo repository.CategoryRepository) *CategoryUsecaseImpl {
	return &CategoryUsecaseImpl{categoryRepo: categoryRepo}
}

func validateCategoryType(t string) bool {
	return validCategoryTypes[strings.ToLower(t)]
}

// List は userID に紐づくカテゴリ一覧を取得する
func (u *CategoryUsecaseImpl) List(ctx context.Context, userID int) ([]*domain.Category, error) {
	return u.categoryRepo.ListByUserID(ctx, userID)
}

// Get は id と userID でカテゴリを取得する
func (u *CategoryUsecaseImpl) Get(ctx context.Context, id, userID int) (*domain.Category, error) {
	cat, err := u.categoryRepo.FindByIDAndUserID(ctx, id, userID)
	if err != nil {
		return nil, err
	}
	if cat == nil {
		return nil, apperrors.ErrCategoryNotFound
	}
	return cat, nil
}

// Create は新規カテゴリを作成する
func (u *CategoryUsecaseImpl) Create(ctx context.Context, userID int, categoryType, name string) (*domain.Category, error) {
	if !validateCategoryType(categoryType) {
		return nil, apperrors.ErrInvalidCategoryType
	}
	if strings.TrimSpace(name) == "" {
		return nil, apperrors.ErrInvalidCategoryName
	}

	typeID, err := u.categoryRepo.GetCategoryTypeIDByName(ctx, strings.ToLower(categoryType))
	if err != nil || typeID == 0 {
		return nil, apperrors.ErrInvalidCategoryType
	}

	return u.categoryRepo.Create(ctx, userID, typeID, strings.TrimSpace(name))
}

// Update はカテゴリを更新する
func (u *CategoryUsecaseImpl) Update(ctx context.Context, id, userID int, categoryType, name string) (*domain.Category, error) {
	cat, err := u.categoryRepo.FindByIDAndUserID(ctx, id, userID)
	if err != nil {
		return nil, err
	}
	if cat == nil {
		return nil, apperrors.ErrCategoryNotFound
	}

	if !validateCategoryType(categoryType) {
		return nil, apperrors.ErrInvalidCategoryType
	}
	if strings.TrimSpace(name) == "" {
		return nil, apperrors.ErrInvalidCategoryName
	}

	typeID, err := u.categoryRepo.GetCategoryTypeIDByName(ctx, strings.ToLower(categoryType))
	if err != nil || typeID == 0 {
		return nil, apperrors.ErrInvalidCategoryType
	}

	return u.categoryRepo.Update(ctx, id, userID, typeID, strings.TrimSpace(name))
}

// Delete はカテゴリを削除する
func (u *CategoryUsecaseImpl) Delete(ctx context.Context, id, userID int) error {
	cat, err := u.categoryRepo.FindByIDAndUserID(ctx, id, userID)
	if err != nil {
		return err
	}
	if cat == nil {
		return apperrors.ErrCategoryNotFound
	}
	return u.categoryRepo.Delete(ctx, id, userID)
}
