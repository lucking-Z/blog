package user

import (
	"blogs/pkg/crypto/md5"
	"blogs/pkg/log"
	"blogs/pkg/resp"
	"blogs/src/model/muser"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type Cuser struct {
}

type loginForm struct {
	Account  string `form:"account" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func (Cuser) Login(ctx *gin.Context) {
	var loginForm loginForm
	if err := ctx.ShouldBind(&loginForm); err != nil {
		log.Info(ctx, "params error", log.Fields{"params": loginForm, "error": err.Error()})
		resp.ResultRender(ctx, resp.ErrParams, nil)
		return
	}

	user, err := muser.GetByAccount(ctx, loginForm.Account)

	if err != nil {
		log.Info(ctx, "get user error", log.Fields{"params": loginForm, "error": err.Error()})
		resp.ResultRender(ctx, resp.ErrFailed, nil)
		return
	}

	if user == nil || user.PassWord != md5.Md5Str(loginForm.Password) {
		log.Info(ctx, "login password error", log.Fields{"params": loginForm, "user": user})
		resp.ResultRender(ctx, resp.ErrPassword, nil)
		return
	}

	sess := sessions.Default(ctx)
	sess.Set("login_user", user)
	if err := sess.Save(); err != nil {
		log.Info(ctx, "login session error", log.Fields{"params": loginForm, "user": user, "error": err.Error()})
		resp.ResultRender(ctx, resp.ErrFailed, nil)
		return
	}

	resp.ResultRender(ctx, nil, nil)
}

func (Cuser) Logout(ctx *gin.Context) {
	sess := sessions.Default(ctx)
	sess.Delete("login_user")
	if err := sess.Save(); err != nil {
		log.Info(ctx, "logout session error", log.Fields{"error": err.Error()})
		resp.ResultRender(ctx, resp.ErrFailed, nil)
		return
	}

	resp.ResultRender(ctx, nil, nil)
}

type createUserForm struct {
	Account  string `form:"account" binding:"required"`
	Name     string `form:"name" binding:"required"`
	Password string `form:"password" binding:"required"`
}

// 创建用户
func (Cuser) Create(ctx *gin.Context) {
	var createUserForm createUserForm
	if err := ctx.ShouldBind(&createUserForm); err != nil {
		log.Info(ctx, "params error", log.Fields{"params": createUserForm, "error": err.Error()})
		resp.ResultRender(ctx, resp.ErrParams, nil)
		return
	}

	if u, _ := muser.GetByAccount(ctx, createUserForm.Account); u != nil {
		log.Info(ctx, "user exists", log.Fields{"params": createUserForm})
		resp.ResultRender(ctx, resp.ErrAccountExists, nil)
		return
	}

	userInfo := &muser.User{}
	userInfo.PassWord = md5.Md5Str(createUserForm.Password)
	userInfo.Account = createUserForm.Account
	userInfo.Name = createUserForm.Name

	id, err := muser.Add(ctx, userInfo)

	if err != nil {
		log.Info(ctx, "create user error", log.Fields{"params": createUserForm})
		resp.ResultRender(ctx, resp.ErrFailed, nil)
		return
	}

	result := map[string]interface{}{
		"id": id,
	}

	resp.ResultRender(ctx, nil, result)
}
