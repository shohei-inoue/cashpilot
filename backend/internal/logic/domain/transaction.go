package domain

// Transaction は transactions テーブルのエンティティ
type Transaction struct {
	ID          int     `json:"id"`
	UserID      int     `json:"user_id"`
	AccountID   int     `json:"account_id"`
	AccountName string  `json:"account_name,omitempty"`
	CategoryID  *int    `json:"category_id,omitempty"`
	CategoryName *string `json:"category_name,omitempty"`
	Amount      int     `json:"amount"`
	Memo        *string `json:"memo,omitempty"`
	OccurredAt  string  `json:"occurred_at"`
	CreatedAt   string  `json:"created_at,omitempty"`
	UpdatedAt   string  `json:"updated_at,omitempty"`
}

// TransactionCreateRequest は取引作成のリクエスト
type TransactionCreateRequest struct {
	AccountID   int     `json:"account_id"`
	CategoryID  *int    `json:"category_id,omitempty"`
	Amount      int     `json:"amount"`
	Memo        *string `json:"memo,omitempty"`
	OccurredAt  string  `json:"occurred_at"`
}

// TransactionUpdateRequest は取引更新のリクエスト
type TransactionUpdateRequest struct {
	AccountID   int     `json:"account_id"`
	CategoryID  *int    `json:"category_id,omitempty"`
	Amount      int     `json:"amount"`
	Memo        *string `json:"memo,omitempty"`
	OccurredAt  string  `json:"occurred_at"`
}

// TransactionResponse は API レスポンス用
type TransactionResponse struct {
	ID          int     `json:"id"`
	AccountID   int     `json:"account_id"`
	AccountName string  `json:"account_name,omitempty"`
	CategoryID  *int    `json:"category_id,omitempty"`
	CategoryName *string `json:"category_name,omitempty"`
	Amount      int     `json:"amount"`
	Memo        *string `json:"memo,omitempty"`
	OccurredAt  string  `json:"occurred_at"`
	CreatedAt   string  `json:"created_at,omitempty"`
	UpdatedAt   string  `json:"updated_at,omitempty"`
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
		OccurredAt:   t.OccurredAt,
		CreatedAt:    t.CreatedAt,
		UpdatedAt:    t.UpdatedAt,
	}
}
