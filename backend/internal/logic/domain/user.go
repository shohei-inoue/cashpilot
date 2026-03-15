package domain

import "time"

// User は users テーブルのエンティティ（GORM + JSON 両対応）
type User struct {
	ID           int       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UUID         string    `gorm:"column:uuid;type:uuid;default:gen_random_uuid()" json:"uuid"`
	Email        string    `gorm:"column:email;uniqueIndex;not null" json:"email"`
	PasswordHash string    `gorm:"column:password_hash" json:"-"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at,omitempty"`
}

// TableName は GORM のテーブル名
func (User) TableName() string { return "users" }

// UserCreateRequest はユーザー作成（サインアップ）のリクエスト
type UserCreateRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UserLoginRequest はログインのリクエスト
type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UserResponse は API レスポンス用（password_hash, uuid は含めない）
type UserResponse struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at,omitempty"`
}

// ToResponse は User を UserResponse に変換する
func (u *User) ToResponse() *UserResponse {
	return &UserResponse{
		ID:        u.ID,
		Email:     u.Email,
		CreatedAt: formatTime(u.CreatedAt),
	}
}
