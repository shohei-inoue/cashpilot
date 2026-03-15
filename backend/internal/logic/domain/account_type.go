package domain

// AccountType は account_types テーブル（口座種別マスタ）
type AccountType struct {
	ID   int    `gorm:"column:id;primaryKey" json:"id"`
	Name string `gorm:"column:name;not null" json:"name"`
}

// TableName は GORM のテーブル名
func (AccountType) TableName() string { return "account_types" }
