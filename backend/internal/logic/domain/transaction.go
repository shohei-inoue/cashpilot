package domain

import "time"

// Transaction は transactions テーブルのエンティティ（GORM + JSON 両対応）
// List/Find で JOIN する場合は account_name, category_name が SELECT で埋まる
type Transaction struct {
	ID           int       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UserID       int       `gorm:"column:user_id;not null;index" json:"user_id"`
	AccountID    int       `gorm:"column:account_id;not null;index" json:"account_id"`
	AccountName  string    `gorm:"column:account_name;->" json:"account_name,omitempty"` // JOIN 時のみ
	CategoryID   *int      `gorm:"column:category_id;index" json:"category_id,omitempty"`
	CategoryName *string   `gorm:"column:category_name;->" json:"category_name,omitempty"` // JOIN 時のみ
	Amount       int       `gorm:"column:amount;not null" json:"amount"`
	Memo         *string   `gorm:"column:memo" json:"memo,omitempty"`
	OccurredAt   time.Time `gorm:"column:occurred_at;not null;index" json:"occurred_at"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at,omitempty"`
	UpdatedAt    time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at,omitempty"`
}

// TableName は GORM のテーブル名
func (Transaction) TableName() string { return "transactions" }

// TransactionCreateRequest は取引作成のリクエスト
type TransactionCreateRequest struct {
	AccountID  int     `json:"account_id"`
	CategoryID *int    `json:"category_id,omitempty"`
	Amount     int     `json:"amount"`
	Memo       *string `json:"memo,omitempty"`
	OccurredAt string  `json:"occurred_at"`
}

// TransactionUpdateRequest は取引更新のリクエスト
type TransactionUpdateRequest struct {
	AccountID  int     `json:"account_id"`
	CategoryID *int    `json:"category_id,omitempty"`
	Amount     int     `json:"amount"`
	Memo       *string `json:"memo,omitempty"`
	OccurredAt string  `json:"occurred_at"`
}

// TransactionResponse は API レスポンス用
type TransactionResponse struct {
	ID           int     `json:"id"`
	AccountID    int     `json:"account_id"`
	AccountName  string  `json:"account_name,omitempty"`
	CategoryID   *int    `json:"category_id,omitempty"`
	CategoryName *string `json:"category_name,omitempty"`
	Amount       int     `json:"amount"`
	Memo         *string `json:"memo,omitempty"`
	OccurredAt   string  `json:"occurred_at"`
	CreatedAt    string  `json:"created_at,omitempty"`
	UpdatedAt    string  `json:"updated_at,omitempty"`
}

// ToResponse は Transaction を TransactionResponse に変換する
func (t *Transaction) ToResponse() *TransactionResponse {
	return &TransactionResponse{
		ID:           t.ID,
		AccountID:    t.AccountID,
		AccountName:  t.AccountName,
		CategoryID:   t.CategoryID,
		CategoryName: t.CategoryName,
		Amount:       t.Amount,
		Memo:         t.Memo,
		OccurredAt:   formatTime(t.OccurredAt),
		CreatedAt:    formatTime(t.CreatedAt),
		UpdatedAt:    formatTime(t.UpdatedAt),
	}
}
