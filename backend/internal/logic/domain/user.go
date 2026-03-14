package domain

// User は users テーブルのエンティティ（DB の全カラム）
type User struct {
	ID           int    `json:"id"`
	UUID         string `json:"uuid"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"` // API には返さない
	CreatedAt    string `json:"created_at,omitempty"`
}

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
		CreatedAt: u.CreatedAt,
	}
}
