package usecase

import (
	"context"
	"regexp"

	"backend/internal/apperrors"
	"backend/internal/jwt"
	"backend/internal/logic/domain"
	"backend/internal/logic/repository"

	"golang.org/x/crypto/bcrypt"
)

const (
	bcryptCost     = 10
	minPasswordLen = 8
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

// AuthUsecase は認証のユースケース
type AuthUsecase interface {
	Signup(ctx context.Context, email, password string) (*domain.User, string, error)
	Login(ctx context.Context, email, password string) (*domain.User, string, error)
}

var _ AuthUsecase = (*AuthUsecaseImpl)(nil)

// AuthUsecaseImpl は AuthUsecase の実装
type AuthUsecaseImpl struct {
	jwtSecret string
	userRepo  repository.UserRepository
}

// NewAuthUsecase は AuthUsecaseImpl を生成する
func NewAuthUsecase(jwtSecret string, userRepo repository.UserRepository) *AuthUsecaseImpl {
	return &AuthUsecaseImpl{jwtSecret: jwtSecret, userRepo: userRepo}
}

// Signup はユーザー登録し、ユーザーとトークンを返す
func (a *AuthUsecaseImpl) Signup(ctx context.Context, email, password string) (*domain.User, string, error) {
	if !emailRegex.MatchString(email) {
		return nil, "", apperrors.ErrInvalidEmail
	}
	if len(password) < minPasswordLen {
		return nil, "", apperrors.ErrInvalidPassword
	}

	existing, err := a.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, "", err
	}
	if existing != nil {
		return nil, "", apperrors.ErrEmailExists
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	if err != nil {
		return nil, "", err
	}

	u, err := a.userRepo.Create(ctx, email, string(hash))
	if err != nil {
		return nil, "", err
	}

	token, err := jwt.Issue(a.jwtSecret, u)
	if err != nil {
		return nil, "", err
	}
	return u, token, nil
}

// Login はログインし、ユーザーとトークンを返す
func (a *AuthUsecaseImpl) Login(ctx context.Context, email, password string) (*domain.User, string, error) {
	u, hash, err := a.userRepo.FindByEmailWithPassword(ctx, email)
	if err != nil {
		return nil, "", err
	}
	if u == nil {
		return nil, "", apperrors.ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return nil, "", apperrors.ErrInvalidCredentials
	}

	token, err := jwt.Issue(a.jwtSecret, u)
	if err != nil {
		return nil, "", err
	}
	return u, token, nil
}
