package domain

// CategoryType は category_types テーブル（カテゴリ種別マスタ）
type CategoryType struct {
	ID   int    `gorm:"column:id;primaryKey" json:"id"`
	Name string `gorm:"column:name;not null" json:"name"`
}

// TableName は GORM のテーブル名
func (CategoryType) TableName() string { return "category_types" }
