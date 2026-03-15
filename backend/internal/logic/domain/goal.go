package domain

import "time"

// Goal は goals テーブルのエンティティ（GORM + JSON 両対応）
type Goal struct {
	ID           int       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UserID       int       `gorm:"column:user_id;not null;index" json:"user_id"`
	Name         string    `gorm:"column:name;not null" json:"name"`
	TargetAmount int       `gorm:"column:target_amount;not null" json:"target_amount"`
	Deadline     *string   `gorm:"column:deadline" json:"deadline,omitempty"` // YYYY-MM-DD
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at,omitempty"`
	UpdatedAt    time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at,omitempty"`
}

// TableName は GORM のテーブル名
func (Goal) TableName() string { return "goals" }

// GoalCreateRequest は目標作成のリクエスト
type GoalCreateRequest struct {
	Name         string  `json:"name"`
	TargetAmount int     `json:"target_amount"`
	Deadline     *string `json:"deadline,omitempty"`
}

// GoalUpdateRequest は目標更新のリクエスト
type GoalUpdateRequest struct {
	Name         string  `json:"name"`
	TargetAmount int     `json:"target_amount"`
	Deadline     *string `json:"deadline,omitempty"`
}

// GoalResponse は API レスポンス用
type GoalResponse struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	TargetAmount int     `json:"target_amount"`
	Deadline     *string `json:"deadline,omitempty"`
	CreatedAt    string  `json:"created_at,omitempty"`
	UpdatedAt    string  `json:"updated_at,omitempty"`
}

// ToResponse は Goal を GoalResponse に変換する
func (g *Goal) ToResponse() *GoalResponse {
	return &GoalResponse{
		ID:           g.ID,
		Name:         g.Name,
		TargetAmount: g.TargetAmount,
		CreatedAt:    formatTime(g.CreatedAt),
		UpdatedAt:    formatTime(g.UpdatedAt),
		Deadline:     g.Deadline,
	}
}
