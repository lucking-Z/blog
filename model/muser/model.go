package muser

import "blogs/model"

type User struct {
	model.BaseModel
	Account  string `gorm:"column:account"`
	Name     string `gorm:"column:name"`
	PassWord string `gorm:"column:password"`
}

func (User) TableName() string {
	return "user"
}
