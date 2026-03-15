package domain

import "time"

// Category は categories テーブルのエンティティ（GORM + JSON 両対応）
type Category struct {
	ID             int           `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UserID         int           `gorm:"column:user_id;not null;index" json:"user_id"`
	CategoryTypeID int           `gorm:"column:category_type_id;not null;index" json:"category_type_id"`
	CategoryType   *CategoryType `gorm:"foreignKey:CategoryTypeID" json:"-"` // Preload 用。API では Type を使用
	Type           string        `gorm:"-" json:"type"`                      // income | expense（Preload 後にリポジトリで設定）
	Name           string        `gorm:"column:name;not null" json:"name"`
	CreatedAt      time.Time     `gorm:"column:created_at;autoCreateTime" json:"created_at,omitempty"`
	UpdatedAt      time.Time     `gorm:"column:updated_at;autoUpdateTime" json:"updated_at,omitempty"`
}

// TableName は GORM のテーブル名
func (Category) TableName() string { return "categories" }

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
	typeName := c.Type
	if typeName == "" && c.CategoryType != nil {
		typeName = c.CategoryType.Name
	}
	return &CategoryResponse{
		ID:        c.ID,
		Type:      typeName,
		Name:      c.Name,
		CreatedAt: formatTime(c.CreatedAt),
		UpdatedAt: formatTime(c.UpdatedAt),
	}
}
