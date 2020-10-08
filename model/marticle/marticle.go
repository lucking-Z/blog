package marticle

import (
	"blogs/utils/db"
	"blogs/utils/log"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//删除文章
func DeleteById(ctx *gin.Context, id uint) error {
	if id == 0 {
		log.Info(ctx, "delete article id is empty", nil)
		return errors.New("delete article id is empty")
	}
	article := &Article{}
	result := db.Instance.Where("id = ?", id).Delete(article)

	return result.Error
}

//添加文章
func Add(ctx *gin.Context, article *Article) (uint, error) {
	if article.Content == "" || article.Title == "" || article.UserId == 0 {
		log.Info(ctx, "add article params error:", log.Fields{"params": article})
		return 0, errors.New(fmt.Sprintf("params error: %v", article))
	}

	result := db.Instance.Create(article)

	if result.Error != nil {
		return 0, result.Error
	}

	return article.ID, nil
}

//获取文章
func GetById(ctx *gin.Context, id uint) (*Article, error) {
	article := &Article{}
	result := db.Instance.First(article, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Info(ctx, "GetById article not found", log.Fields{"params": id})
		return nil, nil
	}

	if result.Error != nil {
		log.Error(ctx, "GetById article error", log.Fields{"params": id, "error":result.Error.Error()})
		return nil, result.Error
	}

	return article, nil
}

//更新文章
func Update(ctx *gin.Context, id uint, title string, content string) error {
	if id == 0 || title == "" || content == "" {
		params := map[string]interface{}{
			"id":      id,
			"title":   title,
			"content": content,
		}
		log.Info(ctx, "update article params error:", log.Fields{"params": params})
		return errors.New(fmt.Sprintf("params error: %v", params))
	}
	article := &Article{}
	article.Content = content
	article.Title = title
	result := db.Instance.Model(article).Where("id = ?", id).Updates(article)
	return result.Error
}
