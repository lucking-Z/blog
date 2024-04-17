package article

import (
	"blogs/pkg/resp"
	"blogs/src/model/marticle"
	"github.com/gin-gonic/gin"
	"strconv"
)

type addArticleForm struct {
	Title   string `form:"title" binding:"required"`
	Content string `form:"content" binding:"required"`
}

func Add(ctx *gin.Context) {
	var articleForm addArticleForm
	if err := ctx.ShouldBind(&articleForm); err != nil {
		resp.ResultRender(ctx, resp.ErrParams, nil)
		return
	}
	article := &marticle.Article{}
	article.Content = articleForm.Content
	article.Title = articleForm.Title
	article.UserId = 1
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

func Update(ctx *gin.Context) {
	var articleForm updateArticleForm
	if err := ctx.ShouldBind(&articleForm); err != nil {
		resp.ResultRender(ctx, resp.ErrParams, nil)
		return
	}

	err := marticle.Update(ctx, articleForm.Id, articleForm.Title, articleForm.Content)
	if err != nil {
		resp.ResultRender(ctx, resp.ErrFailed, nil)
		return
	}
	resp.ResultRender(ctx, nil, nil)
}

func Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.PostForm("id"))
	if id <= 0 {
		resp.ResultRender(ctx, resp.ErrParams, nil)
		return
	}

	err := marticle.DeleteById(ctx, uint(id))
	if err != nil {
		resp.ResultRender(ctx, resp.ErrFailed, nil)
		return
	}
	resp.ResultRender(ctx, nil, nil)
}

func Get(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.DefaultQuery("id", ""))
	if id <= 0 {
		resp.ResultRender(ctx, resp.ErrParams, nil)
		return
	}

	article, err := marticle.GetById(ctx, uint(id))
	if err != nil {
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
