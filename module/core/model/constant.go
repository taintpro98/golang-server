package model

type ConstantModel struct {
	Code  string `json:"code" gorm:"column:code"`
	Value string `json:"value" gorm:"column:value"`
}

func (ConstantModel) TableName() string {
	return "constants"
}
