package jwt

import (
	"fmt"
	"time"

	"backend/internal/logic/domain"

	"github.com/golang-jwt/jwt/v5"
)

const (
	TokenExpiry = 24 * time.Hour
	CookieName  = "token"
)

// Claims は JWT のペイロード
type Claims struct {
	Sub   string `json:"sub"`   // ユーザー ID（文字列）
	Email string `json:"email"`
	jwt.RegisteredClaims
}

// Issue は JWT を発行する
func Issue(secret string, u *domain.User) (string, error) {
	if secret == "" {
		return "", fmt.Errorf("jwt: JWT_SECRET is required")
	}
	now := time.Now()
	claims := Claims{
		Sub:   fmt.Sprintf("%d", u.ID),
		Email: u.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(TokenExpiry)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// Parse は Cookie や文字列から JWT を検証し、sub（user_id）と email を返す
func Parse(secret, tokenString string) (userID int, email string, err error) {
	if secret == "" {
		return 0, "", fmt.Errorf("jwt: JWT_SECRET is required")
	}
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return 0, "", err
	}
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return 0, "", fmt.Errorf("jwt: invalid token")
	}
	var id int
	if _, err := fmt.Sscanf(claims.Sub, "%d", &id); err != nil {
		return 0, "", fmt.Errorf("jwt: invalid sub")
	}
	return id, claims.Email, nil
}
