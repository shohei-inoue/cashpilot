package domain

// Account は accounts テーブルのエンティティ
type Account struct {
	ID            int    `json:"id"`
	UserID        int    `json:"user_id"`
	AccountTypeID int    `json:"account_type_id"`
	AccountType   string `json:"type"` // cash | bank | credit
	Name          string `json:"name"`
	CreatedAt     string `json:"created_at,omitempty"`
	UpdatedAt     string `json:"updated_at,omitempty"`
}

// AccountCreateRequest は口座作成のリクエスト
type AccountCreateRequest struct {
	Type string `json:"type"` // cash | bank | credit
	Name string `json:"name"`
}

// AccountUpdateRequest は口座更新のリクエスト
type AccountUpdateRequest struct {
	Type string `json:"type"` // cash | bank | credit
	Name string `json:"name"`
}

// AccountResponse は API レスポンス用
type AccountResponse struct {
	ID        int    `json:"id"`
	Type      string `json:"type"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

// ToResponse は Account を AccountResponse に変換する
func (a *Account) ToResponse() *AccountResponse {
	return &AccountResponse{
		ID:        a.ID,
		Type:      a.AccountType,
		Name:      a.Name,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
	}
}
