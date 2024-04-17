package article

import (
	"blogs/pkg/db"
	"blogs/pkg/filter"
	"context"
	"gorm.io/gorm"
)

type Repository struct {
	db        *gorm.DB
	tableName string
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db, tableName: "article"}
}

func (a *Repository) Delete(ctx context.Context, f *filter.Chain) error {
	article := &ArticleEntity{}
	clauses, err := db.Build(f.GetExpression(), false)
	if err != nil {
		return err
	}
	return a.db.Table(a.tableName).Clauses(clauses...).Delete(article).Error
}

func (a *Repository) Create(ctx context.Context, article *ArticleEntity) error {
	return a.db.Table(a.tableName).Create(article).Error
}

func (a *Repository) Get(ctx context.Context, f *filter.Chain) (*ArticleEntity, error) {
	article := &ArticleEntity{}
	clauses, err := db.Build(f.GetExpression(), false)
	if err != nil {
		return nil, err
	}
	res := a.db.Table(a.tableName).Clauses(clauses...).Find(article)
	if res.Error != nil {
		return nil, err
	}
	return article, nil
}

func (a *Repository) Update(ctx context.Context, f *filter.Chain, data map[string]interface{}) error {
	clauses, err := db.Build(f.GetExpression(), false)
	if err != nil {
		return err
	}
	res := a.db.Table(a.tableName).Clauses(clauses...).Updates(data)
	if res.Error != nil {
		return err
	}
	return nil
}
