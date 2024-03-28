package marticle

import (
	"time"
)

type Article struct {
	ID        uint64
	Title     string
	Content   string
	UserId    uint64
	Version   int64
	IsDelete  int8
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Article) TableName() string {
	return "article"
}
