package domain

// Goal は goals テーブルのエンティティ
type Goal struct {
	ID            int     `json:"id"`
	UserID        int     `json:"user_id"`
	Name          string  `json:"name"`
	TargetAmount  int     `json:"target_amount"`
	Deadline      *string `json:"deadline,omitempty"` // YYYY-MM-DD
	CreatedAt     string  `json:"created_at,omitempty"`
	UpdatedAt     string  `json:"updated_at,omitempty"`
}

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
		CreatedAt:    g.CreatedAt,
		UpdatedAt:    g.UpdatedAt,
		Deadline:     g.Deadline,
	}
}
