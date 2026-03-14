package apperrors

import "errors"

// ビジネスロジックで返す sentinel エラー
var (
	ErrInvalidEmail       = errors.New("invalid email format")
	ErrInvalidPassword    = errors.New("invalid password")
	ErrEmailExists        = errors.New("email already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserNotFound       = errors.New("user not found")
	ErrAccountNotFound    = errors.New("account not found")
	ErrInvalidAccountType = errors.New("invalid account type")
	ErrInvalidAccountName   = errors.New("account name is required")
	ErrCategoryNotFound     = errors.New("category not found")
	ErrInvalidCategoryType  = errors.New("invalid category type")
	ErrInvalidCategoryName  = errors.New("category name is required")
)
