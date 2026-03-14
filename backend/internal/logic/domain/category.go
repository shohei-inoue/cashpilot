package domain

// Category は categories テーブルのエンティティ
type Category struct {
	ID             int    `json:"id"`
	UserID         int    `json:"user_id"`
	CategoryTypeID int    `json:"category_type_id"`
	CategoryType   string `json:"type"` // income | expense
	Name           string `json:"name"`
	CreatedAt      string `json:"created_at,omitempty"`
	UpdatedAt      string `json:"updated_at,omitempty"`
}

// CategoryCreateRequest はカテゴリ作成のリクエスト
type CategoryCreateRequest struct {
	Type string `json:"type"` // income | expense
	Name string `json:"name"`
}

// CategoryUpdateRequest はカテゴリ更新のリクエスト
type CategoryUpdateRequest struct {
	Type string `json:"type"` // income | expense
	Name string `json:"name"`
}

// CategoryResponse は API レスポンス用
type CategoryResponse struct {
	ID        int    `json:"id"`
	Type      string `json:"type"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

// ToResponse は Category を CategoryResponse に変換する
func (c *Category) ToResponse() *CategoryResponse {
	return &CategoryResponse{
		ID:        c.ID,
		Type:      c.CategoryType,
		Name:      c.Name,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}
