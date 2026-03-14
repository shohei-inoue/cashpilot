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
	ErrTransactionNotFound  = errors.New("transaction not found")
	ErrInvalidAmount        = errors.New("amount must not be zero")
	ErrInvalidAccountRef    = errors.New("account does not belong to user")
	ErrInvalidCategoryRef   = errors.New("category does not belong to user")
	ErrGoalNotFound         = errors.New("goal not found")
	ErrInvalidGoalName      = errors.New("goal name is required")
	ErrInvalidTargetAmount  = errors.New("target amount must be positive")
)
