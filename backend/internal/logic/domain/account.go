package domain

import "time"

// Account は accounts テーブルのエンティティ（GORM + JSON 両対応）
type Account struct {
	ID            int          `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UserID        int          `gorm:"column:user_id;not null;index" json:"user_id"`
	AccountTypeID int          `gorm:"column:account_type_id;not null;index" json:"account_type_id"`
	AccountType   *AccountType `gorm:"foreignKey:AccountTypeID" json:"-"` // Preload 用。API では Type を使用
	Type          string       `gorm:"-" json:"type"`                     // cash | bank | credit（Preload 後にリポジトリで設定）
	Name          string       `gorm:"column:name;not null" json:"name"`
	CreatedAt     time.Time    `gorm:"column:created_at;autoCreateTime" json:"created_at,omitempty"`
	UpdatedAt     time.Time    `gorm:"column:updated_at;autoUpdateTime" json:"updated_at,omitempty"`
}

// TableName は GORM のテーブル名
func (Account) TableName() string { return "accounts" }

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
	typeName := a.Type
	if typeName == "" && a.AccountType != nil {
		typeName = a.AccountType.Name
	}
	return &AccountResponse{
		ID:        a.ID,
		Type:      typeName,
		Name:      a.Name,
		CreatedAt: formatTime(a.CreatedAt),
		UpdatedAt: formatTime(a.UpdatedAt),
	}
}
