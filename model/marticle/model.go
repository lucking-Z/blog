package marticle

import "blogs/model"

type Article struct {
	model.BaseModel
	Title   string `gorm:"column:title"`
	Content string `gorm:"column:content"`
	UserId  uint   `gorm:"column:user_id"`
}

func (Article) TableName() string {
	return "article"
}