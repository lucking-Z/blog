package muser

import "time"

type User struct {
	Account   string `gorm:"column:account"`
	Name      string `gorm:"column:name"`
	PassWord  string `gorm:"column:password"`
	Version   int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (User) TableName() string {
	return "user"
}
