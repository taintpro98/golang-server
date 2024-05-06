package model

type UserModel struct {
	ID    string `json:"id" gorm:"column:id"`
	Phone string `gorm:"column:phone,unique"`
	Email string `gorm:"column:email,unique"`
}

func (UserModel) TableName() string {
	return "users"
}
