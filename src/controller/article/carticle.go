package article

import (
	"blogs/pkg/log"
	"blogs/pkg/resp"
	"blogs/src/model/marticle"
	"github.com/gin-gonic/gin"
	"strconv"
)

type CArticle struct {
}

type addArticleForm struct {
	Title   string `form:"title" binding:"required"`
	Content string `form:"content" binding:"required"`
}

// 添加文章
func (CArticle) Add(ctx *gin.Context) {
	var articleForm addArticleForm
	if err := ctx.ShouldBind(&articleForm); err != nil {
		log.Info(ctx, "params error", log.Fields{"params": articleForm, "error": err.Error()})
		resp.ResultRender(ctx, resp.ErrParams, nil)
		return
	}

	article := &marticle.Article{}
	article.Content = articleForm.Content
	article.Title = articleForm.Title
	article.UserId = 1 //todo 待完善session用户数据
	id, err := marticle.Add(ctx, article)
	if err != nil {
		resp.ResultRender(ctx, resp.ErrFailed, nil)
		return
	}
	result := map[string]interface{}{
		"id": id,
	}
	resp.ResultRender(ctx, nil, result)
}

type updateArticleForm struct {
	Id      uint   `form:"id" binding:"required"`
	Title   string `form:"title" binding:"required"`
	Content string `form:"content" binding:"required"`
}

// 更新文章
func (CArticle) Update(ctx *gin.Context) {
	var articleForm updateArticleForm
	if err := ctx.ShouldBind(&articleForm); err != nil {
		log.Info(ctx, "params error", log.Fields{"params": articleForm, "error": err.Error()})
		resp.ResultRender(ctx, resp.ErrParams, nil)
		return
	}

	err := marticle.Update(ctx, articleForm.Id, articleForm.Title, articleForm.Content)
	if err != nil {
		log.Error(ctx, "update article error", log.Fields{"params": articleForm, "error": err.Error()})
		resp.ResultRender(ctx, resp.ErrFailed, nil)
		return
	}
	resp.ResultRender(ctx, nil, nil)
}

// 删除文章
func (CArticle) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.PostForm("id"))
	if id <= 0 {
		log.Info(ctx, "params error", log.Fields{})
		resp.ResultRender(ctx, resp.ErrParams, nil)
		return
	}

	err := marticle.DeleteById(ctx, uint(id))
	if err != nil {
		log.Error(ctx, "delete article error", log.Fields{"params": id, "error": err.Error()})
		resp.ResultRender(ctx, resp.ErrFailed, nil)
		return
	}
	resp.ResultRender(ctx, nil, nil)
}

// 获取文章
func (CArticle) Get(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.DefaultQuery("id", ""))
	if id <= 0 {
		log.Info(ctx, "params error", log.Fields{"params": map[string]interface{}{"id": id}})
		resp.ResultRender(ctx, resp.ErrParams, nil)
		return
	}

	article, err := marticle.GetById(ctx, uint(id))
	if err != nil {
		log.Error(ctx, "get article error", log.Fields{"params": id, "error": err.Error()})
		resp.ResultRender(ctx, resp.ErrFailed, nil)
		return
	}
	result := make(map[string]interface{})
	if article != nil {
		result["id"] = article.ID
		result["title"] = article.Title
		result["content"] = article.Content
		result["user_id"] = article.UserId
	}
	resp.ResultRender(ctx, nil, result)
}
